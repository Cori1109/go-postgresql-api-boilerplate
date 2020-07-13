package jwt

import (
	"fmt"
	"time"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/config"
	"github.com/dgrijalva/jwt-go"
)

func SignToken(id int) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Unix() + 90*24*60*60,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(config.Dot("JWT_SECRET")))

	fmt.Println(tokenString, "🌹🌹")

	return tokenString, err
}

func CheckToken(tokenString string) bool {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.Dot("JWT_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims["id"], claims["iat"], claims["exp"])
		return true
	}

	fmt.Println(err)
	return false
}
