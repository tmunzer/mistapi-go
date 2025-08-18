// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ResponseAutoZone represents a ResponseAutoZone struct.
type ResponseAutoZone struct {
	// The status for the auto zones service for a given map. enum:
	// * not_started: The auto zones service has not been run on this map or the results were cleared by the user
	// * in_progress: The auto zones service is currently in progress
	// * awaiting_review: The auto zones service has completed and suggested location zones to be added to the map
	// * error: There was an error with the auto zones service
	Status               *ResponseAutoZoneStatusEnum `json:"status,omitempty"`
	Zones                []ResponseAutoZoneZone      `json:"zones,omitempty"`
	AdditionalProperties map[string]interface{}      `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoZone) String() string {
	return fmt.Sprintf(
		"ResponseAutoZone[Status=%v, Zones=%v, AdditionalProperties=%v]",
		r.Status, r.Zones, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoZone.
// It customizes the JSON marshaling process for ResponseAutoZone objects.
func (r ResponseAutoZone) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"status", "zones"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoZone object to a map representation for JSON marshaling.
func (r ResponseAutoZone) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	if r.Status != nil {
		structMap["status"] = r.Status
	}
	if r.Zones != nil {
		structMap["zones"] = r.Zones
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoZone.
// It customizes the JSON unmarshaling process for ResponseAutoZone objects.
func (r *ResponseAutoZone) UnmarshalJSON(input []byte) error {
	var temp tempResponseAutoZone
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "status", "zones")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Status = temp.Status
	r.Zones = temp.Zones
	return nil
}

// tempResponseAutoZone is a temporary struct used for validating the fields of ResponseAutoZone.
type tempResponseAutoZone struct {
	Status *ResponseAutoZoneStatusEnum `json:"status,omitempty"`
	Zones  []ResponseAutoZoneZone      `json:"zones,omitempty"`
}
