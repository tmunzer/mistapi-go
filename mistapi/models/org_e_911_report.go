// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// OrgE911Report represents a OrgE911Report struct.
// E911 AP BSSID report status for the organization
type OrgE911Report struct {
	// Human-readable description of the action taken
	Detail *string `json:"detail,omitempty"`
	// Unix timestamp of when the report file was last generated. Only present when `status` is `available`.
	LastGenerated *int `json:"last_generated,omitempty"`
	// Current status of E911 report generation. enum: `disabled`, `scheduled`, `available`
	Status *OrgE911ReportStatusEnum `json:"status,omitempty"`
	// Presigned URL to download the CSV file. Only present when `status` is `available`.
	Url                  *string                `json:"url,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgE911Report,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgE911Report) String() string {
	return fmt.Sprintf(
		"OrgE911Report[Detail=%v, LastGenerated=%v, Status=%v, Url=%v, AdditionalProperties=%v]",
		o.Detail, o.LastGenerated, o.Status, o.Url, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgE911Report.
// It customizes the JSON marshaling process for OrgE911Report objects.
func (o OrgE911Report) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(o.AdditionalProperties,
		"detail", "last_generated", "status", "url"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(o.toMap())
}

// toMap converts the OrgE911Report object to a map representation for JSON marshaling.
func (o OrgE911Report) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Detail != nil {
		structMap["detail"] = o.Detail
	}
	if o.LastGenerated != nil {
		structMap["last_generated"] = o.LastGenerated
	}
	if o.Status != nil {
		structMap["status"] = o.Status
	}
	if o.Url != nil {
		structMap["url"] = o.Url
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgE911Report.
// It customizes the JSON unmarshaling process for OrgE911Report objects.
func (o *OrgE911Report) UnmarshalJSON(input []byte) error {
	var temp tempOrgE911Report
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "detail", "last_generated", "status", "url")
	if err != nil {
		return err
	}
	o.AdditionalProperties = additionalProperties

	o.Detail = temp.Detail
	o.LastGenerated = temp.LastGenerated
	o.Status = temp.Status
	o.Url = temp.Url
	return nil
}

// tempOrgE911Report is a temporary struct used for validating the fields of OrgE911Report.
type tempOrgE911Report struct {
	Detail        *string                  `json:"detail,omitempty"`
	LastGenerated *int                     `json:"last_generated,omitempty"`
	Status        *OrgE911ReportStatusEnum `json:"status,omitempty"`
	Url           *string                  `json:"url,omitempty"`
}
