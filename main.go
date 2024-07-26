package main

import (
	"net/http"
	"wedding_invitation_service/routes"
)

func main() {
    routes.InitializeRoutes()
    http.ListenAndServe(":8080", nil)
}