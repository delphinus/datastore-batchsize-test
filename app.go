package main

import (
	"net/http"
)

func init() {
	http.Handle("/", confirmMethod(http.MethodGet, index()))
	http.Handle("/create", confirmMethod(http.MethodPost, create()))
	http.Handle("/calc", confirmMethod(http.MethodGet, calcAverageAge()))
}

func confirmMethod(m string, h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != m {
			statusResponse(w, http.StatusMethodNotAllowed)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func statusResponse(w http.ResponseWriter, status int) {
	w.WriteHeader(status)
	_, _ = w.Write(statusText(status))
}

var statusTextMap = make(map[int][]byte)

func statusText(s int) []byte {
	if b, ok := statusTextMap[s]; ok {
		return b
	}
	statusTextMap[s] = []byte(http.StatusText(s))
	return statusTextMap[s]
}
