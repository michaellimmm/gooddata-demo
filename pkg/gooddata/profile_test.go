package gooddata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnMarshalProfile(t *testing.T) {
	jsonString := `
	{
		"organizationId": "cicxeirayj",
		"organizationName": "saul",
		"userId": "5000d6e2-6dac-4a6a-b136-6ba8369d1d64",
		"permissions": [
			"MANAGE",
			"SELF_CREATE_TOKEN"
		],
		"telemetry": {
			"host": "https://matomo.anywhere.gooddata.com/matomo.php",
			"siteId": 3,
			"deploymentId": "deploymentId",
			"organizationHash": "f2c06191e3d0b9494d7de8f79a225702da5ee3c2",
			"userHash": "46af819a0d75934ed2e4f7d0f847e20d806b193a"
		},
		"telemetryConfig": {
			"context": {
				"deploymentId": "deploymentId",
				"organizationHash": "f2c06191e3d0b9494d7de8f79a225702da5ee3c2",
				"userHash": "46af819a0d75934ed2e4f7d0f847e20d806b193a"
			},
			"services": {
				"matomo": {
					"host": "https://matomo.anywhere.gooddata.com/matomo.php",
					"siteId": 3
				},
				"openTelemetry": {
					"host": "https://collector.iad1.prodgdc.com"
				}
			}
		},
		"links": {
			"self": "https://helpful-duck.trial.cloud.gooddata.com/api/v1/profile",
			"organization": "https://helpful-duck.trial.cloud.gooddata.com/api/v1/entities/admin/organizations/cicxeirayj",
			"user": "https://helpful-duck.trial.cloud.gooddata.com/api/v1/entities/users/5000d6e2-6dac-4a6a-b136-6ba8369d1d64"
		},
		"features": {
			"live": {
				"configuration": {
					"host": "https://flags.cloud.gooddata.com",
					"key": "default/09beb6d5-4175-461b-93a8-15245b1be0bb/OjSMKJUXkcIy74nPbAO1WHbsgKzAeCypg3cQazuv"
				},
				"context": {
					"earlyAccess": ""
				}
			}
		},
		"entitlements": [
			{
				"name": "WorkspaceCount",
				"value": "10"
			}
		]
	}
	`

	actual := Profile{}
	err := actual.Unmarshal([]byte(jsonString))
	assert.Nil(t, err)

	expected := Profile{
		UserID:           "5000d6e2-6dac-4a6a-b136-6ba8369d1d64",
		OrganizationID:   "cicxeirayj",
		OrganizationName: "saul",
		Permissions:      []PermissionType{PermissionTypeManage, PermissionTypeSelfCreateToken},
		Telemetry: Telemetry{
			Host:             "https://matomo.anywhere.gooddata.com/matomo.php",
			SiteID:           3,
			DeploymentID:     "deploymentId",
			OrganizationHash: "f2c06191e3d0b9494d7de8f79a225702da5ee3c2",
			UserHash:         "46af819a0d75934ed2e4f7d0f847e20d806b193a",
		},
		TelemetryConfig: TelemetryConfig{
			Context: TelemetryContext{
				DeploymentID:     "deploymentId",
				OrganizationHash: "f2c06191e3d0b9494d7de8f79a225702da5ee3c2",
				UserHash:         "46af819a0d75934ed2e4f7d0f847e20d806b193a",
			},
			Services: TelemetryServices{
				Matomo: MatomoService{
					Host:   "https://matomo.anywhere.gooddata.com/matomo.php",
					SiteID: 3,
				},
				OpenTelemetry: OpenTelemetryService{
					Host: "https://collector.iad1.prodgdc.com",
				},
			},
		},
		Features: Features{
			Live: &LiveFeatures{
				Configuration: LiveFeatureFlagConfiguration{
					Host: "https://flags.cloud.gooddata.com",
					Key:  "default/09beb6d5-4175-461b-93a8-15245b1be0bb/OjSMKJUXkcIy74nPbAO1WHbsgKzAeCypg3cQazuv",
				},
			},
		},
		Entitlements: []Entitlement{
			{
				Name:  "WorkspaceCount",
				Value: "10",
			},
		},
	}
	assert.Equal(t, expected, actual)
}
