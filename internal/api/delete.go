package api

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
)

func DeleteHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Query().Get("key")
		inject.App.Storage.Delete(key)
		w.WriteHeader(http.StatusOK)
	})
}
