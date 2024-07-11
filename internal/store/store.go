package store

import "sync"

var (
	requests  sync.Map
	responses sync.Map
)

func SaveRequest(id string, req interface{}) {
	requests.Store(id, req)
}

func SaveResponse(id string, res interface{}) {
	responses.Store(id, res)
}
