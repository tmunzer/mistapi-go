// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAsyncLicenseDetail represents a ResponseAsyncLicenseDetail struct.
// detail claim status per device
type ResponseAsyncLicenseDetail struct {
	// Device MAC Address
	Mac    *string `json:"mac,omitempty"`
	Status *string `json:"status,omitempty"`
	// Epoch (seconds)
	Timestamp            *float64               `json:"timestamp,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAsyncLicenseDetail,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAsyncLicenseDetail) String() string {
	return fmt.Sprintf(
		"ResponseAsyncLicenseDetail[Mac=%v, Status=%v, Timestamp=%v, AdditionalProperties=%v]",
		r.Mac, r.Status, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAsyncLicenseDetail.
// It customizes the JSON marshaling process for ResponseAsyncLicenseDetail objects.
func (r ResponseAsyncLicenseDetail) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"mac", "status", "timestamp"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAsyncLicenseDetail object to a map representation for JSON marshaling.
func (r ResponseAsyncLicenseDetail) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Mac != nil {
		structMap["mac"] = r.Mac
	}
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	if r.Timestamp != nil {
		structMap["timestamp"] = r.Timestamp
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAsyncLicenseDetail.
// It customizes the JSON unmarshaling process for ResponseAsyncLicenseDetail objects.
func (r *ResponseAsyncLicenseDetail) UnmarshalJSON(input []byte) error {
	var temp tempResponseAsyncLicenseDetail
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "status", "timestamp")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Mac = temp.Mac
	r.Status = temp.Status
	r.Timestamp = temp.Timestamp
	return nil
}

// tempResponseAsyncLicenseDetail is a temporary struct used for validating the fields of ResponseAsyncLicenseDetail.
type tempResponseAsyncLicenseDetail struct {
	Mac       *string  `json:"mac,omitempty"`
	Status    *string  `json:"status,omitempty"`
	Timestamp *float64 `json:"timestamp,omitempty"`
}
