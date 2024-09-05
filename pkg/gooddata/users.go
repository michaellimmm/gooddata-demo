package gooddata

import (
	"fmt"

	"github.com/DataDog/jsonapi"
)

const (
	usersEndpoint = "/api/v1/entities/users"
	userEndpoint  = "/api/v1/entities/users/%s"
)

type User struct {
	ID               string       `jsonapi:"primary,user"`
	AuthenticationID string       `jsonapi:"attribute" json:"authenticationId,omitempty"`
	Email            string       `jsonapi:"attribute" json:"email,omitempty"`
	Firstname        string       `jsonapi:"attribute" json:"firstname,omitempty"`
	LastName         string       `jsonapi:"attribute" json:"lastname,omitempty"`
	UserGroups       []*UserGroup `jsonapi:"relationship" json:"userGroups,omitempty"`
}

func (u *User) Marshal() ([]byte, error) {
	return jsonapi.Marshal(u)
}

func (u *User) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, u)
}

type Users []User

func (u *Users) Marshal() ([]byte, error) {
	return jsonapi.Marshal(u)
}

func (u *Users) Unmarshal(data []byte) error {
	return jsonapi.Unmarshal(data, u)
}

type UserAPI interface {
	ListUsers() (Users, error)
	CreateUser(user User) (User, error)
	GetUser(id string) (User, error)
	DeleteUser(id string) error
	UpdateUser(user User) (User, error)
}

// TODO: add filter, include, page, size, sort, metainclude
// TODO: research about rsql
func (c *gooddataAPI) ListUsers() (Users, error) {
	var result Users

	url, err := c.url(usersEndpoint, nil)
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

func (c *gooddataAPI) CreateUser(user User) (User, error) {
	result := User{}

	url, err := c.url(usersEndpoint, nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("POST", url, header, &user)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}

func (c *gooddataAPI) GetUser(id string) (User, error) {
	result := User{}

	url, err := c.url(fmt.Sprintf(userEndpoint, id), nil)
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

func (c *gooddataAPI) DeleteUser(id string) error {
	url, err := c.url(fmt.Sprintf(userEndpoint, id), nil)
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

func (c *gooddataAPI) UpdateUser(user User) (User, error) {
	result := User{}

	url, err := c.url(fmt.Sprintf(userEndpoint, user.ID), nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	req, err := c.newRequest("PUT", url, header, &user)
	if err != nil {
		return result, err
	}

	err = c.do(req, &result)
	if err != nil {
		return result, err
	}

	return result, nil
}
