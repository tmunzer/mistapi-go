package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OrgSystemEvent represents a OrgSystemEvent struct.
type OrgSystemEvent struct {
    ChangeCat            *string                `json:"change_cat,omitempty"`
    Metadata             *string                `json:"metadata,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Scope                *string                `json:"scope,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSystemEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSystemEvent) String() string {
    return fmt.Sprintf(
    	"OrgSystemEvent[ChangeCat=%v, Metadata=%v, OrgId=%v, Scope=%v, SiteId=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
    	o.ChangeCat, o.Metadata, o.OrgId, o.Scope, o.SiteId, o.Timestamp, o.Type, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSystemEvent.
// It customizes the JSON marshaling process for OrgSystemEvent objects.
func (o OrgSystemEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "change_cat", "metadata", "org_id", "scope", "site_id", "timestamp", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSystemEvent object to a map representation for JSON marshaling.
func (o OrgSystemEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.ChangeCat != nil {
        structMap["change_cat"] = o.ChangeCat
    }
    if o.Metadata != nil {
        structMap["metadata"] = o.Metadata
    }
    if o.OrgId != nil {
        structMap["org_id"] = o.OrgId
    }
    if o.Scope != nil {
        structMap["scope"] = o.Scope
    }
    if o.SiteId != nil {
        structMap["site_id"] = o.SiteId
    }
    if o.Timestamp != nil {
        structMap["timestamp"] = o.Timestamp
    }
    if o.Type != nil {
        structMap["type"] = o.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSystemEvent.
// It customizes the JSON unmarshaling process for OrgSystemEvent objects.
func (o *OrgSystemEvent) UnmarshalJSON(input []byte) error {
    var temp tempOrgSystemEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "change_cat", "metadata", "org_id", "scope", "site_id", "timestamp", "type")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.ChangeCat = temp.ChangeCat
    o.Metadata = temp.Metadata
    o.OrgId = temp.OrgId
    o.Scope = temp.Scope
    o.SiteId = temp.SiteId
    o.Timestamp = temp.Timestamp
    o.Type = temp.Type
    return nil
}

// tempOrgSystemEvent is a temporary struct used for validating the fields of OrgSystemEvent.
type tempOrgSystemEvent  struct {
    ChangeCat *string    `json:"change_cat,omitempty"`
    Metadata  *string    `json:"metadata,omitempty"`
    OrgId     *uuid.UUID `json:"org_id,omitempty"`
    Scope     *string    `json:"scope,omitempty"`
    SiteId    *uuid.UUID `json:"site_id,omitempty"`
    Timestamp *float64   `json:"timestamp,omitempty"`
    Type      *string    `json:"type,omitempty"`
}
