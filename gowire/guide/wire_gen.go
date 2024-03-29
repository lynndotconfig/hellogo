// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"context"
	"github.com/lynndotconfig/hellogo/gowire/guide/foobar"
)

// Injectors from foobarbaz.go:

func initializeBaz(ctx context.Context) (foobar.Baz, error) {
	foo := foobar.ProvideFoo()
	bar := foobar.ProvideBar(foo)
	baz, err := foobar.ProvideBaz(ctx, bar)
	if err != nil {
		return foobar.Baz{}, err
	}
	return baz, nil
}

// foobarbaz.go:

func main() {
}
