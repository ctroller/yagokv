package server

import (
	"log"
	"net/http"

	"github.com/ctroller/yagokv/internal/api"
	"github.com/ctroller/yagokv/internal/inject"
	"github.com/ctroller/yagokv/pkg/kvs"
)

func Setup() {
	inject.App = inject.Application{
		Storage: kvs.NewStorage(16),
	}

	http.HandleFunc("/api/v1/kvs/", api.ApiKvsHandler())

	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
