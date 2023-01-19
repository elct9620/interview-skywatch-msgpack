package main

import (
	"fmt"

	"github.com/elct9620/interview-skywatch-msgpack/pkg/msgpack"
)

func main() {
	data, _ := msgpack.Marshal(map[string]int{"count": 10})
	fmt.Printf("From Map: %v\n", data)

	data, _ = msgpack.FromJSON([]byte(`{"count": 10}`))
	fmt.Printf("From JSON: %v\n", data)
}
