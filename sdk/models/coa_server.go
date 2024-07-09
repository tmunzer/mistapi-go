package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// CoaServer represents a CoaServer struct.
// CoA Server
type CoaServer struct {
    // whether to disable Event-Timestamp Check
    DisableEventTimestampCheck *bool          `json:"disable_event_timestamp_check,omitempty"`
    Enabled                    *bool          `json:"enabled,omitempty"`
    Ip                         string         `json:"ip"`
    Port                       *int           `json:"port,omitempty"`
    Secret                     string         `json:"secret"`
    AdditionalProperties       map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for CoaServer.
// It customizes the JSON marshaling process for CoaServer objects.
func (c CoaServer) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the CoaServer object to a map representation for JSON marshaling.
func (c CoaServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.DisableEventTimestampCheck != nil {
        structMap["disable_event_timestamp_check"] = c.DisableEventTimestampCheck
    }
    if c.Enabled != nil {
        structMap["enabled"] = c.Enabled
    }
    structMap["ip"] = c.Ip
    if c.Port != nil {
        structMap["port"] = c.Port
    }
    structMap["secret"] = c.Secret
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoaServer.
// It customizes the JSON unmarshaling process for CoaServer objects.
func (c *CoaServer) UnmarshalJSON(input []byte) error {
    var temp coaServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "disable_event_timestamp_check", "enabled", "ip", "port", "secret")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.DisableEventTimestampCheck = temp.DisableEventTimestampCheck
    c.Enabled = temp.Enabled
    c.Ip = *temp.Ip
    c.Port = temp.Port
    c.Secret = *temp.Secret
    return nil
}

// coaServer is a temporary struct used for validating the fields of CoaServer.
type coaServer  struct {
    DisableEventTimestampCheck *bool   `json:"disable_event_timestamp_check,omitempty"`
    Enabled                    *bool   `json:"enabled,omitempty"`
    Ip                         *string `json:"ip"`
    Port                       *int    `json:"port,omitempty"`
    Secret                     *string `json:"secret"`
}

func (c *coaServer) validate() error {
    var errs []string
    if c.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `Coa_Server`")
    }
    if c.Secret == nil {
        errs = append(errs, "required field `secret` is missing for type `Coa_Server`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}