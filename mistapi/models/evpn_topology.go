// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// EvpnTopology represents a EvpnTopology struct.
type EvpnTopology struct {
    // When the object has been created, in epoch
    CreatedTime          *float64                            `json:"created_time,omitempty"`
    // EVPN Options
    EvpnOptions          *EvpnOptions                        `json:"evpn_options,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID                          `json:"id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                            `json:"modified_time,omitempty"`
    Name                 *string                             `json:"name,omitempty"`
    OrgId                *uuid.UUID                          `json:"org_id,omitempty"`
    Overwrite            *bool                               `json:"overwrite,omitempty"`
    // Property key is the pod number
    PodNames             map[string]string                   `json:"pod_names,omitempty"`
    SiteId               *uuid.UUID                          `json:"site_id,omitempty"`
    // Property key is the switch mac
    SwitchConfigs        map[string]EvpnTopologySwitchConfig `json:"switch_configs,omitempty"`
    Switches             []EvpnTopologySwitch                `json:"switches"`
    AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopology,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopology) String() string {
    return fmt.Sprintf(
    	"EvpnTopology[CreatedTime=%v, EvpnOptions=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Overwrite=%v, PodNames=%v, SiteId=%v, SwitchConfigs=%v, Switches=%v, AdditionalProperties=%v]",
    	e.CreatedTime, e.EvpnOptions, e.Id, e.ModifiedTime, e.Name, e.OrgId, e.Overwrite, e.PodNames, e.SiteId, e.SwitchConfigs, e.Switches, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopology.
// It customizes the JSON marshaling process for EvpnTopology objects.
func (e EvpnTopology) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "created_time", "evpn_options", "id", "modified_time", "name", "org_id", "overwrite", "pod_names", "site_id", "switch_configs", "switches"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopology object to a map representation for JSON marshaling.
func (e EvpnTopology) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.CreatedTime != nil {
        structMap["created_time"] = e.CreatedTime
    }
    if e.EvpnOptions != nil {
        structMap["evpn_options"] = e.EvpnOptions.toMap()
    }
    if e.Id != nil {
        structMap["id"] = e.Id
    }
    if e.ModifiedTime != nil {
        structMap["modified_time"] = e.ModifiedTime
    }
    if e.Name != nil {
        structMap["name"] = e.Name
    }
    if e.OrgId != nil {
        structMap["org_id"] = e.OrgId
    }
    if e.Overwrite != nil {
        structMap["overwrite"] = e.Overwrite
    }
    if e.PodNames != nil {
        structMap["pod_names"] = e.PodNames
    }
    if e.SiteId != nil {
        structMap["site_id"] = e.SiteId
    }
    if e.SwitchConfigs != nil {
        structMap["switch_configs"] = e.SwitchConfigs
    }
    structMap["switches"] = e.Switches
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopology.
// It customizes the JSON unmarshaling process for EvpnTopology objects.
func (e *EvpnTopology) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopology
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "evpn_options", "id", "modified_time", "name", "org_id", "overwrite", "pod_names", "site_id", "switch_configs", "switches")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.CreatedTime = temp.CreatedTime
    e.EvpnOptions = temp.EvpnOptions
    e.Id = temp.Id
    e.ModifiedTime = temp.ModifiedTime
    e.Name = temp.Name
    e.OrgId = temp.OrgId
    e.Overwrite = temp.Overwrite
    e.PodNames = temp.PodNames
    e.SiteId = temp.SiteId
    e.SwitchConfigs = temp.SwitchConfigs
    e.Switches = *temp.Switches
    return nil
}

// tempEvpnTopology is a temporary struct used for validating the fields of EvpnTopology.
type tempEvpnTopology  struct {
    CreatedTime   *float64                            `json:"created_time,omitempty"`
    EvpnOptions   *EvpnOptions                        `json:"evpn_options,omitempty"`
    Id            *uuid.UUID                          `json:"id,omitempty"`
    ModifiedTime  *float64                            `json:"modified_time,omitempty"`
    Name          *string                             `json:"name,omitempty"`
    OrgId         *uuid.UUID                          `json:"org_id,omitempty"`
    Overwrite     *bool                               `json:"overwrite,omitempty"`
    PodNames      map[string]string                   `json:"pod_names,omitempty"`
    SiteId        *uuid.UUID                          `json:"site_id,omitempty"`
    SwitchConfigs map[string]EvpnTopologySwitchConfig `json:"switch_configs,omitempty"`
    Switches      *[]EvpnTopologySwitch               `json:"switches"`
}

func (e *tempEvpnTopology) validate() error {
    var errs []string
    if e.Switches == nil {
        errs = append(errs, "required field `switches` is missing for type `evpn_topology`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
