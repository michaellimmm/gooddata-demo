package gooddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshalUserGroup(t *testing.T) {
	jsonString := `
	{
		"data": {
			"id": "ug_tenant",
			"type": "userGroup",
			"attributes": {
				"name": "tenant_group"
			}
		},
		"links": {
			"self": "https://helpful-duck.trial.cloud.gooddata.com/api/v1/entities/userGroups/ug_tenant"
		}
	}
	`

	actual := UserGroup{}
	err := actual.Unmarshal([]byte(jsonString))
	assert.Nil(t, err)

	expected := UserGroup{
		ID:   "ug_tenant",
		Name: "tenant_group",
	}
	assert.Equal(t, expected, actual)
}

func TestUnmarshalUserGroups(t *testing.T) {
	jsonString := `
	{
		"data": [
			{
				"id": "adminGroup",
				"type": "userGroup",
				"attributes": {},
				"links": {
					"self": "https://marvel.cloud.gooddata.com/api/v1/entities/userGroups/adminGroup"
				}
			},
			{
				"id": "ug_asgard",
				"type": "userGroup",
				"attributes": {
					"name": "asgard"
				},
				"links": {
					"self": "https://marvel.cloud.gooddata.com/api/v1/entities/userGroups/ug_tenant"
				}
			}
		],
		"links": {
			"self": "https://marvel.cloud.gooddata.com/api/v1/entities/userGroups?page=0&size=20",
			"next": "https://marvel.cloud.gooddata.com/api/v1/entities/userGroups?page=1&size=20"
		}
	}
	`

	var actual UserGroups
	err := actual.Unmarshal([]byte(jsonString))
	assert.Nil(t, err)

	expected := UserGroups{
		{
			ID: "adminGroup",
		},
		{
			ID:   "ug_asgard",
			Name: "asgard",
		},
	}
	assert.Equal(t, expected, actual)
}
