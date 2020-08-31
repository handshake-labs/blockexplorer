package handler

import (
	"encoding"
	"errors"
	"reflect"
	"strconv"
)

func unmarshalParam(t []byte, v reflect.Value) error {
	if i, ok := v.Addr().Interface().(encoding.TextUnmarshaler); ok {
		return i.UnmarshalText(t)
	}
	switch v.Kind() {
	case reflect.String:
		v.SetString(string(t))
		return nil
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(string(t), 10, v.Type().Bits())
		v.SetInt(n)
		return err
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		n, err := strconv.ParseUint(string(t), 10, v.Type().Bits())
		v.SetUint(n)
		return err
	}
	return errors.New("text unmarshal: unsupported type")
}
