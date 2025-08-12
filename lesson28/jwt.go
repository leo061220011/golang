package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type User struct {
	UserId   int
	UserName string
}
type Claims struct {
	jwt.RegisteredClaims
	UserId   int
	UserName string
}

const TOKEN_EXPIRES_AT = time.Hour * 2
const SECRET_KEY = "supersecretkey"

func BuildJWTString() (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXPIRES_AT)),
		},
		UserId:   1,
		UserName: "Oleg",
	})
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
func GetData(tokenString string) User {
	claims := &Claims{}
	jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	user := User{UserId: claims.UserId, UserName: claims.UserName}
	return user
}
func main() {
	tokenString, err := BuildJWTString()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(tokenString)
	fmt.Println(GetData(tokenString))
}
