// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookWifiUnconnRawEventPacket represents a WebhookWifiUnconnRawEventPacket struct.
type WebhookWifiUnconnRawEventPacket struct {
	Band                 *string                `json:"band,omitempty"`
	Rssi                 *int                   `json:"rssi,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookWifiUnconnRawEventPacket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookWifiUnconnRawEventPacket) String() string {
	return fmt.Sprintf(
		"WebhookWifiUnconnRawEventPacket[Band=%v, Rssi=%v, AdditionalProperties=%v]",
		w.Band, w.Rssi, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookWifiUnconnRawEventPacket.
// It customizes the JSON marshaling process for WebhookWifiUnconnRawEventPacket objects.
func (w WebhookWifiUnconnRawEventPacket) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"band", "rssi"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookWifiUnconnRawEventPacket object to a map representation for JSON marshaling.
func (w WebhookWifiUnconnRawEventPacket) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookWifiUnconnRawEventPacket.
// It customizes the JSON unmarshaling process for WebhookWifiUnconnRawEventPacket objects.
func (w *WebhookWifiUnconnRawEventPacket) UnmarshalJSON(input []byte) error {
	var temp tempWebhookWifiUnconnRawEventPacket
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

// tempWebhookWifiUnconnRawEventPacket is a temporary struct used for validating the fields of WebhookWifiUnconnRawEventPacket.
type tempWebhookWifiUnconnRawEventPacket struct {
	Band *string `json:"band,omitempty"`
	Rssi *int    `json:"rssi,omitempty"`
}
