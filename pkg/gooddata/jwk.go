package gooddata

import (
	"encoding/json"
	"fmt"

	"github.com/DataDog/jsonapi"
)

const (
	jwkEndpoint  = "/api/v1/entities/jwks/%s"
	jwksEndpoint = "/api/v1/entities/jwks"
)

type JWKAPI interface {
	ListJWKs() (JWKs, error)
	CreateJWK(jwk JWK) (JWK, error)
	GetJWK(id string) (JWK, error)
	UpdateJWK(jwk JWK) (JWK, error)
	DeleteJWK(id string) error
}

type RsaSpecification struct {
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	Use string `json:"use"`
	N   string `json:"n"`
	E   string `json:"e"`
	Kid string `json:"kid"`
}

func (rsa *RsaSpecification) Marshal() ([]byte, error) {
	return json.Marshal(rsa)
}

func (rsa *RsaSpecification) Unmarshal(data []byte) error {
	return json.Unmarshal(data, rsa)
}

type JWK struct {
	ID      string           `jsonapi:"primary,jwk"`
	Content RsaSpecification `jsonapi:"attribute" json:"content"`
}

func (j *JWK) Marshal() ([]byte, error) {
	return jsonapi.Marshal(j)
}

func (j *JWK) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, j)
}

type JWKs []JWK

func (j *JWKs) Marshal() ([]byte, error) {
	return jsonapi.Marshal(j)
}

func (j *JWKs) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, j)
}

func (c *gooddataAPI) ListJWKs() (JWKs, error) {
	var result JWKs

	url, err := c.url(jwksEndpoint, nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("GET", url, header, nil)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) CreateJWK(jwk JWK) (JWK, error) {
	result := JWK{}

	url, err := c.url(jwksEndpoint, nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("POST", url, header, &jwk)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) GetJWK(id string) (JWK, error) {
	var result JWK

	url, err := c.url(fmt.Sprintf(jwkEndpoint, id), nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("GET", url, header, nil)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) UpdateJWK(jwk JWK) (JWK, error) {
	result := JWK{}

	url, err := c.url(fmt.Sprintf(jwkEndpoint, jwk.ID), nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("PUT", url, header, &jwk)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) DeleteJWK(id string) error {
	url, err := c.url(fmt.Sprintf(jwkEndpoint, id), nil)
	if err != nil {
		return err
	}

	header := make(map[string]string)
	req, err := c.newRequest("DELETE", url, header, nil)
	if err != nil {
		return err
	}

	err = c.do(req, nil)
	if err != nil {
		return err
	}

	return nil
}
