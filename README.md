# Test for BatchSize property for querying Datastore

Datastore shows an error below in tracing if you are querying large amount entities and do not use `BatchSize` property.

| Issue | Description | Recommendation |
|---|---|---|
| Many datastore.next() calls. | Your app made 193 remote procedure calls to datastore.next() while processing this request. This was likely due the use of 20 as query batch size. | Increase the value of query batch size to reduce the number of datastore.next() calls |

This repository shows that and its solution.

## Usage

```sh
# deploy this app
gcloud app deploy --version 1

# index always returns OK
curl 'https://datastore-batchsize-test-dot-example-com.appspot.com'
OK

# create creates 5000 entities
curl -X POST 'https://datastore-batchsize-test-dot-example-com.appspot.com/create'
OK

# get all entities without BatchSize() (it takes a bit long time)
curl 'https://datastore-batchsize-test-dot-example-com.appspot.com/calc'
OK

# when with BatchSize(), it finishes faster!
curl 'https://datastore-batchsize-test-dot-example-com.appspot.com/calc?batchsize=1'
OK
```
