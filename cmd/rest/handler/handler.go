package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"reflect"

	"github.com/handshake-labs/blockexplorer/cmd/rest/actions"
	"github.com/handshake-labs/blockexplorer/pkg/db"
)

type Handler struct {
	db  *db.Queries
	els map[string]func(http.ResponseWriter, *http.Request)
}

func NewHandler(db *db.Queries, acs map[string]interface{}) *Handler {
	els := make(map[string]func(http.ResponseWriter, *http.Request))
	for name, ac := range acs {
		tparams := reflect.TypeOf(ac).In(1).Elem()
		vf := reflect.ValueOf(ac)
		els[name] = func(w http.ResponseWriter, r *http.Request) {
			query := r.URL.Query()
			vparams := reflect.New(tparams)
			for i := 0; i < tparams.NumField(); i++ {
				q := query[tparams.Field(i).Tag.Get("json")]
				if len(q) != 1 {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
				if err := unmarshalParam([]byte(q[0]), vparams.Elem().Field(i)); err != nil {
					w.WriteHeader(http.StatusBadRequest)
					return
				}
			}
			result := vf.Call([]reflect.Value{
				reflect.ValueOf(actions.NewContext(db, r)),
				vparams,
			})
			if !result[1].IsNil() {
				log.Println(result[1].Interface())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			data, err := json.Marshal(result[0].Interface())
			if err != nil {
				log.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(data)
		}
	}
	return &Handler{db, els}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	name := r.URL.Path
	if name == "/" {
		w.WriteHeader(http.StatusOK)
		return
	}
	if e, ok := h.els[name]; ok {
		e(w, r)
		return
	}
	w.WriteHeader(http.StatusNotFound)
}
