package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTOption describes the jwt extra data
type JWTOption struct {
	Key string
	Val any
}

// WithJWTOption returns the option from k/v
func WithJWTOption(key string, val any) JWTOption {
	return JWTOption{
		Key: key,
		Val: val,
	}
}

// GenJWT 生成JWT
func GenJWT(secretKey string, exp int64, opt ...JWTOption) (string, error) {
	iat := time.Now().Unix() // jwt 以秒为过期时间

	claims := make(jwt.MapClaims)
	claims["exp"] = iat + exp
	claims["iat"] = iat

	for _, v := range opt {
		claims[v.Key] = v.Val
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secretKey))
}
