package main

import (
	"encoding/binary"
	"testing"

	pb "./pb"
	"github.com/golang/protobuf/proto"
)

func BenchmarkProtocolBuffersMarshal(b *testing.B) {
	testData := &pb.SampleProtobuffData{
		StringField: "A relatively big string data",
		IntField:    1234567,
		ArrayField:  []float32{123.456, 3.14159, 2.718},
	}

	for i := 0; i < b.N; i++ {
		proto.Marshal(testData)
	}

	val, _ := proto.Marshal(testData)
	b.ReportMetric(float64(binary.Size(val)), "bytes")
}

func BenchmarkProtocolBuffersUnmarshal(b *testing.B) {
	var marshaledData []byte = []byte{10, 28, 65, 32, 114, 101, 108, 97, 116, 105, 118, 101, 108, 121, 32, 98, 105, 103, 32, 115, 116, 114, 105, 110, 103, 32, 100, 97, 116, 97, 16, 135, 173, 75, 26, 12, 121, 233, 246, 66, 208, 15, 73, 64, 182, 243, 45, 64}
	actualData := &pb.SampleProtobuffData{}

	for i := 0; i < b.N; i++ {
		proto.Unmarshal(marshaledData, actualData)
	}
}
