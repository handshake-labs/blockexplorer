package go2ts

import (
	"io"
	"strings"
)

type Renderer struct {
	writer   io.Writer
	registry *Registry
}

func NewRenderer(writer io.Writer, registry *Registry) *Renderer {
	return &Renderer{writer, registry}
}

func (r *Renderer) Write(s string) {
	r.writer.Write([]byte(s))
}

func (r *Renderer) RenderType(tid int) {
	if name := r.registry.GetName(tid); name != "" {
		r.Write(name)
		return
	}
	r.RenderDesc(r.registry.GetDesc(tid))
}

func (r *Renderer) RenderDesc(d Desc) {
	switch dd := d.(type) {
	case *Number:
		r.Write("number")
	case *String:
		r.Write("string")
	case *Boolean:
		if dd.True && dd.False {
			r.Write("boolean")
		} else if dd.True {
			r.Write("true")
		} else if dd.False {
			r.Write("false")
		} else {
			r.Write("never")
		}
	case *Dict:
		r.Write("{ [key: string]: ")
		r.RenderType(dd.Elem)
		r.Write(" }")
	case *Array:
		r.Write("Array<")
		r.RenderType(dd.Elem)
		r.Write(">")
	case *Maybe:
		r.RenderType(dd.Elem)
		r.Write(" | null")
	case *Record:
		r.Write("{\n")
		for _, prop := range dd.Props {
			name := prop.Name
			name = strings.Replace(name, `\`, `\\`, -1)
			name = strings.Replace(name, `"`, `\"`, -1)
			r.Write(`"` + name + `"`)
			if prop.Optional {
				r.Write("?: ")
			} else {
				r.Write(": ")
			}
			r.RenderType(prop.Elem)
			r.Write("\n")
		}
		r.Write("}")
	default:
		r.Write("never")
	}
}

func (r *Renderer) Render() {
	for _, tid := range r.registry.GetNamed() {
		r.Write("export type " + r.registry.GetName(tid) + " = ")
		r.RenderDesc(r.registry.GetDesc(tid))
		r.Write("\n")
	}
}
