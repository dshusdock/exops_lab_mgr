package jwtauthsvc

import (
	"fmt"
	"time"

	"github.com/go-chi/jwtauth/v5"
	// "github.com/lestrrat-go/jwx/jwt"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// var secretKey = []byte("secret-key")

var tokenAuth *jwtauth.JWTAuth


func init() {
	fmt.Println("Init Token service")
	tokenAuth = jwtauth.New("HS256", []byte("somethingotherthansecrett"), nil, jwt.WithAcceptableSkew(30*time.Second))
	
	// For debugging/example purposes, we generate and print
	// a sample jwt token with claims `user_id:123` here:
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123, "exp": time.Now().Add(60*time.Second).Unix()})
	// fmt.Println("DEBUG: a sample jwt is ", time.Now().Add(60*time.Second).Unix() )
	// fmt.Printf("DEBUG: a sample jwt is %s\n\n", tokenString)

}

func GetToken() *jwtauth.JWTAuth {
	return tokenAuth
}

func CreateToken(username string) (string, error) {
	// _, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": username, "exp": time.Now().Add(300*time.Second).Unix()})
	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": username})	
	return tokenString, nil
}


