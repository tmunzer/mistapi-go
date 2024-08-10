package models

import (
    "encoding/json"
)

// Snmpv3ConfigNotifyFilterItem represents a Snmpv3ConfigNotifyFilterItem struct.
type Snmpv3ConfigNotifyFilterItem struct {
    Contents             []Snmpv3ConfigNotifyFilterItemContent `json:"contents,omitempty"`
    ProfileName          *string                               `json:"profile_name,omitempty"`
    AdditionalProperties map[string]any                        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigNotifyFilterItem.
// It customizes the JSON marshaling process for Snmpv3ConfigNotifyFilterItem objects.
func (s Snmpv3ConfigNotifyFilterItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigNotifyFilterItem object to a map representation for JSON marshaling.
func (s Snmpv3ConfigNotifyFilterItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "contents", "profile_name")
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
