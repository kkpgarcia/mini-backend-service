package api

import (
	"mini-backend-service/api/auth"
	"net/http"
)

func RegisterControllers() {
	authCont := auth.NewAuthController()

	//Sign in
	http.Handle("/signin", *authCont)
}
