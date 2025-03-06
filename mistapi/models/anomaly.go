package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Anomaly represents a Anomaly struct.
// Anomaly
type Anomaly struct {
    Events               []string               `json:"events"`
    Since                *float64               `json:"since,omitempty"`
    SleBaseline          float64                `json:"sle_baseline"`
    SleDeviation         float64                `json:"sle_deviation"`
    // Epoch (seconds)
    Timestamp            float64                `json:"timestamp"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Anomaly,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a Anomaly) String() string {
    return fmt.Sprintf(
    	"Anomaly[Events=%v, Since=%v, SleBaseline=%v, SleDeviation=%v, Timestamp=%v, AdditionalProperties=%v]",
    	a.Events, a.Since, a.SleBaseline, a.SleDeviation, a.Timestamp, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Anomaly.
// It customizes the JSON marshaling process for Anomaly objects.
func (a Anomaly) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "events", "since", "sle_baseline", "sle_deviation", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the Anomaly object to a map representation for JSON marshaling.
func (a Anomaly) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    structMap["events"] = a.Events
    if a.Since != nil {
        structMap["since"] = a.Since
    }
    structMap["sle_baseline"] = a.SleBaseline
    structMap["sle_deviation"] = a.SleDeviation
    structMap["timestamp"] = a.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Anomaly.
// It customizes the JSON unmarshaling process for Anomaly objects.
func (a *Anomaly) UnmarshalJSON(input []byte) error {
    var temp tempAnomaly
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "events", "since", "sle_baseline", "sle_deviation", "timestamp")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Events = *temp.Events
    a.Since = temp.Since
    a.SleBaseline = *temp.SleBaseline
    a.SleDeviation = *temp.SleDeviation
    a.Timestamp = *temp.Timestamp
    return nil
}

// tempAnomaly is a temporary struct used for validating the fields of Anomaly.
type tempAnomaly  struct {
    Events       *[]string `json:"events"`
    Since        *float64  `json:"since,omitempty"`
    SleBaseline  *float64  `json:"sle_baseline"`
    SleDeviation *float64  `json:"sle_deviation"`
    Timestamp    *float64  `json:"timestamp"`
}

func (a *tempAnomaly) validate() error {
    var errs []string
    if a.Events == nil {
        errs = append(errs, "required field `events` is missing for type `anomaly`")
    }
    if a.SleBaseline == nil {
        errs = append(errs, "required field `sle_baseline` is missing for type `anomaly`")
    }
    if a.SleDeviation == nil {
        errs = append(errs, "required field `sle_deviation` is missing for type `anomaly`")
    }
    if a.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `anomaly`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
