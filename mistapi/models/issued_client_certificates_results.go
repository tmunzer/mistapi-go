// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// IssuedClientCertificatesResults represents a IssuedClientCertificatesResults struct.
type IssuedClientCertificatesResults struct {
    Results              []IssuedClientCertificate `json:"results,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for IssuedClientCertificatesResults,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IssuedClientCertificatesResults) String() string {
    return fmt.Sprintf(
    	"IssuedClientCertificatesResults[Results=%v, AdditionalProperties=%v]",
    	i.Results, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IssuedClientCertificatesResults.
// It customizes the JSON marshaling process for IssuedClientCertificatesResults objects.
func (i IssuedClientCertificatesResults) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IssuedClientCertificatesResults object to a map representation for JSON marshaling.
func (i IssuedClientCertificatesResults) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Results != nil {
        structMap["results"] = i.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IssuedClientCertificatesResults.
// It customizes the JSON unmarshaling process for IssuedClientCertificatesResults objects.
func (i *IssuedClientCertificatesResults) UnmarshalJSON(input []byte) error {
    var temp tempIssuedClientCertificatesResults
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "results")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.Results = temp.Results
    return nil
}

// tempIssuedClientCertificatesResults is a temporary struct used for validating the fields of IssuedClientCertificatesResults.
type tempIssuedClientCertificatesResults  struct {
    Results []IssuedClientCertificate `json:"results,omitempty"`
}
