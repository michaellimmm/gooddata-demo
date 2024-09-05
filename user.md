# Users

ref: [manage-users](https://www.gooddata.com/docs/cloud/manage-organization/manage-users/)

## Create User

```bash
curl $HOST_URL/api/v1/entities/users \
-H "Authorization: Bearer $API_TOKEN" \
-H "Content-Type: application/vnd.gooddata.api+json" \
-X POST \
-d '{
    "data": {
        "id": "<userId>",
        "type": "user",
        "attributes": {
        "authenticationId": "<subClaim>",
        "email": "<email>",
        "firstname": "<firstName>",
        "lastname": "<lastName>"
        },
        "relationships": {
        "userGroups": {
            "data": [{
            "id": "<userGroupId>",
            "type": "userGroup"
            }]
        }
        }
    }
}'
```

notes:

- *authenticationId* is optional

## Create a User Group

To effectively set up a user group through the API, you'll need to complete three key actions:

-

### Create a new user group

```bash
curl $HOST_URL/api/v1/entities/userGroups \
-H "Content-Type: application/vnd.gooddata.api+json" \
-H "Accept: application/vnd.gooddata.api+json" \
-H "Authorization: Bearer $API_TOKEN" \
-X POST \
-d '{
    "data": {
        "type": "userGroup",
        "id": "<userGroupId>",
        "attributes": {
        "name": "<userGroupDisplayName>"
        }
    }
}'
```

### Link the user group to the desired workspaces

```bash
curl $HOST_URL/api/v1/actions/workspaces/<workspaceId>/managePermissions/ \
    -H "Authorization: Bearer $API_TOKEN" \
    -H "Content-Type: application/json" \
    -X POST \
    -d '[
    {
        "assigneeIdentifier": {
            "id": "<userGroupId>",
            "type": "userGroup"
        },
        "permissions": ["VIEW"]
    }
]'
```

#### Supported Permissions

- `VIEW` -> user can view dashboard that have been shared to them
- `ANALYZE` and `EXPORT` -> either permission gives you the same level of access as the VIEW permission. Additionally:
  - `ANALYZE` -> user can also create, edit or delete dashboard and visualization, and view the LDM and metrics.
  - `EXPORT` -> user can view and export dashboard to PDF files and tabluar data from visualization to XLSX and CSV files.
    - The `EXPORT` permission has more granular sub-permissions:
      - `EXPORT_PDF` -> user can only view and export dashboards to PDF files.
      - `EXPORT_TABULAR` -> user can only view and export tabluar data from visualizations to XLSX and CSV files.
- `MANAGE` -> covers VIEW, ANALYZE, and EXPORT permissions. Additionally, a user can create, edit, or delete the logical data model and metrics, and access all dashboards and edit their dashboard permissions without limitations.

Enum: `"MANAGE"`, `"ANALYZE"`, `"EXPORT"`, `"EXPORT_TABULAR"`, `"EXPORT_PDF"`, `"VIEW"`

##### Workspace Permissions

there are two types of permission:

- `permissions` are tied to a specific workspace and define what a suer can do with that one specific workspace.
- `hierarchyPermissions` are tied to a specific workspace and define what a user can do with that specific workspace _and all of its child workspaces_.

### Add users to the user group

```bash
curl $HOST_URL/api/v1/actions/userManagement/userGroups/<userGroupId>/addMembers \
-H "Authorization: Bearer $API_TOKEN" \
-H "Content-Type: application/vnd.gooddata.api+json" \
-X POST \
-d '{
    "data": {
        "members": [
            {
                "id": "<userId>"
            }
        ]
    }
}'
```
