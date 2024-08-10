package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RadiusAuthServer represents a RadiusAuthServer struct.
// Authentication Server
type RadiusAuthServer struct {
    // ip / hostname of RADIUS server
    Host                 string                   `json:"host"`
    KeywrapEnabled       *bool                    `json:"keywrap_enabled,omitempty"`
    // enum: `ascii`, `hex`
    KeywrapFormat        *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek           *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack          *string                  `json:"keywrap_mack,omitempty"`
    // Auth port of RADIUS server
    Port                 *int                     `json:"port,omitempty"`
    // secret of RADIUS server
    Secret               string                   `json:"secret"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RadiusAuthServer.
// It customizes the JSON marshaling process for RadiusAuthServer objects.
func (r RadiusAuthServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusAuthServer object to a map representation for JSON marshaling.
func (r RadiusAuthServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
        structMap["port"] = r.Port
    }
    structMap["secret"] = r.Secret
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusAuthServer.
// It customizes the JSON unmarshaling process for RadiusAuthServer objects.
func (r *RadiusAuthServer) UnmarshalJSON(input []byte) error {
    var temp tempRadiusAuthServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "host", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "secret")
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

// tempRadiusAuthServer is a temporary struct used for validating the fields of RadiusAuthServer.
type tempRadiusAuthServer  struct {
    Host           *string                  `json:"host"`
    KeywrapEnabled *bool                    `json:"keywrap_enabled,omitempty"`
    KeywrapFormat  *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek     *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack    *string                  `json:"keywrap_mack,omitempty"`
    Port           *int                     `json:"port,omitempty"`
    Secret         *string                  `json:"secret"`
}

func (r *tempRadiusAuthServer) validate() error {
    var errs []string
    if r.Host == nil {
        errs = append(errs, "required field `host` is missing for type `radius_auth_server`")
    }
    if r.Secret == nil {
        errs = append(errs, "required field `secret` is missing for type `radius_auth_server`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
