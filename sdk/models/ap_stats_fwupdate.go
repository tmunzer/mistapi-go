package models

import (
    "encoding/json"
)

// ApStatsFwupdate represents a ApStatsFwupdate struct.
type ApStatsFwupdate struct {
    Progress             *int           `json:"progress,omitempty"`
    Status               *string        `json:"status,omitempty"`
    StatusId             *int           `json:"status_id,omitempty"`
    Timestamp            *int           `json:"timestamp,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ApStatsFwupdate.
// It customizes the JSON marshaling process for ApStatsFwupdate objects.
func (a ApStatsFwupdate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the ApStatsFwupdate object to a map representation for JSON marshaling.
func (a ApStatsFwupdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Progress != nil {
        structMap["progress"] = a.Progress
    }
    if a.Status != nil {
        structMap["status"] = a.Status
    }
    if a.StatusId != nil {
        structMap["status_id"] = a.StatusId
    }
    if a.Timestamp != nil {
        structMap["timestamp"] = a.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ApStatsFwupdate.
// It customizes the JSON unmarshaling process for ApStatsFwupdate objects.
func (a *ApStatsFwupdate) UnmarshalJSON(input []byte) error {
    var temp apStatsFwupdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "progress", "status", "status_id", "timestamp")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Progress = temp.Progress
    a.Status = temp.Status
    a.StatusId = temp.StatusId
    a.Timestamp = temp.Timestamp
    return nil
}

// apStatsFwupdate is a temporary struct used for validating the fields of ApStatsFwupdate.
type apStatsFwupdate  struct {
    Progress  *int    `json:"progress,omitempty"`
    Status    *string `json:"status,omitempty"`
    StatusId  *int    `json:"status_id,omitempty"`
    Timestamp *int    `json:"timestamp,omitempty"`
}
