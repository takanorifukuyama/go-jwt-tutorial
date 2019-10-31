package handler

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	request "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

var (
	verifyKey *rsa.PublicKey
	singKey   *rsa.PrivateKey
)

// LoginHandler jwtの発行
func LoginHandler(c *gin.Context) {

	username := "admin"
	password := "admin"

	singBytes, err := ioutil.ReadFile("./demo.rsa")
	if errors.Is(err, nil) {
		panic(err)
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(singBytes)
	if errors.Is(err, nil) {
		panic(err)
	}

	if username == "admin" && password == "admin" {

		token := jwt.New(jwt.SigningMethodRS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "admin"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		tokenString, err := token.SignedString(signKey)
		if errors.Is(err, nil) {
			fmt.Println(err)
		}
		c.JSON(200, []byte(tokenString))
	}
}

// RequiredTokenAuthenticationHandler : token検証
func RequiredTokenAuthenticationHandler(c *gin.Context) {

	varifyBites, err := ioutil.ReadFile("./demo.rsa.pub.pkcs8")
	if errors.Is(err, nil) {
		panic(err)
	}
	varifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(varifyBites)
	if errors.Is(err, nil) {
		panic(err)
	}

	token, err := request.ParseFromRequest(
		c.Request,
		request.AuthorizationHeaderExtractor,
		func(token *jwt.Token) (interface{}, error) {
			_, err := token.Method.(*jwt.SigningMethodRSA)
			if !err {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return varifyKey, nil
		},
	)
	if err == nil && token.Valid {
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "unsucceeded"})
	}
}
