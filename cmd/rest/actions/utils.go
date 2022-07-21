package actions

import (
	"encoding"
	"errors"
	"reflect"
	"strconv"
)

func ptrToStruct(t reflect.Type) reflect.Type {
	if t.Kind() != reflect.Ptr {
		panic("must be a pointer")
	}
	t = t.Elem()
	if t.Kind() != reflect.Struct {
		panic("must be a struct")
	}
	return t
}

func paramIntoValue(p string, v reflect.Value) error {
	if i, ok := v.Addr().Interface().(encoding.TextUnmarshaler); ok {
		return i.UnmarshalText([]byte(p))
	}
	switch v.Kind() {
	case reflect.String:
		v.SetString(p)
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(p, 10, v.Type().Bits())
		if n < 0 {
			return errors.New("negative value")
		}
		v.SetInt(n)
		return err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(p, 10, v.Type().Bits())
		v.SetUint(n)
		return err
	}
	return errors.New("unsupported type")
}
