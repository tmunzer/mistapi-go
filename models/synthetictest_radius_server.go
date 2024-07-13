package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SynthetictestRadiusServer represents a SynthetictestRadiusServer struct.
type SynthetictestRadiusServer struct {
    // Specify the password associated with the username
    Password             string         `json:"password"`
    // Specify the access profile associated with the subscriber
    Profile              *string        `json:"profile,omitempty"`
    // Specify the subscriber username to test
    User                 string         `json:"user"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestRadiusServer.
// It customizes the JSON marshaling process for SynthetictestRadiusServer objects.
func (s SynthetictestRadiusServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestRadiusServer object to a map representation for JSON marshaling.
func (s SynthetictestRadiusServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["password"] = s.Password
    if s.Profile != nil {
        structMap["profile"] = s.Profile
    }
    structMap["user"] = s.User
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestRadiusServer.
// It customizes the JSON unmarshaling process for SynthetictestRadiusServer objects.
func (s *SynthetictestRadiusServer) UnmarshalJSON(input []byte) error {
    var temp synthetictestRadiusServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "password", "profile", "user")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Password = *temp.Password
    s.Profile = temp.Profile
    s.User = *temp.User
    return nil
}

// synthetictestRadiusServer is a temporary struct used for validating the fields of SynthetictestRadiusServer.
type synthetictestRadiusServer  struct {
    Password *string `json:"password"`
    Profile  *string `json:"profile,omitempty"`
    User     *string `json:"user"`
}

func (s *synthetictestRadiusServer) validate() error {
    var errs []string
    if s.Password == nil {
        errs = append(errs, "required field `password` is missing for type `Synthetictest_Radius_Server`")
    }
    if s.User == nil {
        errs = append(errs, "required field `user` is missing for type `Synthetictest_Radius_Server`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
