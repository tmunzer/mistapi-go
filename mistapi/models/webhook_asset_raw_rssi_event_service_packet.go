// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookAssetRawRssiEventServicePacket represents a WebhookAssetRawRssiEventServicePacket struct.
type WebhookAssetRawRssiEventServicePacket struct {
	// optional, data from service data
	ServiceData Optional[string] `json:"service_data"`
	// optional, UUID from service data
	ServiceUuid          Optional[string]       `json:"service_uuid"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookAssetRawRssiEventServicePacket,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookAssetRawRssiEventServicePacket) String() string {
	return fmt.Sprintf(
		"WebhookAssetRawRssiEventServicePacket[ServiceData=%v, ServiceUuid=%v, AdditionalProperties=%v]",
		w.ServiceData, w.ServiceUuid, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookAssetRawRssiEventServicePacket.
// It customizes the JSON marshaling process for WebhookAssetRawRssiEventServicePacket objects.
func (w WebhookAssetRawRssiEventServicePacket) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"service_data", "service_uuid"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookAssetRawRssiEventServicePacket object to a map representation for JSON marshaling.
func (w WebhookAssetRawRssiEventServicePacket) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.ServiceData.IsValueSet() {
		if w.ServiceData.Value() != nil {
			structMap["service_data"] = w.ServiceData.Value()
		} else {
			structMap["service_data"] = nil
		}
	}
	if w.ServiceUuid.IsValueSet() {
		if w.ServiceUuid.Value() != nil {
			structMap["service_uuid"] = w.ServiceUuid.Value()
		} else {
			structMap["service_uuid"] = nil
		}
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookAssetRawRssiEventServicePacket.
// It customizes the JSON unmarshaling process for WebhookAssetRawRssiEventServicePacket objects.
func (w *WebhookAssetRawRssiEventServicePacket) UnmarshalJSON(input []byte) error {
	var temp tempWebhookAssetRawRssiEventServicePacket
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "service_data", "service_uuid")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.ServiceData = temp.ServiceData
	w.ServiceUuid = temp.ServiceUuid
	return nil
}

// tempWebhookAssetRawRssiEventServicePacket is a temporary struct used for validating the fields of WebhookAssetRawRssiEventServicePacket.
type tempWebhookAssetRawRssiEventServicePacket struct {
	ServiceData Optional[string] `json:"service_data"`
	ServiceUuid Optional[string] `json:"service_uuid"`
}
