package gooddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TesUnMarshallUser(t *testing.T) {
	jsonString := `
	{
		"data": {
			"id": "user1",
			"type": "user",
			"attributes": {
				"authenticationId": "12345",
				"firstname": "Gal",
				"lastname": "Gadot",
				"email": "gal.gadot@marvel.com"
			}
		},
		"links": {
			"self": "https://marvel.cloud.gooddata.com/api/v1/entities/users/user1"
		}
	}`

	actual := User{}
	err := actual.Unmarshal([]byte(jsonString))
	if err != nil {
		t.Errorf("failed Unmarshal payload: %s, err: %+v", jsonString, err)
	}

	expected := User{
		ID:               "user1",
		AuthenticationID: "12345",
		Firstname:        "Gal",
		LastName:         "Gadot",
		Email:            "gal.gadot@marvel.com",
	}

	assert.Equal(t, expected, actual)
}

func TestUnMarshallUsers(t *testing.T) {
	jsonString := `
	{
		"data": [
			{
				"id": "user1",
				"type": "user",
				"attributes": {
					"authenticationId": "12345",
					"firstname": "Gal",
					"lastname": "Gadot",
					"email": "gal.gadot@marvel.com"
				},
				"links": {
					"self": "https://marvel.cloud.gooddata.com/api/v1/entities/users/user1"
				}
			},
			{
				"id": "user2",
				"type": "user",
				"attributes": {
					"authenticationId": "12346",
					"firstname": "Brie",
					"lastname": "Larson",
					"email": "brie.larson@marvel.com"
				},
				"links": {
					"self": "https://marvel.cloud.gooddata.com/api/v1/entities/users/user1"
				}
			}
		],
		"links": {
			"self": "https://marvel.cloud.gooddata.com/api/v1/entities/users?page=0&size=20",
			"next": "https://marvel.cloud.gooddata.com/api/v1/entities/users?page=1&size=20"
		}
	}
	`

	var actual Users
	err := actual.Unmarshal([]byte(jsonString))
	if err != nil {
		t.Errorf("failed Unmarshal payload: %s, err: %+v", jsonString, err)
	}

	expected := Users{
		{
			ID:               "user1",
			AuthenticationID: "12345",
			Firstname:        "Gal",
			LastName:         "Gadot",
			Email:            "gal.gadot@marvel.com",
		},
		{
			ID:               "user2",
			AuthenticationID: "12346",
			Firstname:        "Brie",
			LastName:         "Larson",
			Email:            "brie.larson@marvel.com",
		},
	}

	assert.Equal(t, expected, actual)
}
