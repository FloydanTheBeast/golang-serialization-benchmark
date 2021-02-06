package main

import (
	"encoding/binary"
	"testing"

	binarySerializer "github.com/kelindar/binary"
)

func BenchmarkBinaryMarshal(b *testing.B) {
	testData := SampleData{"A relatively big string data", 1234567, []float32{123.456, 3.14159, 2.718}}

	for i := 0; i < b.N; i++ {
		binarySerializer.Marshal(testData)
	}

	val, _ := binarySerializer.Marshal(testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkBinaryUnmarshal(b *testing.B) {
	var marshaledData []byte = []byte{28, 65, 32, 114, 101, 108, 97, 116, 105, 118, 101, 108, 121, 32, 98, 105, 103, 32, 115, 116, 114, 105, 110, 103, 32, 100, 97, 116, 97, 142, 218, 150, 1, 3, 121, 233, 246, 66, 208, 15, 73, 64, 182, 243, 45, 64}
	var actualData SampleData

	for i := 0; i < b.N; i++ {
		binarySerializer.Unmarshal(marshaledData, &actualData)
	}
}
