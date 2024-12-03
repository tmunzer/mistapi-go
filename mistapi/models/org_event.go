package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// OrgEvent represents a OrgEvent struct.
type OrgEvent struct {
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Text                 *string                `json:"text,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgEvent.
// It customizes the JSON marshaling process for OrgEvent objects.
func (o OrgEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "org_id", "text", "timestamp", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgEvent object to a map representation for JSON marshaling.
func (o OrgEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.Text != nil {
        structMap["text"] = o.Text
    }
    if o.Timestamp != nil {
        structMap["timestamp"] = o.Timestamp
    }
    if o.Type != nil {
        structMap["type"] = o.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgEvent.
// It customizes the JSON unmarshaling process for OrgEvent objects.
func (o *OrgEvent) UnmarshalJSON(input []byte) error {
    var temp tempOrgEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "org_id", "text", "timestamp", "type")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.OrgId = temp.OrgId
    o.Text = temp.Text
    o.Timestamp = temp.Timestamp
    o.Type = temp.Type
    return nil
}

// tempOrgEvent is a temporary struct used for validating the fields of OrgEvent.
type tempOrgEvent  struct {
    OrgId     *uuid.UUID `json:"org_id,omitempty"`
    Text      *string    `json:"text,omitempty"`
    Timestamp *float64   `json:"timestamp,omitempty"`
    Type      *string    `json:"type,omitempty"`
}
