package inject

import "github.com/ctroller/yagokv/pkg/kvs"

type Application struct {
	Storage *kvs.Storage
}

var App Application
