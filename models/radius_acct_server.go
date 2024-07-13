package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RadiusAcctServer represents a RadiusAcctServer struct.
type RadiusAcctServer struct {
    // ip / hostname of RADIUS server
    Host                 string                   `json:"host"`
    KeywrapEnabled       *bool                    `json:"keywrap_enabled,omitempty"`
    KeywrapFormat        *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek           *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack          *string                  `json:"keywrap_mack,omitempty"`
    // Acct port of RADIUS server
    Port                 *int                     `json:"port,omitempty"`
    // secret of RADIUS server
    Secret               string                   `json:"secret"`
    AdditionalProperties map[string]any           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RadiusAcctServer.
// It customizes the JSON marshaling process for RadiusAcctServer objects.
func (r RadiusAcctServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusAcctServer object to a map representation for JSON marshaling.
func (r RadiusAcctServer) toMap() map[string]any {
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

// UnmarshalJSON implements the json.Unmarshaler interface for RadiusAcctServer.
// It customizes the JSON unmarshaling process for RadiusAcctServer objects.
func (r *RadiusAcctServer) UnmarshalJSON(input []byte) error {
    var temp radiusAcctServer
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

// radiusAcctServer is a temporary struct used for validating the fields of RadiusAcctServer.
type radiusAcctServer  struct {
    Host           *string                  `json:"host"`
    KeywrapEnabled *bool                    `json:"keywrap_enabled,omitempty"`
    KeywrapFormat  *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek     *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack    *string                  `json:"keywrap_mack,omitempty"`
    Port           *int                     `json:"port,omitempty"`
    Secret         *string                  `json:"secret"`
}

func (r *radiusAcctServer) validate() error {
    var errs []string
    if r.Host == nil {
        errs = append(errs, "required field `host` is missing for type `Radius_Acct_Server`")
    }
    if r.Secret == nil {
        errs = append(errs, "required field `secret` is missing for type `Radius_Acct_Server`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
