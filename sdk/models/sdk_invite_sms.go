package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SdkInviteSms represents a SdkInviteSms struct.
type SdkInviteSms struct {
    Number               string         `json:"number"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SdkInviteSms.
// It customizes the JSON marshaling process for SdkInviteSms objects.
func (s SdkInviteSms) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SdkInviteSms object to a map representation for JSON marshaling.
func (s SdkInviteSms) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["number"] = s.Number
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SdkInviteSms.
// It customizes the JSON unmarshaling process for SdkInviteSms objects.
func (s *SdkInviteSms) UnmarshalJSON(input []byte) error {
    var temp sdkInviteSms
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "number")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Number = *temp.Number
    return nil
}

// sdkInviteSms is a temporary struct used for validating the fields of SdkInviteSms.
type sdkInviteSms  struct {
    Number *string `json:"number"`
}

func (s *sdkInviteSms) validate() error {
    var errs []string
    if s.Number == nil {
        errs = append(errs, "required field `number` is missing for type `Sdk_Invite_Sms`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
