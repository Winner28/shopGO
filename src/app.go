package main

import (
	"handlers"
	// psql
	_ "github.com/lib/pq"
)

func main() {
	handlers.Init()
}
