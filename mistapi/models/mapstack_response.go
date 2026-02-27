// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MapstackResponse represents a MapstackResponse struct.
// Map Stack response object
type MapstackResponse struct {
	// When the object has been created, in epoch
	CreatedTime *float64 `json:"created_time,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// The name of the map stack
	Name                 *string                `json:"name,omitempty"`
	OrgId                *uuid.UUID             `json:"org_id,omitempty"`
	SiteId               *uuid.UUID             `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MapstackResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MapstackResponse) String() string {
	return fmt.Sprintf(
		"MapstackResponse[CreatedTime=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, SiteId=%v, AdditionalProperties=%v]",
		m.CreatedTime, m.Id, m.ModifiedTime, m.Name, m.OrgId, m.SiteId, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MapstackResponse.
// It customizes the JSON marshaling process for MapstackResponse objects.
func (m MapstackResponse) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"created_time", "id", "modified_time", "name", "org_id", "site_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MapstackResponse object to a map representation for JSON marshaling.
func (m MapstackResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.CreatedTime != nil {
		structMap["created_time"] = m.CreatedTime
	}
	if m.Id != nil {
		structMap["id"] = m.Id
	}
	if m.ModifiedTime != nil {
		structMap["modified_time"] = m.ModifiedTime
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.OrgId != nil {
		structMap["org_id"] = m.OrgId
	}
	if m.SiteId != nil {
		structMap["site_id"] = m.SiteId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MapstackResponse.
// It customizes the JSON unmarshaling process for MapstackResponse objects.
func (m *MapstackResponse) UnmarshalJSON(input []byte) error {
	var temp tempMapstackResponse
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "id", "modified_time", "name", "org_id", "site_id")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.CreatedTime = temp.CreatedTime
	m.Id = temp.Id
	m.ModifiedTime = temp.ModifiedTime
	m.Name = temp.Name
	m.OrgId = temp.OrgId
	m.SiteId = temp.SiteId
	return nil
}

// tempMapstackResponse is a temporary struct used for validating the fields of MapstackResponse.
type tempMapstackResponse struct {
	CreatedTime  *float64   `json:"created_time,omitempty"`
	Id           *uuid.UUID `json:"id,omitempty"`
	ModifiedTime *float64   `json:"modified_time,omitempty"`
	Name         *string    `json:"name,omitempty"`
	OrgId        *uuid.UUID `json:"org_id,omitempty"`
	SiteId       *uuid.UUID `json:"site_id,omitempty"`
}
