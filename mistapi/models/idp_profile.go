// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// IdpProfile represents a IdpProfile struct.
type IdpProfile struct {
    // enum: `critical`, `standard`, `strict`
    BaseProfile          *IdpProfileBaseProfileEnum `json:"base_profile,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                   `json:"created_time,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID                 `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                   `json:"modified_time,omitempty"`
    Name                 *string                    `json:"name,omitempty"`
    OrgId                *uuid.UUID                 `json:"org_id,omitempty"`
    Overwrites           []IdpProfileOverwrite      `json:"overwrites,omitempty"`
    AdditionalProperties map[string]interface{}     `json:"_"`
}

// String implements the fmt.Stringer interface for IdpProfile,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IdpProfile) String() string {
    return fmt.Sprintf(
    	"IdpProfile[BaseProfile=%v, CreatedTime=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Overwrites=%v, AdditionalProperties=%v]",
    	i.BaseProfile, i.CreatedTime, i.Id, i.ModifiedTime, i.Name, i.OrgId, i.Overwrites, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IdpProfile.
// It customizes the JSON marshaling process for IdpProfile objects.
func (i IdpProfile) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "base_profile", "created_time", "id", "modified_time", "name", "org_id", "overwrites"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IdpProfile object to a map representation for JSON marshaling.
func (i IdpProfile) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.BaseProfile != nil {
        structMap["base_profile"] = i.BaseProfile
    }
    if i.CreatedTime != nil {
        structMap["created_time"] = i.CreatedTime
    }
    if i.Id != nil {
        structMap["id"] = i.Id
    }
    if i.ModifiedTime != nil {
        structMap["modified_time"] = i.ModifiedTime
    }
    if i.Name != nil {
        structMap["name"] = i.Name
    }
    if i.OrgId != nil {
        structMap["org_id"] = i.OrgId
    }
    if i.Overwrites != nil {
        structMap["overwrites"] = i.Overwrites
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IdpProfile.
// It customizes the JSON unmarshaling process for IdpProfile objects.
func (i *IdpProfile) UnmarshalJSON(input []byte) error {
    var temp tempIdpProfile
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "base_profile", "created_time", "id", "modified_time", "name", "org_id", "overwrites")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.BaseProfile = temp.BaseProfile
    i.CreatedTime = temp.CreatedTime
    i.Id = temp.Id
    i.ModifiedTime = temp.ModifiedTime
    i.Name = temp.Name
    i.OrgId = temp.OrgId
    i.Overwrites = temp.Overwrites
    return nil
}

// tempIdpProfile is a temporary struct used for validating the fields of IdpProfile.
type tempIdpProfile  struct {
    BaseProfile  *IdpProfileBaseProfileEnum `json:"base_profile,omitempty"`
    CreatedTime  *float64                   `json:"created_time,omitempty"`
    Id           *uuid.UUID                 `json:"id,omitempty"`
    ModifiedTime *float64                   `json:"modified_time,omitempty"`
    Name         *string                    `json:"name,omitempty"`
    OrgId        *uuid.UUID                 `json:"org_id,omitempty"`
    Overwrites   []IdpProfileOverwrite      `json:"overwrites,omitempty"`
}
