// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ApMqtt represents a ApMqtt struct.
// MQTT broker publishing settings for an AP; use `mqtt_topic` on individual AssetFilter entries to specify which MQTT topic each matching BLE advertisement is forwarded to
type ApMqtt struct {
	// MQTT broker hostname or IP address; required when `enabled` is `true`
	BrokerHost *string `json:"broker_host,omitempty"`
	// MQTT broker port; defaults to `1883` for `tcp` and `8883` for `ssl`
	BrokerPort *int `json:"broker_port,omitempty"`
	// MQTT broker transport protocol. enum: `ssl`, `tcp`
	BrokerProto *ApMqttBrokerProtoEnum `json:"broker_proto,omitempty"`
	// Whether to enable MQTT publishing
	Enabled *bool `json:"enabled,omitempty"`
	// Payload format for MQTT published messages. enum: `json`, `raw`
	Format *ApMqttFormatEnum `json:"format,omitempty"`
	// Optional MQTT password; masked in GET responses
	Password *string `json:"password,omitempty"`
	// Optional MQTT username
	Username             *string                `json:"username,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ApMqtt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a ApMqtt) String() string {
	return fmt.Sprintf(
		"ApMqtt[BrokerHost=%v, BrokerPort=%v, BrokerProto=%v, Enabled=%v, Format=%v, Password=%v, Username=%v, AdditionalProperties=%v]",
		a.BrokerHost, a.BrokerPort, a.BrokerProto, a.Enabled, a.Format, a.Password, a.Username, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ApMqtt.
// It customizes the JSON marshaling process for ApMqtt objects.
func (a ApMqtt) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(a.AdditionalProperties,
		"broker_host", "broker_port", "broker_proto", "enabled", "format", "password", "username"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(a.toMap())
}

// toMap converts the ApMqtt object to a map representation for JSON marshaling.
func (a ApMqtt) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, a.AdditionalProperties)
	if a.BrokerHost != nil {
		structMap["broker_host"] = a.BrokerHost
	}
	if a.BrokerPort != nil {
		structMap["broker_port"] = a.BrokerPort
	}
	if a.BrokerProto != nil {
		structMap["broker_proto"] = a.BrokerProto
	}
	if a.Enabled != nil {
		structMap["enabled"] = a.Enabled
	}
	if a.Format != nil {
		structMap["format"] = a.Format
	}
	if a.Password != nil {
		structMap["password"] = a.Password
	}
	if a.Username != nil {
		structMap["username"] = a.Username
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApMqtt.
// It customizes the JSON unmarshaling process for ApMqtt objects.
func (a *ApMqtt) UnmarshalJSON(input []byte) error {
	var temp tempApMqtt
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "broker_host", "broker_port", "broker_proto", "enabled", "format", "password", "username")
	if err != nil {
		return err
	}
	a.AdditionalProperties = additionalProperties

	a.BrokerHost = temp.BrokerHost
	a.BrokerPort = temp.BrokerPort
	a.BrokerProto = temp.BrokerProto
	a.Enabled = temp.Enabled
	a.Format = temp.Format
	a.Password = temp.Password
	a.Username = temp.Username
	return nil
}

// tempApMqtt is a temporary struct used for validating the fields of ApMqtt.
type tempApMqtt struct {
	BrokerHost  *string                `json:"broker_host,omitempty"`
	BrokerPort  *int                   `json:"broker_port,omitempty"`
	BrokerProto *ApMqttBrokerProtoEnum `json:"broker_proto,omitempty"`
	Enabled     *bool                  `json:"enabled,omitempty"`
	Format      *ApMqttFormatEnum      `json:"format,omitempty"`
	Password    *string                `json:"password,omitempty"`
	Username    *string                `json:"username,omitempty"`
}
