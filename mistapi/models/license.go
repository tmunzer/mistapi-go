package models

import (
    "encoding/json"
)

// License represents a License struct.
// License
type License struct {
    Amendments           []LicenseAmendment     `json:"amendments,omitempty"`
    // Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses entitled.
    Entitled             map[string]int         `json:"entitled,omitempty"`
    Licenses             []LicenseSub           `json:"licenses,omitempty"`
    // Property key is license type (e.g. SUB-MAN) and Property value is the number of licenses consumed.
    Summary              map[string]int         `json:"summary,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for License.
// It customizes the JSON marshaling process for License objects.
func (l License) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(l.AdditionalProperties,
        "amendments", "entitled", "licenses", "summary"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(l.toMap())
}

// toMap converts the License object to a map representation for JSON marshaling.
func (l License) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, l.AdditionalProperties)
    if l.Amendments != nil {
        structMap["amendments"] = l.Amendments
    }
    if l.Entitled != nil {
        structMap["entitled"] = l.Entitled
    }
    if l.Licenses != nil {
        structMap["licenses"] = l.Licenses
    }
    if l.Summary != nil {
        structMap["summary"] = l.Summary
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for License.
// It customizes the JSON unmarshaling process for License objects.
func (l *License) UnmarshalJSON(input []byte) error {
    var temp tempLicense
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "amendments", "entitled", "licenses", "summary")
    if err != nil {
    	return err
    }
    l.AdditionalProperties = additionalProperties
    
    l.Amendments = temp.Amendments
    l.Entitled = temp.Entitled
    l.Licenses = temp.Licenses
    l.Summary = temp.Summary
    return nil
}

// tempLicense is a temporary struct used for validating the fields of License.
type tempLicense  struct {
    Amendments []LicenseAmendment `json:"amendments,omitempty"`
    Entitled   map[string]int     `json:"entitled,omitempty"`
    Licenses   []LicenseSub       `json:"licenses,omitempty"`
    Summary    map[string]int     `json:"summary,omitempty"`
}
