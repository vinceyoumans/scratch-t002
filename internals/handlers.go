package internals

import (
	"github.com/golang-jwt/jwt/v5"
	"html/template"
	"net/http"
	su "t002/setup"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	// Simulate user login
	claims := NewClaims("john_doe")

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, "Failed to create token", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "jwt_token",
		Value:    tokenString,
		HttpOnly: true,
		Secure:   true, // Set to false in development
	})

	http.Redirect(w, r, "/protected", http.StatusSeeOther)
}

func Landing(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.ParseFiles("theme/lo001/templates/landing.html"))
	t := template.Must(template.ParseFiles(su.UrlToTemplate + "templates/landing.html"))

	t.Execute(w, nil)
}

func LandingAuth(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.ParseFiles("theme/lo001/templates/landing-auth.html"))
	t := template.Must(template.ParseFiles(su.UrlToTemplate + "templates/landing-auth.html"))

	t.Execute(w, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	//t := template.Must(template.ParseFiles("theme/lo001/templates/login.html"))
	t := template.Must(template.ParseFiles(su.UrlToTemplate + "templates/login.html"))

	t.Execute(w, nil)
}
