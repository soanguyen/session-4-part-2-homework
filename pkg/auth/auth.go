package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var mySigningKey = []byte("ct-secret-key")

func GenerateToken(username string, expireDuration time.Duration) (string, error) {
	// Create the Claims
	claims := &jwt.RegisteredClaims{
		Issuer:    "ct-backend-course",
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expireDuration)),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		return "", err
	}

	return ss, nil
}

func VerifyToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.NewValidationError("unexpected signing method", jwt.ValidationErrorSignatureInvalid)
		}

		return mySigningKey, nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username, ok := claims["sub"].(string)
		if !ok {
			return "", jwt.NewValidationError("invalid subject claim", jwt.ValidationErrorClaimsInvalid)
		}
		return username, nil
	}

	return "", jwt.NewValidationError("invalid token", jwt.ValidationErrorClaimsInvalid)
}
