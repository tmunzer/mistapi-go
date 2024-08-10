package models

import (
    "encoding/json"
)

// AclPolicy represents a AclPolicy struct.
// - for GBP-based policy, all src_tags and dst_tags have to be gbp-based
// - for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to
type AclPolicy struct {
    // - for GBP-based policy, all src_tags and dst_tags have to be gbp-based
    // - for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to
    Actions              []AclPolicyAction `json:"actions,omitempty"`
    Name                 *string           `json:"name,omitempty"`
    // - for GBP-based policy, all src_tags and dst_tags have to be gbp-based
    // - for ACL-based policy, `network` is required in either the source or destination so that we know where to attach the policy to
    SrcTags              []string          `json:"src_tags,omitempty"`
    AdditionalProperties map[string]any    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AclPolicy.
// It customizes the JSON marshaling process for AclPolicy objects.
func (a AclPolicy) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(a.toMap())
}

// toMap converts the AclPolicy object to a map representation for JSON marshaling.
func (a AclPolicy) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, a.AdditionalProperties)
    if a.Actions != nil {
        structMap["actions"] = a.Actions
    }
    if a.Name != nil {
        structMap["name"] = a.Name
    }
    if a.SrcTags != nil {
        structMap["src_tags"] = a.SrcTags
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AclPolicy.
// It customizes the JSON unmarshaling process for AclPolicy objects.
func (a *AclPolicy) UnmarshalJSON(input []byte) error {
    var temp tempAclPolicy
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "actions", "name", "src_tags")
    if err != nil {
    	return err
    }
    
    a.AdditionalProperties = additionalProperties
    a.Actions = temp.Actions
    a.Name = temp.Name
    a.SrcTags = temp.SrcTags
    return nil
}

// tempAclPolicy is a temporary struct used for validating the fields of AclPolicy.
type tempAclPolicy  struct {
    Actions []AclPolicyAction `json:"actions,omitempty"`
    Name    *string           `json:"name,omitempty"`
    SrcTags []string          `json:"src_tags,omitempty"`
}
