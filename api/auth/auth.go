package auth

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type authController struct {
}

var jwtKey = []byte("my_secret_key")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func NewAuthController() *authController {
	return &authController{}
}

func (ac authController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/signin" {
		switch r.Method {
		case http.MethodGet:
			ac.signIn(w, r)
		default:
			w.WriteHeader(http.StatusNotImplemented)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}

/**
*	Device Registration?
 */

/**
*	Base Registration
 */
func (ac *authController) register(w http.ResponseWriter, r *http.Request) {

}

/**
*	Sign in using Username and Password
 */
func (ac *authController) signIn(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	err := json.NewDecoder(r.Body).Decode(&creds)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Should be getting from graphql
	expectedPassword, ok := users[creds.Username]

	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
