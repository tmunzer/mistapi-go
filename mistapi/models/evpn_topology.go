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
    // EVPN Options
    EvpnOptions          *EvpnOptions           `json:"evpn_options,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Overwrite            *bool                  `json:"overwrite,omitempty"`
    // Property key is the pod number
    PodNames             map[string]string      `json:"pod_names,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    Switches             []EvpnTopologySwitch   `json:"switches"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopology,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopology) String() string {
    return fmt.Sprintf(
    	"EvpnTopology[EvpnOptions=%v, Id=%v, Name=%v, OrgId=%v, Overwrite=%v, PodNames=%v, SiteId=%v, Switches=%v, AdditionalProperties=%v]",
    	e.EvpnOptions, e.Id, e.Name, e.OrgId, e.Overwrite, e.PodNames, e.SiteId, e.Switches, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopology.
// It customizes the JSON marshaling process for EvpnTopology objects.
func (e EvpnTopology) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "evpn_options", "id", "name", "org_id", "overwrite", "pod_names", "site_id", "switches"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopology object to a map representation for JSON marshaling.
func (e EvpnTopology) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    if e.EvpnOptions != nil {
        structMap["evpn_options"] = e.EvpnOptions.toMap()
    }
    if e.Id != nil {
        structMap["id"] = e.Id
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "evpn_options", "id", "name", "org_id", "overwrite", "pod_names", "site_id", "switches")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.EvpnOptions = temp.EvpnOptions
    e.Id = temp.Id
    e.Name = temp.Name
    e.OrgId = temp.OrgId
    e.Overwrite = temp.Overwrite
    e.PodNames = temp.PodNames
    e.SiteId = temp.SiteId
    e.Switches = *temp.Switches
    return nil
}

// tempEvpnTopology is a temporary struct used for validating the fields of EvpnTopology.
type tempEvpnTopology  struct {
    EvpnOptions *EvpnOptions          `json:"evpn_options,omitempty"`
    Id          *uuid.UUID            `json:"id,omitempty"`
    Name        *string               `json:"name,omitempty"`
    OrgId       *uuid.UUID            `json:"org_id,omitempty"`
    Overwrite   *bool                 `json:"overwrite,omitempty"`
    PodNames    map[string]string     `json:"pod_names,omitempty"`
    SiteId      *uuid.UUID            `json:"site_id,omitempty"`
    Switches    *[]EvpnTopologySwitch `json:"switches"`
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
