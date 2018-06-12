package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func index(w http.ResponseWriter, r *http.Request) {
	statusResponse(w, http.StatusOK)
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		statusResponse(w, http.StatusBadRequest)
		return
	}
	ctx := appengine.NewContext(r)
	if err := createEntities(ctx); err != nil {
		log.Errorf(ctx, "%v", err)
		statusResponse(w, http.StatusInternalServerError)
		return
	}
	statusResponse(w, http.StatusOK)
}

func calcAverageAge(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		statusResponse(w, http.StatusBadRequest)
		return
	}
	var avg float64
	var err error
	ctx := appengine.NewContext(r)
	if err = r.ParseForm(); err != nil {
		log.Warningf(ctx, "%v", err)
		return
	}
	if r.FormValue("batchsize") == "" {
		avg, err = averageAge(ctx)
	} else {
		avg, err = averageAgeWithBatchSize(ctx)
	}
	if err != nil {
		log.Errorf(ctx, "%v", err)
		statusResponse(w, http.StatusInternalServerError)
		return
	}
	log.Infof(ctx, "average: %.2f", avg)
	statusResponse(w, http.StatusOK)
}
