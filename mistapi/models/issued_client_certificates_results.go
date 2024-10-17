package models

import (
	"encoding/json"
)

// IssuedClientCertificatesResults represents a IssuedClientCertificatesResults struct.
type IssuedClientCertificatesResults struct {
	Results              []IssuedClientCertificate `json:"results,omitempty"`
	AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for IssuedClientCertificatesResults.
// It customizes the JSON marshaling process for IssuedClientCertificatesResults objects.
func (i IssuedClientCertificatesResults) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(i.toMap())
}

// toMap converts the IssuedClientCertificatesResults object to a map representation for JSON marshaling.
func (i IssuedClientCertificatesResults) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, i.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "results")
	if err != nil {
		return err
	}

	i.AdditionalProperties = additionalProperties
	i.Results = temp.Results
	return nil
}

// tempIssuedClientCertificatesResults is a temporary struct used for validating the fields of IssuedClientCertificatesResults.
type tempIssuedClientCertificatesResults struct {
	Results []IssuedClientCertificate `json:"results,omitempty"`
}
