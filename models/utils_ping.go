package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// UtilsPing represents a UtilsPing struct.
type UtilsPing struct {
    Count                *int               `json:"count,omitempty"`
    // Interface through which packet needs to egress
    EgressInterface      *string            `json:"egress_interface,omitempty"`
    Host                 string             `json:"host"`
    // only for HA
    Node                 *HaClusterNodeEnum `json:"node,omitempty"`
    Size                 *int               `json:"size,omitempty"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for UtilsPing.
// It customizes the JSON marshaling process for UtilsPing objects.
func (u UtilsPing) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UtilsPing object to a map representation for JSON marshaling.
func (u UtilsPing) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    if u.Count != nil {
        structMap["count"] = u.Count
    }
    if u.EgressInterface != nil {
        structMap["egress_interface"] = u.EgressInterface
    }
    structMap["host"] = u.Host
    if u.Node != nil {
        structMap["node"] = u.Node
    }
    if u.Size != nil {
        structMap["size"] = u.Size
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsPing.
// It customizes the JSON unmarshaling process for UtilsPing objects.
func (u *UtilsPing) UnmarshalJSON(input []byte) error {
    var temp utilsPing
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "count", "egress_interface", "host", "node", "size")
    if err != nil {
    	return err
    }
    
    u.AdditionalProperties = additionalProperties
    u.Count = temp.Count
    u.EgressInterface = temp.EgressInterface
    u.Host = *temp.Host
    u.Node = temp.Node
    u.Size = temp.Size
    return nil
}

// utilsPing is a temporary struct used for validating the fields of UtilsPing.
type utilsPing  struct {
    Count           *int               `json:"count,omitempty"`
    EgressInterface *string            `json:"egress_interface,omitempty"`
    Host            *string            `json:"host"`
    Node            *HaClusterNodeEnum `json:"node,omitempty"`
    Size            *int               `json:"size,omitempty"`
}

func (u *utilsPing) validate() error {
    var errs []string
    if u.Host == nil {
        errs = append(errs, "required field `host` is missing for type `Utils_Ping`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}