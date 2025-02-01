package inject

import "github.com/ctroller/yagokv/internal/kvs"

type Application struct {
	Storage *kvs.Storage
}

var App Application
