package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// OrgDevice represents a OrgDevice struct.
type OrgDevice struct {
    Mac                  string         `json:"mac"`
    Name                 string         `json:"name"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgDevice.
// It customizes the JSON marshaling process for OrgDevice objects.
func (o OrgDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgDevice object to a map representation for JSON marshaling.
func (o OrgDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    structMap["mac"] = o.Mac
    structMap["name"] = o.Name
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgDevice.
// It customizes the JSON unmarshaling process for OrgDevice objects.
func (o *OrgDevice) UnmarshalJSON(input []byte) error {
    var temp tempOrgDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "name")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.Mac = *temp.Mac
    o.Name = *temp.Name
    return nil
}

// tempOrgDevice is a temporary struct used for validating the fields of OrgDevice.
type tempOrgDevice  struct {
    Mac  *string `json:"mac"`
    Name *string `json:"name"`
}

func (o *tempOrgDevice) validate() error {
    var errs []string
    if o.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `org_device`")
    }
    if o.Name == nil {
        errs = append(errs, "required field `name` is missing for type `org_device`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
