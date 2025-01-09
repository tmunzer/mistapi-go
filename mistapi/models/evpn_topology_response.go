package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// EvpnTopologyResponse represents a EvpnTopologyResponse struct.
type EvpnTopologyResponse struct {
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
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EvpnTopologyResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EvpnTopologyResponse) String() string {
    return fmt.Sprintf(
    	"EvpnTopologyResponse[EvpnOptions=%v, Id=%v, Name=%v, OrgId=%v, Overwrite=%v, PodNames=%v, SiteId=%v, AdditionalProperties=%v]",
    	e.EvpnOptions, e.Id, e.Name, e.OrgId, e.Overwrite, e.PodNames, e.SiteId, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EvpnTopologyResponse.
// It customizes the JSON marshaling process for EvpnTopologyResponse objects.
func (e EvpnTopologyResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "evpn_options", "id", "name", "org_id", "overwrite", "pod_names", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EvpnTopologyResponse object to a map representation for JSON marshaling.
func (e EvpnTopologyResponse) toMap() map[string]any {
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "evpn_options", "id", "name", "org_id", "overwrite", "pod_names", "site_id")
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
    return nil
}

// tempEvpnTopologyResponse is a temporary struct used for validating the fields of EvpnTopologyResponse.
type tempEvpnTopologyResponse  struct {
    EvpnOptions *EvpnOptions      `json:"evpn_options,omitempty"`
    Id          *uuid.UUID        `json:"id,omitempty"`
    Name        *string           `json:"name,omitempty"`
    OrgId       *uuid.UUID        `json:"org_id,omitempty"`
    Overwrite   *bool             `json:"overwrite,omitempty"`
    PodNames    map[string]string `json:"pod_names,omitempty"`
    SiteId      *uuid.UUID        `json:"site_id,omitempty"`
}
