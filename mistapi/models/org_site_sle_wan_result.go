package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// OrgSiteSleWanResult represents a OrgSiteSleWanResult struct.
type OrgSiteSleWanResult struct {
	ApplicationHealth    *float64       `json:"application_health,omitempty"`
	GatewayHealth        *float64       `json:"gateway-health,omitempty"`
	NumClients           *float64       `json:"num_clients,omitempty"`
	NumGateways          *float64       `json:"num_gateways,omitempty"`
	SiteId               uuid.UUID      `json:"site_id"`
	WanLinkHealth        *float64       `json:"wan-link-health,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSiteSleWanResult.
// It customizes the JSON marshaling process for OrgSiteSleWanResult objects.
func (o OrgSiteSleWanResult) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgSiteSleWanResult object to a map representation for JSON marshaling.
func (o OrgSiteSleWanResult) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.ApplicationHealth != nil {
		structMap["application_health"] = o.ApplicationHealth
	}
	if o.GatewayHealth != nil {
		structMap["gateway-health"] = o.GatewayHealth
	}
	if o.NumClients != nil {
		structMap["num_clients"] = o.NumClients
	}
	if o.NumGateways != nil {
		structMap["num_gateways"] = o.NumGateways
	}
	structMap["site_id"] = o.SiteId
	if o.WanLinkHealth != nil {
		structMap["wan-link-health"] = o.WanLinkHealth
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSiteSleWanResult.
// It customizes the JSON unmarshaling process for OrgSiteSleWanResult objects.
func (o *OrgSiteSleWanResult) UnmarshalJSON(input []byte) error {
	var temp tempOrgSiteSleWanResult
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "application_health", "gateway-health", "num_clients", "num_gateways", "site_id", "wan-link-health")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.ApplicationHealth = temp.ApplicationHealth
	o.GatewayHealth = temp.GatewayHealth
	o.NumClients = temp.NumClients
	o.NumGateways = temp.NumGateways
	o.SiteId = *temp.SiteId
	o.WanLinkHealth = temp.WanLinkHealth
	return nil
}

// tempOrgSiteSleWanResult is a temporary struct used for validating the fields of OrgSiteSleWanResult.
type tempOrgSiteSleWanResult struct {
	ApplicationHealth *float64   `json:"application_health,omitempty"`
	GatewayHealth     *float64   `json:"gateway-health,omitempty"`
	NumClients        *float64   `json:"num_clients,omitempty"`
	NumGateways       *float64   `json:"num_gateways,omitempty"`
	SiteId            *uuid.UUID `json:"site_id"`
	WanLinkHealth     *float64   `json:"wan-link-health,omitempty"`
}

func (o *tempOrgSiteSleWanResult) validate() error {
	var errs []string
	if o.SiteId == nil {
		errs = append(errs, "required field `site_id` is missing for type `org_site_sle_wan_result`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
