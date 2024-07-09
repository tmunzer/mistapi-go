package models

import (
    "encoding/json"
)

// NacPortalSsoRoleMatching represents a NacPortalSsoRoleMatching struct.
type NacPortalSsoRoleMatching struct {
    Assigned             *string        `json:"assigned,omitempty"`
    Match                *string        `json:"match,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for NacPortalSsoRoleMatching.
// It customizes the JSON marshaling process for NacPortalSsoRoleMatching objects.
func (n NacPortalSsoRoleMatching) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(n.toMap())
}

// toMap converts the NacPortalSsoRoleMatching object to a map representation for JSON marshaling.
func (n NacPortalSsoRoleMatching) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, n.AdditionalProperties)
    if n.Assigned != nil {
        structMap["assigned"] = n.Assigned
    }
    if n.Match != nil {
        structMap["match"] = n.Match
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for NacPortalSsoRoleMatching.
// It customizes the JSON unmarshaling process for NacPortalSsoRoleMatching objects.
func (n *NacPortalSsoRoleMatching) UnmarshalJSON(input []byte) error {
    var temp nacPortalSsoRoleMatching
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "assigned", "match")
    if err != nil {
    	return err
    }
    
    n.AdditionalProperties = additionalProperties
    n.Assigned = temp.Assigned
    n.Match = temp.Match
    return nil
}

// nacPortalSsoRoleMatching is a temporary struct used for validating the fields of NacPortalSsoRoleMatching.
type nacPortalSsoRoleMatching  struct {
    Assigned *string `json:"assigned,omitempty"`
    Match    *string `json:"match,omitempty"`
}
