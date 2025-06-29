package utils

import (
	"context"
	"errors"
	"rbac-service/domain"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var userRepo domain.UserRepository

func NewUserRepo(repo domain.UserRepository) {
	userRepo = repo
}

// JWT 密鑰
var jwtKey = []byte("jwt_for_rcba_login")

// GenerateJWTToken 生成 JWT token
func GenerateJWTToken(username string, roles []string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     roles,
		"exp":      jwt.NewNumericDate(time.Now().Add(time.Hour * 2)),
	}

	// jwt 加密方式
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", errors.New("token generation failed")
	}

	return tokenString, nil
}

// CompareJWTToken 比對 jwt 是否相同
func CompareJWTToken(
	ctx context.Context,
	username string,
	tokenNeedToCheck string,
) (bool, error) {
	user, err := userRepo.GetByUsername(ctx, username)
	if err != nil {
		return false, errors.New("user not found")
	}

	return tokenNeedToCheck == user.Jwt, nil
}

// ParseJWTToken 解析 JWT token
func ParseJWTToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// IsTokenExpired 檢查 token 是否過期
func IsTokenExpired(tokenString string) bool {
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 直接檢查過期時間
		return time.Now().After(time.Unix(int64(claims["exp"].(float64)), 0))
	}
	return true
}
