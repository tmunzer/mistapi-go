package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// DiscoveredSwitch represents a DiscoveredSwitch struct.
type DiscoveredSwitch struct {
    Adopted              *bool                `json:"adopted,omitempty"`
    ApRedundancy         *ApRedundancy        `json:"ap_redundancy,omitempty"`
    Aps                  []DiscoveredSwitchAp `json:"aps,omitempty"`
    ChassisId            []string             `json:"chassis_id,omitempty"`
    ForSite              *bool                `json:"for_site,omitempty"`
    Model                *string              `json:"model,omitempty"`
    OrgId                *uuid.UUID           `json:"org_id,omitempty"`
    SiteId               *uuid.UUID           `json:"site_id,omitempty"`
    SystemDesc           *string              `json:"system_desc,omitempty"`
    SystemName           *string              `json:"system_name,omitempty"`
    Timestamp            *float64             `json:"timestamp,omitempty"`
    Vendor               *string              `json:"vendor,omitempty"`
    Version              *string              `json:"version,omitempty"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for DiscoveredSwitch.
// It customizes the JSON marshaling process for DiscoveredSwitch objects.
func (d DiscoveredSwitch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the DiscoveredSwitch object to a map representation for JSON marshaling.
func (d DiscoveredSwitch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Adopted != nil {
        structMap["adopted"] = d.Adopted
    }
    if d.ApRedundancy != nil {
        structMap["ap_redundancy"] = d.ApRedundancy.toMap()
    }
    if d.Aps != nil {
        structMap["aps"] = d.Aps
    }
    if d.ChassisId != nil {
        structMap["chassis_id"] = d.ChassisId
    }
    if d.ForSite != nil {
        structMap["for_site"] = d.ForSite
    }
    if d.Model != nil {
        structMap["model"] = d.Model
    }
    if d.OrgId != nil {
        structMap["org_id"] = d.OrgId
    }
    if d.SiteId != nil {
        structMap["site_id"] = d.SiteId
    }
    if d.SystemDesc != nil {
        structMap["system_desc"] = d.SystemDesc
    }
    if d.SystemName != nil {
        structMap["system_name"] = d.SystemName
    }
    if d.Timestamp != nil {
        structMap["timestamp"] = d.Timestamp
    }
    if d.Vendor != nil {
        structMap["vendor"] = d.Vendor
    }
    if d.Version != nil {
        structMap["version"] = d.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DiscoveredSwitch.
// It customizes the JSON unmarshaling process for DiscoveredSwitch objects.
func (d *DiscoveredSwitch) UnmarshalJSON(input []byte) error {
    var temp discoveredSwitch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "adopted", "ap_redundancy", "aps", "chassis_id", "for_site", "model", "org_id", "site_id", "system_desc", "system_name", "timestamp", "vendor", "version")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.Adopted = temp.Adopted
    d.ApRedundancy = temp.ApRedundancy
    d.Aps = temp.Aps
    d.ChassisId = temp.ChassisId
    d.ForSite = temp.ForSite
    d.Model = temp.Model
    d.OrgId = temp.OrgId
    d.SiteId = temp.SiteId
    d.SystemDesc = temp.SystemDesc
    d.SystemName = temp.SystemName
    d.Timestamp = temp.Timestamp
    d.Vendor = temp.Vendor
    d.Version = temp.Version
    return nil
}

// discoveredSwitch is a temporary struct used for validating the fields of DiscoveredSwitch.
type discoveredSwitch  struct {
    Adopted      *bool                `json:"adopted,omitempty"`
    ApRedundancy *ApRedundancy        `json:"ap_redundancy,omitempty"`
    Aps          []DiscoveredSwitchAp `json:"aps,omitempty"`
    ChassisId    []string             `json:"chassis_id,omitempty"`
    ForSite      *bool                `json:"for_site,omitempty"`
    Model        *string              `json:"model,omitempty"`
    OrgId        *uuid.UUID           `json:"org_id,omitempty"`
    SiteId       *uuid.UUID           `json:"site_id,omitempty"`
    SystemDesc   *string              `json:"system_desc,omitempty"`
    SystemName   *string              `json:"system_name,omitempty"`
    Timestamp    *float64             `json:"timestamp,omitempty"`
    Vendor       *string              `json:"vendor,omitempty"`
    Version      *string              `json:"version,omitempty"`
}