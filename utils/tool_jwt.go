package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	viper "github.com/spf13/viper"
)

// MyClaims
type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

// SecretKey
var SecretKey = []byte(viper.GetString("jwt.serect"))
var jwtExpireDuration = time.Hour * 2

// get jwt expire duration
func GetJwtExpireDuration() time.Duration {
	return jwtExpireDuration
}

// GenerateToken
func GenerateToken(username string) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Username:         username,
	}
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(jwtExpireDuration))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

// ParseToken
func ParseToken(tokenString string) (*MyClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*MyClaims)
	if !ok {
		return nil, err
	}
	return claims, nil
}

// VerifyToken vaild 验证合法性 是否被篡改
func VerifyToken(tokenString string) (bool, error) {
	_, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

// RefreshToken
func RefreshToken(tokenString string) (string, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return "", err
	}
	return GenerateToken(claims.Username)
}
