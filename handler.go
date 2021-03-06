package main

import (
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/log"
)

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		statusResponse(w, http.StatusOK)
	})
}

func create() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := appengine.NewContext(r)
		if err := createEntities(ctx); err != nil {
			log.Errorf(ctx, "%v", err)
			statusResponse(w, http.StatusInternalServerError)
			return
		}
		statusResponse(w, http.StatusOK)
	})
}

func calcAverageAge() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	})
}
