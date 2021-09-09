package main

import (
	"testing"
	"time"
)

func TestUpdateAverage(t *testing.T) {
	want := ResponseStats{Average: 100, Total: 2, TotalTime: 200}
	now := time.Now()
	after100ms := now.Add(100 * time.Microsecond)
	actual := ResponseStats{Average: 0, Total: 0, TotalTime: 0}
	actual.updateAverage(now, after100ms)
	actual.updateAverage(now, after100ms)
	if want != actual {
		t.Fatal("Average did not match expected value ")
	}
}

func TestToJson(t *testing.T) {
	stat := ResponseStats{Average: 100, Total: 1, TotalTime: 100}
	rsw, _ := stat.toJson()
	want := "{\"total\":1,\"average\":100}"
	if want != string(rsw) {
		t.Fatal("ToJson output did not match expected output")
	}
}
