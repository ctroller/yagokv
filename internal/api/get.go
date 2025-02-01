package api

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	key := getKeyFromPath(r)
	val, err := inject.App.Storage.Get(key)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(val))
}
