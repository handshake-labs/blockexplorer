package actions

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/handshake-labs/blockexplorer/pkg/db"
)

type Context struct {
	context.Context
	db *db.Queries
}

type Action struct {
	Params   reflect.Type
	Result   reflect.Type
	Function reflect.Value
}

func NewAction(function interface{}) *Action {
	t := reflect.TypeOf(function)
	if t.NumIn() != 2 || t.NumOut() != 2 {
		panic("must have two inputs and two outputs")
	}
	if t.In(0) != reflect.TypeOf(&Context{}) {
		panic("must be the Context")
	}
	if !t.Out(1).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		panic("must be an error")
	}
	return &Action{
		ptrToStruct(t.In(1)),
		ptrToStruct(t.Out(0)),
		reflect.ValueOf(function),
	}
}

func (a *Action) parseParams(r *http.Request) (*reflect.Value, error) {
	query := r.URL.Query()
	params := reflect.New(a.Params)
	for i := 0; i < a.Params.NumField(); i++ {
		value := query[a.Params.Field(i).Tag.Get("json")]
		if len(value) != 1 {
			return nil, fmt.Errorf("bad params value %v", value)
		}
		if err := paramIntoValue(value[0], params.Elem().Field(i)); err != nil {
			return nil, err
		}
	}
	return &params, nil
}

func (a *Action) writeResult(w http.ResponseWriter, result interface{}) {
	data, err := json.Marshal(result)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func (a *Action) BuildHandlerFunc(db *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := a.parseParams(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		out := a.Function.Call([]reflect.Value{
			reflect.ValueOf(&Context{r.Context(), db}),
			*params,
		})
		if !out[1].IsNil() {
			log.Println(out[1].Interface().(error))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		a.writeResult(w, out[0].Interface())
	}
}
