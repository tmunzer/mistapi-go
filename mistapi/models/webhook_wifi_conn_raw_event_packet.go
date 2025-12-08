// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookWifiConnRawEventPacket represents a WebhookWifiConnRawEventPacket struct.
type WebhookWifiConnRawEventPacket struct {
	Band                 *string                `json:"band,omitempty"`
	Rssi                 *int                   `json:"rssi,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookWifiConnRawEventPacket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookWifiConnRawEventPacket) String() string {
	return fmt.Sprintf(
		"WebhookWifiConnRawEventPacket[Band=%v, Rssi=%v, AdditionalProperties=%v]",
		w.Band, w.Rssi, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookWifiConnRawEventPacket.
// It customizes the JSON marshaling process for WebhookWifiConnRawEventPacket objects.
func (w WebhookWifiConnRawEventPacket) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"band", "rssi"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookWifiConnRawEventPacket object to a map representation for JSON marshaling.
func (w WebhookWifiConnRawEventPacket) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.Band != nil {
		structMap["band"] = w.Band
	}
	if w.Rssi != nil {
		structMap["rssi"] = w.Rssi
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookWifiConnRawEventPacket.
// It customizes the JSON unmarshaling process for WebhookWifiConnRawEventPacket objects.
func (w *WebhookWifiConnRawEventPacket) UnmarshalJSON(input []byte) error {
	var temp tempWebhookWifiConnRawEventPacket
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "band", "rssi")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Band = temp.Band
	w.Rssi = temp.Rssi
	return nil
}

// tempWebhookWifiConnRawEventPacket is a temporary struct used for validating the fields of WebhookWifiConnRawEventPacket.
type tempWebhookWifiConnRawEventPacket struct {
	Band *string `json:"band,omitempty"`
	Rssi *int    `json:"rssi,omitempty"`
}
