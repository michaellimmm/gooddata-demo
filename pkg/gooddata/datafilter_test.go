package gooddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnMarshalUserDataFilter(t *testing.T) {
	jsonString := `
	{
		"data": {
			"type": "userDataFilter",
			"id": "udf_tenant1",
			"attributes": {
				"maql": "{label/TENANT_ID} = \"tenant1\"",
				"title": "tenant1 filter",
				"description": "tenant filter for tenant1"
			},
			"relationships": {
				"user": {
					"data": {
						"id": "tenant1",
						"type": "user"
					}
				}
			}
		}
	}
	`

	actual := UserDataFilter{}
	err := actual.Unmarshal([]byte(jsonString))
	assert.Nil(t, err)

	expected := UserDataFilter{
		ID:          "udf_tenant1",
		Maql:        "{label/TENANT_ID} = \"tenant1\"",
		Title:       "tenant1 filter",
		Description: "tenant filter for tenant1",
		User: &User{
			ID: "tenant1",
		},
	}
	assert.Equal(t, expected, actual)
}

func TestMarshalUserDataFilter(t *testing.T) {
	data := UserDataFilter{
		ID:          "udf_tenant1",
		Maql:        "{label/TENANT_ID} = \"tenant1\"",
		Title:       "tenant1 filter",
		Description: "tenant filter for tenant1",
		User: &User{
			ID: "tenant1",
		},
	}

	actual, err := data.Marshal()
	assert.Nil(t, err)

	expected := `{"data":{"id":"udf_tenant1","type":"userDataFilter","attributes":{"description":"tenant filter for tenant1","maql":"{label/TENANT_ID} = \"tenant1\"","title":"tenant1 filter"},"relationships":{"user":{"data":{"id":"tenant1","type":"user"}}}}}`

	assert.Equal(t, expected, string(actual))
}
