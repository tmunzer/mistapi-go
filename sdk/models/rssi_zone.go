package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// RssiZone represents a RssiZone struct.
// RSSI Zone
type RssiZone struct {
    CreatedTime          *float64         `json:"created_time,omitempty"`
    // List of devices and the respective RSSI values to be considered in the zone
    Devices              []RssiZoneDevice `json:"devices"`
    ForSite              *bool            `json:"for_site,omitempty"`
    Id                   *uuid.UUID       `json:"id,omitempty"`
    ModifiedTime         *float64         `json:"modified_time,omitempty"`
    // The name of the zone
    Name                 *string          `json:"name,omitempty"`
    OrgId                *uuid.UUID       `json:"org_id,omitempty"`
    SiteId               *uuid.UUID       `json:"site_id,omitempty"`
    AdditionalProperties map[string]any   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RssiZone.
// It customizes the JSON marshaling process for RssiZone objects.
func (r RssiZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RssiZone object to a map representation for JSON marshaling.
func (r RssiZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    if r.CreatedTime != nil {
        structMap["created_time"] = r.CreatedTime
    }
    structMap["devices"] = r.Devices
    if r.ForSite != nil {
        structMap["for_site"] = r.ForSite
    }
    if r.Id != nil {
        structMap["id"] = r.Id
    }
    if r.ModifiedTime != nil {
        structMap["modified_time"] = r.ModifiedTime
    }
    if r.Name != nil {
        structMap["name"] = r.Name
    }
    if r.OrgId != nil {
        structMap["org_id"] = r.OrgId
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RssiZone.
// It customizes the JSON unmarshaling process for RssiZone objects.
func (r *RssiZone) UnmarshalJSON(input []byte) error {
    var temp rssiZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "devices", "for_site", "id", "modified_time", "name", "org_id", "site_id")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.CreatedTime = temp.CreatedTime
    r.Devices = *temp.Devices
    r.ForSite = temp.ForSite
    r.Id = temp.Id
    r.ModifiedTime = temp.ModifiedTime
    r.Name = temp.Name
    r.OrgId = temp.OrgId
    r.SiteId = temp.SiteId
    return nil
}

// rssiZone is a temporary struct used for validating the fields of RssiZone.
type rssiZone  struct {
    CreatedTime  *float64          `json:"created_time,omitempty"`
    Devices      *[]RssiZoneDevice `json:"devices"`
    ForSite      *bool             `json:"for_site,omitempty"`
    Id           *uuid.UUID        `json:"id,omitempty"`
    ModifiedTime *float64          `json:"modified_time,omitempty"`
    Name         *string           `json:"name,omitempty"`
    OrgId        *uuid.UUID        `json:"org_id,omitempty"`
    SiteId       *uuid.UUID        `json:"site_id,omitempty"`
}

func (r *rssiZone) validate() error {
    var errs []string
    if r.Devices == nil {
        errs = append(errs, "required field `devices` is missing for type `Rssi_Zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
