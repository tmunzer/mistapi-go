package models

import (
    "encoding/json"
)

// ResponseDeviceBiosUpgrade represents a ResponseDeviceBiosUpgrade struct.
type ResponseDeviceBiosUpgrade struct {
    Status               *string        `json:"status,omitempty"`
    Timestamp            *int           `json:"timestamp,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceBiosUpgrade.
// It customizes the JSON marshaling process for ResponseDeviceBiosUpgrade objects.
func (r ResponseDeviceBiosUpgrade) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceBiosUpgrade object to a map representation for JSON marshaling.
func (r ResponseDeviceBiosUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceBiosUpgrade.
// It customizes the JSON unmarshaling process for ResponseDeviceBiosUpgrade objects.
func (r *ResponseDeviceBiosUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceBiosUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "status", "timestamp")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Status = temp.Status
    r.Timestamp = temp.Timestamp
    return nil
}

// tempResponseDeviceBiosUpgrade is a temporary struct used for validating the fields of ResponseDeviceBiosUpgrade.
type tempResponseDeviceBiosUpgrade  struct {
    Status    *string `json:"status,omitempty"`
    Timestamp *int    `json:"timestamp,omitempty"`
}
