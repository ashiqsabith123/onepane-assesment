package main

import (
	"fmt"
	"net/http"

	"github.com/ashiqsabith123/onepane-assesment/routes"
)

func main() {

	routes.SetupRoutes()

	fmt.Println("Server runnig at port: 8080")
	err := http.ListenAndServe(":8089", nil)
	if err != nil {
		fmt.Println("Failed to start server")
		return
	}

}
