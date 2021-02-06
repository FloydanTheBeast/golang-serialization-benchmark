package main

import (
	"encoding/binary"
	"testing"

	"github.com/hamba/avro"
)

var Schema = `{
		"type": "record",
		"name": "simple",
		"namespace": "org.hamba.avro",
		"fields" : [
			{"name": "stringfield", "type": "string"},
			{"name": "intfield", "type": "int"},
			{"name": "arrayfield", "type": {
				"type": "array",
				"items": "float"
			}}
		]
	}`

func BenchmarkAvroMarshal(b *testing.B) {
	schema, _ := avro.Parse(Schema)
	testData := SampleData{"A relatively big string data", 1234567, []float32{123.456, 3.14159, 2.718}}

	for i := 0; i < b.N; i++ {
		avro.Marshal(schema, &testData)
	}

	val, _ := avro.Marshal(schema, testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkAvroUnmarhsal(b *testing.B) {
	schema, _ := avro.Parse(Schema)
	marshaledData := []byte{56, 65, 32, 114, 101, 108, 97, 116, 105, 118, 101, 108, 121, 32, 98, 105, 103, 32, 115, 116, 114, 105, 110, 103, 32, 100, 97, 116, 97, 142, 218, 150, 1, 5, 24, 121, 233, 246, 66, 208, 15, 73, 64, 182, 243, 45, 64, 0}
	var actualData SampleData

	for i := 0; i < b.N; i++ {
		avro.Unmarshal(schema, marshaledData, &actualData)
	}
}
