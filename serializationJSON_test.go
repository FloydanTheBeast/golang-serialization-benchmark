package main

import (
	"encoding/binary"
	"encoding/json"
	"testing"
)

func BenchmarkJSONMarshal(b *testing.B) {
	testData := SampleData{"A relatively big string data", 1234567, []float32{123.456, 3.14159, 2.718}}

	for i := 0; i < b.N; i++ {
		json.Marshal(testData)
	}

	val, _ := json.Marshal(testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkJSONUnmarhsal(b *testing.B) {
	marshaledData := []byte(`{"StringField":"A relatively big string data","IntField":1234567,"ArrayField":[123.456,3.14159,2.718]}`)
	var actualData SampleData

	for i := 0; i < b.N; i++ {
		json.Unmarshal(marshaledData, &actualData)
	}
}
