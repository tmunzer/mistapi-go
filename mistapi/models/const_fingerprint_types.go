package models

import (
    "encoding/json"
    "fmt"
)

// ConstFingerprintTypes represents a ConstFingerprintTypes struct.
type ConstFingerprintTypes struct {
    Family               []string               `json:"family,omitempty"`
    Mfg                  []string               `json:"mfg,omitempty"`
    Model                []string               `json:"model,omitempty"`
    Os                   []string               `json:"os,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstFingerprintTypes,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstFingerprintTypes) String() string {
    return fmt.Sprintf(
    	"ConstFingerprintTypes[Family=%v, Mfg=%v, Model=%v, Os=%v, AdditionalProperties=%v]",
    	c.Family, c.Mfg, c.Model, c.Os, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstFingerprintTypes.
// It customizes the JSON marshaling process for ConstFingerprintTypes objects.
func (c ConstFingerprintTypes) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "family", "mfg", "model", "os"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstFingerprintTypes object to a map representation for JSON marshaling.
func (c ConstFingerprintTypes) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Family != nil {
        structMap["family"] = c.Family
    }
    if c.Mfg != nil {
        structMap["mfg"] = c.Mfg
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.Os != nil {
        structMap["os"] = c.Os
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstFingerprintTypes.
// It customizes the JSON unmarshaling process for ConstFingerprintTypes objects.
func (c *ConstFingerprintTypes) UnmarshalJSON(input []byte) error {
    var temp tempConstFingerprintTypes
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "family", "mfg", "model", "os")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Family = temp.Family
    c.Mfg = temp.Mfg
    c.Model = temp.Model
    c.Os = temp.Os
    return nil
}

// tempConstFingerprintTypes is a temporary struct used for validating the fields of ConstFingerprintTypes.
type tempConstFingerprintTypes  struct {
    Family []string `json:"family,omitempty"`
    Mfg    []string `json:"mfg,omitempty"`
    Model  []string `json:"model,omitempty"`
    Os     []string `json:"os,omitempty"`
}
