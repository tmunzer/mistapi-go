package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// Secpolicy represents a Secpolicy struct.
// Security Policy is designed to audit / catch discripancies between “what’s intended to be running” versus “what’s actually running” in a network. Many big organizations have separated Security and IT team (for good reasons). Each site can be assigned a security policy. Whenever an AP is provisioned, the configuration will be checked against the security policy. Any violations will be flagged in Device Config History where you can search for the when and where the violation occurs.
type Secpolicy struct {
    CreatedTime          *float64       `json:"created_time,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    ModifiedTime         *float64       `json:"modified_time,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Wlans                []Wlan         `json:"wlans,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Secpolicy.
// It customizes the JSON marshaling process for Secpolicy objects.
func (s Secpolicy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Secpolicy object to a map representation for JSON marshaling.
func (s Secpolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    if s.Id != nil {
        structMap["id"] = s.Id
    }
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Wlans != nil {
        structMap["wlans"] = s.Wlans
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Secpolicy.
// It customizes the JSON unmarshaling process for Secpolicy objects.
func (s *Secpolicy) UnmarshalJSON(input []byte) error {
    var temp tempSecpolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "id", "modified_time", "name", "org_id", "site_id", "wlans")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.CreatedTime = temp.CreatedTime
    s.Id = temp.Id
    s.ModifiedTime = temp.ModifiedTime
    s.Name = temp.Name
    s.OrgId = temp.OrgId
    s.SiteId = temp.SiteId
    s.Wlans = temp.Wlans
    return nil
}

// tempSecpolicy is a temporary struct used for validating the fields of Secpolicy.
type tempSecpolicy  struct {
    CreatedTime  *float64   `json:"created_time,omitempty"`
    Id           *uuid.UUID `json:"id,omitempty"`
    ModifiedTime *float64   `json:"modified_time,omitempty"`
    Name         *string    `json:"name,omitempty"`
    OrgId        *uuid.UUID `json:"org_id,omitempty"`
    SiteId       *uuid.UUID `json:"site_id,omitempty"`
    Wlans        []Wlan     `json:"wlans,omitempty"`
}
