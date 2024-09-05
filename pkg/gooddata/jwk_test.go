package gooddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshalJWK(t *testing.T) {
	jsonString := `
	{
		"data": {
			"id": "id1",
			"type": "jwk",
			"attributes": {
				"content": {
					"kty": "RSA",
					"alg": "RS256",
					"use": "sig",
					"n": "n",
					"e": "AQAB",
					"kid": "kid"
				}
			}
		},
		"links": {
			"self": "https://gooddata.com/api/v1/entities/jwks/id1"
		}
	}
	`

	actual := JWK{}
	err := actual.Unmarshal([]byte(jsonString))
	assert.Nil(t, err)

	expected := JWK{
		ID: "id1",
		Content: RsaSpecification{
			Kty: "RSA",
			Alg: "RS256",
			Use: "sig",
			N:   "n",
			E:   "AQAB",
			Kid: "kid",
		},
	}

	assert.Equal(t, expected, actual)
}

func TestUnMarshalJWK(t *testing.T) {
	data := JWK{
		ID: "id1",
		Content: RsaSpecification{
			Kty: "RSA",
			Alg: "RS256",
			Use: "sig",
			N:   "n",
			E:   "AQAB",
			Kid: "kid",
		},
	}

	actual, err := data.Marshal()
	assert.Nil(t, err)

	expected := `{"data":{"id":"id1","type":"jwk","attributes":{"content":{"kty":"RSA","alg":"RS256","use":"sig","n":"n","e":"AQAB","kid":"kid"}}}}`
	assert.Equal(t, expected, string(actual))
}

func TestListJWKs(t *testing.T) {
	api, err := NewGooddataAPI("https://helpful-duck.trial.cloud.gooddata.com/", "NTAwMGQ2ZTItNmRhYy00YTZhLWIxMzYtNmJhODM2OWQxZDY0Om5ld25ldzp6ZHAvTkVCdUp5NW9LbndUK05hb1drYkRuYjZtTlNTZQ==")
	assert.Nil(t, err)

	res, err := api.ListJWKs()
	assert.Nil(t, err)

	t.Logf("%+v", res)
}
