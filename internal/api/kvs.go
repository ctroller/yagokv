package api

import (
	"net/http"

	"github.com/ctroller/yagokv/internal/inject"
)

func ApiKvsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getHandler(w, r)
		case http.MethodPut:
		case http.MethodPost:
			setHandler(w, r)
		case http.MethodDelete:
			deleteHandler(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	}
}

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

func deleteHandler(w http.ResponseWriter, r *http.Request) {
	key := getKeyFromPath(r)
	inject.App.Storage.Delete(key)
	w.WriteHeader(http.StatusOK)
}
