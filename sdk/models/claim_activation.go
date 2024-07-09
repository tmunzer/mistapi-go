package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ClaimActivation represents a ClaimActivation struct.
type ClaimActivation struct {
    // activation code
    Code                 string          `json:"code"`
    DeviceType           *DeviceTypeEnum `json:"device_type,omitempty"`
    // what to claim
    Type                 ClaimTypeEnum   `json:"type"`
    AdditionalProperties map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClaimActivation.
// It customizes the JSON marshaling process for ClaimActivation objects.
func (c ClaimActivation) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClaimActivation object to a map representation for JSON marshaling.
func (c ClaimActivation) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["code"] = c.Code
    if c.DeviceType != nil {
        structMap["device_type"] = c.DeviceType
    }
    structMap["type"] = c.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClaimActivation.
// It customizes the JSON unmarshaling process for ClaimActivation objects.
func (c *ClaimActivation) UnmarshalJSON(input []byte) error {
    var temp claimActivation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "code", "device_type", "type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Code = *temp.Code
    c.DeviceType = temp.DeviceType
    c.Type = *temp.Type
    return nil
}

// claimActivation is a temporary struct used for validating the fields of ClaimActivation.
type claimActivation  struct {
    Code       *string         `json:"code"`
    DeviceType *DeviceTypeEnum `json:"device_type,omitempty"`
    Type       *ClaimTypeEnum  `json:"type"`
}

func (c *claimActivation) validate() error {
    var errs []string
    if c.Code == nil {
        errs = append(errs, "required field `code` is missing for type `Claim_Activation`")
    }
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Claim_Activation`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
