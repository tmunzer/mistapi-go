package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// AclPolicyAction represents a AclPolicyAction struct.
type AclPolicyAction struct {
    // enum: `allow`, `deny`
    Action               *AllowDenyEnum         `json:"action,omitempty"`
    DstTag               string                 `json:"dst_tag"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AclPolicyAction.
// It customizes the JSON marshaling process for AclPolicyAction objects.
func (a AclPolicyAction) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "action", "dst_tag"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AclPolicyAction object to a map representation for JSON marshaling.
func (a AclPolicyAction) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Action != nil {
        structMap["action"] = a.Action
    }
    structMap["dst_tag"] = a.DstTag
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
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "action", "dst_tag")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.Action = temp.Action
    a.DstTag = *temp.DstTag
    return nil
}

// tempAclPolicyAction is a temporary struct used for validating the fields of AclPolicyAction.
type tempAclPolicyAction  struct {
    Action *AllowDenyEnum `json:"action,omitempty"`
    DstTag *string        `json:"dst_tag"`
}

func (a *tempAclPolicyAction) validate() error {
    var errs []string
    if a.DstTag == nil {
        errs = append(errs, "required field `dst_tag` is missing for type `acl_policy_action`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
