// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// LicenseAmendment represents a LicenseAmendment struct.
type LicenseAmendment struct {
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	EndTime     *int     `json:"end_time,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime   *float64 `json:"modified_time,omitempty"`
	Quantity       *int     `json:"quantity,omitempty"`
	StartTime      *int     `json:"start_time,omitempty"`
	SubscriptionId *string  `json:"subscription_id,omitempty"`
	// Type of license. The list of supported license type can be retrieve with the [List License Type]($e/Constants%20Definitions/listLicenseTypes) API request.
	Type                 *string                `json:"type,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for LicenseAmendment,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (l LicenseAmendment) String() string {
	return fmt.Sprintf(
		"LicenseAmendment[CreatedTime=%v, EndTime=%v, Id=%v, ModifiedTime=%v, Quantity=%v, StartTime=%v, SubscriptionId=%v, Type=%v, AdditionalProperties=%v]",
		l.CreatedTime, l.EndTime, l.Id, l.ModifiedTime, l.Quantity, l.StartTime, l.SubscriptionId, l.Type, l.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for LicenseAmendment.
// It customizes the JSON marshaling process for LicenseAmendment objects.
func (l LicenseAmendment) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(l.AdditionalProperties,
		"created_time", "end_time", "id", "modified_time", "quantity", "start_time", "subscription_id", "type"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(l.toMap())
}

// toMap converts the LicenseAmendment object to a map representation for JSON marshaling.
func (l LicenseAmendment) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, l.AdditionalProperties)
	if l.CreatedTime != nil {
		structMap["created_time"] = l.CreatedTime
	}
	if l.EndTime != nil {
		structMap["end_time"] = l.EndTime
	}
	if l.Id != nil {
		structMap["id"] = l.Id
	}
	if l.ModifiedTime != nil {
		structMap["modified_time"] = l.ModifiedTime
	}
	if l.Quantity != nil {
		structMap["quantity"] = l.Quantity
	}
	if l.StartTime != nil {
		structMap["start_time"] = l.StartTime
	}
	if l.SubscriptionId != nil {
		structMap["subscription_id"] = l.SubscriptionId
	}
	if l.Type != nil {
		structMap["type"] = l.Type
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for LicenseAmendment.
// It customizes the JSON unmarshaling process for LicenseAmendment objects.
func (l *LicenseAmendment) UnmarshalJSON(input []byte) error {
	var temp tempLicenseAmendment
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "end_time", "id", "modified_time", "quantity", "start_time", "subscription_id", "type")
	if err != nil {
		return err
	}
	l.AdditionalProperties = additionalProperties

	l.CreatedTime = temp.CreatedTime
	l.EndTime = temp.EndTime
	l.Id = temp.Id
	l.ModifiedTime = temp.ModifiedTime
	l.Quantity = temp.Quantity
	l.StartTime = temp.StartTime
	l.SubscriptionId = temp.SubscriptionId
	l.Type = temp.Type
	return nil
}

// tempLicenseAmendment is a temporary struct used for validating the fields of LicenseAmendment.
type tempLicenseAmendment struct {
	CreatedTime    *float64   `json:"created_time,omitempty"`
	EndTime        *int       `json:"end_time,omitempty"`
	Id             *uuid.UUID `json:"id,omitempty"`
	ModifiedTime   *float64   `json:"modified_time,omitempty"`
	Quantity       *int       `json:"quantity,omitempty"`
	StartTime      *int       `json:"start_time,omitempty"`
	SubscriptionId *string    `json:"subscription_id,omitempty"`
	Type           *string    `json:"type,omitempty"`
}
