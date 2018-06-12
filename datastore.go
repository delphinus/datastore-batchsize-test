package main

import (
	"context"
	"math/rand"

	"google.golang.org/appengine/datastore"
)

const (
	entityNum = 5000
	maxAge    = 100
	batchSize = 500
	kind      = "datastore-batchsize-test.User"
)

// User is a sample user struct
type User struct {
	Age int8 `datastore:",noindex"`
}

func createEntities(ctx context.Context) error {
	keys := make([]*datastore.Key, entityNum)
	users := make([]*User, entityNum)
	for i := 0; i < entityNum; i++ {
		keys[i] = datastore.NewIncompleteKey(ctx, kind, nil)
		users[i] = &User{
			Age: int8(rand.Intn(maxAge)),
		}
	}
	if _, err := datastore.PutMulti(ctx, keys, users); err != nil {
		return err
	}
	return nil
}

func averageAge(ctx context.Context) (float64, error) {
	q := datastore.NewQuery(kind)
	return doQuery(ctx, q)
}

func averageAgeWithBatchSize(ctx context.Context) (float64, error) {
	q := datastore.NewQuery(kind).BatchSize(batchSize)
	return doQuery(ctx, q)
}

func doQuery(ctx context.Context, q *datastore.Query) (float64, error) {
	users := make([]*User, entityNum)
	if _, err := q.GetAll(ctx, &users); err != nil {
		return 0, err
	}
	var sum int64
	for _, u := range users {
		sum += int64(u.Age)
	}
	return float64(sum) / float64(len(users)), nil
}
