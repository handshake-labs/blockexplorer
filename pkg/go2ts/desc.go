package go2ts

type Desc interface {
}

type Number struct{}

type String struct{}

type Boolean struct {
	True  bool
	False bool
}

type Dict struct {
	Elem int
}

type Array struct {
	Elem int
}

type Maybe struct {
	Elem int
}

type Property struct {
	Name     string
	Elem     int
	Optional bool
}

type Record struct {
	Props []Property
}
