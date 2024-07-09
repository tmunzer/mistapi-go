package models

import (
    "encoding/json"
)

// Snmpv3ConfigTargetAddressItem represents a Snmpv3ConfigTargetAddressItem struct.
type Snmpv3ConfigTargetAddressItem struct {
    Address              *string        `json:"address,omitempty"`
    AddressMask          *string        `json:"address_mask,omitempty"`
    Port                 *int           `json:"port,omitempty"`
    // <refer to notify tag, can be multiple with blank
    TagList              *string        `json:"tag_list,omitempty"`
    TargetAddressName    *string        `json:"target_address_name,omitempty"`
    // refer to notify target parameters name
    TargetParameters     *string        `json:"target_parameters,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigTargetAddressItem.
// It customizes the JSON marshaling process for Snmpv3ConfigTargetAddressItem objects.
func (s Snmpv3ConfigTargetAddressItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigTargetAddressItem object to a map representation for JSON marshaling.
func (s Snmpv3ConfigTargetAddressItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Address != nil {
        structMap["address"] = s.Address
    }
    if s.AddressMask != nil {
        structMap["address_mask"] = s.AddressMask
    }
    if s.Port != nil {
        structMap["port"] = s.Port
    }
    if s.TagList != nil {
        structMap["tag_list"] = s.TagList
    }
    if s.TargetAddressName != nil {
        structMap["target_address_name"] = s.TargetAddressName
    }
    if s.TargetParameters != nil {
        structMap["target_parameters"] = s.TargetParameters
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigTargetAddressItem.
// It customizes the JSON unmarshaling process for Snmpv3ConfigTargetAddressItem objects.
func (s *Snmpv3ConfigTargetAddressItem) UnmarshalJSON(input []byte) error {
    var temp snmpv3ConfigTargetAddressItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "address", "address_mask", "port", "tag_list", "target_address_name", "target_parameters")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Address = temp.Address
    s.AddressMask = temp.AddressMask
    s.Port = temp.Port
    s.TagList = temp.TagList
    s.TargetAddressName = temp.TargetAddressName
    s.TargetParameters = temp.TargetParameters
    return nil
}

// snmpv3ConfigTargetAddressItem is a temporary struct used for validating the fields of Snmpv3ConfigTargetAddressItem.
type snmpv3ConfigTargetAddressItem  struct {
    Address           *string `json:"address,omitempty"`
    AddressMask       *string `json:"address_mask,omitempty"`
    Port              *int    `json:"port,omitempty"`
    TagList           *string `json:"tag_list,omitempty"`
    TargetAddressName *string `json:"target_address_name,omitempty"`
    TargetParameters  *string `json:"target_parameters,omitempty"`
}
