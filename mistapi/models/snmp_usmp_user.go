package models

import (
    "encoding/json"
)

// SnmpUsmpUser represents a SnmpUsmpUser struct.
type SnmpUsmpUser struct {
    // Not required if `authentication_type`==`authentication_none`
    // include alphabetic, numeric, and special characters, but it cannot include control characters.
    AuthenticationPassword *string                             `json:"authentication_password,omitempty"`
    // sha224, sha256, sha384, sha512 are supported in 21.1 and newer release. enum: `authentication_md5`, `authentication_none`, `authentication_sha`, `authentication_sha224`, `authentication_sha256`, `authentication_sha384`, `authentication_sha512`
    AuthenticationType     *SnmpUsmpUserAuthenticationTypeEnum `json:"authentication_type,omitempty"`
    // Not required if `encryption_type`==`privacy-none`
    // include alphabetic, numeric, and special characters, but it cannot include control characters
    EncryptionPassword     *string                             `json:"encryption_password,omitempty"`
    // enum: `privacy-3des`, `privacy-aes128`, `privacy-des`, `privacy-none`
    EncryptionType         *SnmpUsmpUserEncryptionTypeEnum     `json:"encryption_type,omitempty"`
    Name                   *string                             `json:"name,omitempty"`
    AdditionalProperties   map[string]interface{}              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpUsmpUser.
// It customizes the JSON marshaling process for SnmpUsmpUser objects.
func (s SnmpUsmpUser) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "authentication_password", "authentication_type", "encryption_password", "encryption_type", "name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpUsmpUser object to a map representation for JSON marshaling.
func (s SnmpUsmpUser) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpUsmpUser.
// It customizes the JSON unmarshaling process for SnmpUsmpUser objects.
func (s *SnmpUsmpUser) UnmarshalJSON(input []byte) error {
    var temp tempSnmpUsmpUser
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

// tempSnmpUsmpUser is a temporary struct used for validating the fields of SnmpUsmpUser.
type tempSnmpUsmpUser  struct {
    AuthenticationPassword *string                             `json:"authentication_password,omitempty"`
    AuthenticationType     *SnmpUsmpUserAuthenticationTypeEnum `json:"authentication_type,omitempty"`
    EncryptionPassword     *string                             `json:"encryption_password,omitempty"`
    EncryptionType         *SnmpUsmpUserEncryptionTypeEnum     `json:"encryption_type,omitempty"`
    Name                   *string                             `json:"name,omitempty"`
}
