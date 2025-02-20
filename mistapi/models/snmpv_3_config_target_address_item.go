package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// Snmpv3ConfigTargetAddressItem represents a Snmpv3ConfigTargetAddressItem struct.
type Snmpv3ConfigTargetAddressItem struct {
    Address              string                 `json:"address"`
    AddressMask          string                 `json:"address_mask"`
    Port                 Optional[string]       `json:"port"`
    // Refer to notify tag, can be multiple with blank
    TagList              *string                `json:"tag_list,omitempty"`
    TargetAddressName    string                 `json:"target_address_name"`
    // Refer to notify target parameters name
    TargetParameters     *string                `json:"target_parameters,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Snmpv3ConfigTargetAddressItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s Snmpv3ConfigTargetAddressItem) String() string {
    return fmt.Sprintf(
    	"Snmpv3ConfigTargetAddressItem[Address=%v, AddressMask=%v, Port=%v, TagList=%v, TargetAddressName=%v, TargetParameters=%v, AdditionalProperties=%v]",
    	s.Address, s.AddressMask, s.Port, s.TagList, s.TargetAddressName, s.TargetParameters, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Snmpv3ConfigTargetAddressItem.
// It customizes the JSON marshaling process for Snmpv3ConfigTargetAddressItem objects.
func (s Snmpv3ConfigTargetAddressItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "address", "address_mask", "port", "tag_list", "target_address_name", "target_parameters"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the Snmpv3ConfigTargetAddressItem object to a map representation for JSON marshaling.
func (s Snmpv3ConfigTargetAddressItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["address"] = s.Address
    structMap["address_mask"] = s.AddressMask
    if s.Port.IsValueSet() {
        if s.Port.Value() != nil {
            structMap["port"] = s.Port.Value()
        } else {
            structMap["port"] = nil
        }
    }
    if s.TagList != nil {
        structMap["tag_list"] = s.TagList
    }
    structMap["target_address_name"] = s.TargetAddressName
    if s.TargetParameters != nil {
        structMap["target_parameters"] = s.TargetParameters
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Snmpv3ConfigTargetAddressItem.
// It customizes the JSON unmarshaling process for Snmpv3ConfigTargetAddressItem objects.
func (s *Snmpv3ConfigTargetAddressItem) UnmarshalJSON(input []byte) error {
    var temp tempSnmpv3ConfigTargetAddressItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "address", "address_mask", "port", "tag_list", "target_address_name", "target_parameters")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Address = *temp.Address
    s.AddressMask = *temp.AddressMask
    s.Port = temp.Port
    s.TagList = temp.TagList
    s.TargetAddressName = *temp.TargetAddressName
    s.TargetParameters = temp.TargetParameters
    return nil
}

// tempSnmpv3ConfigTargetAddressItem is a temporary struct used for validating the fields of Snmpv3ConfigTargetAddressItem.
type tempSnmpv3ConfigTargetAddressItem  struct {
    Address           *string          `json:"address"`
    AddressMask       *string          `json:"address_mask"`
    Port              Optional[string] `json:"port"`
    TagList           *string          `json:"tag_list,omitempty"`
    TargetAddressName *string          `json:"target_address_name"`
    TargetParameters  *string          `json:"target_parameters,omitempty"`
}

func (s *tempSnmpv3ConfigTargetAddressItem) validate() error {
    var errs []string
    if s.Address == nil {
        errs = append(errs, "required field `address` is missing for type `snmpv3_config_target_address_item`")
    }
    if s.AddressMask == nil {
        errs = append(errs, "required field `address_mask` is missing for type `snmpv3_config_target_address_item`")
    }
    if s.TargetAddressName == nil {
        errs = append(errs, "required field `target_address_name` is missing for type `snmpv3_config_target_address_item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
