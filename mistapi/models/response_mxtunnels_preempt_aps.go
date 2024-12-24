package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseMxtunnelsPreemptAps represents a ResponseMxtunnelsPreemptAps struct.
type ResponseMxtunnelsPreemptAps struct {
    PreemptedAps         []string               `json:"preempted_aps"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseMxtunnelsPreemptAps,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseMxtunnelsPreemptAps) String() string {
    return fmt.Sprintf(
    	"ResponseMxtunnelsPreemptAps[PreemptedAps=%v, AdditionalProperties=%v]",
    	r.PreemptedAps, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseMxtunnelsPreemptAps.
// It customizes the JSON marshaling process for ResponseMxtunnelsPreemptAps objects.
func (r ResponseMxtunnelsPreemptAps) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "preempted_aps"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseMxtunnelsPreemptAps object to a map representation for JSON marshaling.
func (r ResponseMxtunnelsPreemptAps) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["preempted_aps"] = r.PreemptedAps
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseMxtunnelsPreemptAps.
// It customizes the JSON unmarshaling process for ResponseMxtunnelsPreemptAps objects.
func (r *ResponseMxtunnelsPreemptAps) UnmarshalJSON(input []byte) error {
    var temp tempResponseMxtunnelsPreemptAps
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "preempted_aps")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.PreemptedAps = *temp.PreemptedAps
    return nil
}

// tempResponseMxtunnelsPreemptAps is a temporary struct used for validating the fields of ResponseMxtunnelsPreemptAps.
type tempResponseMxtunnelsPreemptAps  struct {
    PreemptedAps *[]string `json:"preempted_aps"`
}

func (r *tempResponseMxtunnelsPreemptAps) validate() error {
    var errs []string
    if r.PreemptedAps == nil {
        errs = append(errs, "required field `preempted_aps` is missing for type `response_mxtunnels_preempt_aps`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
