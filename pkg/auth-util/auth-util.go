package authutil

import (
	"context"
	"encoding/base64"
	"errors"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var utcNow = func() time.Time {
	return time.Now().UTC()
}

func GenerateAuthenticationToken(ctx context.Context, userID string, ttl time.Duration) (string, error) {
	select {
	case <-ctx.Done():
		return "", ctx.Err()
	default:
	}

	now := utcNow()

	claims := jwt.RegisteredClaims{
		Issuer:    os.Getenv("ISSUER"),
		Audience:  jwt.ClaimStrings{os.Getenv("AUDIENCE")},
		Subject:   userID,
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
		ExpiresAt: jwt.NewNumericDate(now.Add(ttl)),
	}

	secret, err := jwtSecret()
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}
	return signed, nil
}

func ParseAndValidateJWT(headerToken string) (*jwt.RegisteredClaims, error) {
	token := strings.TrimSpace(headerToken)
	if strings.HasPrefix(strings.ToLower(token), "bearer ") {
		token = strings.TrimSpace(token[7:])
	}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithIssuedAt(),
		jwt.WithAudience(os.Getenv("AUDIENCE")),
		jwt.WithIssuer(os.Getenv("ISSUER")),
		jwt.WithLeeway(30*time.Second),
	)

	claims := &jwt.RegisteredClaims{}
	_, err := parser.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return jwtSecret()
	})
	if err != nil {
		return nil, err
	}
	return claims, nil
}

func jwtSecret() ([]byte, error) {
	s := strings.TrimSpace(os.Getenv("JWT_SECRET_KEY"))
	if s == "" {
		return nil, errors.New("missing JWT_SECRET_KEY")
	}

	if b, err := base64.StdEncoding.DecodeString(s); err == nil && len(b) > 0 {
		return b, nil
	}

	return []byte(s), nil
}
