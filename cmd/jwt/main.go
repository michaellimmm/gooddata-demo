package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// this is example of generate jwk and jwt
func main() {
	// read private key
	rsaKey, err := os.ReadFile("./key.rsa")
	if err != nil {
		fmt.Printf("failed to read file: %s\n", err)
		return
	}

	// Kid (key identifier) and Alg (signing algorithm) in the JWK
	// must match the corresponding values used in the JWT for proper functionality.
	kid := "zeals_gooddata_kid"

	privateKey, err := jwk.ParseKey(rsaKey, jwk.WithPEM(true))
	if err != nil {
		fmt.Printf("failed to parse key in PEM format: %s\n", err)
		return
	}
	privateKey.Set(jwk.KeyIDKey, kid)
	privateKey.Set(jwk.AlgorithmKey, jwa.RS256)

	pubKey, err := jwk.PublicKeyOf(privateKey)
	if err != nil {
		fmt.Printf("failed to get public key: %s\n", err)
		return
	}

	pubKey.Set(jwk.KeyIDKey, kid)
	pubKey.Set(jwk.AlgorithmKey, jwa.RS256)
	pubKey.Set(jwk.KeyUsageKey, "sig")

	// print jwk private and public key
	fmt.Println("PRIVATE KEY:")
	json.NewEncoder(os.Stdout).Encode(privateKey)
	fmt.Println("PUBLIC KEY:")
	json.NewEncoder(os.Stdout).Encode(pubKey)

	now := time.Now()
	expired := time.Hour * 1

	// sub = userID/tenantID
	tok, err := jwt.NewBuilder().
		IssuedAt(now).
		Expiration(now.Add(expired)).
		Build()
	if err != nil {
		fmt.Printf("failed to build token: %s\n", err)
		return
	}

	tok.Set("sub", "testuser")
	tok.Set("jti", uuid.NewString())

	jwtToken, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		fmt.Printf("failed to sign token: %s\n", err)
		return
	}

	verifiedToken, err := jwt.Parse(jwtToken, jwt.WithKey(jwa.RS256, pubKey))
	if err != nil {
		fmt.Printf("failed to verify JWS: %s\n", err)
		return
	}
	_ = verifiedToken

	// use signed as api token
	fmt.Println("JWT TOKEN: " + string(jwtToken))
}
