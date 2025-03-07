package internals

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	// Simulate user login
	//claims := internals.NewClaims("john_doe")
	claims := NewClaims("john_doe")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//tokenString, err := token.SignedString(internals.jwtKey)
	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "Bearer %s", tokenString)
}
