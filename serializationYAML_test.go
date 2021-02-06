package main

import (
	"encoding/binary"
	"testing"

	"gopkg.in/yaml.v2"
)

const YAMLMarshaledData string = `str: A relatively big string data
int: 1234567
array:
- 123.456
- 3.14159
- 2.718
`

func BenchmarkYAMLMarshal(b *testing.B) {
	testData := SampleData{"A relatively big string data", 1234567, []float32{123.456, 3.14159, 2.718}}

	for i := 0; i < b.N; i++ {
		yaml.Marshal(testData)
	}

	val, _ := yaml.Marshal(&testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkYAMLUnmarhsal(b *testing.B) {
	marshaledData := []byte(YAMLMarshaledData)

	var actualData SampleData

	for i := 0; i < b.N; i++ {
		yaml.Unmarshal(marshaledData, &actualData)
	}
}
