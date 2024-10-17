package models

import (
	"encoding/json"
)

// ConstGatewayApplicationsDefinition represents a ConstGatewayApplicationsDefinition struct.
type ConstGatewayApplicationsDefinition struct {
	AppId                *bool          `json:"app_id,omitempty"`
	Key                  *string        `json:"key,omitempty"`
	Name                 *string        `json:"name,omitempty"`
	SsrAppId             *bool          `json:"ssr_app_id,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstGatewayApplicationsDefinition.
// It customizes the JSON marshaling process for ConstGatewayApplicationsDefinition objects.
func (c ConstGatewayApplicationsDefinition) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(c.toMap())
}

// toMap converts the ConstGatewayApplicationsDefinition object to a map representation for JSON marshaling.
func (c ConstGatewayApplicationsDefinition) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, c.AdditionalProperties)
	if c.AppId != nil {
		structMap["app_id"] = c.AppId
	}
	if c.Key != nil {
		structMap["key"] = c.Key
	}
	if c.Name != nil {
		structMap["name"] = c.Name
	}
	if c.SsrAppId != nil {
		structMap["ssr_app_id"] = c.SsrAppId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstGatewayApplicationsDefinition.
// It customizes the JSON unmarshaling process for ConstGatewayApplicationsDefinition objects.
func (c *ConstGatewayApplicationsDefinition) UnmarshalJSON(input []byte) error {
	var temp tempConstGatewayApplicationsDefinition
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "app_id", "key", "name", "ssr_app_id")
	if err != nil {
		return err
	}

	c.AdditionalProperties = additionalProperties
	c.AppId = temp.AppId
	c.Key = temp.Key
	c.Name = temp.Name
	c.SsrAppId = temp.SsrAppId
	return nil
}

// tempConstGatewayApplicationsDefinition is a temporary struct used for validating the fields of ConstGatewayApplicationsDefinition.
type tempConstGatewayApplicationsDefinition struct {
	AppId    *bool   `json:"app_id,omitempty"`
	Key      *string `json:"key,omitempty"`
	Name     *string `json:"name,omitempty"`
	SsrAppId *bool   `json:"ssr_app_id,omitempty"`
}
