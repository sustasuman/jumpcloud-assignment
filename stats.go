package main

import (
	"encoding/json"
	"sync"
	"time"
)

type ResponseStats struct {
	sync.RWMutex
	Total     int64 `json:"total"`
	Average   int64 `json:"average"`
	TotalTime int64 `json:"-"`
}

func (responseStats *ResponseStats) countAvarage(start time.Time, end time.Time) {
	responseStats.RLock()
	defer responseStats.RUnlock()
	elapsed := end.Sub(start).Microseconds()
	responseStats.TotalTime = responseStats.TotalTime + elapsed
	responseStats.Total = responseStats.Total + 1
	responseStats.Average = responseStats.TotalTime / responseStats.Total
}

func (responseStats *ResponseStats) toJson() ([]byte, error) {
	responseStats.Lock()
	response := struct {
		Total   int64 `json:"total"`
		Average int64 `json:"average"`
	}{responseStats.Total, responseStats.Average}
	responseStats.Unlock()

	return json.Marshal(response)
}
