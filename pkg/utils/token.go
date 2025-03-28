package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"www.marawa.com/microservice_service/internal/infra/config"
)

func VerifyToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.ConfigInstance.AuthConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token is invalid")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("Could not parse claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("Could not parse exp")
	}

	if int64(exp) < time.Now().Unix() {
		return nil, errors.New("Token is expired")
	}

	return claims, nil
}
