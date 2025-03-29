package util

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var (
	AccessTokenExpire  = 15 * time.Minute   // Access Token 有效期
	RefreshTokenExpire = 7 * 24 * time.Hour // Refresh Token 有效期
	SecretKey          = []byte("jkjhjasdjh2bg1hjnbsxfjkhjk")
)

func GenTokens(userId int64) (accessToken, refreshToken string, err error) {
	// Access Token
	accessClaims := jwt.MapClaims{
		KeyUserId: userId,
		KeyExp:    time.Now().Add(AccessTokenExpire).Unix(),
		KeyType:   KeyAccess,
	}
	accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(SecretKey)
	if err != nil {
		return "", "", err
	}

	// Refresh Token
	refreshClaims := jwt.MapClaims{
		KeyUserId: userId,
		KeyExp:    time.Now().Add(AccessTokenExpire).Unix(),
		KeyType:   KeyRefresh,
	}
	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(SecretKey)

	return accessToken, refreshToken, err
}

func ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return token.Claims.(jwt.MapClaims), nil
}

func GenRefreshKey(userId int64) string {
	return TokenPrefix + fmt.Sprintf(":%d", userId)
}

func GenBlackListKey(accessToken string) string {
	return BlackListPrefix + ":" + accessToken
}
