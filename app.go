package main

import (
	"net/http"
)

func init() {
	http.Handle("/", index())
	http.Handle("/create", create())
	http.Handle("/calc", calcAverageAge())
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
