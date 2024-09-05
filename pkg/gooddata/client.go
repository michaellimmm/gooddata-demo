package gooddata

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"path"

	"github.com/google/go-querystring/query"
)

type gooddataAPI struct {
	baseUrl    *url.URL
	apiToken   string
	httpClient *http.Client
}

type GooddataAPI interface {
	UserAPI
	DataFilterAPI
	PermissionsAPI
	JWKAPI
	ProfileAPI
	UserGroupAPI
}

type GooddataAPIOption func(*gooddataAPI) error

func NewGooddataAPI(baseUrl string, apiToken string, options ...GooddataAPIOption) (GooddataAPI, error) {
	if baseUrl == "" {
		return nil, errors.New("baseUrl can not be empty")
	}

	if apiToken == "" {
		return nil, errors.New("apiToken can not be empty")
	}

	u, err := url.ParseRequestURI(baseUrl)
	if err != nil {
		return nil, err
	}

	c := &gooddataAPI{
		baseUrl:    u,
		apiToken:   apiToken,
		httpClient: http.DefaultClient,
	}

	for _, option := range options {
		err := option(c)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func WithHttpClient(httpClient *http.Client) GooddataAPIOption {
	return func(ga *gooddataAPI) error {
		ga.httpClient = httpClient
		return nil
	}
}

func (c *gooddataAPI) newRequest(method, endpoint string, header map[string]string, body Serializer) (*http.Request, error) {
	var payload io.Reader
	if body != nil {
		buf, err := body.Marshal()
		if err != nil {
			return nil, err
		}

		payload = bytes.NewBuffer(buf)
	}

	req, err := http.NewRequest(method, endpoint, payload)
	if err != nil {
		return nil, err
	}

	if header != nil && header["Content-Type"] != "" {
		req.Header.Set("Content-Type", header["Content-Type"])
	} else {
		req.Header.Set("Content-Type", "application/vnd.gooddata.api+json")
	}

	req.Header.Set("Authorization", "Bearer "+c.apiToken)

	return req, nil
}

func (c *gooddataAPI) do(req *http.Request, resp Serializer) error {
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	var bodyBytes []byte
	if res.Body != nil {
		bodyBytes, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}
	}

	res.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if res.StatusCode/100 != 2 {
		return fmt.Errorf("unexpected status code: %d, %s", res.StatusCode, string(bodyBytes))
	}

	if resp != nil && bodyBytes != nil {
		err = resp.Unmarshal(bodyBytes)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *gooddataAPI) url(endpointPath string, options interface{}) (string, error) {
	newBasePath := *c.baseUrl
	newBasePath.Path = path.Join(newBasePath.Path, endpointPath)

	if options != nil {
		optionsQuery, err := query.Values(options)
		if err != nil {
			return "", err
		}

		for k, values := range newBasePath.Query() {
			for _, v := range values {
				optionsQuery.Add(k, v)
			}
		}
		newBasePath.RawQuery = optionsQuery.Encode()
	}

	return newBasePath.String(), nil
}
