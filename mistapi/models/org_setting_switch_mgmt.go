package models

import (
    "encoding/json"
)

// OrgSettingSwitchMgmt represents a OrgSettingSwitchMgmt struct.
type OrgSettingSwitchMgmt struct {
    // If the field is set in both site/setting and org/setting, the value from site/setting will be used.
    ApAffinityThreshold  *int           `json:"ap_affinity_threshold,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingSwitchMgmt.
// It customizes the JSON marshaling process for OrgSettingSwitchMgmt objects.
func (o OrgSettingSwitchMgmt) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingSwitchMgmt object to a map representation for JSON marshaling.
func (o OrgSettingSwitchMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, o.AdditionalProperties)
    if o.ApAffinityThreshold != nil {
        structMap["ap_affinity_threshold"] = o.ApAffinityThreshold
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingSwitchMgmt.
// It customizes the JSON unmarshaling process for OrgSettingSwitchMgmt objects.
func (o *OrgSettingSwitchMgmt) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingSwitchMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_affinity_threshold")
    if err != nil {
    	return err
    }
    
    o.AdditionalProperties = additionalProperties
    o.ApAffinityThreshold = temp.ApAffinityThreshold
    return nil
}

// tempOrgSettingSwitchMgmt is a temporary struct used for validating the fields of OrgSettingSwitchMgmt.
type tempOrgSettingSwitchMgmt  struct {
    ApAffinityThreshold *int `json:"ap_affinity_threshold,omitempty"`
}
