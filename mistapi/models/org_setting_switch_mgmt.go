package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingSwitchMgmt represents a OrgSettingSwitchMgmt struct.
type OrgSettingSwitchMgmt struct {
    // If the field is set in both site/setting and org/setting, the value from site/setting will be used.
    ApAffinityThreshold  *int                   `json:"ap_affinity_threshold,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingSwitchMgmt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingSwitchMgmt) String() string {
    return fmt.Sprintf(
    	"OrgSettingSwitchMgmt[ApAffinityThreshold=%v, AdditionalProperties=%v]",
    	o.ApAffinityThreshold, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingSwitchMgmt.
// It customizes the JSON marshaling process for OrgSettingSwitchMgmt objects.
func (o OrgSettingSwitchMgmt) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "ap_affinity_threshold"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingSwitchMgmt object to a map representation for JSON marshaling.
func (o OrgSettingSwitchMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_affinity_threshold")
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
