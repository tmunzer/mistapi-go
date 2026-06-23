// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// MarvisClient represents a MarvisClient struct.
// Marvis Client configuration profile
type MarvisClient struct {
	// Whether this Marvis Client profile is disabled
	Disabled *bool `json:"disabled,omitempty"`
	// In MDM, add `--enrollment_url <enrollment_url>` to the install command
	EnrollmentUrl *string `json:"enrollment_url,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id *uuid.UUID `json:"id,omitempty"`
	// Location collection settings for Marvis Client
	Location *MarvisClientLocation `json:"location,omitempty"`
	// Display name for the Marvis Client profile
	Name *string `json:"name,omitempty"`
	// Synthetic test settings for Marvis Client
	SyntheticTest *MarvisClientSyntheticTest `json:"synthetic_test,omitempty"`
	// Note: some stats are not collected when it's not connected to Mist infrastructure
	Telemetry            *MarvisClientTelemetry `json:"telemetry,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClient) String() string {
	return fmt.Sprintf(
		"MarvisClient[Disabled=%v, EnrollmentUrl=%v, Id=%v, Location=%v, Name=%v, SyntheticTest=%v, Telemetry=%v, AdditionalProperties=%v]",
		m.Disabled, m.EnrollmentUrl, m.Id, m.Location, m.Name, m.SyntheticTest, m.Telemetry, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClient.
// It customizes the JSON marshaling process for MarvisClient objects.
func (m MarvisClient) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(m.AdditionalProperties,
		"disabled", "enrollment_url", "id", "location", "name", "synthetic_test", "telemetry"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(m.toMap())
}

// toMap converts the MarvisClient object to a map representation for JSON marshaling.
func (m MarvisClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, m.AdditionalProperties)
	if m.Disabled != nil {
		structMap["disabled"] = m.Disabled
	}
	if m.EnrollmentUrl != nil {
		structMap["enrollment_url"] = m.EnrollmentUrl
	}
	if m.Id != nil {
		structMap["id"] = m.Id
	}
	if m.Location != nil {
		structMap["location"] = m.Location.toMap()
	}
	if m.Name != nil {
		structMap["name"] = m.Name
	}
	if m.SyntheticTest != nil {
		structMap["synthetic_test"] = m.SyntheticTest.toMap()
	}
	if m.Telemetry != nil {
		structMap["telemetry"] = m.Telemetry.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClient.
// It customizes the JSON unmarshaling process for MarvisClient objects.
func (m *MarvisClient) UnmarshalJSON(input []byte) error {
	var temp tempMarvisClient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "enrollment_url", "id", "location", "name", "synthetic_test", "telemetry")
	if err != nil {
		return err
	}
	m.AdditionalProperties = additionalProperties

	m.Disabled = temp.Disabled
	m.EnrollmentUrl = temp.EnrollmentUrl
	m.Id = temp.Id
	m.Location = temp.Location
	m.Name = temp.Name
	m.SyntheticTest = temp.SyntheticTest
	m.Telemetry = temp.Telemetry
	return nil
}

// tempMarvisClient is a temporary struct used for validating the fields of MarvisClient.
type tempMarvisClient struct {
	Disabled      *bool                      `json:"disabled,omitempty"`
	EnrollmentUrl *string                    `json:"enrollment_url,omitempty"`
	Id            *uuid.UUID                 `json:"id,omitempty"`
	Location      *MarvisClientLocation      `json:"location,omitempty"`
	Name          *string                    `json:"name,omitempty"`
	SyntheticTest *MarvisClientSyntheticTest `json:"synthetic_test,omitempty"`
	Telemetry     *MarvisClientTelemetry     `json:"telemetry,omitempty"`
}
