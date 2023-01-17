package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(email string, isUser bool) string
	ValidateToken(token string) (*jwt.Token, error)
}

type authCustomClims struct {
	Name string `json:"Name"`
	User bool   `json:"user"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issure    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issure:    "Lahiru",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRETE")
	if secret == "" {
		secret = "secrete"
	}
	return secret
}

func (Service *jwtServices) GenerateToken(email string, isUser bool) string {
	claims := &authCustomClims{
		email, isUser, jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			Issuer:    Service.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	new_token, err := token.SignedString([]byte(Service.secretKey))

	if err != nil {
		panic(err)
	}

	return new_token
}

func (Services *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid Token", token.Header["alg"])
		}
		return []byte(Services.secretKey), nil
	})
}
