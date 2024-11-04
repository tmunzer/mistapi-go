package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ResponseWiredCoa represents a ResponseWiredCoa struct.
type ResponseWiredCoa struct {
    DeviceMac            *string        `json:"device_mac,omitempty"`
    PortId               *string        `json:"port_id,omitempty"`
    Session              *uuid.UUID     `json:"session,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseWiredCoa.
// It customizes the JSON marshaling process for ResponseWiredCoa objects.
func (r ResponseWiredCoa) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseWiredCoa object to a map representation for JSON marshaling.
func (r ResponseWiredCoa) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.DeviceMac != nil {
        structMap["device_mac"] = r.DeviceMac
    }
    if r.PortId != nil {
        structMap["port_id"] = r.PortId
    }
    if r.Session != nil {
        structMap["session"] = r.Session
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseWiredCoa.
// It customizes the JSON unmarshaling process for ResponseWiredCoa objects.
func (r *ResponseWiredCoa) UnmarshalJSON(input []byte) error {
    var temp tempResponseWiredCoa
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "port_id", "session")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.DeviceMac = temp.DeviceMac
    r.PortId = temp.PortId
    r.Session = temp.Session
    return nil
}

// tempResponseWiredCoa is a temporary struct used for validating the fields of ResponseWiredCoa.
type tempResponseWiredCoa  struct {
    DeviceMac *string    `json:"device_mac,omitempty"`
    PortId    *string    `json:"port_id,omitempty"`
    Session   *uuid.UUID `json:"session,omitempty"`
}
