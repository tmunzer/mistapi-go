// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// IotproxyVisionline represents a IotproxyVisionline struct.
// Visionline integration settings for IoT proxy
type IotproxyVisionline struct {
	// Access ID for the Visionline service
	AccessId *string `json:"access_id,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty"`
	// Hostname or IP of the Visionline collector
	Host *string `json:"host,omitempty"`
	// Password for the Visionline service
	Password *string `json:"password,omitempty"`
	// TCP port of the Visionline collector
	Port *int `json:"port,omitempty"`
	// Username for the Visionline service
	Username             *string                `json:"username,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IotproxyVisionline,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IotproxyVisionline) String() string {
	return fmt.Sprintf(
		"IotproxyVisionline[AccessId=%v, Enabled=%v, Host=%v, Password=%v, Port=%v, Username=%v, AdditionalProperties=%v]",
		i.AccessId, i.Enabled, i.Host, i.Password, i.Port, i.Username, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IotproxyVisionline.
// It customizes the JSON marshaling process for IotproxyVisionline objects.
func (i IotproxyVisionline) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(i.AdditionalProperties,
		"access_id", "enabled", "host", "password", "port", "username"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(i.toMap())
}

// toMap converts the IotproxyVisionline object to a map representation for JSON marshaling.
func (i IotproxyVisionline) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, i.AdditionalProperties)
	if i.AccessId != nil {
		structMap["access_id"] = i.AccessId
	}
	if i.Enabled != nil {
		structMap["enabled"] = i.Enabled
	}
	if i.Host != nil {
		structMap["host"] = i.Host
	}
	if i.Password != nil {
		structMap["password"] = i.Password
	}
	if i.Port != nil {
		structMap["port"] = i.Port
	}
	if i.Username != nil {
		structMap["username"] = i.Username
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IotproxyVisionline.
// It customizes the JSON unmarshaling process for IotproxyVisionline objects.
func (i *IotproxyVisionline) UnmarshalJSON(input []byte) error {
	var temp tempIotproxyVisionline
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "access_id", "enabled", "host", "password", "port", "username")
	if err != nil {
		return err
	}
	i.AdditionalProperties = additionalProperties

	i.AccessId = temp.AccessId
	i.Enabled = temp.Enabled
	i.Host = temp.Host
	i.Password = temp.Password
	i.Port = temp.Port
	i.Username = temp.Username
	return nil
}

// tempIotproxyVisionline is a temporary struct used for validating the fields of IotproxyVisionline.
type tempIotproxyVisionline struct {
	AccessId *string `json:"access_id,omitempty"`
	Enabled  *bool   `json:"enabled,omitempty"`
	Host     *string `json:"host,omitempty"`
	Password *string `json:"password,omitempty"`
	Port     *int    `json:"port,omitempty"`
	Username *string `json:"username,omitempty"`
}
