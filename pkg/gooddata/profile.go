package gooddata

import "encoding/json"

const (
	profileEndpoint = "/api/v1/profile"
)

type Entitlement struct {
	Expiry string `json:"expiry,omitempty"`
	Name   string `json:"name"`
	Value  string `json:"value,omitempty"`
}

type Features struct {
	Live *LiveFeatures `json:"live,omitempty"`
}

type LiveFeatures struct {
	Context       FeatureFlagsContext          `json:"context"`
	Configuration LiveFeatureFlagConfiguration `json:"configuration"`
}

type FeatureFlagsContext struct {
	EarlyAccess   string                       `json:"earlyAccess"`
	Configuration LiveFeatureFlagConfiguration `json:"configuration"`
}

type LiveFeatureFlagConfiguration struct {
	Host string `json:"host"`
	Key  string `json:"key"`
}

type Telemetry struct {
	DeploymentID     string `json:"deploymentId"`
	Host             string `json:"host"`
	OrganizationHash string `json:"organizationHash"`
	SiteID           int32  `json:"siteId"`
	UserHash         string `json:"userHash"`
}

type TelemetryConfig struct {
	Context  TelemetryContext  `json:"context"`
	Services TelemetryServices `json:"services"`
}

type TelemetryContext struct {
	DeploymentID     string `json:"deploymentId"`
	OrganizationHash string `json:"organizationHash"`
	UserHash         string `json:"userHash"`
}

type TelemetryServices struct {
	Matomo        MatomoService        `json:"matomo"`
	OpenTelemetry OpenTelemetryService `json:"openTelemetry"`
}

type MatomoService struct {
	Host   string `json:"host"`
	SiteID int32  `json:"siteId"`
}

type OpenTelemetryService struct {
	Host string `json:"host"`
}

type Profile struct {
	UserID           string           `json:"userId"`
	Name             string           `json:"name,omitempty"`
	OrganizationID   string           `json:"organizationId"`
	OrganizationName string           `json:"organizationName"`
	Permissions      []PermissionType `json:"permissions"`
	Telemetry        Telemetry        `json:"telemetry"`
	TelemetryConfig  TelemetryConfig  `json:"telemetryConfig"`
	Features         Features         `json:"features"`
	Entitlements     []Entitlement    `json:"entitlements"`
}

func (p *Profile) Marshal() ([]byte, error) {
	return json.Marshal(p)
}

func (p *Profile) Unmarshal(data []byte) error {
	return json.Unmarshal(data, p)
}

type ProfileAPI interface {
	GetProfile() (Profile, error)
}

func (c *gooddataAPI) GetProfile() (Profile, error) {
	result := Profile{}

	url, err := c.url(profileEndpoint, nil)
	if err != nil {
		return result, err
	}

	header := make(map[string]string)
	header["Content-Type"] = "application/json"

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
