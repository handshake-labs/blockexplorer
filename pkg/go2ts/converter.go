package go2ts

import "io"

type Converter struct {
	registry  *Registry
	extractor *Extractor
	renderer  *Renderer
}

func NewConverter(writer io.Writer) *Converter {
	registry := NewRegistry()
	extractor := NewExtractor(registry)
	renderer := NewRenderer(writer, registry)
	return &Converter{registry, extractor, renderer}
}

func (c *Converter) Register(v interface{}) int {
	return c.extractor.Register(v)
}

func (c *Converter) RegisterWithName(v interface{}, name string) int {
	return c.extractor.RegisterWithName(v, name)
}

func (c *Converter) RegisterDesc(d Desc) int {
	return c.extractor.RegisterDesc(d)
}

func (c *Converter) RegisterDescWithName(d Desc, name string) int {
	return c.extractor.RegisterDescWithName(d, name)
}

func (c *Converter) Render() {
	c.registry.CollapseNils()
	c.registry.GenerateNames()
	c.renderer.Render()
}
