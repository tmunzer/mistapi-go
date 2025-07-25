// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SnmpUsmUser represents a SnmpUsmUser struct.
type SnmpUsmUser struct {
    // Not required if `authentication_type`==`authentication-none`. Include alphabetic, numeric, and special characters, but it cannot include control characters.
    AuthenticationPassword *string                            `json:"authentication_password,omitempty"`
    // sha224, sha256, sha384, sha512 are supported in 21.1 and newer release. enum: `authentication-md5`, `authentication-none`, `authentication-sha`, `authentication-sha224`, `authentication-sha256`, `authentication-sha384`, `authentication-sha512`
    AuthenticationType     *SnmpUsmUserAuthenticationTypeEnum `json:"authentication_type,omitempty"`
    // Not required if `encryption_type`==`privacy-none`. Include alphabetic, numeric, and special characters, but it cannot include control characters
    EncryptionPassword     *string                            `json:"encryption_password,omitempty"`
    // enum: `privacy-3des`, `privacy-aes128`, `privacy-des`, `privacy-none`
    EncryptionType         *SnmpUsmUserEncryptionTypeEnum     `json:"encryption_type,omitempty"`
    Name                   *string                            `json:"name,omitempty"`
    AdditionalProperties   map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpUsmUser,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpUsmUser) String() string {
    return fmt.Sprintf(
    	"SnmpUsmUser[AuthenticationPassword=%v, AuthenticationType=%v, EncryptionPassword=%v, EncryptionType=%v, Name=%v, AdditionalProperties=%v]",
    	s.AuthenticationPassword, s.AuthenticationType, s.EncryptionPassword, s.EncryptionType, s.Name, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpUsmUser.
// It customizes the JSON marshaling process for SnmpUsmUser objects.
func (s SnmpUsmUser) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "authentication_password", "authentication_type", "encryption_password", "encryption_type", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpUsmUser object to a map representation for JSON marshaling.
func (s SnmpUsmUser) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AuthenticationPassword != nil {
        structMap["authentication_password"] = s.AuthenticationPassword
    }
    if s.AuthenticationType != nil {
        structMap["authentication_type"] = s.AuthenticationType
    }
    if s.EncryptionPassword != nil {
        structMap["encryption_password"] = s.EncryptionPassword
    }
    if s.EncryptionType != nil {
        structMap["encryption_type"] = s.EncryptionType
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpUsmUser.
// It customizes the JSON unmarshaling process for SnmpUsmUser objects.
func (s *SnmpUsmUser) UnmarshalJSON(input []byte) error {
    var temp tempSnmpUsmUser
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authentication_password", "authentication_type", "encryption_password", "encryption_type", "name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AuthenticationPassword = temp.AuthenticationPassword
    s.AuthenticationType = temp.AuthenticationType
    s.EncryptionPassword = temp.EncryptionPassword
    s.EncryptionType = temp.EncryptionType
    s.Name = temp.Name
    return nil
}

// tempSnmpUsmUser is a temporary struct used for validating the fields of SnmpUsmUser.
type tempSnmpUsmUser  struct {
    AuthenticationPassword *string                            `json:"authentication_password,omitempty"`
    AuthenticationType     *SnmpUsmUserAuthenticationTypeEnum `json:"authentication_type,omitempty"`
    EncryptionPassword     *string                            `json:"encryption_password,omitempty"`
    EncryptionType         *SnmpUsmUserEncryptionTypeEnum     `json:"encryption_type,omitempty"`
    Name                   *string                            `json:"name,omitempty"`
}
