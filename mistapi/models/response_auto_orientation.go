package models

import (
    "encoding/json"
)

// ResponseAutoOrientation represents a ResponseAutoOrientation struct.
type ResponseAutoOrientation struct {
    // The state of auto orient for a given map derived from an Enum. enum: `Enqueued`, `Not Started`, `Oriented`
    State                *AutoOrientationStateEnum `json:"state,omitempty"`
    // Time when auto orient process was last queued for this map
    TimeQueued           *float64                  `json:"time_queued,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoOrientation.
// It customizes the JSON marshaling process for ResponseAutoOrientation objects.
func (r ResponseAutoOrientation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoOrientation object to a map representation for JSON marshaling.
func (r ResponseAutoOrientation) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.State != nil {
        structMap["state"] = r.State
    }
    if r.TimeQueued != nil {
        structMap["time_queued"] = r.TimeQueued
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoOrientation.
// It customizes the JSON unmarshaling process for ResponseAutoOrientation objects.
func (r *ResponseAutoOrientation) UnmarshalJSON(input []byte) error {
    var temp responseAutoOrientation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "state", "time_queued")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.State = temp.State
    r.TimeQueued = temp.TimeQueued
    return nil
}

// responseAutoOrientation is a temporary struct used for validating the fields of ResponseAutoOrientation.
type responseAutoOrientation  struct {
    State      *AutoOrientationStateEnum `json:"state,omitempty"`
    TimeQueued *float64                  `json:"time_queued,omitempty"`
}
