package models

import (
    "encoding/json"
)

// OspfAreasNetwork represents a OspfAreasNetwork struct.
// Property key is the network name. Networks to participate in an OSPF area
type OspfAreasNetwork struct {
    // Required if `auth_type`==`md5`. Property key is the key number
    AuthKeys               map[string]string                 `json:"auth_keys,omitempty"`
    // Required if `auth_type`==`password`, the password, max length is 8
    AuthPassword           *string                           `json:"auth_password,omitempty"`
    // auth type. enum: `md5`, `none`, `password`
    AuthType               *OspfAreaNetworkAuthTypeEnum      `json:"auth_type,omitempty"`
    BfdMinimumInterval     *int                              `json:"bfd_minimum_interval,omitempty"`
    DeadInterval           *int                              `json:"dead_interval,omitempty"`
    ExportPolicy           *string                           `json:"export_policy,omitempty"`
    HelloInterval          *int                              `json:"hello_interval,omitempty"`
    ImportPolicy           *string                           `json:"import_policy,omitempty"`
    // interface type (nbma = non-broadcast multi-access). enum: `broadcast`, `nbma`, `p2mp`, `p2p`
    InterfaceType          *OspfAreaNetworkInterfaceTypeEnum `json:"interface_type,omitempty"`
    Metric                 Optional[int]                     `json:"metric"`
    // by default, we'll re-advertise all learned OSPF routes toward overlay
    NoReadvertiseToOverlay *bool                             `json:"no_readvertise_to_overlay,omitempty"`
    // whether to send OSPF-Hello
    Passive                *bool                             `json:"passive,omitempty"`
    AdditionalProperties   map[string]any                    `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OspfAreasNetwork.
// It customizes the JSON marshaling process for OspfAreasNetwork objects.
func (o OspfAreasNetwork) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OspfAreasNetwork object to a map representation for JSON marshaling.
func (o OspfAreasNetwork) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.AuthKeys != nil {
        structMap["auth_keys"] = o.AuthKeys
    }
    if o.AuthPassword != nil {
        structMap["auth_password"] = o.AuthPassword
    }
    if o.AuthType != nil {
        structMap["auth_type"] = o.AuthType
    }
    if o.BfdMinimumInterval != nil {
        structMap["bfd_minimum_interval"] = o.BfdMinimumInterval
    }
    if o.DeadInterval != nil {
        structMap["dead_interval"] = o.DeadInterval
    }
    if o.ExportPolicy != nil {
        structMap["export_policy"] = o.ExportPolicy
    }
    if o.HelloInterval != nil {
        structMap["hello_interval"] = o.HelloInterval
    }
    if o.ImportPolicy != nil {
        structMap["import_policy"] = o.ImportPolicy
    }
    if o.InterfaceType != nil {
        structMap["interface_type"] = o.InterfaceType
    }
    if o.Metric.IsValueSet() {
        if o.Metric.Value() != nil {
            structMap["metric"] = o.Metric.Value()
        } else {
            structMap["metric"] = nil
        }
    }
    if o.NoReadvertiseToOverlay != nil {
        structMap["no_readvertise_to_overlay"] = o.NoReadvertiseToOverlay
    }
    if o.Passive != nil {
        structMap["passive"] = o.Passive
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OspfAreasNetwork.
// It customizes the JSON unmarshaling process for OspfAreasNetwork objects.
func (o *OspfAreasNetwork) UnmarshalJSON(input []byte) error {
    var temp tempOspfAreasNetwork
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "auth_keys", "auth_password", "auth_type", "bfd_minimum_interval", "dead_interval", "export_policy", "hello_interval", "import_policy", "interface_type", "metric", "no_readvertise_to_overlay", "passive")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.AuthKeys = temp.AuthKeys
    o.AuthPassword = temp.AuthPassword
    o.AuthType = temp.AuthType
    o.BfdMinimumInterval = temp.BfdMinimumInterval
    o.DeadInterval = temp.DeadInterval
    o.ExportPolicy = temp.ExportPolicy
    o.HelloInterval = temp.HelloInterval
    o.ImportPolicy = temp.ImportPolicy
    o.InterfaceType = temp.InterfaceType
    o.Metric = temp.Metric
    o.NoReadvertiseToOverlay = temp.NoReadvertiseToOverlay
    o.Passive = temp.Passive
    return nil
}

// tempOspfAreasNetwork is a temporary struct used for validating the fields of OspfAreasNetwork.
type tempOspfAreasNetwork  struct {
    AuthKeys               map[string]string                 `json:"auth_keys,omitempty"`
    AuthPassword           *string                           `json:"auth_password,omitempty"`
    AuthType               *OspfAreaNetworkAuthTypeEnum      `json:"auth_type,omitempty"`
    BfdMinimumInterval     *int                              `json:"bfd_minimum_interval,omitempty"`
    DeadInterval           *int                              `json:"dead_interval,omitempty"`
    ExportPolicy           *string                           `json:"export_policy,omitempty"`
    HelloInterval          *int                              `json:"hello_interval,omitempty"`
    ImportPolicy           *string                           `json:"import_policy,omitempty"`
    InterfaceType          *OspfAreaNetworkInterfaceTypeEnum `json:"interface_type,omitempty"`
    Metric                 Optional[int]                     `json:"metric"`
    NoReadvertiseToOverlay *bool                             `json:"no_readvertise_to_overlay,omitempty"`
    Passive                *bool                             `json:"passive,omitempty"`
}
