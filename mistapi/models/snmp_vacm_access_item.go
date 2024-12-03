package models

import (
    "encoding/json"
)

// SnmpVacmAccessItem represents a SnmpVacmAccessItem struct.
type SnmpVacmAccessItem struct {
    GroupName            *string                            `json:"group_name,omitempty"`
    PrefixList           []SnmpVacmAccessItemPrefixListItem `json:"prefix_list,omitempty"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpVacmAccessItem.
// It customizes the JSON marshaling process for SnmpVacmAccessItem objects.
func (s SnmpVacmAccessItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "group_name", "prefix_list"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpVacmAccessItem object to a map representation for JSON marshaling.
func (s SnmpVacmAccessItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.GroupName != nil {
        structMap["group_name"] = s.GroupName
    }
    if s.PrefixList != nil {
        structMap["prefix_list"] = s.PrefixList
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpVacmAccessItem.
// It customizes the JSON unmarshaling process for SnmpVacmAccessItem objects.
func (s *SnmpVacmAccessItem) UnmarshalJSON(input []byte) error {
    var temp tempSnmpVacmAccessItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "group_name", "prefix_list")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.GroupName = temp.GroupName
    s.PrefixList = temp.PrefixList
    return nil
}

// tempSnmpVacmAccessItem is a temporary struct used for validating the fields of SnmpVacmAccessItem.
type tempSnmpVacmAccessItem  struct {
    GroupName  *string                            `json:"group_name,omitempty"`
    PrefixList []SnmpVacmAccessItemPrefixListItem `json:"prefix_list,omitempty"`
}
