package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseMxtunnelsPreemptAps represents a ResponseMxtunnelsPreemptAps struct.
type ResponseMxtunnelsPreemptAps struct {
    PreemptedAps         []string       `json:"preempted_aps"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseMxtunnelsPreemptAps.
// It customizes the JSON marshaling process for ResponseMxtunnelsPreemptAps objects.
func (r ResponseMxtunnelsPreemptAps) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMxtunnelsPreemptAps object to a map representation for JSON marshaling.
func (r ResponseMxtunnelsPreemptAps) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["preempted_aps"] = r.PreemptedAps
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMxtunnelsPreemptAps.
// It customizes the JSON unmarshaling process for ResponseMxtunnelsPreemptAps objects.
func (r *ResponseMxtunnelsPreemptAps) UnmarshalJSON(input []byte) error {
    var temp responseMxtunnelsPreemptAps
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "preempted_aps")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.PreemptedAps = *temp.PreemptedAps
    return nil
}

// responseMxtunnelsPreemptAps is a temporary struct used for validating the fields of ResponseMxtunnelsPreemptAps.
type responseMxtunnelsPreemptAps  struct {
    PreemptedAps *[]string `json:"preempted_aps"`
}

func (r *responseMxtunnelsPreemptAps) validate() error {
    var errs []string
    if r.PreemptedAps == nil {
        errs = append(errs, "required field `preempted_aps` is missing for type `Response_Mxtunnels_Preempt_Aps`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
