// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// MxclusterRadsecAuthServer represents a MxclusterRadsecAuthServer struct.
type MxclusterRadsecAuthServer struct {
    // IP / hostname of RADIUS server
    Host                 *string                                           `json:"host,omitempty"`
    // Whether to enable inband status check
    InbandStatusCheck    *bool                                             `json:"inband_status_check,omitempty"`
    // Inband status interval, in seconds
    InbandStatusInterval *int                                              `json:"inband_status_interval,omitempty"`
    // If used for Mist APs, enable keywrap algorithm. Default is false
    KeywrapEnabled       *bool                                             `json:"keywrap_enabled,omitempty"`
    // if used for Mist APs. enum: `ascii`, `hex`
    KeywrapFormat        Optional[MxclusterRadAuthServerKeywrapFormatEnum] `json:"keywrap_format"`
    // If used for Mist APs, encryption key
    KeywrapKek           *string                                           `json:"keywrap_kek,omitempty"`
    // If used for Mist APs, Message Authentication Code Key
    KeywrapMack          *string                                           `json:"keywrap_mack,omitempty"`
    // Auth port of RADIUS server
    Port                 *int                                              `json:"port,omitempty"`
    // Secret of RADIUS server
    Secret               *string                                           `json:"secret,omitempty"`
    // List of ssids that will use this server if match_ssid is true and match is found
    Ssids                []string                                          `json:"ssids,omitempty"`
    AdditionalProperties map[string]interface{}                            `json:"_"`
}

// String implements the fmt.Stringer interface for MxclusterRadsecAuthServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxclusterRadsecAuthServer) String() string {
    return fmt.Sprintf(
    	"MxclusterRadsecAuthServer[Host=%v, InbandStatusCheck=%v, InbandStatusInterval=%v, KeywrapEnabled=%v, KeywrapFormat=%v, KeywrapKek=%v, KeywrapMack=%v, Port=%v, Secret=%v, Ssids=%v, AdditionalProperties=%v]",
    	m.Host, m.InbandStatusCheck, m.InbandStatusInterval, m.KeywrapEnabled, m.KeywrapFormat, m.KeywrapKek, m.KeywrapMack, m.Port, m.Secret, m.Ssids, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxclusterRadsecAuthServer.
// It customizes the JSON marshaling process for MxclusterRadsecAuthServer objects.
func (m MxclusterRadsecAuthServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "host", "inband_status_check", "inband_status_interval", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "secret", "ssids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxclusterRadsecAuthServer object to a map representation for JSON marshaling.
func (m MxclusterRadsecAuthServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Host != nil {
        structMap["host"] = m.Host
    }
    if m.InbandStatusCheck != nil {
        structMap["inband_status_check"] = m.InbandStatusCheck
    }
    if m.InbandStatusInterval != nil {
        structMap["inband_status_interval"] = m.InbandStatusInterval
    }
    if m.KeywrapEnabled != nil {
        structMap["keywrap_enabled"] = m.KeywrapEnabled
    }
    if m.KeywrapFormat.IsValueSet() {
        if m.KeywrapFormat.Value() != nil {
            structMap["keywrap_format"] = m.KeywrapFormat.Value()
        } else {
            structMap["keywrap_format"] = nil
        }
    }
    if m.KeywrapKek != nil {
        structMap["keywrap_kek"] = m.KeywrapKek
    }
    if m.KeywrapMack != nil {
        structMap["keywrap_mack"] = m.KeywrapMack
    }
    if m.Port != nil {
        structMap["port"] = m.Port
    }
    if m.Secret != nil {
        structMap["secret"] = m.Secret
    }
    if m.Ssids != nil {
        structMap["ssids"] = m.Ssids
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxclusterRadsecAuthServer.
// It customizes the JSON unmarshaling process for MxclusterRadsecAuthServer objects.
func (m *MxclusterRadsecAuthServer) UnmarshalJSON(input []byte) error {
    var temp tempMxclusterRadsecAuthServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "host", "inband_status_check", "inband_status_interval", "keywrap_enabled", "keywrap_format", "keywrap_kek", "keywrap_mack", "port", "secret", "ssids")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Host = temp.Host
    m.InbandStatusCheck = temp.InbandStatusCheck
    m.InbandStatusInterval = temp.InbandStatusInterval
    m.KeywrapEnabled = temp.KeywrapEnabled
    m.KeywrapFormat = temp.KeywrapFormat
    m.KeywrapKek = temp.KeywrapKek
    m.KeywrapMack = temp.KeywrapMack
    m.Port = temp.Port
    m.Secret = temp.Secret
    m.Ssids = temp.Ssids
    return nil
}

// tempMxclusterRadsecAuthServer is a temporary struct used for validating the fields of MxclusterRadsecAuthServer.
type tempMxclusterRadsecAuthServer  struct {
    Host                 *string                                           `json:"host,omitempty"`
    InbandStatusCheck    *bool                                             `json:"inband_status_check,omitempty"`
    InbandStatusInterval *int                                              `json:"inband_status_interval,omitempty"`
    KeywrapEnabled       *bool                                             `json:"keywrap_enabled,omitempty"`
    KeywrapFormat        Optional[MxclusterRadAuthServerKeywrapFormatEnum] `json:"keywrap_format"`
    KeywrapKek           *string                                           `json:"keywrap_kek,omitempty"`
    KeywrapMack          *string                                           `json:"keywrap_mack,omitempty"`
    Port                 *int                                              `json:"port,omitempty"`
    Secret               *string                                           `json:"secret,omitempty"`
    Ssids                []string                                          `json:"ssids,omitempty"`
}
