// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// SkyatpList represents a SkyatpList struct.
type SkyatpList struct {
    Domains              []SkyatpListDomain     `json:"domains,omitempty"`
    Ip                   []SkyatpListIp         `json:"ip,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SkyatpList,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SkyatpList) String() string {
    return fmt.Sprintf(
    	"SkyatpList[Domains=%v, Ip=%v, AdditionalProperties=%v]",
    	s.Domains, s.Ip, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SkyatpList.
// It customizes the JSON marshaling process for SkyatpList objects.
func (s SkyatpList) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "domains", "ip"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SkyatpList object to a map representation for JSON marshaling.
func (s SkyatpList) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Domains != nil {
        structMap["domains"] = s.Domains
    }
    if s.Ip != nil {
        structMap["ip"] = s.Ip
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SkyatpList.
// It customizes the JSON unmarshaling process for SkyatpList objects.
func (s *SkyatpList) UnmarshalJSON(input []byte) error {
    var temp tempSkyatpList
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "domains", "ip")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Domains = temp.Domains
    s.Ip = temp.Ip
    return nil
}

// tempSkyatpList is a temporary struct used for validating the fields of SkyatpList.
type tempSkyatpList  struct {
    Domains []SkyatpListDomain `json:"domains,omitempty"`
    Ip      []SkyatpListIp     `json:"ip,omitempty"`
}
