// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ConstGatewayApplicationsDefinition represents a ConstGatewayApplicationsDefinition struct.
type ConstGatewayApplicationsDefinition struct {
    AppId                *bool                  `json:"app_id,omitempty"`
    Key                  *string                `json:"key,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    SsrAppId             *bool                  `json:"ssr_app_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstGatewayApplicationsDefinition,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstGatewayApplicationsDefinition) String() string {
    return fmt.Sprintf(
    	"ConstGatewayApplicationsDefinition[AppId=%v, Key=%v, Name=%v, SsrAppId=%v, AdditionalProperties=%v]",
    	c.AppId, c.Key, c.Name, c.SsrAppId, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstGatewayApplicationsDefinition.
// It customizes the JSON marshaling process for ConstGatewayApplicationsDefinition objects.
func (c ConstGatewayApplicationsDefinition) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "app_id", "key", "name", "ssr_app_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstGatewayApplicationsDefinition object to a map representation for JSON marshaling.
func (c ConstGatewayApplicationsDefinition) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "app_id", "key", "name", "ssr_app_id")
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
type tempConstGatewayApplicationsDefinition  struct {
    AppId    *bool   `json:"app_id,omitempty"`
    Key      *string `json:"key,omitempty"`
    Name     *string `json:"name,omitempty"`
    SsrAppId *bool   `json:"ssr_app_id,omitempty"`
}
