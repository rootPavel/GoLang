package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("TOPSECRETKEY")
var users = map[string]string{
	"Dima":  "1234",
	"Katya": "Ekaterina2005",
	"Vasya": "Good12345",
}

type LoginRequest struct {
	Username string
	Password string
}

func publicRoute(c *gin.Context) {
	c.JSON(200, gin.H{
		"page": "This is Public Page",
	})
}

func loginRoute(c *gin.Context) {
	var credo LoginRequest
	if err := c.ShouldBindJSON(&credo); err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	correctPass, exists := users[credo.Username]
	if !exists || correctPass != credo.Password {
		c.JSON(401, gin.H{
			"error": "bad data",
		})
		return
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Username": credo.Username,
		"Since":    time.Now().Unix(),
		"Until":    time.Now().Add(2 * time.Hour).Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"token": tokenString,
	})
}

func secretRoute(c *gin.Context) {
	var authHeader string
	authHeader = c.GetHeader("Authorization")
	if authHeader == "" {
		c.JSON(403, gin.H{
			"info": "U r not authorized!",
		})
		return
	}
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Sign error")
		}
		return secretKey, nil
	})
	if err != nil {
		c.JSON(401, gin.H{
			"error": "Bad token",
			"info":  err.Error(),
		})
		return
	}
	if climes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.JSON(200, gin.H{
			"secret": "This is secret page",
			"user":   climes["Username"],
		})
	} else {
		c.JSON(401, gin.H{
			"error": err.Error(),
		})
	}

}

func main() {
	r := gin.Default()
	r.GET("/public", publicRoute)
	r.GET("/secret", secretRoute)
	r.POST("/login", loginRoute)
	r.Run(":8080")
}
