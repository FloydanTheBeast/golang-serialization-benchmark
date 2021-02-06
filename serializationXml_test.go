package main

import (
	"encoding/binary"
	"encoding/xml"
	"testing"
)

func BenchmarkXMLMarshal(b *testing.B) {
	testData := SampleData{"A relatively big string data", 1234567, []float32{123.456, 3.14159, 2.718}}

	for i := 0; i < b.N; i++ {
		xml.Marshal(testData)
	}

	val, _ := xml.Marshal(testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkXMLUnmarshal(b *testing.B) {
	marshaledData := []byte("`<SampleData><StringField>A relatively big string data</StringField><IntField>1234567</IntField><ArrayField>123.456</ArrayField><ArrayField>3.14159</ArrayField><ArrayField>2.718</ArrayField></SampleData>`")
	var actualData SampleData

	for i := 0; i < b.N; i++ {
		xml.Unmarshal(marshaledData, &actualData)
	}
}
