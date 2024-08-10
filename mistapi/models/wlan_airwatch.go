package models

import (
    "encoding/json"
)

// WlanAirwatch represents a WlanAirwatch struct.
// airwatch wlan settings
type WlanAirwatch struct {
    // API Key
    ApiKey               *string        `json:"api_key,omitempty"`
    // console URL
    ConsoleUrl           *string        `json:"console_url,omitempty"`
    Enabled              *bool          `json:"enabled,omitempty"`
    // password
    Password             *string        `json:"password,omitempty"`
    // username
    Username             *string        `json:"username,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanAirwatch.
// It customizes the JSON marshaling process for WlanAirwatch objects.
func (w WlanAirwatch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanAirwatch object to a map representation for JSON marshaling.
func (w WlanAirwatch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.ApiKey != nil {
        structMap["api_key"] = w.ApiKey
    }
    if w.ConsoleUrl != nil {
        structMap["console_url"] = w.ConsoleUrl
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.Password != nil {
        structMap["password"] = w.Password
    }
    if w.Username != nil {
        structMap["username"] = w.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanAirwatch.
// It customizes the JSON unmarshaling process for WlanAirwatch objects.
func (w *WlanAirwatch) UnmarshalJSON(input []byte) error {
    var temp tempWlanAirwatch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "api_key", "console_url", "enabled", "password", "username")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.ApiKey = temp.ApiKey
    w.ConsoleUrl = temp.ConsoleUrl
    w.Enabled = temp.Enabled
    w.Password = temp.Password
    w.Username = temp.Username
    return nil
}

// tempWlanAirwatch is a temporary struct used for validating the fields of WlanAirwatch.
type tempWlanAirwatch  struct {
    ApiKey     *string `json:"api_key,omitempty"`
    ConsoleUrl *string `json:"console_url,omitempty"`
    Enabled    *bool   `json:"enabled,omitempty"`
    Password   *string `json:"password,omitempty"`
    Username   *string `json:"username,omitempty"`
}
