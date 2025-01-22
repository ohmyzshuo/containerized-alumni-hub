package content

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"os"
	"time"
)

var jwtSecret = os.Getenv("JWT_SECRET")
var secretKey = []byte(jwtSecret)

type TokenClaims struct {
	AlumniID  uint `json:"alumni_id"`
	ContentID uint `json:"content_id"`
	Status    uint `json:"status"`
	jwt.StandardClaims
}

func GenerateSignedToken(alumniID uint, contentID uint, status uint) (string, error) {
	claims := TokenClaims{
		AlumniID:  alumniID,
		ContentID: contentID,
		Status:    status,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * 14 * time.Hour).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedStr, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	log.Println(signedStr)
	return signedStr, nil
}

func ParseAndValidateToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, errors.New("invalid token: " + err.Error())
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token claims")
	}

	//	if claims.ExpiresAt < time.Now().Unix() {
	//		return nil, errors.New("token has expired")
	//	}

	return claims, nil
}
