package main

import (
	"github.com/google/wire"
	"github.com/lynndotconfig/hellogo/gowire/guide/foobar"
)

var MegaSet = wire.NewSet(foobar.SuperSet)
