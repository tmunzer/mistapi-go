package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// OrgSettingMgmt represents a OrgSettingMgmt struct.
// management-related properties
type OrgSettingMgmt struct {
    // list of Mist Tunnels
    MxtunnelIds          []uuid.UUID            `json:"mxtunnel_ids,omitempty"`
    // whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel
    UseMxtunnel          *bool                  `json:"use_mxtunnel,omitempty"`
    // whether to use wxtunnel for mgmt connectivity
    UseWxtunnel          *bool                  `json:"use_wxtunnel,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingMgmt,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingMgmt) String() string {
    return fmt.Sprintf(
    	"OrgSettingMgmt[MxtunnelIds=%v, UseMxtunnel=%v, UseWxtunnel=%v, AdditionalProperties=%v]",
    	o.MxtunnelIds, o.UseMxtunnel, o.UseWxtunnel, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingMgmt.
// It customizes the JSON marshaling process for OrgSettingMgmt objects.
func (o OrgSettingMgmt) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "mxtunnel_ids", "use_mxtunnel", "use_wxtunnel"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingMgmt object to a map representation for JSON marshaling.
func (o OrgSettingMgmt) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.MxtunnelIds != nil {
        structMap["mxtunnel_ids"] = o.MxtunnelIds
    }
    if o.UseMxtunnel != nil {
        structMap["use_mxtunnel"] = o.UseMxtunnel
    }
    if o.UseWxtunnel != nil {
        structMap["use_wxtunnel"] = o.UseWxtunnel
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingMgmt.
// It customizes the JSON unmarshaling process for OrgSettingMgmt objects.
func (o *OrgSettingMgmt) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingMgmt
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mxtunnel_ids", "use_mxtunnel", "use_wxtunnel")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.MxtunnelIds = temp.MxtunnelIds
    o.UseMxtunnel = temp.UseMxtunnel
    o.UseWxtunnel = temp.UseWxtunnel
    return nil
}

// tempOrgSettingMgmt is a temporary struct used for validating the fields of OrgSettingMgmt.
type tempOrgSettingMgmt  struct {
    MxtunnelIds []uuid.UUID `json:"mxtunnel_ids,omitempty"`
    UseMxtunnel *bool       `json:"use_mxtunnel,omitempty"`
    UseWxtunnel *bool       `json:"use_wxtunnel,omitempty"`
}
