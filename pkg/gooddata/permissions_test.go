package gooddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnMarshalManagePermissions(t *testing.T) {
	jsonString := `
	[
		{
			"assigneeIdentifier": {
				"id": "user_id",
				"type": "user"
			},
			"permissions": [
				"MANAGE"
			]
		}
	]
	`

	var actual ManagePermissions
	err := actual.Unmarshal([]byte(jsonString))
	assert.Nil(t, err)

	expected := ManagePermissions{
		{
			AssigneeIdentifier: AssigneeIdentifier{
				ID:                     "user_id",
				AssigneeIdentifierType: AssigneeIdentifierTypeUser,
			},
			Permissions: []PermissionType{PermissionTypeManage},
		},
	}

	assert.Equal(t, expected, actual)
}

func TestMarshalManagePermissions(t *testing.T) {
	data := ManagePermissions{
		{
			AssigneeIdentifier: AssigneeIdentifier{
				ID:                     "user_id",
				AssigneeIdentifierType: AssigneeIdentifierTypeUser,
			},
			Permissions: []PermissionType{PermissionTypeManage},
		},
	}
	actual, err := data.Marshal()
	assert.Nil(t, err)

	expected := `[{"assigneeIdentifier":{"id":"user_id","type":"user"},"permissions":["MANAGE"]}]`
	assert.Equal(t, expected, string(actual))
}
