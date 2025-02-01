package api

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
)

func setHandler(w http.ResponseWriter, r *http.Request) {
	key := getKeyFromPath(r)
	val := r.URL.Query().Get("val")

	err := inject.App.Storage.Set(key, val)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
