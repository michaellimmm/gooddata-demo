package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"os"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwk"
)

const (
	filename = "test-key"
	kid      = "zeals_gooddata_kid_test"
)

func main() {
	// Generate RSA key.
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}

	// Extract public component.
	pub := key.Public()

	// Encode private key to PKCS#1 ASN.1 PEM.
	keyPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(key),
		},
	)

	// Encode public key to PKCS#1 ASN.1 PEM.
	pubPEM := pem.EncodeToMemory(
		&pem.Block{
			Type:  "RSA PUBLIC KEY",
			Bytes: x509.MarshalPKCS1PublicKey(pub.(*rsa.PublicKey)),
		},
	)

	// Write private key to file.
	if err := os.WriteFile(filename+".rsa", keyPEM, 0700); err != nil {
		panic(err)
	}

	// Write public key to file.
	if err := os.WriteFile(filename+".rsa.pub", pubPEM, 0755); err != nil {
		panic(err)
	}

	privateKey, err := jwk.FromRaw(key)
	if err != nil {
		panic(err)
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
	fmt.Println()
	fmt.Println("PUBLIC KEY:")
	json.NewEncoder(os.Stdout).Encode(pubKey)
}
