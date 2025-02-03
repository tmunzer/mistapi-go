package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Zone represents a Zone struct.
// Zone
type Zone struct {
    // When the object has been created, in epoch
    CreatedTime          *float64               `json:"created_time,omitempty"`
    ForSite              *bool                  `json:"for_site,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    // Map where this zone is defined
    MapId                *uuid.UUID             `json:"map_id,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64               `json:"modified_time,omitempty"`
    // The name of the zone
    Name                 *string                `json:"name,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // Vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area
    Vertices             []ZoneVertex           `json:"vertices,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for Zone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (z Zone) String() string {
    return fmt.Sprintf(
    	"Zone[CreatedTime=%v, ForSite=%v, Id=%v, MapId=%v, ModifiedTime=%v, Name=%v, OrgId=%v, SiteId=%v, Vertices=%v, AdditionalProperties=%v]",
    	z.CreatedTime, z.ForSite, z.Id, z.MapId, z.ModifiedTime, z.Name, z.OrgId, z.SiteId, z.Vertices, z.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Zone.
// It customizes the JSON marshaling process for Zone objects.
func (z Zone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(z.AdditionalProperties,
        "created_time", "for_site", "id", "map_id", "modified_time", "name", "org_id", "site_id", "vertices"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(z.toMap())
}

// toMap converts the Zone object to a map representation for JSON marshaling.
func (z Zone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, z.AdditionalProperties)
    if z.CreatedTime != nil {
        structMap["created_time"] = z.CreatedTime
    }
    if z.ForSite != nil {
        structMap["for_site"] = z.ForSite
    }
    if z.Id != nil {
        structMap["id"] = z.Id
    }
    if z.MapId != nil {
        structMap["map_id"] = z.MapId
    }
    if z.ModifiedTime != nil {
        structMap["modified_time"] = z.ModifiedTime
    }
    if z.Name != nil {
        structMap["name"] = z.Name
    }
    if z.OrgId != nil {
        structMap["org_id"] = z.OrgId
    }
    if z.SiteId != nil {
        structMap["site_id"] = z.SiteId
    }
    if z.Vertices != nil {
        structMap["vertices"] = z.Vertices
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Zone.
// It customizes the JSON unmarshaling process for Zone objects.
func (z *Zone) UnmarshalJSON(input []byte) error {
    var temp tempZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "for_site", "id", "map_id", "modified_time", "name", "org_id", "site_id", "vertices")
    if err != nil {
    	return err
    }
    z.AdditionalProperties = additionalProperties
    
    z.CreatedTime = temp.CreatedTime
    z.ForSite = temp.ForSite
    z.Id = temp.Id
    z.MapId = temp.MapId
    z.ModifiedTime = temp.ModifiedTime
    z.Name = temp.Name
    z.OrgId = temp.OrgId
    z.SiteId = temp.SiteId
    z.Vertices = temp.Vertices
    return nil
}

// tempZone is a temporary struct used for validating the fields of Zone.
type tempZone  struct {
    CreatedTime  *float64     `json:"created_time,omitempty"`
    ForSite      *bool        `json:"for_site,omitempty"`
    Id           *uuid.UUID   `json:"id,omitempty"`
    MapId        *uuid.UUID   `json:"map_id,omitempty"`
    ModifiedTime *float64     `json:"modified_time,omitempty"`
    Name         *string      `json:"name,omitempty"`
    OrgId        *uuid.UUID   `json:"org_id,omitempty"`
    SiteId       *uuid.UUID   `json:"site_id,omitempty"`
    Vertices     []ZoneVertex `json:"vertices,omitempty"`
}
