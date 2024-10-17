package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// Zone represents a Zone struct.
// Zone
type Zone struct {
	CreatedTime *float64   `json:"created_time,omitempty"`
	ForSite     *bool      `json:"for_site,omitempty"`
	Id          *uuid.UUID `json:"id,omitempty"`
	// map where this zone is defined
	MapId        *uuid.UUID `json:"map_id,omitempty"`
	ModifiedTime *float64   `json:"modified_time,omitempty"`
	// The name of the zone
	Name   *string    `json:"name,omitempty"`
	OrgId  *uuid.UUID `json:"org_id,omitempty"`
	SiteId *uuid.UUID `json:"site_id,omitempty"`
	// vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area
	Vertices             []ZoneVertex   `json:"vertices,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Zone.
// It customizes the JSON marshaling process for Zone objects.
func (z Zone) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(z.toMap())
}

// toMap converts the Zone object to a map representation for JSON marshaling.
func (z Zone) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, z.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "created_time", "for_site", "id", "map_id", "modified_time", "name", "org_id", "site_id", "vertices")
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
type tempZone struct {
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
