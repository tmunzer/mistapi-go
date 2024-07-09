package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseOrgDevices represents a ResponseOrgDevices struct.
type ResponseOrgDevices struct {
    Results              []OrgDevice    `json:"results"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgDevices.
// It customizes the JSON marshaling process for ResponseOrgDevices objects.
func (r ResponseOrgDevices) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgDevices object to a map representation for JSON marshaling.
func (r ResponseOrgDevices) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["results"] = r.Results
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgDevices.
// It customizes the JSON unmarshaling process for ResponseOrgDevices objects.
func (r *ResponseOrgDevices) UnmarshalJSON(input []byte) error {
    var temp responseOrgDevices
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "results")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Results = *temp.Results
    return nil
}

// responseOrgDevices is a temporary struct used for validating the fields of ResponseOrgDevices.
type responseOrgDevices  struct {
    Results *[]OrgDevice `json:"results"`
}

func (r *responseOrgDevices) validate() error {
    var errs []string
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `Response_Org_Devices`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
