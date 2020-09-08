package go2ts

import (
	"encoding"
	"reflect"
)

type Extractor struct {
	registry *Registry
}

func NewExtractor(registry *Registry) *Extractor {
	return &Extractor{registry}
}

var textMarshalerType = reflect.TypeOf((*encoding.TextMarshaler)(nil)).Elem()

func (e *Extractor) Register(v interface{}) int {
	if t, ok := v.(reflect.Type); ok {
		return e.RegisterType(t)
	}
	if v, ok := v.(reflect.Value); ok {
		return e.RegisterType(v.Type())
	}
	return e.RegisterType(reflect.TypeOf(v))
}

func (e *Extractor) RegisterWithName(v interface{}, name string) int {
	tid := e.Register(v)
	e.registry.SetName(tid, name)
	return tid
}

func (e *Extractor) RegisterDesc(d Desc) int {
	tid, _ := e.registry.Register(nil)
	e.registry.SetDesc(tid, d)
	return tid
}

func (e *Extractor) RegisterDescWithName(d Desc, name string) int {
	tid := e.RegisterDesc(d)
	e.registry.SetName(tid, name)
	return tid
}

func (e *Extractor) RegisterType(t reflect.Type) int {
	tid, ok := e.registry.Register(t)
	if ok {
		e.registry.SetDesc(tid, e.ExtractType(t))
	}
	return tid
}

func (e *Extractor) ExtractType(t reflect.Type) Desc {
	if t.Implements(textMarshalerType) {
		return &String{}
	}
	switch t.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
		reflect.Float32, reflect.Float64:
		return &Number{}
	case reflect.String:
		return &String{}
	case reflect.Bool:
		return &Boolean{true, true}
	case reflect.Map:
		switch t.Key().Kind() {
		case reflect.String,
			reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
			reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		default:
			if !t.Key().Implements(textMarshalerType) {
				return nil
			}
		}
		return &Dict{e.RegisterType(t.Elem())}
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			p := reflect.PtrTo(t.Elem())
			if p.Implements(textMarshalerType) {
				return &String{}
			}
		}
		return &Array{e.RegisterType(t.Elem())}
	case reflect.Array:
		return &Array{e.RegisterType(t.Elem())}
	case reflect.Struct:
		return e.ExtractStruct(t)
	case reflect.Ptr:
		tid := e.RegisterType(t.Elem())
		if d, ok := e.registry.GetDesc(tid).(*Maybe); ok {
			return d
		}
		return &Maybe{tid}
	}
	return nil
}

func (e *Extractor) ExtractStruct(t reflect.Type) Desc {
	props := make([]Property, 0)
	pntag := make([]bool, 0)
	for i := 0; i < t.NumField(); i++ {
		sf := t.Field(i)
		if sf.PkgPath != "" {
			continue
		}
		tag := ParseTag("json", sf.Tag)
		if tag == nil {
			continue
		}
		name := tag.Name()
		if name == "" {
			pntag = append(pntag, false)
			name = sf.Name
		} else {
			pntag = append(pntag, true)
		}
		var tid int = -1
		if tag.Contains("string") {
			switch sf.Type.Kind() {
			case reflect.Bool,
				reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
				reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr,
				reflect.Float32, reflect.Float64,
				reflect.String:
				tid = e.RegisterDesc(&String{})
			}
		}
		optional := false
		if tid == -1 {
			tid = e.RegisterType(sf.Type)
			if tag.Contains("omitempty") {
				d := e.registry.GetDesc(tid)
				if _, ok := d.(*Record); !ok {
					optional = true
					if dd, ok := d.(*Maybe); ok {
						tid = dd.Elem
						d = e.registry.GetDesc(tid)
					}
					if dd, ok := d.(*Boolean); ok {
						tid = e.RegisterDesc(&Boolean{dd.True, false})
					}
				}
			}
		}
		props = append(props, Property{name, tid, optional})
	}
	fprops := make([]Property, 0)
F:
	for i1, p1 := range props {
		for i2, p2 := range props {
			if i1 != i2 && p1.Name == p2.Name && pntag[i2] {
				continue F
			}
		}
		fprops = append(fprops, p1)
	}
	return &Record{fprops}
}
