package auth

import (
	"encoding/base64"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtKey, _ = base64.URLEncoding.DecodeString("demo21")
)

// get token from Header
func ExtractToken(r *http.Request) string {
	tokenHeader := r.Header.Get("Authorization")
	if tokenHeader == "" {
		return ""
	}
	splitted := strings.Split(tokenHeader, " ")

	if len(splitted) != 2 {
		return ""
	}
	tokenpath := splitted[1]

	return tokenpath
}

func IsAuthorized(tokenpath string) (map[string]interface{}, error) {

	token, err := jwt.Parse(tokenpath, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Can't authorized token")
		}
		return jwtKey, nil
	})

	if err != nil {
		return nil, errors.New("Can't authorized token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Can't authorized token")
}
