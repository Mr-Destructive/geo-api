package main

import (
	"fmt"

	"github.com/coast-nav-api/infrastructure"
)

func main() {
	fmt.Println("Hello")
	noSqlHandler, err := infrastructure.NewMongoDBHandler()
	if err != nil {
		fmt.Println(err)
	}
	infrastructure.HandleRequest(noSqlHandler)

}
