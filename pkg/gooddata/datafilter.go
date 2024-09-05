package gooddata

import (
	"fmt"

	"github.com/DataDog/jsonapi"
)

const (
	userDataFiltersEndpoint = "/api/v1/entities/workspaces/%s/userDataFilters"
	userDataFilterEndpoint  = "/api/v1/entities/workspaces/%s/userDataFilters/%s"
)

type UserDataFilter struct {
	ID                string     `jsonapi:"primary,userDataFilter"`
	AreRelationsValid bool       `jsonapi:"attribute" json:"areRelationsValid,omitempty"`
	Description       string     `jsonapi:"attribute" json:"description,omitempty"`
	Maql              string     `jsonapi:"attribute" json:"maql"`
	Tags              []string   `jsonapi:"attribute" json:"tags,omitempty"`
	Title             string     `jsonapi:"attribute" json:"title,omitempty"`
	User              *User      `jsonapi:"relationship" json:"user,omitempty"`
	UserGroup         *UserGroup `jsonapi:"relationship" json:"userGroup,omitempty"`
}

func (udf *UserDataFilter) Marshal() ([]byte, error) {
	return jsonapi.Marshal(udf)
}

func (udf *UserDataFilter) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, udf)
}

type UserDataFilters []UserDataFilter

func (udf *UserDataFilters) Marshal() ([]byte, error) {
	return jsonapi.Marshal(udf)
}

func (udf *UserDataFilters) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, udf)
}

type DataFilterAPI interface {
	ListUserDataFilters(workspaceID string) ([]UserDataFilter, error)
	CreateUserDataFilter(workspaceID string, udf UserDataFilter) (UserDataFilter, error)
	DeleteUserDataFilter(workspaceID string, udfID string) error
	GetUserDataFilter(workspaceID string, udfID string) (UserDataFilter, error)
	UpdateUserDataFilter(workspaceID string, udf UserDataFilter) (UserDataFilter, error)
}

func (c *gooddataAPI) ListUserDataFilters(workspaceID string) ([]UserDataFilter, error) {
	var result UserDataFilters

	url, err := c.url(fmt.Sprintf(userDataFiltersEndpoint, workspaceID), nil)
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

func (c *gooddataAPI) CreateUserDataFilter(workspaceID string, udf UserDataFilter) (UserDataFilter, error) {
	result := UserDataFilter{}

	url, err := c.url(fmt.Sprintf(userDataFiltersEndpoint, workspaceID), nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("POST", url, header, &udf)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) DeleteUserDataFilter(workspaceID string, udfID string) error {
	url, err := c.url(fmt.Sprintf(userDataFilterEndpoint, workspaceID, udfID), nil)
	if err != nil {
		return err
	}

	header := make(map[string]string)
	req, err := c.newRequest("DELETE", url, header, nil)
	if err != nil {
		return err
	}

	return c.do(req, nil)
}

func (c *gooddataAPI) GetUserDataFilter(workspaceID string, udfID string) (UserDataFilter, error) {
	result := UserDataFilter{}

	url, err := c.url(fmt.Sprintf(userDataFilterEndpoint, workspaceID, udfID), nil)
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

func (c *gooddataAPI) UpdateUserDataFilter(workspaceID string, udf UserDataFilter) (UserDataFilter, error) {
	result := UserDataFilter{}

	url, err := c.url(fmt.Sprintf(userDataFilterEndpoint, workspaceID, udf.ID), nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("PUT", url, header, &udf)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
