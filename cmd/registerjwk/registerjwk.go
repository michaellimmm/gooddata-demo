package main

import (
	"encoding/json"
	"fmt"
	"github/michaellimmm/gooddata-demo/pkg/gooddata"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

// register public jwk to gooddata
func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("failed to load env var, err: %+v", err)
		return
	}

	// Notes: Kid/jwk.KeyIDKey (key identifier) and Alg/jwk.AlgorithmKey (signing algorithm) in the JWK
	// must match the corresponding values used in the JWT for proper functionality.
	baseUrl := os.Getenv("GOODDATA_BASEURL")
	accessToken := os.Getenv("GOODDATA_ACCESSTOKEN")
	kid := "zeals_gooddata_kid"
	jwkID := "zeals_jwk"

	// read private key
	rsaKey, err := os.ReadFile("./key.rsa")
	if err != nil {
		fmt.Printf("failed to read file, err: %+v\n", err)
		return
	}

	privateKey, err := jwk.ParseKey([]byte(rsaKey), jwk.WithPEM(true))
	if err != nil {
		fmt.Printf("failed to parse key in PEM format, err: %+v\n", err)
		return
	}
	privateKey.Set(jwk.KeyIDKey, kid)
	privateKey.Set(jwk.AlgorithmKey, jwa.RS256)

	pubKey, err := jwk.PublicKeyOf(privateKey)
	if err != nil {
		fmt.Printf("failed to get public key, err: %+v\n", err)
		return
	}

	pubKey.Set(jwk.KeyIDKey, kid)           // required by gooddata jwk RsaSpecification
	pubKey.Set(jwk.AlgorithmKey, jwa.RS256) // required by gooddata jwk RsaSpecification
	pubKey.Set(jwk.KeyUsageKey, "sig")      // required by gooddata jwk RsaSpecification

	// print jwk private and public key
	json.NewEncoder(os.Stdout).Encode(privateKey)
	json.NewEncoder(os.Stdout).Encode(pubKey)

	pubKeyBytes, err := json.Marshal(pubKey)
	if err != nil {
		fmt.Printf("failed to get marshal public key, err: %+v\n", err)
		return
	}

	content := gooddata.RsaSpecification{}
	err = content.Unmarshal(pubKeyBytes)
	if err != nil {
		fmt.Printf("failed to get unmarshal public key, err: %+v\n", err)
		return
	}

	data := gooddata.JWK{
		ID:      jwkID,
		Content: content,
	}

	api, err := gooddata.NewGooddataAPI(baseUrl, accessToken)
	if err != nil {
		fmt.Printf("failed to initiate gooddata api, err: %+v\n", err)
		return
	}

	res, err := api.CreateJWK(data)
	if err != nil {
		fmt.Printf("failed to create jwk, err: %+v\n", err)
		return
	}

	fmt.Printf("result: %+v\n", res)

	// verify jwk token
	user := gooddata.User{
		ID:        "testuser",
		Firstname: "test",
		LastName:  "user",
	}
	newUser, err := api.CreateUser(user)
	if err != nil {
		fmt.Printf("failed to create user, err: %+v", err)
		return
	}

	now := time.Now()
	expired := time.Hour * 1
	tok, err := jwt.NewBuilder().
		IssuedAt(now).
		Expiration(now.Add(expired)).
		Build()
	if err != nil {
		fmt.Printf("failed to build token, err: %+v\n", err)
		return
	}

	tok.Set(jwk.KeyIDKey, kid)
	tok.Set(jwk.AlgorithmKey, jwa.RS256)
	tok.Set("sub", newUser.ID)
	tok.Set("jti", "jwk_test_id")

	jwtToken, err := jwt.Sign(tok, jwt.WithKey(jwa.RS256, privateKey))
	if err != nil {
		fmt.Printf("failed to sign token, err: +%v\n", err)
		return
	}

	_, err = jwt.Parse(jwtToken, jwt.WithKey(jwa.RS256, pubKey))
	if err != nil {
		fmt.Printf("failed to verify JWS, err: %+v\n", err)
		return
	}

	api2, err := gooddata.NewGooddataAPI(baseUrl, string(jwtToken))
	if err != nil {
		fmt.Printf("failed to initialize gooddata api, err: %+v\n", err)
		return
	}

	profile, err := api2.GetProfile()
	if err != nil {
		fmt.Printf("failed get user: %s\n", err)
		return
	}

	result, err := profile.Marshal()
	if err != nil {
		fmt.Printf("failed result user: %s\n", err)
		return
	}

	fmt.Println(string(result))
}
