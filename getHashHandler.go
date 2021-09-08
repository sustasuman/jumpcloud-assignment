package main

import (
	"net/http"
	"strconv"
)

func GetHashHttp(w http.ResponseWriter, r *http.Request) {
	id := GetField(r, 0)
	idNum, error := strconv.ParseInt(id, 10, 32)
	if error != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	encodedString := CounterMap.Get(int(idNum))
	if encodedString == "" {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(encodedString))
	}
}
