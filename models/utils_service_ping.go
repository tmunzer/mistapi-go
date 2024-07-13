package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsServicePing represents a UtilsServicePing struct.
type UtilsServicePing struct {
    Count                *int           `json:"count,omitempty"`
    Host                 string         `json:"host"`
    // ping packet takes the same path as the service
    Service              string         `json:"service"`
    Size                 *int           `json:"size,omitempty"`
    // tenant context in which the packet is sent
    Tenant               *string        `json:"tenant,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsServicePing.
// It customizes the JSON marshaling process for UtilsServicePing objects.
func (u UtilsServicePing) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsServicePing object to a map representation for JSON marshaling.
func (u UtilsServicePing) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Count != nil {
        structMap["count"] = u.Count
    }
    structMap["host"] = u.Host
    structMap["service"] = u.Service
    if u.Size != nil {
        structMap["size"] = u.Size
    }
    if u.Tenant != nil {
        structMap["tenant"] = u.Tenant
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsServicePing.
// It customizes the JSON unmarshaling process for UtilsServicePing objects.
func (u *UtilsServicePing) UnmarshalJSON(input []byte) error {
    var temp utilsServicePing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "count", "host", "service", "size", "tenant")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Count = temp.Count
    u.Host = *temp.Host
    u.Service = *temp.Service
    u.Size = temp.Size
    u.Tenant = temp.Tenant
    return nil
}

// utilsServicePing is a temporary struct used for validating the fields of UtilsServicePing.
type utilsServicePing  struct {
    Count   *int    `json:"count,omitempty"`
    Host    *string `json:"host"`
    Service *string `json:"service"`
    Size    *int    `json:"size,omitempty"`
    Tenant  *string `json:"tenant,omitempty"`
}

func (u *utilsServicePing) validate() error {
    var errs []string
    if u.Host == nil {
        errs = append(errs, "required field `host` is missing for type `Utils_Service_Ping`")
    }
    if u.Service == nil {
        errs = append(errs, "required field `service` is missing for type `Utils_Service_Ping`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
