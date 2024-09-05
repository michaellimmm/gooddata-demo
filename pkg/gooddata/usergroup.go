package gooddata

import (
	"fmt"

	"github.com/DataDog/jsonapi"
)

const (
	userGroupsEndpoint = "/api/v1/entities/userGroups"
	userGroupEndpoint  = "/api/v1/entities/userGroups/%s"
)

type UserGroup struct {
	ID   string `jsonapi:"primary,userGroup"`
	Name string `jsonapi:"attribute" json:"name,omitempty"`
}

func (u *UserGroup) Marshal() ([]byte, error) {
	return jsonapi.Marshal(u)
}

func (u *UserGroup) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, u)
}

type UserGroups []UserGroup

func (u *UserGroups) Marshal() ([]byte, error) {
	return jsonapi.Marshal(u)
}

func (u *UserGroups) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, u)
}

type UserGroupAPI interface {
	ListUserGroups() (UserGroups, error)
	CreateUserGroup(userGroup UserGroup) (UserGroup, error)
	GetUserGroup(id string) (UserGroup, error)
	DeleteUserGroup(id string) error
	UpdateUserGroup(userGroup UserGroup) (UserGroup, error)
}

func (c *gooddataAPI) ListUserGroups() (UserGroups, error) {
	var result UserGroups

	url, err := c.url(userGroupsEndpoint, nil)
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

func (c *gooddataAPI) CreateUserGroup(userGroup UserGroup) (UserGroup, error) {
	result := UserGroup{}

	url, err := c.url(userGroupsEndpoint, nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("POST", url, header, &userGroup)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) GetUserGroup(id string) (UserGroup, error) {
	result := UserGroup{}

	url, err := c.url(fmt.Sprintf(userGroupEndpoint, id), nil)
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

func (c *gooddataAPI) DeleteUserGroup(id string) error {
	url, err := c.url(fmt.Sprintf(userGroupEndpoint, id), nil)
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

func (c *gooddataAPI) UpdateUserGroup(userGroup UserGroup) (UserGroup, error) {
	result := UserGroup{}

	url, err := c.url(fmt.Sprintf(userGroupEndpoint, userGroup.ID), nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("PUT", url, header, &userGroup)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
