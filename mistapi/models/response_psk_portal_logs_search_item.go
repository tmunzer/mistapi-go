package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ResponsePskPortalLogsSearchItem represents a ResponsePskPortalLogsSearchItem struct.
type ResponsePskPortalLogsSearchItem struct {
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Message              *string                `json:"message,omitempty"`
    NameId               *string                `json:"name_id,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    PskId                *uuid.UUID             `json:"psk_id,omitempty"`
    PskName              *string                `json:"psk_name,omitempty"`
    PskportalId          *uuid.UUID             `json:"pskportal_id,omitempty"`
    Timestamp            *float64               `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponsePskPortalLogsSearchItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponsePskPortalLogsSearchItem) String() string {
    return fmt.Sprintf(
    	"ResponsePskPortalLogsSearchItem[Id=%v, Message=%v, NameId=%v, OrgId=%v, PskId=%v, PskName=%v, PskportalId=%v, Timestamp=%v, AdditionalProperties=%v]",
    	r.Id, r.Message, r.NameId, r.OrgId, r.PskId, r.PskName, r.PskportalId, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponsePskPortalLogsSearchItem.
// It customizes the JSON marshaling process for ResponsePskPortalLogsSearchItem objects.
func (r ResponsePskPortalLogsSearchItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "id", "message", "name_id", "org_id", "psk_id", "psk_name", "pskportal_id", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponsePskPortalLogsSearchItem object to a map representation for JSON marshaling.
func (r ResponsePskPortalLogsSearchItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.Message != nil {
        structMap["message"] = r.Message
    }
    if r.NameId != nil {
        structMap["name_id"] = r.NameId
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.PskId != nil {
        structMap["psk_id"] = r.PskId
    }
    if r.PskName != nil {
        structMap["psk_name"] = r.PskName
    }
    if r.PskportalId != nil {
        structMap["pskportal_id"] = r.PskportalId
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponsePskPortalLogsSearchItem.
// It customizes the JSON unmarshaling process for ResponsePskPortalLogsSearchItem objects.
func (r *ResponsePskPortalLogsSearchItem) UnmarshalJSON(input []byte) error {
    var temp tempResponsePskPortalLogsSearchItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "message", "name_id", "org_id", "psk_id", "psk_name", "pskportal_id", "timestamp")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Id = temp.Id
    r.Message = temp.Message
    r.NameId = temp.NameId
    r.OrgId = temp.OrgId
    r.PskId = temp.PskId
    r.PskName = temp.PskName
    r.PskportalId = temp.PskportalId
    r.Timestamp = temp.Timestamp
    return nil
}

// tempResponsePskPortalLogsSearchItem is a temporary struct used for validating the fields of ResponsePskPortalLogsSearchItem.
type tempResponsePskPortalLogsSearchItem  struct {
    Id          *uuid.UUID `json:"id,omitempty"`
    Message     *string    `json:"message,omitempty"`
    NameId      *string    `json:"name_id,omitempty"`
    OrgId       *uuid.UUID `json:"org_id,omitempty"`
    PskId       *uuid.UUID `json:"psk_id,omitempty"`
    PskName     *string    `json:"psk_name,omitempty"`
    PskportalId *uuid.UUID `json:"pskportal_id,omitempty"`
    Timestamp   *float64   `json:"timestamp,omitempty"`
}
