package models

import (
    "encoding/json"
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
    AdditionalProperties map[string]any               `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for VirtualChassisUpdate.
// It customizes the JSON marshaling process for VirtualChassisUpdate objects.
func (v VirtualChassisUpdate) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(v.toMap())
}

// toMap converts the VirtualChassisUpdate object to a map representation for JSON marshaling.
func (v VirtualChassisUpdate) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, v.AdditionalProperties)
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
    var temp virtualChassisUpdate
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "member", "members", "new-member", "op")
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

// virtualChassisUpdate is a temporary struct used for validating the fields of VirtualChassisUpdate.
type virtualChassisUpdate  struct {
    Member    *int                         `json:"member,omitempty"`
    Members   []VirtualChassisMemberUpdate `json:"members,omitempty"`
    NewMember *int                         `json:"new-member,omitempty"`
    Op        *VirtualChassisUpdateOpEnum  `json:"op,omitempty"`
}
