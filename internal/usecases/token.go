package usecases

import (
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type TokenKey struct {
	Kid string
	Sub string
}

func GenerateToken(jsonRSAPrivateKey string, payload TokenKey) (string, error) {
	privateKey, err := jwk.ParseKey([]byte(jsonRSAPrivateKey))
	if err != nil {
		return "", err
	}
	privateKey.Set(jwk.KeyIDKey, payload.Kid)
	privateKey.Set(jwk.AlgorithmKey, jwa.RS256)

	now := time.Now()
	expired := time.Hour * 1
	tok, err := jwt.NewBuilder().
		IssuedAt(now).
		Expiration(now.Add(expired)).
		Build()
	if err != nil {
		return "", err
	}

	tok.Set("sub", payload.Sub)
	tok.Set("jti", uuid.NewString())

	jwtToken, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		return "", err
	}

	return string(jwtToken), nil
}
