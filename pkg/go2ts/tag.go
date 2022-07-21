package go2ts

import (
	"reflect"
	"strings"
)

type Tag []string

func ParseTag(key string, st reflect.StructTag) *Tag {
	t := st.Get(key)
	if t == "-" {
		return nil
	}
	s := strings.Split(t, ",")
	return (*Tag)(&s)
}

func (tag *Tag) Name() string {
	return (*tag)[0]
}

func (tag *Tag) Contains(key string) bool {
	for i := 1; i < len(*tag); i++ {
		if (*tag)[i] == key {
			return true
		}
	}
	return false
}
