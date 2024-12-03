package models

import (
    "encoding/json"
)

// ResponseNacCrlFiles represents a ResponseNacCrlFiles struct.
type ResponseNacCrlFiles struct {
    Results              []NacCrlFile           `json:"results,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseNacCrlFiles.
// It customizes the JSON marshaling process for ResponseNacCrlFiles objects.
func (r ResponseNacCrlFiles) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseNacCrlFiles object to a map representation for JSON marshaling.
func (r ResponseNacCrlFiles) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Results != nil {
        structMap["results"] = r.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseNacCrlFiles.
// It customizes the JSON unmarshaling process for ResponseNacCrlFiles objects.
func (r *ResponseNacCrlFiles) UnmarshalJSON(input []byte) error {
    var temp tempResponseNacCrlFiles
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Results = temp.Results
    return nil
}

// tempResponseNacCrlFiles is a temporary struct used for validating the fields of ResponseNacCrlFiles.
type tempResponseNacCrlFiles  struct {
    Results []NacCrlFile `json:"results,omitempty"`
}
