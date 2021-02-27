package main

import (
	"fmt"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/routes"
)

func main() {
	router := routes.MapRoutes()
	router.Run(fmt.Sprintf(":%v", "8080"))
}
