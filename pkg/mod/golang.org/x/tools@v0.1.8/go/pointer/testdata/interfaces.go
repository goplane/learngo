//go:build ignore
// +build ignore

package main

type I interface {
	f()
}

type C int

func (*C) f() {}

type D struct{ ptr *int }

func (D) f() {}

type E struct{}

func (*E) f() {}

var a, b int

var unknown bool // defeat dead-code elimination

func interface1() {
	var i interface{} = &a
	var j interface{} = D{&b}
	k := j
	if unknown {
		k = i
	}

	print(i) // @types *int
	print(j) // @types D
	print(k) // @types *int | D

	print(i.(*int)) // @pointsto command-line-arguments.a
	print(j.(*int)) // @pointsto
	print(k.(*int)) // @pointsto command-line-arguments.a

	print(i.(D).ptr) // @pointsto
	print(j.(D).ptr) // @pointsto command-line-arguments.b
	print(k.(D).ptr) // @pointsto command-line-arguments.b
}

func interface2() {
	var i I = (*C)(&a)
	var j I = D{&a}
	k := j
	if unknown {
		k = i
	}

	print(i) // @types *C
	print(j) // @types D
	print(k) // @types *C | D
	print(k) // @pointsto makeinterface:command-line-arguments.D | makeinterface:*command-line-arguments.C

	k.f()
	// @calls command-line-arguments.interface2 -> (*command-line-arguments.C).f
	// @calls command-line-arguments.interface2 -> (command-line-arguments.D).f

	print(i.(*C))    // @pointsto command-line-arguments.a
	print(j.(D).ptr) // @pointsto command-line-arguments.a
	print(k.(*C))    // @pointsto command-line-arguments.a

	switch x := k.(type) {
	case *C:
		print(x) // @pointsto command-line-arguments.a
	case D:
		print(x.ptr) // @pointsto command-line-arguments.a
	case *E:
		print(x) // @pointsto
	}
}

func interface3() {
	// There should be no backflow of concrete types from the type-switch to x.
	var x interface{} = 0
	print(x) // @types int
	switch x.(type) {
	case int:
	case string:
	}
}

func interface4() {
	var i interface{} = D{&a}
	if unknown {
		i = 123
	}

	print(i) // @types int | D

	j := i.(I)       // interface narrowing type-assertion
	print(j)         // @types D
	print(j.(D).ptr) // @pointsto command-line-arguments.a

	var l interface{} = j // interface widening assignment.
	print(l)              // @types D
	print(l.(D).ptr)      // @pointsto command-line-arguments.a

	m := j.(interface{}) // interface widening type-assertion.
	print(m)             // @types D
	print(m.(D).ptr)     // @pointsto command-line-arguments.a
}

// Interface method calls and value flow:

type J interface {
	f(*int) *int
}

type P struct {
	x int
}

func (p *P) f(pi *int) *int {
	print(p)  // @pointsto p@i5p:6
	print(pi) // @pointsto i@i5i:6
	return &p.x
}

func interface5() {
	var p P // @line i5p
	var j J = &p
	var i int      // @line i5i
	print(j.f(&i)) // @pointsto p.x@i5p:6
	print(&i)      // @pointsto i@i5i:6

	print(j) // @pointsto makeinterface:*command-line-arguments.P
}

// @calls command-line-arguments.interface5 -> (*command-line-arguments.P).f

func interface6() {
	f := I.f
	print(f) // @pointsto (command-line-arguments.I).f$thunk
	f(new(struct{ D }))
}

// @calls command-line-arguments.interface6 -> (command-line-arguments.I).f$thunk
// @calls (command-line-arguments.I).f$thunk -> (*struct{command-line-arguments.D}).f

func main() {
	interface1()
	interface2()
	interface3()
	interface4()
	interface5()
	interface6()
}
