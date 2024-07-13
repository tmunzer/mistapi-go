package models

import (
    "encoding/json"
)

// ResponseOrgSuppressAlarm represents a ResponseOrgSuppressAlarm struct.
type ResponseOrgSuppressAlarm struct {
    Results              []ResponseOrgSuppressAlarmItem `json:"results,omitempty"`
    AdditionalProperties map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgSuppressAlarm.
// It customizes the JSON marshaling process for ResponseOrgSuppressAlarm objects.
func (r ResponseOrgSuppressAlarm) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgSuppressAlarm object to a map representation for JSON marshaling.
func (r ResponseOrgSuppressAlarm) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgSuppressAlarm.
// It customizes the JSON unmarshaling process for ResponseOrgSuppressAlarm objects.
func (r *ResponseOrgSuppressAlarm) UnmarshalJSON(input []byte) error {
    var temp responseOrgSuppressAlarm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "results")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Results = temp.Results
    return nil
}

// responseOrgSuppressAlarm is a temporary struct used for validating the fields of ResponseOrgSuppressAlarm.
type responseOrgSuppressAlarm  struct {
    Results []ResponseOrgSuppressAlarmItem `json:"results,omitempty"`
}
