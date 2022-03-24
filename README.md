# Staty

Very basic state management library for Go

## Installation

```shell
go get -v github.com/Streamer272/staty
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/Streamer272/staty"
)

func main() {
	store := staty.NewStore()

	store.Set("value1", "string")
	store.Set("value2", 1)
	store.Set("value3", true)

	fmt.Printf("value1: %v\n", store.Get("value1"))
	fmt.Printf("value2: %v\n", store.Get("value2"))
	fmt.Printf("value3: %v\n", store.Get("value3"))
}
```

## Contribution

This project is open to any contribution, just create a pull request,
and I will have a look at it as soon as I can.

## License

This project is licensed under [MIT](https://github.com/Streamer272/staty/blob/master/LICENSE) license.
