package middleware

import (
	"errors"
	"net/http"
	"strings"
	"time"

	//usr "xm/pkg/user"

	jwt "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("supersecretkey")

type JWTClaim struct {
	Password string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

//var AuthenticatedUser usr.User

func GenerateJWT(email, password string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &JWTClaim{
		Email:    email,
		Password: password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//	tokenString, err := token.SignedString([]byte(confg.DefaulConfig.JwtSecret))

	tokenString, err := token.SignedString([]byte(jwtKey))
	return tokenString, err
}

func JWTAuth(original func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header["Authorization"]
		if authHeader == nil {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		authHeaderParts := strings.Split(authHeader[0], " ")
		if len(authHeaderParts) != 2 || strings.ToLower(authHeaderParts[0]) != "bearer" {
			http.Error(w, "not authorized", http.StatusUnauthorized)
			return
		}

		ok, _ := validateToken(authHeaderParts[1])
		if ok {
			original(w, r)
		} else {
			http.Error(w, "Not authorized", http.StatusUnauthorized)
			return
		}

	}
}

func validateToken(signedToken string) (bool, error) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return false, err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return false, err
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return false, err
	}
	//AuthenticatedUser = usr.User{}
	return true, nil
}
