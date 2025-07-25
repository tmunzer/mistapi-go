// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RadiusAcctServer represents a RadiusAcctServer struct.
type RadiusAcctServer struct {
    // IP/ hostname of RADIUS server
    Host                 string                   `json:"host"`
    KeywrapEnabled       *bool                    `json:"keywrap_enabled,omitempty"`
    // enum: `ascii`, `hex`
    KeywrapFormat        *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek           *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack          *string                  `json:"keywrap_mack,omitempty"`
    // Radius Auth Port, value from 1 to 65535, default is 1813
    Port                 *RadiusAcctPort          `json:"port,omitempty"`
    // Secret of RADIUS server
    Secret               string                   `json:"secret"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for RadiusAcctServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadiusAcctServer) String() string {
    return fmt.Sprintf(
    	"RadiusAcctServer[Host=%v, KeywrapEnabled=%v, KeywrapFormat=%v, KeywrapKek=%v, KeywrapMack=%v, Port=%v, Secret=%v, AdditionalProperties=%v]",
    	r.Host, r.KeywrapEnabled, r.KeywrapFormat, r.KeywrapKek, r.KeywrapMack, r.Port, r.Secret, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RadiusAcctServer.
// It customizes the JSON marshaling process for RadiusAcctServer objects.
func (r RadiusAcctServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "host", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusAcctServer object to a map representation for JSON marshaling.
func (r RadiusAcctServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["host"] = r.Host
    if r.KeywrapEnabled != nil {
        structMap["keywrap_enabled"] = r.KeywrapEnabled
    }
    if r.KeywrapFormat != nil {
        structMap["keywrap_format"] = r.KeywrapFormat
    }
    if r.KeywrapKek != nil {
        structMap["keywrap_kek"] = r.KeywrapKek
    }
    if r.KeywrapMack != nil {
        structMap["keywrap_mack"] = r.KeywrapMack
    }
    if r.Port != nil {
        structMap["port"] = r.Port.toMap()
    }
    structMap["secret"] = r.Secret
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusAcctServer.
// It customizes the JSON unmarshaling process for RadiusAcctServer objects.
func (r *RadiusAcctServer) UnmarshalJSON(input []byte) error {
    var temp tempRadiusAcctServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "secret")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Host = *temp.Host
    r.KeywrapEnabled = temp.KeywrapEnabled
    r.KeywrapFormat = temp.KeywrapFormat
    r.KeywrapKek = temp.KeywrapKek
    r.KeywrapMack = temp.KeywrapMack
    r.Port = temp.Port
    r.Secret = *temp.Secret
    return nil
}

// tempRadiusAcctServer is a temporary struct used for validating the fields of RadiusAcctServer.
type tempRadiusAcctServer  struct {
    Host           *string                  `json:"host"`
    KeywrapEnabled *bool                    `json:"keywrap_enabled,omitempty"`
    KeywrapFormat  *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek     *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack    *string                  `json:"keywrap_mack,omitempty"`
    Port           *RadiusAcctPort          `json:"port,omitempty"`
    Secret         *string                  `json:"secret"`
}

func (r *tempRadiusAcctServer) validate() error {
    var errs []string
    if r.Host == nil {
        errs = append(errs, "required field `host` is missing for type `radius_acct_server`")
    }
    if r.Secret == nil {
        errs = append(errs, "required field `secret` is missing for type `radius_acct_server`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
