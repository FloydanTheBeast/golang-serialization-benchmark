# Сравнение различных методов сериализации на языке Go

### Тестируемая структура

```Go
type SampleData struct {
	StringField string    `avro:"stringfield"`
	IntField    int       `avro:"intfield"`
	ArrayField  []float32 `avro:"arrayfield"`
}
```

### Результаты

```
BenchmarkAvroMarshal-4                	 4761921	       229 ns/op	        48.0 bytes
BenchmarkAvroUnmarhsal-4              	 5217372	       233 ns/op
BenchmarkBinaryMarshal-4              	 3150049	       383 ns/op	        46.0 bytes
BenchmarkBinaryUnmarshal-4            	 3141084	       378 ns/op
BenchmarkJSONMarshal-4                	 1454550	       814 ns/op	       102 bytes
BenchmarkJSONUnmarhsal-4              	  600002	      2025 ns/op
BenchmarkMessagepackMarshal-4         	 1591530	       764 ns/op	        83.0 bytes
BenchmarkMessagepackUnmarshal-4       	 1247398	       943 ns/op
BenchmarkProtocolBuffersMarshal-4     	 5970151	       201 ns/op	        48.0 bytes
BenchmarkProtocolBuffersUnmarshal-4   	 4799544	       248 ns/op
BenchmarkXMLMarshal-4                 	  307724	      4031 ns/op	       202 bytes
BenchmarkXMLUnmarshal-4               	  109089	     11211 ns/op
BenchmarkYAMLMarshal-4                	  149996	      8021 ns/op	       100 bytes
BenchmarkYAMLUnmarhsal-4              	  125017	      9534 ns/op
```

Результаты также доступны в файле [benchmark-results](benchmark-results.txt)

> Работа выполнена для курса по сервис-ориентированным архитектурам <br/>
> Абдельсалам Шади Мазен, студент программной инженерии **НИУ ВШЭ** группы **БПИ191**
