package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// CoaServer represents a CoaServer struct.
// CoA Server
type CoaServer struct {
    // Whether to disable Event-Timestamp Check
    DisableEventTimestampCheck *bool                  `json:"disable_event_timestamp_check,omitempty"`
    Enabled                    *bool                  `json:"enabled,omitempty"`
    Ip                         string                 `json:"ip"`
    // CoA Port, value from 1 to 65535, default is 3799
    Port                       *CoaPort               `json:"port,omitempty"`
    Secret                     string                 `json:"secret"`
    AdditionalProperties       map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for CoaServer,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c CoaServer) String() string {
    return fmt.Sprintf(
    	"CoaServer[DisableEventTimestampCheck=%v, Enabled=%v, Ip=%v, Port=%v, Secret=%v, AdditionalProperties=%v]",
    	c.DisableEventTimestampCheck, c.Enabled, c.Ip, c.Port, c.Secret, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for CoaServer.
// It customizes the JSON marshaling process for CoaServer objects.
func (c CoaServer) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "disable_event_timestamp_check", "enabled", "ip", "port", "secret"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the CoaServer object to a map representation for JSON marshaling.
func (c CoaServer) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.DisableEventTimestampCheck != nil {
        structMap["disable_event_timestamp_check"] = c.DisableEventTimestampCheck
    }
    if c.Enabled != nil {
        structMap["enabled"] = c.Enabled
    }
    structMap["ip"] = c.Ip
    if c.Port != nil {
        structMap["port"] = c.Port.toMap()
    }
    structMap["secret"] = c.Secret
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for CoaServer.
// It customizes the JSON unmarshaling process for CoaServer objects.
func (c *CoaServer) UnmarshalJSON(input []byte) error {
    var temp tempCoaServer
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_event_timestamp_check", "enabled", "ip", "port", "secret")
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

// tempCoaServer is a temporary struct used for validating the fields of CoaServer.
type tempCoaServer  struct {
    DisableEventTimestampCheck *bool    `json:"disable_event_timestamp_check,omitempty"`
    Enabled                    *bool    `json:"enabled,omitempty"`
    Ip                         *string  `json:"ip"`
    Port                       *CoaPort `json:"port,omitempty"`
    Secret                     *string  `json:"secret"`
}

func (c *tempCoaServer) validate() error {
    var errs []string
    if c.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `coa_server`")
    }
    if c.Secret == nil {
        errs = append(errs, "required field `secret` is missing for type `coa_server`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
