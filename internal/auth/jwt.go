package auth

import (
    "time"

    "github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("your-secret-key") // 🔐 Replace in production

func GenerateJWT(userID string) (string, error) {
    claims := jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString(jwtKey)
}

func ValidateJWT(tokenString string) (string, error) {
    token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        return jwtKey, nil
    })
    if err != nil || !token.Valid {
        return "", err
    }
    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        return claims["user_id"].(string), nil
    }
    return "", jwt.ErrSignatureInvalid
}