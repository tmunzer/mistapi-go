package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// UserMac represents a UserMac struct.
type UserMac struct {
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Labels               []string       `json:"labels,omitempty"`
    // only non-local-admin MAC is accepted
    Mac                  *string        `json:"mac,omitempty"`
    Notes                *string        `json:"notes,omitempty"`
    Vlan                 *int           `json:"vlan,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UserMac.
// It customizes the JSON marshaling process for UserMac objects.
func (u UserMac) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UserMac object to a map representation for JSON marshaling.
func (u UserMac) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Id != nil {
        structMap["id"] = u.Id
    }
    if u.Labels != nil {
        structMap["labels"] = u.Labels
    }
    if u.Mac != nil {
        structMap["mac"] = u.Mac
    }
    if u.Notes != nil {
        structMap["notes"] = u.Notes
    }
    if u.Vlan != nil {
        structMap["vlan"] = u.Vlan
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserMac.
// It customizes the JSON unmarshaling process for UserMac objects.
func (u *UserMac) UnmarshalJSON(input []byte) error {
    var temp userMac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "labels", "mac", "notes", "vlan")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Id = temp.Id
    u.Labels = temp.Labels
    u.Mac = temp.Mac
    u.Notes = temp.Notes
    u.Vlan = temp.Vlan
    return nil
}

// userMac is a temporary struct used for validating the fields of UserMac.
type userMac  struct {
    Id     *uuid.UUID `json:"id,omitempty"`
    Labels []string   `json:"labels,omitempty"`
    Mac    *string    `json:"mac,omitempty"`
    Notes  *string    `json:"notes,omitempty"`
    Vlan   *int       `json:"vlan,omitempty"`
}
