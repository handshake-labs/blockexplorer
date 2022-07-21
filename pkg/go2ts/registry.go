package go2ts

import (
	"reflect"
	"sort"
)

type Registry struct {
	types map[reflect.Type]int
	descs []Desc
	names map[int]string
}

func NewRegistry() *Registry {
	return &Registry{
		types: make(map[reflect.Type]int, 0),
		descs: make([]Desc, 0),
		names: make(map[int]string, 0),
	}
}

func (r *Registry) Register(t reflect.Type) (int, bool) {
	if t != nil {
		if tid, ok := r.types[t]; ok {
			return tid, false
		}
		r.types[t] = len(r.descs)
	}
	r.descs = append(r.descs, nil)
	return len(r.descs) - 1, true
}

func (r *Registry) SetDesc(tid int, d Desc) {
	r.descs[tid] = d
}

func (r *Registry) GetDesc(tid int) Desc {
	return r.descs[tid]
}

func (r *Registry) SetName(tid int, name string) bool {
	for _, n := range r.names {
		if n == name {
			return false
		}
	}
	r.names[tid] = name
	return true
}

func (r *Registry) GetName(tid int) string {
	return r.names[tid]
}

func (r *Registry) GetNamed() []int {
	tids := make([]int, 0, len(r.names))
	for tid, _ := range r.names {
		tids = append(tids, tid)
	}
	sort.IntSlice(tids).Sort()
	return tids
}

func (r *Registry) CollapseNils() {
	c := true
	for c {
		c = false
		for tid, d := range r.descs {
			n := false
			switch dd := d.(type) {
			case *Boolean:
				n = !dd.True && !dd.False
			case *Dict:
				n = r.descs[dd.Elem] == nil
			case *Array:
				n = r.descs[dd.Elem] == nil
			case *Maybe:
				n = r.descs[dd.Elem] == nil
			}
			if n {
				r.descs[tid] = nil
				c = true
			}
		}
	}
	for _, d := range r.descs {
		switch dd := d.(type) {
		case *Record:
			props := make([]Property, 0)
			for _, prop := range dd.Props {
				if r.descs[prop.Elem] != nil {
					props = append(props, prop)
				}
			}
			dd.Props = props
		}
	}
}

func (r *Registry) GenerateNames() {
	for t, tid := range r.types {
		// TODO: make stable
		if t != nil && t.PkgPath() != "" && r.descs[tid] != nil && r.names[tid] == "" {
			r.SetName(tid, t.Name())
		}
	}
}
