package routes

import (
	"net/http"

	"github.com/ashiqsabith123/onepane-assesment/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/", handlers.GetPostDetails)
}
