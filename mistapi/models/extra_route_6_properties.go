package models

import (
    "encoding/json"
)

// ExtraRoute6Properties represents a ExtraRoute6Properties struct.
type ExtraRoute6Properties struct {
    // this takes precedence
    Discard              *bool                                                   `json:"discard,omitempty"`
    Metric               Optional[int]                                           `json:"metric"`
    NextQualified        map[string]ExtraRoute6PropertiesNextQualifiedProperties `json:"next_qualified,omitempty"`
    NoResolve            *bool                                                   `json:"no_resolve,omitempty"`
    Preference           Optional[int]                                           `json:"preference"`
    // next-hop IP Address
    Via                  *string                                                 `json:"via,omitempty"`
    AdditionalProperties map[string]any                                          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ExtraRoute6Properties.
// It customizes the JSON marshaling process for ExtraRoute6Properties objects.
func (e ExtraRoute6Properties) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(e.toMap())
}

// toMap converts the ExtraRoute6Properties object to a map representation for JSON marshaling.
func (e ExtraRoute6Properties) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, e.AdditionalProperties)
    if e.Discard != nil {
        structMap["discard"] = e.Discard
    }
    if e.Metric.IsValueSet() {
        if e.Metric.Value() != nil {
            structMap["metric"] = e.Metric.Value()
        } else {
            structMap["metric"] = nil
        }
    }
    if e.NextQualified != nil {
        structMap["next_qualified"] = e.NextQualified
    }
    if e.NoResolve != nil {
        structMap["no_resolve"] = e.NoResolve
    }
    if e.Preference.IsValueSet() {
        if e.Preference.Value() != nil {
            structMap["preference"] = e.Preference.Value()
        } else {
            structMap["preference"] = nil
        }
    }
    if e.Via != nil {
        structMap["via"] = e.Via
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ExtraRoute6Properties.
// It customizes the JSON unmarshaling process for ExtraRoute6Properties objects.
func (e *ExtraRoute6Properties) UnmarshalJSON(input []byte) error {
    var temp extraRoute6Properties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "discard", "metric", "next_qualified", "no_resolve", "preference", "via")
    if err != nil {
    	return err
    }
    
    e.AdditionalProperties = additionalProperties
    e.Discard = temp.Discard
    e.Metric = temp.Metric
    e.NextQualified = temp.NextQualified
    e.NoResolve = temp.NoResolve
    e.Preference = temp.Preference
    e.Via = temp.Via
    return nil
}

// extraRoute6Properties is a temporary struct used for validating the fields of ExtraRoute6Properties.
type extraRoute6Properties  struct {
    Discard       *bool                                                   `json:"discard,omitempty"`
    Metric        Optional[int]                                           `json:"metric"`
    NextQualified map[string]ExtraRoute6PropertiesNextQualifiedProperties `json:"next_qualified,omitempty"`
    NoResolve     *bool                                                   `json:"no_resolve,omitempty"`
    Preference    Optional[int]                                           `json:"preference"`
    Via           *string                                                 `json:"via,omitempty"`
}
