package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// EventsClientWan represents a EventsClientWan struct.
type EventsClientWan struct {
    When                 *string        `json:"When,omitempty"`
    EvType               *string        `json:"ev_type,omitempty"`
    Metadata             *interface{}   `json:"metadata,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    RandomMac            *bool          `json:"random_mac,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    Text                 *string        `json:"text,omitempty"`
    Wcid                 *uuid.UUID     `json:"wcid,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventsClientWan.
// It customizes the JSON marshaling process for EventsClientWan objects.
func (e EventsClientWan) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the EventsClientWan object to a map representation for JSON marshaling.
func (e EventsClientWan) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.When != nil {
        structMap["When"] = e.When
    }
    if e.EvType != nil {
        structMap["ev_type"] = e.EvType
    }
    if e.Metadata != nil {
        structMap["metadata"] = e.Metadata
    }
    if e.OrgId != nil {
        structMap["org_id"] = e.OrgId
    }
    if e.RandomMac != nil {
        structMap["random_mac"] = e.RandomMac
    }
    if e.SiteId != nil {
        structMap["site_id"] = e.SiteId
    }
    if e.Text != nil {
        structMap["text"] = e.Text
    }
    if e.Wcid != nil {
        structMap["wcid"] = e.Wcid
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventsClientWan.
// It customizes the JSON unmarshaling process for EventsClientWan objects.
func (e *EventsClientWan) UnmarshalJSON(input []byte) error {
    var temp eventsClientWan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "When", "ev_type", "metadata", "org_id", "random_mac", "site_id", "text", "wcid")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.When = temp.When
    e.EvType = temp.EvType
    e.Metadata = temp.Metadata
    e.OrgId = temp.OrgId
    e.RandomMac = temp.RandomMac
    e.SiteId = temp.SiteId
    e.Text = temp.Text
    e.Wcid = temp.Wcid
    return nil
}

// eventsClientWan is a temporary struct used for validating the fields of EventsClientWan.
type eventsClientWan  struct {
    When      *string      `json:"When,omitempty"`
    EvType    *string      `json:"ev_type,omitempty"`
    Metadata  *interface{} `json:"metadata,omitempty"`
    OrgId     *uuid.UUID   `json:"org_id,omitempty"`
    RandomMac *bool        `json:"random_mac,omitempty"`
    SiteId    *uuid.UUID   `json:"site_id,omitempty"`
    Text      *string      `json:"text,omitempty"`
    Wcid      *uuid.UUID   `json:"wcid,omitempty"`
}