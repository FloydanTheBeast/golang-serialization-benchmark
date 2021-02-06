package main

type SampleData struct {
	StringField string    `avro:"stringfield"`
	IntField    int       `avro:"intfield"`
	ArrayField  []float32 `avro:"arrayfield"`
}

func main() {}
