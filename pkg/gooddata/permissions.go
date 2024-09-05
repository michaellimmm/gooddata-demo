package gooddata

import (
	"encoding/json"
	"fmt"
)

const (
	managePermissionsForWorkspaceAPI = "/api/v1/actions/workspaces/%s/managePermissions"
)

type PermissionType string

const (
	PermissionTypeManage          PermissionType = "MANAGE"
	PermissionTypeAnalyze         PermissionType = "ANALYZE"
	PermissionTypeExport          PermissionType = "EXPORT"
	PermissionTypeExportTabular   PermissionType = "EXPORT_TABULAR"
	PermissionTypeExportPdf       PermissionType = "EXPORT_PDF"
	PermissionTypeView            PermissionType = "VIEW"
	PermissionTypeSelfCreateToken PermissionType = "SELF_CREATE_TOKEN"
)

type AssigneeIdentifierType string

const (
	AssigneeIdentifierTypeUser      AssigneeIdentifierType = "user"
	AssigneeIdentifierTypeUserGroup AssigneeIdentifierType = "userGroup"
)

type AssigneeIdentifier struct {
	ID                     string                 `json:"id"`
	AssigneeIdentifierType AssigneeIdentifierType `json:"type"`
}

type ManagePermission struct {
	AssigneeIdentifier   AssigneeIdentifier `json:"assigneeIdentifier"`
	HierarchyPermissions []PermissionType   `json:"hierarchyPermissions,omitempty"`
	Permissions          []PermissionType   `json:"permissions"`
}

type ManagePermissions []ManagePermission

func (mp *ManagePermissions) Marshal() ([]byte, error) {
	return json.Marshal(mp)
}

func (mp *ManagePermissions) Unmarshal(data []byte) error {
	return json.Unmarshal(data, mp)
}

type PermissionsAPI interface {
	ManagePermissionWorkspace(workspaceID string, permissions ManagePermissions) error
}

func (c *gooddataAPI) ManagePermissionWorkspace(workspaceID string, permissions ManagePermissions) error {
	url, err := c.url(fmt.Sprintf(managePermissionsForWorkspaceAPI, workspaceID), nil)
	if err != nil {
		return nil
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	req, err := c.newRequest("POST", url, header, &permissions)
	if err != nil {
		return err
	}

	return c.do(req, nil)
}
