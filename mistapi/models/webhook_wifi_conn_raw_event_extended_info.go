// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// WebhookWifiConnRawEventExtendedInfo represents a WebhookWifiConnRawEventExtendedInfo struct.
type WebhookWifiConnRawEventExtendedInfo struct {
	FrameCtrl            *int                   `json:"frame_ctrl,omitempty"`
	Payload              *string                `json:"payload,omitempty"`
	SequenceCtrl         *int                   `json:"sequence_ctrl,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookWifiConnRawEventExtendedInfo,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookWifiConnRawEventExtendedInfo) String() string {
	return fmt.Sprintf(
		"WebhookWifiConnRawEventExtendedInfo[FrameCtrl=%v, Payload=%v, SequenceCtrl=%v, AdditionalProperties=%v]",
		w.FrameCtrl, w.Payload, w.SequenceCtrl, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookWifiConnRawEventExtendedInfo.
// It customizes the JSON marshaling process for WebhookWifiConnRawEventExtendedInfo objects.
func (w WebhookWifiConnRawEventExtendedInfo) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"frame_ctrl", "payload", "sequence_ctrl"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookWifiConnRawEventExtendedInfo object to a map representation for JSON marshaling.
func (w WebhookWifiConnRawEventExtendedInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	if w.FrameCtrl != nil {
		structMap["frame_ctrl"] = w.FrameCtrl
	}
	if w.Payload != nil {
		structMap["payload"] = w.Payload
	}
	if w.SequenceCtrl != nil {
		structMap["sequence_ctrl"] = w.SequenceCtrl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookWifiConnRawEventExtendedInfo.
// It customizes the JSON unmarshaling process for WebhookWifiConnRawEventExtendedInfo objects.
func (w *WebhookWifiConnRawEventExtendedInfo) UnmarshalJSON(input []byte) error {
	var temp tempWebhookWifiConnRawEventExtendedInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "frame_ctrl", "payload", "sequence_ctrl")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.FrameCtrl = temp.FrameCtrl
	w.Payload = temp.Payload
	w.SequenceCtrl = temp.SequenceCtrl
	return nil
}

// tempWebhookWifiConnRawEventExtendedInfo is a temporary struct used for validating the fields of WebhookWifiConnRawEventExtendedInfo.
type tempWebhookWifiConnRawEventExtendedInfo struct {
	FrameCtrl    *int    `json:"frame_ctrl,omitempty"`
	Payload      *string `json:"payload,omitempty"`
	SequenceCtrl *int    `json:"sequence_ctrl,omitempty"`
}
