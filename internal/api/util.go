package api

import "net/http"

func getKeyFromPath(r *http.Request) string {
	return r.URL.Path[len("/api/"):]
}
