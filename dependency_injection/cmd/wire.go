//go:build wireinject
// +build wireinject

package main

import "learngo/dependency_injection/pkg"
import "github.com/google/wire"

func InitializeEvent(phrase string) (pkg.Event, error) {
	panic(wire.Build(pkg.Providers))
}
