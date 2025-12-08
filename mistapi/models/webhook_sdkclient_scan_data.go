// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// WebhookSdkclientScanData represents a WebhookSdkclientScanData struct.
// Sample of the `sdkclient-scan-data` webhook payload.
type WebhookSdkclientScanData struct {
	Events []WebhookSdkclientScanDataEvent `json:"events"`
	// enum: `sdkclient-scan-data`
	Topic                string                 `json:"topic"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookSdkclientScanData,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookSdkclientScanData) String() string {
	return fmt.Sprintf(
		"WebhookSdkclientScanData[Events=%v, Topic=%v, AdditionalProperties=%v]",
		w.Events, w.Topic, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookSdkclientScanData.
// It customizes the JSON marshaling process for WebhookSdkclientScanData objects.
func (w WebhookSdkclientScanData) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(w.AdditionalProperties,
		"events", "topic"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(w.toMap())
}

// toMap converts the WebhookSdkclientScanData object to a map representation for JSON marshaling.
func (w WebhookSdkclientScanData) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, w.AdditionalProperties)
	structMap["events"] = w.Events
	structMap["topic"] = w.Topic
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSdkclientScanData.
// It customizes the JSON unmarshaling process for WebhookSdkclientScanData objects.
func (w *WebhookSdkclientScanData) UnmarshalJSON(input []byte) error {
	var temp tempWebhookSdkclientScanData
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "events", "topic")
	if err != nil {
		return err
	}
	w.AdditionalProperties = additionalProperties

	w.Events = *temp.Events
	w.Topic = *temp.Topic
	return nil
}

// tempWebhookSdkclientScanData is a temporary struct used for validating the fields of WebhookSdkclientScanData.
type tempWebhookSdkclientScanData struct {
	Events *[]WebhookSdkclientScanDataEvent `json:"events"`
	Topic  *string                          `json:"topic"`
}

func (w *tempWebhookSdkclientScanData) validate() error {
	var errs []string
	if w.Events == nil {
		errs = append(errs, "required field `events` is missing for type `webhook_sdkclient_scan_data`")
	}
	if w.Topic == nil {
		errs = append(errs, "required field `topic` is missing for type `webhook_sdkclient_scan_data`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
