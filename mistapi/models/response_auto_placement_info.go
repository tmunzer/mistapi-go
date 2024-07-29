package models

import (
    "encoding/json"
)

// ResponseAutoPlacementInfo represents a ResponseAutoPlacementInfo struct.
type ResponseAutoPlacementInfo struct {
    // time when autoplacement completed or was manually stopped
    EndTime              *float64                     `json:"end_time,omitempty"`
    // (Only when inprogress) estimate of the time to completion
    EstTimeLeft          *float64                     `json:"est_time_left,omitempty"`
    // time when autoplacement process was last queued for this map
    StartTime            *float64                     `json:"start_time,omitempty"`
    // the status of autoplacement for a given map. enum: `done`, `error`, `inprogress`, `pending`
    Status               *AutoPlacementInfoStatusEnum `json:"status,omitempty"`
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoPlacementInfo.
// It customizes the JSON marshaling process for ResponseAutoPlacementInfo objects.
func (r ResponseAutoPlacementInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoPlacementInfo object to a map representation for JSON marshaling.
func (r ResponseAutoPlacementInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.EndTime != nil {
        structMap["end_time"] = r.EndTime
    }
    if r.EstTimeLeft != nil {
        structMap["est_time_left"] = r.EstTimeLeft
    }
    if r.StartTime != nil {
        structMap["start_time"] = r.StartTime
    }
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoPlacementInfo.
// It customizes the JSON unmarshaling process for ResponseAutoPlacementInfo objects.
func (r *ResponseAutoPlacementInfo) UnmarshalJSON(input []byte) error {
    var temp responseAutoPlacementInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "end_time", "est_time_left", "start_time", "status")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.EndTime = temp.EndTime
    r.EstTimeLeft = temp.EstTimeLeft
    r.StartTime = temp.StartTime
    r.Status = temp.Status
    return nil
}

// responseAutoPlacementInfo is a temporary struct used for validating the fields of ResponseAutoPlacementInfo.
type responseAutoPlacementInfo  struct {
    EndTime     *float64                     `json:"end_time,omitempty"`
    EstTimeLeft *float64                     `json:"est_time_left,omitempty"`
    StartTime   *float64                     `json:"start_time,omitempty"`
    Status      *AutoPlacementInfoStatusEnum `json:"status,omitempty"`
}
