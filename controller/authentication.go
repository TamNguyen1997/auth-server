package controller

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang-jwt/jwt"
	"main.go/model"
)

type Authenticator struct {
	privateKey []byte
}

type AuthenticationRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewAuthenticator() *Authenticator {
	key, _ := os.ReadFile("resources/keys/jwtRS256.key")
	return &Authenticator{
		privateKey: key,
	}
}

func (auth *Authenticator) login(ctx *gin.Context) {
	var body AuthenticationRequest
	if err := ctx.ShouldBindBodyWith(&body, binding.JSON); err != nil {
		ctx.IndentedJSON(http.StatusUnauthorized, err)
		return
	}

	token := auth.generateToken()
	ctx.IndentedJSON(http.StatusOK, token)
}

func (auth *Authenticator) generateToken() *string {
	pKey, _ := jwt.ParseRSAPrivateKeyFromPEM(auth.privateKey)
	claims := &model.BasicJwt{
		Authorities: []string{},
		Exp:         time.Now().Add(time.Hour).Unix(),
	}
	jwt := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	token, err := jwt.SignedString(pKey)
	if err != nil {
		fmt.Println("error while generating token", err)
		return nil
	}
	return &token
}

func (auth *Authenticator) AddRoute(server *gin.Engine) {
	route := server.Group("/login")
	{
		route.POST("", auth.login)
	}
}
