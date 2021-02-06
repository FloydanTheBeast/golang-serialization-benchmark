package main

import (
	"encoding/binary"
	"testing"

	"github.com/vmihailenco/msgpack"
)

func BenchmarkMessagepackMarshal(b *testing.B) {
	testData := SampleData{"A relatively big string data", 1234567, []float32{123.456, 3.14159, 2.718}}

	for i := 0; i < b.N; i++ {
		msgpack.Marshal(testData)
	}

	val, _ := msgpack.Marshal(testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkMessagepackUnmarshal(b *testing.B) {
	marshaledData := []byte{131, 171, 83, 116, 114, 105, 110, 103, 70, 105, 101, 108, 100, 188, 65, 32, 114,
		101, 108, 97, 116, 105, 118, 101, 108, 121, 32, 98, 105, 103, 32, 115, 116, 114, 105, 110,
		103, 32, 100, 97, 116, 97, 168, 73, 110, 116, 70, 105, 101, 108, 100, 206, 0, 18, 214, 135,
		170, 65, 114, 114, 97, 121, 70, 105, 101, 108, 100, 147, 202, 66, 246, 233, 121, 202, 64, 73,
		15, 208, 202, 64, 45, 243, 182}

	var actualData SampleData

	for i := 0; i < b.N; i++ {
		msgpack.Unmarshal(marshaledData, &actualData)
	}
}
