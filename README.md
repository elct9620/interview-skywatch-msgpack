(Interview) Skywatch Msgpack
===

[![Go](https://github.com/elct9620/interview-skywatch-msgpack/actions/workflows/go.yml/badge.svg)](https://github.com/elct9620/interview-skywatch-msgpack/actions/workflows/go.yml)

This is Skywatch Interview Assignment to implement a msgpack encoder.

## Usage

Install package

```bash
go get https://github.com/elct9620/interview-skywatch-msgpack
```

Include package to encode from json or struct

```go
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
```

## Capability

| Type            | Supported |
|-----------------|-----------|
| nil             | ✅
| true            | ✅
| false           | ✅
| positive fixint | ✅
| negative fixint | ✅
| uint8           | ✅
| uint16          | ✅
| uint32          | ✅
| uint64          | ✅
| int8            | ✅
| int16           | ✅
| int32           | ✅
| int64           | ✅
| float32         | ✅
| float64         | ✅
| fixstr          | ✅
| str8            | ✅
| str16           | ✅
| str32           | ✅
| fixarray        | ✅
| array16         | ✅
| array32         | ✅
| fixmap          | ✅
| map16           | ✅
| map32           | ✅
| bin8            | ❌
| bin16           | ❌
| bin32           | ❌
| fixext1         | ❌
| fixext2         | ❌
| fixext4         | ❌
| fixext8         | ❌
| fixext16        | ❌
| ext8            | ❌
| ext16           | ❌
| ext32           | ❌

## Limitation

### Unoptimized Size

For example, the JSON to Msgpack never uses `float32` because `json.Unmarshal` will use `float64` by default in Go

> The behavior is follow by msgpack playground generated data from JSON

### No strict type check

The package may overflow or use incorrect type due to the implementation only checking it by the `reflect` package but no more checks.

### No `bin` and `ext` support

The JSON didn't support the `bin` or `ext` without extra configuration therefore not implement it.
