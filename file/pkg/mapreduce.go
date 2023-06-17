package pkg

import "strings"

type KeyValue struct {
	Key   string
	Value string
}

type MapReduce struct {
}

func NewMapReduce() *MapReduce {
	return &MapReduce{}
}

func (m *MapReduce) Map(data []string) []KeyValue {
	var keyValues []KeyValue
	for _, record := range data {
		split := strings.Split(record, ",")
		keyValues = append(keyValues, KeyValue{split[0], "1"})
	}
	return keyValues
}

func (m *MapReduce) Reduce(key string, values []string) KeyValue {
	return KeyValue{key, string(len(values))}
}

//data 传入的文件分块
func (m *MapReduce) MapReduce(data []string, mapper func(string) []KeyValue, reducer func(string, []string) KeyValue) {

}
