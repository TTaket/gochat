package utils

import "github.com/golang-jwt/jwt/v5"

// MyClaims
type MyClaims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

// SecretKey
var SecretKey = []byte("my_sercret_key")

// GenerateToken
func GenerateToken(username string) (string, error) {
	claims := MyClaims{
		RegisteredClaims: jwt.RegisteredClaims{},
		Username:         username,
	}
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

// VerifyToken
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
