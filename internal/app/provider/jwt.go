package provider

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Token response struct
type Token struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	IssuedAt    int64  `json:"issued_at"`
}

// Claims for JWT payload
type Claims struct {
	UserID int
	jwt.StandardClaims
}

// NewClaims factory
func NewClaims() *Claims {
	return &Claims{}
}

// GenerateToken ...
func GenerateToken(id int) (Token, error) {
	exp, err := strconv.Atoi(os.Getenv("TOKEN_EXPIRATION_DURATION"))
	if err != nil {
		return Token{}, err
	}

	claims := &Claims{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			Issuer:    os.Getenv("APPLICATION_NAME"),
			ExpiresAt: time.Now().Add(time.Duration(exp) * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod(os.Getenv("SIGNING_METHOD")), claims)
	signedToken, _ := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	tk := Token{
		AccessToken: signedToken,
		TokenType:   "Bearer",
		IssuedAt:    time.Now().Unix(),
	}

	return tk, nil
}

// ParseAndValidateToken check token from header
func ParseAndValidateToken(tokenHeader string, c *Claims) error {
	if tokenHeader == "" {
		return errors.New("Missing token header")
	}

	splitted := strings.Split(tokenHeader, " ")
	if len(splitted) != 2 {
		return errors.New("Invalid token header")
	}

	tokenPart := splitted[1]
	token, err := jwt.ParseWithClaims(tokenPart, c, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return errors.New("Failed to parse token")
	}

	if !token.Valid {
		return errors.New("Invalid token")
	}

	return nil
}
