package main

import "github.com/google/wire"

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func ProviderMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World"
	return b
}

type Bar string

func ProvideBar(f Fooer) string {
	return f.Foo()
}

var Set = wire.NewSet(
	ProviderMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)),
	ProvideBar)
