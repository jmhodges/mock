package orig

type Adder interface {
	Inc(x int64)
}

type Thing struct {
	a Adder
}

func NewThing(a Adder) Thing {
	return Thing{a}
}

func (t Thing) CallInc(x int64) {
	t.a.Inc(x)
}

type HasVariadic interface {
	AllBoolVar(foo ...bool)
	AllInterfaceVar(foo ...interface{})
	SomeBoolVar(s string, foo ...bool)
	SomeInterfaceVar(s string, foo ...interface{})
}
