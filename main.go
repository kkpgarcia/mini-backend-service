package main

import (
	"mini-backend-service/api"
	"net/http"
)

func main() {
	api.RegisterControllers()
	http.ListenAndServe(":3000", nil)
}
