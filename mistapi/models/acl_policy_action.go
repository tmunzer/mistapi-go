package models

import (
    "encoding/json"
)

// AclPolicyAction represents a AclPolicyAction struct.
type AclPolicyAction struct {
    // enum: `allow`, `deny`
    Action               *AllowDenyEnum `json:"action,omitempty"`
    DstTag               *string        `json:"dst_tag,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AclPolicyAction.
// It customizes the JSON marshaling process for AclPolicyAction objects.
func (a AclPolicyAction) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AclPolicyAction object to a map representation for JSON marshaling.
func (a AclPolicyAction) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Action != nil {
        structMap["action"] = a.Action
    }
    if a.DstTag != nil {
        structMap["dst_tag"] = a.DstTag
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AclPolicyAction.
// It customizes the JSON unmarshaling process for AclPolicyAction objects.
func (a *AclPolicyAction) UnmarshalJSON(input []byte) error {
    var temp tempAclPolicyAction
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "action", "dst_tag")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Action = temp.Action
    a.DstTag = temp.DstTag
    return nil
}

// tempAclPolicyAction is a temporary struct used for validating the fields of AclPolicyAction.
type tempAclPolicyAction  struct {
    Action *AllowDenyEnum `json:"action,omitempty"`
    DstTag *string        `json:"dst_tag,omitempty"`
}
