package server

import (
	"fmt"
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
	"github.com/ctroller/yagokv/internal/kvs"
)

func Setup() {
	inject.App = inject.Application{
		Storage: kvs.NewStorage(16),
	}

	inject.App.Storage.Set("foo", "bar")
	inject.App.Storage.Set("bar", "42")

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		val, err := inject.App.Storage.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error: %v", err)
			return
		}

		fmt.Fprintf(w, "%v", val)
	})

	http.ListenAndServe(":8080", nil)
}
