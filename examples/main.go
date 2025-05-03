package main

import (
	"fmt"

	"github.com/pizzacream/enx"
	_ "github.com/pizzacream/enx/autoload"
)

type MyStruct struct {
	Value string
}

func main() {
	port := enx.GetInt("PORT", 3000)
	maxConnections := enx.GetInt("MAX_CONNECTIONS", 10)
	myStruct := enx.MustGetStruct[MyStruct]("MY_STRUCT")

	fmt.Println(port, maxConnections, myStruct)
}
