package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// EvpnTopologyResponse represents a EvpnTopologyResponse struct.
type EvpnTopologyResponse struct {
    // when the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    // EVPN Options
    EvpnOptions          *EvpnOptions           `json:"evpn_options,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // when the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    Overwrite            *bool                  `json:"overwrite,omitempty"`
    // Property key is the pod number
    PodNames             map[string]string      `json:"pod_names,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopologyResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopologyResponse) String() string {
    return fmt.Sprintf(
    	"EvpnTopologyResponse[CreatedTime=%v, EvpnOptions=%v, Id=%v, ModifiedTime=%v, Name=%v, OrgId=%v, Overwrite=%v, PodNames=%v, SiteId=%v, AdditionalProperties=%v]",
    	e.CreatedTime, e.EvpnOptions, e.Id, e.ModifiedTime, e.Name, e.OrgId, e.Overwrite, e.PodNames, e.SiteId, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologyResponse.
// It customizes the JSON marshaling process for EvpnTopologyResponse objects.
func (e EvpnTopologyResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "created_time", "evpn_options", "id", "modified_time", "name", "org_id", "overwrite", "pod_names", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologyResponse object to a map representation for JSON marshaling.
func (e EvpnTopologyResponse) toMap() map[string]any {
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EvpnTopologyResponse.
// It customizes the JSON unmarshaling process for EvpnTopologyResponse objects.
func (e *EvpnTopologyResponse) UnmarshalJSON(input []byte) error {
    var temp tempEvpnTopologyResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "evpn_options", "id", "modified_time", "name", "org_id", "overwrite", "pod_names", "site_id")
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
    return nil
}

// tempEvpnTopologyResponse is a temporary struct used for validating the fields of EvpnTopologyResponse.
type tempEvpnTopologyResponse  struct {
    CreatedTime  *float64          `json:"created_time,omitempty"`
    EvpnOptions  *EvpnOptions      `json:"evpn_options,omitempty"`
    Id           *uuid.UUID        `json:"id,omitempty"`
    ModifiedTime *float64          `json:"modified_time,omitempty"`
    Name         *string           `json:"name,omitempty"`
    OrgId        *uuid.UUID        `json:"org_id,omitempty"`
    Overwrite    *bool             `json:"overwrite,omitempty"`
    PodNames     map[string]string `json:"pod_names,omitempty"`
    SiteId       *uuid.UUID        `json:"site_id,omitempty"`
}
