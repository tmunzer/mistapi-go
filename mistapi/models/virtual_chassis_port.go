package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// VirtualChassisPort represents a VirtualChassisPort struct.
type VirtualChassisPort struct {
    Members              []ConfigVcPortMember            `json:"members"`
    // enum: `delete`, `set`
    Op                   VirtualChassisPortOperationEnum `json:"op"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisPort.
// It customizes the JSON marshaling process for VirtualChassisPort objects.
func (v VirtualChassisPort) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisPort object to a map representation for JSON marshaling.
func (v VirtualChassisPort) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
    structMap["members"] = v.Members
    structMap["op"] = v.Op
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VirtualChassisPort.
// It customizes the JSON unmarshaling process for VirtualChassisPort objects.
func (v *VirtualChassisPort) UnmarshalJSON(input []byte) error {
    var temp tempVirtualChassisPort
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "members", "op")
    if err != nil {
    	return err
    }
    
    v.AdditionalProperties = additionalProperties
    v.Members = *temp.Members
    v.Op = *temp.Op
    return nil
}

// tempVirtualChassisPort is a temporary struct used for validating the fields of VirtualChassisPort.
type tempVirtualChassisPort  struct {
    Members *[]ConfigVcPortMember            `json:"members"`
    Op      *VirtualChassisPortOperationEnum `json:"op"`
}

func (v *tempVirtualChassisPort) validate() error {
    var errs []string
    if v.Members == nil {
        errs = append(errs, "required field `members` is missing for type `virtual_chassis_port`")
    }
    if v.Op == nil {
        errs = append(errs, "required field `op` is missing for type `virtual_chassis_port`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
