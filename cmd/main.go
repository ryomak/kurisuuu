package main

import (
	"github.com/ryomak/kurisuuu/adapter"
)

func main() {
	r := adapter.New()
	r.Route()
	r.Run("8080")
}
