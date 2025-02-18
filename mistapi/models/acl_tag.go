package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// AclTag represents a AclTag struct.
// Resource tags (`type`==`resource` or `type`==`gbp_resource`) can only be used in `dst_tags`
type AclTag struct {
    // Required if
    // - `type`==`dynamic_gbp` (gbp_tag received from RADIUS)
    // - `type`==`gbp_resource`
    // - `type`==`static_gbp` (applying gbp tag against matching conditions)
    GbpTag               *int                   `json:"gbp_tag,omitempty"`
    // Required if
    // - `type`==`mac`
    // - `type`==`static_gbp` if from matching mac
    Macs                 []string               `json:"macs,omitempty"`
    // If:
    // * `type`==`mac` (optional. default is `any`)
    // * `type`==`subnet` (optional. default is `any`)
    // * `type`==`network`
    // * `type`==`resource` (optional. default is `any`)
    // * `type`==`static_gbp` if from matching network (vlan)'
    Network              *string                `json:"network,omitempty"`
    // Required if:
    // * `type`==`radius_group`
    // * `type`==`static_gbp`
    // if from matching radius_group
    RadiusGroup          *string                `json:"radius_group,omitempty"`
    // If `type`==`resource` or `type`==`gbp_resource`. Empty means unrestricted, i.e. any
    Specs                []AclTagSpec           `json:"specs,omitempty"`
    // If
    // - `type`==`subnet`
    // - `type`==`resource` (optional. default is `any`)
    // - `type`==`static_gbp` if from matching subnet
    Subnets              []string               `json:"subnets,omitempty"`
    // enum:
    // * `any`: matching anything not identified
    // * `dynamic_gbp`: from the gbp_tag received from RADIUS
    // * `gbp_resource`: can only be used in `dst_tags`
    // * `mac`
    // * `network`
    // * `radius_group`
    // * `resource`: can only be used in `dst_tags`
    // * `static_gbp`: applying gbp tag against matching conditions
    // * `subnet`'
    Type                 AclTagTypeEnum         `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for AclTag,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (a AclTag) String() string {
    return fmt.Sprintf(
    	"AclTag[GbpTag=%v, Macs=%v, Network=%v, RadiusGroup=%v, Specs=%v, Subnets=%v, Type=%v, AdditionalProperties=%v]",
    	a.GbpTag, a.Macs, a.Network, a.RadiusGroup, a.Specs, a.Subnets, a.Type, a.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for AclTag.
// It customizes the JSON marshaling process for AclTag objects.
func (a AclTag) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "gbp_tag", "macs", "network", "radius_group", "specs", "subnets", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AclTag object to a map representation for JSON marshaling.
func (a AclTag) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.GbpTag != nil {
        structMap["gbp_tag"] = a.GbpTag
    }
    if a.Macs != nil {
        structMap["macs"] = a.Macs
    }
    if a.Network != nil {
        structMap["network"] = a.Network
    }
    if a.RadiusGroup != nil {
        structMap["radius_group"] = a.RadiusGroup
    }
    if a.Specs != nil {
        structMap["specs"] = a.Specs
    }
    if a.Subnets != nil {
        structMap["subnets"] = a.Subnets
    }
    structMap["type"] = a.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AclTag.
// It customizes the JSON unmarshaling process for AclTag objects.
func (a *AclTag) UnmarshalJSON(input []byte) error {
    var temp tempAclTag
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "gbp_tag", "macs", "network", "radius_group", "specs", "subnets", "type")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.GbpTag = temp.GbpTag
    a.Macs = temp.Macs
    a.Network = temp.Network
    a.RadiusGroup = temp.RadiusGroup
    a.Specs = temp.Specs
    a.Subnets = temp.Subnets
    a.Type = *temp.Type
    return nil
}

// tempAclTag is a temporary struct used for validating the fields of AclTag.
type tempAclTag  struct {
    GbpTag      *int            `json:"gbp_tag,omitempty"`
    Macs        []string        `json:"macs,omitempty"`
    Network     *string         `json:"network,omitempty"`
    RadiusGroup *string         `json:"radius_group,omitempty"`
    Specs       []AclTagSpec    `json:"specs,omitempty"`
    Subnets     []string        `json:"subnets,omitempty"`
    Type        *AclTagTypeEnum `json:"type"`
}

func (a *tempAclTag) validate() error {
    var errs []string
    if a.Type == nil {
        errs = append(errs, "required field `type` is missing for type `acl_tag`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
