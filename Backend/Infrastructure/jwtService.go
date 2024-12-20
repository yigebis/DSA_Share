package Infrastructure

import (
	"DSAShare/Error"
	"DSAShare/UseCase"

	"github.com/dgrijalva/jwt-go"
)

type TokenService struct{
	JwtSecret string
}

func NewTokenService(jwtSecret string) UseCase.ITokenService {
	return &TokenService{JwtSecret: jwtSecret}
}

// GenerateToken implements UseCase.ITokenService.
func (ts *TokenService) GenerateToken(email string, userName string, expiryDuration int64) (string, error) {
	claims := jwt.MapClaims{
		"email" : email,
		"user_name" : userName,
		"expiryDuration" : expiryDuration,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString([]byte(ts.JwtSecret))
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}


// ValidateToken implements UseCase.ITokenService.
func (ts *TokenService) ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(ts.JwtSecret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.MapClaims)
	if ok && token.Valid {
		return *claims, nil
	}
	
	return nil, Error.ErrInvalidToken
}
