// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// VirtualChassisUpdate represents a VirtualChassisUpdate struct.
// Virtual Chassis
type VirtualChassisUpdate struct {
    // Only if `op`==`renumber`
    Member               *int                         `json:"member,omitempty"`
    Members              []VirtualChassisMemberUpdate `json:"members,omitempty"`
    // Only if `op`==`renumber`
    NewMember            *int                         `json:"new-member,omitempty"`
    // enum: `add`, `preprovision`, `remove`, `renumber`
    Op                   *VirtualChassisUpdateOpEnum  `json:"op,omitempty"`
    AdditionalProperties map[string]interface{}       `json:"_"`
}

// String implements the fmt.Stringer interface for VirtualChassisUpdate,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (v VirtualChassisUpdate) String() string {
    return fmt.Sprintf(
    	"VirtualChassisUpdate[Member=%v, Members=%v, NewMember=%v, Op=%v, AdditionalProperties=%v]",
    	v.Member, v.Members, v.NewMember, v.Op, v.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisUpdate.
// It customizes the JSON marshaling process for VirtualChassisUpdate objects.
func (v VirtualChassisUpdate) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(v.AdditionalProperties,
        "member", "members", "new-member", "op"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisUpdate object to a map representation for JSON marshaling.
func (v VirtualChassisUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, v.AdditionalProperties)
    if v.Member != nil {
        structMap["member"] = v.Member
    }
    if v.Members != nil {
        structMap["members"] = v.Members
    }
    if v.NewMember != nil {
        structMap["new-member"] = v.NewMember
    }
    if v.Op != nil {
        structMap["op"] = v.Op
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for VirtualChassisUpdate.
// It customizes the JSON unmarshaling process for VirtualChassisUpdate objects.
func (v *VirtualChassisUpdate) UnmarshalJSON(input []byte) error {
    var temp tempVirtualChassisUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "member", "members", "new-member", "op")
    if err != nil {
    	return err
    }
    v.AdditionalProperties = additionalProperties
    
    v.Member = temp.Member
    v.Members = temp.Members
    v.NewMember = temp.NewMember
    v.Op = temp.Op
    return nil
}

// tempVirtualChassisUpdate is a temporary struct used for validating the fields of VirtualChassisUpdate.
type tempVirtualChassisUpdate  struct {
    Member    *int                         `json:"member,omitempty"`
    Members   []VirtualChassisMemberUpdate `json:"members,omitempty"`
    NewMember *int                         `json:"new-member,omitempty"`
    Op        *VirtualChassisUpdateOpEnum  `json:"op,omitempty"`
}
