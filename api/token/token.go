package token

import (
	"auth_service/config"
	pb "auth_service/genproto"
	"errors"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	_ "github.com/form3tech-oss/jwt-go"
	"time"
)

var secret_key = config.Load().TokenKey

type Token struct {
	AccessToken  string
	RefreshToken string
}

func GENERATEJWTToken(user *pb.LoginResponse) *Token {
	AccessToken := jwt.New(jwt.SigningMethodHS256)

	claims := AccessToken.Claims.(jwt.MapClaims)
	claims["user_name"] = user.UserName
	claims["password"] = user.Password
	claims["email"] = user.Email
	claims["iat"] = time.Now().Unix()
	claims["exp"] = time.Now().Add(3 * time.Hour).Unix()

	access, err := AccessToken.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("error  access singed my_secret key")
	}
	RefreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaim := RefreshToken.Claims.(jwt.MapClaims)
	refreshClaim["user_name"] = user.UserName
	refreshClaim["password"] = user.Password
	refreshClaim["email"] = user.Email
	refreshClaim["iat"] = time.Now().Unix()
	refreshClaim["exp"] = time.Now().Add(24 * time.Hour)
	refresh, err := RefreshToken.SignedString([]byte(secret_key))
	if err != nil {
		fmt.Println("error  refresh singed my_secret key")
	}
	return &Token{
		AccessToken:  access,
		RefreshToken: refresh,
	}

}
func ExtractClaim(tokenStr string) (jwt.MapClaims, error) {
	var (
		token *jwt.Token
		err   error
	)

	keyFunc := func(token *jwt.Token) (interface{}, error) {
		return []byte(secret_key), nil
	}
	token, err = jwt.Parse(tokenStr, keyFunc)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !(ok && token.Valid) {
		return nil, errors.New("invalid token")
	}

	return claims, nil
}
