package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ClaimActivation represents a ClaimActivation struct.
type ClaimActivation struct {
    // whether to do a async claim process
    Async                *bool                  `json:"async,omitempty"`
    // activation code
    Code                 string                 `json:"code"`
    // enum: `ap`, `gateway`, `switch`
    DeviceType           *DeviceTypeEnum        `json:"device_type,omitempty"`
    // what to claim. enum: `all`, `inventory`, `license`
    Type                 ClaimTypeEnum          `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClaimActivation.
// It customizes the JSON marshaling process for ClaimActivation objects.
func (c ClaimActivation) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "async", "code", "device_type", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ClaimActivation object to a map representation for JSON marshaling.
func (c ClaimActivation) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Async != nil {
        structMap["async"] = c.Async
    }
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
    var temp tempClaimActivation
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "async", "code", "device_type", "type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Async = temp.Async
    c.Code = *temp.Code
    c.DeviceType = temp.DeviceType
    c.Type = *temp.Type
    return nil
}

// tempClaimActivation is a temporary struct used for validating the fields of ClaimActivation.
type tempClaimActivation  struct {
    Async      *bool           `json:"async,omitempty"`
    Code       *string         `json:"code"`
    DeviceType *DeviceTypeEnum `json:"device_type,omitempty"`
    Type       *ClaimTypeEnum  `json:"type"`
}

func (c *tempClaimActivation) validate() error {
    var errs []string
    if c.Code == nil {
        errs = append(errs, "required field `code` is missing for type `claim_activation`")
    }
    if c.Type == nil {
        errs = append(errs, "required field `type` is missing for type `claim_activation`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
