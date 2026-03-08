package helper

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenPair struct {
	AccessToken   string
	RefreshToken  string
	AccessExpire  int64 // Timestamp en segundos cuando caduca el access token
	RefreshExpire int64 // Timestamp en segundos cuando caduca el refresh token
}

// GenerateJwt - mantener para compatibilidad
func GenerateJwt(user_id, email, role string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user_id
	claims["email"] = email
	claims["role"] = role
	claims["exp"] = time.Now().Add(time.Hour).Unix() // Access Token: 1 hora

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// GenerateAccessToken - generar solo el access token y retornar token + expiración
func GenerateAccessToken(userID, email, role string) (string, int64, error) {
	accessExpireTime := time.Now().Add(time.Hour)
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["sub"] = userID
	accessClaims["email"] = email
	accessClaims["role"] = role
	accessClaims["type"] = "access"
	accessClaims["exp"] = accessExpireTime.Unix()

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", 0, err
	}

	return accessTokenString, accessExpireTime.Unix(), nil
}

// GenerateTokenPair - generar ambos tokens (Access + Refresh)
func GenerateTokenPair(userID, email, role string) (*TokenPair, error) {
	// Access Token: 1 hora
	accessExpireTime := time.Now().Add(time.Hour)
	accessToken := jwt.New(jwt.SigningMethodHS256)
	accessClaims := accessToken.Claims.(jwt.MapClaims)
	accessClaims["sub"] = userID
	accessClaims["email"] = email
	accessClaims["role"] = role
	accessClaims["type"] = "access"
	accessClaims["exp"] = accessExpireTime.Unix()

	accessTokenString, err := accessToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	// Refresh Token: 7 días
	refreshExpireTime := time.Now().Add(7 * 24 * time.Hour)
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	refreshClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshClaims["sub"] = userID
	refreshClaims["email"] = email
	refreshClaims["type"] = "refresh"
	refreshClaims["exp"] = refreshExpireTime.Unix()

	refreshTokenString, err := refreshToken.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &TokenPair{
		AccessToken:   accessTokenString,
		RefreshToken:  refreshTokenString,
		AccessExpire:  accessExpireTime.Unix(),
		RefreshExpire: refreshExpireTime.Unix(),
	}, nil
}

// ValidateToken - validar tanto access como refresh tokens
func ValidateToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token claims")
	}

	return claims, nil
}
