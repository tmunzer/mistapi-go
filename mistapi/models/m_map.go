// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// Map represents a Map struct.
// Map
type Map struct {
    // When the object has been created, in epoch
    CreatedTime          *float64                 `json:"created_time,omitempty"`
    // Name/val pair objects for location engine to use
    Flags                map[string]int           `json:"flags,omitempty"`
    ForSite              *bool                    `json:"for_site,omitempty"`
    // When type=image, height of the image map
    Height               *int                     `json:"height,omitempty"`
    HeightM              *float64                 `json:"height_m,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID               `json:"id,omitempty"`
    // When type=google, latitude / longitude of the bottom-right corner
    LatlngBr             *LatlngBr                `json:"latlng_br,omitempty"`
    // When type=google, latitude / longitude of the top-left corner
    LatlngTl             *LatlngTl                `json:"latlng_tl,omitempty"`
    // Whether this map is considered locked down
    Locked               *bool                    `json:"locked,omitempty"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                 `json:"modified_time,omitempty"`
    // The name of the map
    Name                 *string                  `json:"name,omitempty"`
    OccupancyLimit       *int                     `json:"occupancy_limit,omitempty"`
    OrgId                *uuid.UUID               `json:"org_id,omitempty"`
    // Orientation of the map, 0 means up is north, 90 means up is west
    Orientation          *int                     `json:"orientation,omitempty"`
    // User-annotated X origin, pixels
    OriginX              *int                     `json:"origin_x,omitempty"`
    // User-annotated Y origin, pixels
    OriginY              *int                     `json:"origin_y,omitempty"`
    // When type=image, pixels per meter
    Ppm                  *float64                 `json:"ppm,omitempty"`
    SiteId               *uuid.UUID               `json:"site_id,omitempty"`
    // Sitesurvey_path
    SitesurveyPath       []MapSitesurveyPathItems `json:"sitesurvey_path,omitempty"`
    // When type=image, the url for the thumbnail image / preview
    ThumbnailUrl         *string                  `json:"thumbnail_url,omitempty"`
    // enum: `google`, `image`
    Type                 *MapTypeEnum             `json:"type,omitempty"`
    // When type=image, the url
    Url                  *string                  `json:"url,omitempty"`
    // if `type`==`google`. enum: `hybrid`, `roadmap`, `satellite`, `terrain`
    View                 Optional[MapViewEnum]    `json:"view"`
    // JSON blob for wall definition (same format as wayfinding_path)
    WallPath             *MapWallPath             `json:"wall_path,omitempty"`
    // Properties related to wayfinding
    Wayfinding           *MapWayfinding           `json:"wayfinding,omitempty"`
    // JSON blob for wayfinding (using Dijkstra’s algorithm)
    WayfindingPath       *MapWayfindingPath       `json:"wayfinding_path,omitempty"`
    // When type=image, width of the image map
    Width                *int                     `json:"width,omitempty"`
    WidthM               *float64                 `json:"width_m,omitempty"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// String implements the fmt.Stringer interface for Map,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m Map) String() string {
    return fmt.Sprintf(
    	"Map[CreatedTime=%v, Flags=%v, ForSite=%v, Height=%v, HeightM=%v, Id=%v, LatlngBr=%v, LatlngTl=%v, Locked=%v, ModifiedTime=%v, Name=%v, OccupancyLimit=%v, OrgId=%v, Orientation=%v, OriginX=%v, OriginY=%v, Ppm=%v, SiteId=%v, SitesurveyPath=%v, ThumbnailUrl=%v, Type=%v, Url=%v, View=%v, WallPath=%v, Wayfinding=%v, WayfindingPath=%v, Width=%v, WidthM=%v, AdditionalProperties=%v]",
    	m.CreatedTime, m.Flags, m.ForSite, m.Height, m.HeightM, m.Id, m.LatlngBr, m.LatlngTl, m.Locked, m.ModifiedTime, m.Name, m.OccupancyLimit, m.OrgId, m.Orientation, m.OriginX, m.OriginY, m.Ppm, m.SiteId, m.SitesurveyPath, m.ThumbnailUrl, m.Type, m.Url, m.View, m.WallPath, m.Wayfinding, m.WayfindingPath, m.Width, m.WidthM, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for Map.
// It customizes the JSON marshaling process for Map objects.
func (m Map) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "created_time", "flags", "for_site", "height", "height_m", "id", "latlng_br", "latlng_tl", "locked", "modified_time", "name", "occupancy_limit", "org_id", "orientation", "origin_x", "origin_y", "ppm", "site_id", "sitesurvey_path", "thumbnail_url", "type", "url", "view", "wall_path", "wayfinding", "wayfinding_path", "width", "width_m"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the Map object to a map representation for JSON marshaling.
func (m Map) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.CreatedTime != nil {
        structMap["created_time"] = m.CreatedTime
    }
    if m.Flags != nil {
        structMap["flags"] = m.Flags
    }
    if m.ForSite != nil {
        structMap["for_site"] = m.ForSite
    }
    if m.Height != nil {
        structMap["height"] = m.Height
    }
    if m.HeightM != nil {
        structMap["height_m"] = m.HeightM
    }
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.LatlngBr != nil {
        structMap["latlng_br"] = m.LatlngBr.toMap()
    }
    if m.LatlngTl != nil {
        structMap["latlng_tl"] = m.LatlngTl.toMap()
    }
    if m.Locked != nil {
        structMap["locked"] = m.Locked
    }
    if m.ModifiedTime != nil {
        structMap["modified_time"] = m.ModifiedTime
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.OccupancyLimit != nil {
        structMap["occupancy_limit"] = m.OccupancyLimit
    }
    if m.OrgId != nil {
        structMap["org_id"] = m.OrgId
    }
    if m.Orientation != nil {
        structMap["orientation"] = m.Orientation
    }
    if m.OriginX != nil {
        structMap["origin_x"] = m.OriginX
    }
    if m.OriginY != nil {
        structMap["origin_y"] = m.OriginY
    }
    if m.Ppm != nil {
        structMap["ppm"] = m.Ppm
    }
    if m.SiteId != nil {
        structMap["site_id"] = m.SiteId
    }
    if m.SitesurveyPath != nil {
        structMap["sitesurvey_path"] = m.SitesurveyPath
    }
    if m.ThumbnailUrl != nil {
        structMap["thumbnail_url"] = m.ThumbnailUrl
    }
    if m.Type != nil {
        structMap["type"] = m.Type
    }
    if m.Url != nil {
        structMap["url"] = m.Url
    }
    if m.View.IsValueSet() {
        if m.View.Value() != nil {
            structMap["view"] = m.View.Value()
        } else {
            structMap["view"] = nil
        }
    }
    if m.WallPath != nil {
        structMap["wall_path"] = m.WallPath.toMap()
    }
    if m.Wayfinding != nil {
        structMap["wayfinding"] = m.Wayfinding.toMap()
    }
    if m.WayfindingPath != nil {
        structMap["wayfinding_path"] = m.WayfindingPath.toMap()
    }
    if m.Width != nil {
        structMap["width"] = m.Width
    }
    if m.WidthM != nil {
        structMap["width_m"] = m.WidthM
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Map.
// It customizes the JSON unmarshaling process for Map objects.
func (m *Map) UnmarshalJSON(input []byte) error {
    var temp tempMap
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "created_time", "flags", "for_site", "height", "height_m", "id", "latlng_br", "latlng_tl", "locked", "modified_time", "name", "occupancy_limit", "org_id", "orientation", "origin_x", "origin_y", "ppm", "site_id", "sitesurvey_path", "thumbnail_url", "type", "url", "view", "wall_path", "wayfinding", "wayfinding_path", "width", "width_m")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.CreatedTime = temp.CreatedTime
    m.Flags = temp.Flags
    m.ForSite = temp.ForSite
    m.Height = temp.Height
    m.HeightM = temp.HeightM
    m.Id = temp.Id
    m.LatlngBr = temp.LatlngBr
    m.LatlngTl = temp.LatlngTl
    m.Locked = temp.Locked
    m.ModifiedTime = temp.ModifiedTime
    m.Name = temp.Name
    m.OccupancyLimit = temp.OccupancyLimit
    m.OrgId = temp.OrgId
    m.Orientation = temp.Orientation
    m.OriginX = temp.OriginX
    m.OriginY = temp.OriginY
    m.Ppm = temp.Ppm
    m.SiteId = temp.SiteId
    m.SitesurveyPath = temp.SitesurveyPath
    m.ThumbnailUrl = temp.ThumbnailUrl
    m.Type = temp.Type
    m.Url = temp.Url
    m.View = temp.View
    m.WallPath = temp.WallPath
    m.Wayfinding = temp.Wayfinding
    m.WayfindingPath = temp.WayfindingPath
    m.Width = temp.Width
    m.WidthM = temp.WidthM
    return nil
}

// tempMap is a temporary struct used for validating the fields of Map.
type tempMap  struct {
    CreatedTime    *float64                 `json:"created_time,omitempty"`
    Flags          map[string]int           `json:"flags,omitempty"`
    ForSite        *bool                    `json:"for_site,omitempty"`
    Height         *int                     `json:"height,omitempty"`
    HeightM        *float64                 `json:"height_m,omitempty"`
    Id             *uuid.UUID               `json:"id,omitempty"`
    LatlngBr       *LatlngBr                `json:"latlng_br,omitempty"`
    LatlngTl       *LatlngTl                `json:"latlng_tl,omitempty"`
    Locked         *bool                    `json:"locked,omitempty"`
    ModifiedTime   *float64                 `json:"modified_time,omitempty"`
    Name           *string                  `json:"name,omitempty"`
    OccupancyLimit *int                     `json:"occupancy_limit,omitempty"`
    OrgId          *uuid.UUID               `json:"org_id,omitempty"`
    Orientation    *int                     `json:"orientation,omitempty"`
    OriginX        *int                     `json:"origin_x,omitempty"`
    OriginY        *int                     `json:"origin_y,omitempty"`
    Ppm            *float64                 `json:"ppm,omitempty"`
    SiteId         *uuid.UUID               `json:"site_id,omitempty"`
    SitesurveyPath []MapSitesurveyPathItems `json:"sitesurvey_path,omitempty"`
    ThumbnailUrl   *string                  `json:"thumbnail_url,omitempty"`
    Type           *MapTypeEnum             `json:"type,omitempty"`
    Url            *string                  `json:"url,omitempty"`
    View           Optional[MapViewEnum]    `json:"view"`
    WallPath       *MapWallPath             `json:"wall_path,omitempty"`
    Wayfinding     *MapWayfinding           `json:"wayfinding,omitempty"`
    WayfindingPath *MapWayfindingPath       `json:"wayfinding_path,omitempty"`
    Width          *int                     `json:"width,omitempty"`
    WidthM         *float64                 `json:"width_m,omitempty"`
}
