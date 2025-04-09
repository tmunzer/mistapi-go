package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RadiusAuthServer represents a RadiusAuthServer struct.
// Authentication Server
type RadiusAuthServer struct {
    // IP/ hostname of RADIUS server
    Host                        string                   `json:"host"`
    KeywrapEnabled              *bool                    `json:"keywrap_enabled,omitempty"`
    // enum: `ascii`, `hex`
    KeywrapFormat               *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek                  *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack                 *string                  `json:"keywrap_mack,omitempty"`
    // Radius Auth Port, value from 1 to 65535, default is 1812
    Port                        *RadiusAuthPort          `json:"port,omitempty"`
    // Whether to require Message-Authenticator in requests
    RequireMessageAuthenticator *bool                    `json:"require_message_authenticator,omitempty"`
    // Secret of RADIUS server
    Secret                      string                   `json:"secret"`
    AdditionalProperties        map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for RadiusAuthServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RadiusAuthServer) String() string {
    return fmt.Sprintf(
    	"RadiusAuthServer[Host=%v, KeywrapEnabled=%v, KeywrapFormat=%v, KeywrapKek=%v, KeywrapMack=%v, Port=%v, RequireMessageAuthenticator=%v, Secret=%v, AdditionalProperties=%v]",
    	r.Host, r.KeywrapEnabled, r.KeywrapFormat, r.KeywrapKek, r.KeywrapMack, r.Port, r.RequireMessageAuthenticator, r.Secret, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RadiusAuthServer.
// It customizes the JSON marshaling process for RadiusAuthServer objects.
func (r RadiusAuthServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "host", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "require_message_authenticator", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RadiusAuthServer object to a map representation for JSON marshaling.
func (r RadiusAuthServer) toMap() map[string]any {
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
    if r.RequireMessageAuthenticator != nil {
        structMap["require_message_authenticator"] = r.RequireMessageAuthenticator
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "require_message_authenticator", "secret")
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
    r.RequireMessageAuthenticator = temp.RequireMessageAuthenticator
    r.Secret = *temp.Secret
    return nil
}

// tempRadiusAuthServer is a temporary struct used for validating the fields of RadiusAuthServer.
type tempRadiusAuthServer  struct {
    Host                        *string                  `json:"host"`
    KeywrapEnabled              *bool                    `json:"keywrap_enabled,omitempty"`
    KeywrapFormat               *RadiusKeywrapFormatEnum `json:"keywrap_format,omitempty"`
    KeywrapKek                  *string                  `json:"keywrap_kek,omitempty"`
    KeywrapMack                 *string                  `json:"keywrap_mack,omitempty"`
    Port                        *RadiusAuthPort          `json:"port,omitempty"`
    RequireMessageAuthenticator *bool                    `json:"require_message_authenticator,omitempty"`
    Secret                      *string                  `json:"secret"`
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
    return errors.New(strings.Join (errs, "\n"))
}
