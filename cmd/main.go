package main

import (
	"github.com/ryomak/kurisuuu/router"
)

func main() {
	r := router.New()
	r.Route("8080")
}
