package api

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
)

func GetHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		val, err := inject.App.Storage.Get(key)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(val))
	})
}
