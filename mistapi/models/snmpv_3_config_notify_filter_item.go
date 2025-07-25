// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// Snmpv3ConfigNotifyFilterItem represents a Snmpv3ConfigNotifyFilterItem struct.
type Snmpv3ConfigNotifyFilterItem struct {
    Contents             []Snmpv3ConfigNotifyFilterItemContent `json:"contents,omitempty"`
    ProfileName          *string                               `json:"profile_name,omitempty"`
    AdditionalProperties map[string]interface{}                `json:"_"`
}

// String implements the fmt.Stringer interface for Snmpv3ConfigNotifyFilterItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Snmpv3ConfigNotifyFilterItem) String() string {
    return fmt.Sprintf(
    	"Snmpv3ConfigNotifyFilterItem[Contents=%v, ProfileName=%v, AdditionalProperties=%v]",
    	s.Contents, s.ProfileName, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigNotifyFilterItem.
// It customizes the JSON marshaling process for Snmpv3ConfigNotifyFilterItem objects.
func (s Snmpv3ConfigNotifyFilterItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "contents", "profile_name"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigNotifyFilterItem object to a map representation for JSON marshaling.
func (s Snmpv3ConfigNotifyFilterItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Contents != nil {
        structMap["contents"] = s.Contents
    }
    if s.ProfileName != nil {
        structMap["profile_name"] = s.ProfileName
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigNotifyFilterItem.
// It customizes the JSON unmarshaling process for Snmpv3ConfigNotifyFilterItem objects.
func (s *Snmpv3ConfigNotifyFilterItem) UnmarshalJSON(input []byte) error {
    var temp tempSnmpv3ConfigNotifyFilterItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "contents", "profile_name")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Contents = temp.Contents
    s.ProfileName = temp.ProfileName
    return nil
}

// tempSnmpv3ConfigNotifyFilterItem is a temporary struct used for validating the fields of Snmpv3ConfigNotifyFilterItem.
type tempSnmpv3ConfigNotifyFilterItem  struct {
    Contents    []Snmpv3ConfigNotifyFilterItemContent `json:"contents,omitempty"`
    ProfileName *string                               `json:"profile_name,omitempty"`
}
