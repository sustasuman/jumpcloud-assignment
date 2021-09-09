package main

import (
	"encoding/json"
	"net/http"
	"sync"
	"time"
)

var ResponseStat = ResponseStats{Average: 0.00, Total: 0}

type ResponseStats struct {
	sync.RWMutex
	Total     int64 `json:"total"`
	Average   int64 `json:"average"`
	TotalTime int64 `json:"-"`
}

func (responseStats *ResponseStats) updateAverage(start time.Time, end time.Time) {
	responseStats.Lock()
	defer responseStats.Unlock()
	elapsed := end.Sub(start).Microseconds()
	responseStats.TotalTime = responseStats.TotalTime + elapsed
	responseStats.Total = responseStats.Total + 1
	responseStats.Average = responseStats.TotalTime / responseStats.Total
}

func (responseStats *ResponseStats) toJson() ([]byte, error) {
	responseStats.RLock()
	defer responseStats.RUnlock()
	response := struct {
		Total   int64 `json:"total"`
		Average int64 `json:"average"`
	}{responseStats.Total, responseStats.Average}
	return json.Marshal(response)
}

func getStats(w http.ResponseWriter, r *http.Request) {
	response, _ := ResponseStat.toJson()
	w.Write(response)
}
