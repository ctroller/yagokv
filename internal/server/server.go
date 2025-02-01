package server

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/api"
	"github.com/ctroller/yagokv/internal/inject"
	"github.com/ctroller/yagokv/internal/kvs"
)

func Setup() {
	inject.App = inject.Application{
		Storage: kvs.NewStorage(16),
	}

	http.HandleFunc("/api/get", api.GetHandler())
	http.HandleFunc("/api/set", api.SetHandler())
	http.HandleFunc("/api/delete", api.DeleteHandler())

	http.ListenAndServe(":8080", nil)
}
