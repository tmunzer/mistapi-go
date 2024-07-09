package models

import (
    "encoding/json"
)

// ResponseNacCrlFiles represents a ResponseNacCrlFiles struct.
type ResponseNacCrlFiles struct {
    Results              []NacCrlFile   `json:"results,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseNacCrlFiles.
// It customizes the JSON marshaling process for ResponseNacCrlFiles objects.
func (r ResponseNacCrlFiles) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseNacCrlFiles object to a map representation for JSON marshaling.
func (r ResponseNacCrlFiles) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseNacCrlFiles.
// It customizes the JSON unmarshaling process for ResponseNacCrlFiles objects.
func (r *ResponseNacCrlFiles) UnmarshalJSON(input []byte) error {
    var temp responseNacCrlFiles
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

// responseNacCrlFiles is a temporary struct used for validating the fields of ResponseNacCrlFiles.
type responseNacCrlFiles  struct {
    Results []NacCrlFile `json:"results,omitempty"`
}
