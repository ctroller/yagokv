package api

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
)

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := getKeyFromPath(r)
	inject.App.Storage.Delete(key)
	w.WriteHeader(http.StatusOK)
}
