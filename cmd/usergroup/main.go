package main

import (
	"fmt"
	"github/michaellimmm/gooddata-demo/pkg/gooddata"
	"os"

	"github.com/joho/godotenv"
)

// this script create usergroup and assign permission to usergroup

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("failed to load env var, err: %+v", err)
		return
	}

	baseUrl := os.Getenv("GOODDATA_BASEURL")
	accessToken := os.Getenv("GOODDATA_ACCESSTOKEN")
	workspaceID := os.Getenv("WORKSPACE_ID")

	api, err := gooddata.NewGooddataAPI(baseUrl, accessToken)
	if err != nil {
		fmt.Printf("failed to initialize gooddataAPI, err: %+v", err)
		return
	}

	userGroup := gooddata.UserGroup{
		ID:   "ug_tenant",
		Name: "tenant_group",
	}

	usergroups, err := api.CreateUserGroup(userGroup)
	if err != nil {
		fmt.Printf("faile to get list of usergroups, err: %+v", err)
		return
	}

	fmt.Printf("result: %+v\n", usergroups)

	permissions := gooddata.ManagePermissions{
		{
			AssigneeIdentifier: gooddata.AssigneeIdentifier{
				ID:                     usergroups.ID,
				AssigneeIdentifierType: gooddata.AssigneeIdentifierTypeUserGroup,
			},
			Permissions: []gooddata.PermissionType{gooddata.PermissionTypeExport},
		},
	}
	err = api.ManagePermissionWorkspace(workspaceID, permissions)
	if err != nil {
		fmt.Printf("faile to assign permission to usergroup, err: %+v", err)
		return
	}
}
