// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"
	"github.com/google/wire"
	"github.com/lynndotconfig/hellogo/gowire/guide/foobar"
)

func initializeBaz(ctx context.Context) (foobar.Baz, error) {
	wire.Build(MegaSet)
	return foobar.Baz{}, nil
}

func main() {
}