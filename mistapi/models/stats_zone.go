package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsZone represents a StatsZone struct.
// Zone statistics
type StatsZone struct {
    // ble asset wait time right now
    AssetsWaits          *StatsZoneAssetsWaits     `json:"assets_waits,omitempty"`
    // client wait time right now
    ClientsWaits         *StatsZoneClientsWaits    `json:"clients_waits,omitempty"`
    CreatedTime          *float64                  `json:"created_time,omitempty"`
    // id of the zone
    Id                   uuid.UUID                 `json:"id"`
    // map_id of the zone
    MapId                uuid.UUID                 `json:"map_id"`
    ModifiedTime         *float64                  `json:"modified_time,omitempty"`
    // name of the zone
    Name                 string                    `json:"name"`
    // number of assets
    NumAssets            *int                      `json:"num_assets,omitempty"`
    // number of wifi clients (unconnected + connected)
    NumClients           *int                      `json:"num_clients,omitempty"`
    // number of sdk clients
    NumSdkclients        *int                      `json:"num_sdkclients,omitempty"`
    OccupancyLimit       *int                      `json:"occupancy_limit,omitempty"`
    OrgId                *uuid.UUID                `json:"org_id,omitempty"`
    // sdkclient wait time right now
    SdkclientsWaits      *StatsZoneSdkclientsWaits `json:"sdkclients_waits,omitempty"`
    SiteId               *uuid.UUID                `json:"site_id,omitempty"`
    // vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area
    Vertices             []ZoneVertex              `json:"vertices,omitempty"`
    VerticesM            []ZoneVertexM             `json:"vertices_m,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsZone.
// It customizes the JSON marshaling process for StatsZone objects.
func (s StatsZone) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsZone object to a map representation for JSON marshaling.
func (s StatsZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AssetsWaits != nil {
        structMap["assets_waits"] = s.AssetsWaits.toMap()
    }
    if s.ClientsWaits != nil {
        structMap["clients_waits"] = s.ClientsWaits.toMap()
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    structMap["id"] = s.Id
    structMap["map_id"] = s.MapId
    if s.ModifiedTime != nil {
        structMap["modified_time"] = s.ModifiedTime
    }
    structMap["name"] = s.Name
    if s.NumAssets != nil {
        structMap["num_assets"] = s.NumAssets
    }
    if s.NumClients != nil {
        structMap["num_clients"] = s.NumClients
    }
    if s.NumSdkclients != nil {
        structMap["num_sdkclients"] = s.NumSdkclients
    }
    if s.OccupancyLimit != nil {
        structMap["occupancy_limit"] = s.OccupancyLimit
    }
    if s.OrgId != nil {
        structMap["org_id"] = s.OrgId
    }
    if s.SdkclientsWaits != nil {
        structMap["sdkclients_waits"] = s.SdkclientsWaits.toMap()
    }
    if s.SiteId != nil {
        structMap["site_id"] = s.SiteId
    }
    if s.Vertices != nil {
        structMap["vertices"] = s.Vertices
    }
    if s.VerticesM != nil {
        structMap["vertices_m"] = s.VerticesM
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsZone.
// It customizes the JSON unmarshaling process for StatsZone objects.
func (s *StatsZone) UnmarshalJSON(input []byte) error {
    var temp tempStatsZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "assets_waits", "clients_waits", "created_time", "id", "map_id", "modified_time", "name", "num_assets", "num_clients", "num_sdkclients", "occupancy_limit", "org_id", "sdkclients_waits", "site_id", "vertices", "vertices_m")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AssetsWaits = temp.AssetsWaits
    s.ClientsWaits = temp.ClientsWaits
    s.CreatedTime = temp.CreatedTime
    s.Id = *temp.Id
    s.MapId = *temp.MapId
    s.ModifiedTime = temp.ModifiedTime
    s.Name = *temp.Name
    s.NumAssets = temp.NumAssets
    s.NumClients = temp.NumClients
    s.NumSdkclients = temp.NumSdkclients
    s.OccupancyLimit = temp.OccupancyLimit
    s.OrgId = temp.OrgId
    s.SdkclientsWaits = temp.SdkclientsWaits
    s.SiteId = temp.SiteId
    s.Vertices = temp.Vertices
    s.VerticesM = temp.VerticesM
    return nil
}

// tempStatsZone is a temporary struct used for validating the fields of StatsZone.
type tempStatsZone  struct {
    AssetsWaits     *StatsZoneAssetsWaits     `json:"assets_waits,omitempty"`
    ClientsWaits    *StatsZoneClientsWaits    `json:"clients_waits,omitempty"`
    CreatedTime     *float64                  `json:"created_time,omitempty"`
    Id              *uuid.UUID                `json:"id"`
    MapId           *uuid.UUID                `json:"map_id"`
    ModifiedTime    *float64                  `json:"modified_time,omitempty"`
    Name            *string                   `json:"name"`
    NumAssets       *int                      `json:"num_assets,omitempty"`
    NumClients      *int                      `json:"num_clients,omitempty"`
    NumSdkclients   *int                      `json:"num_sdkclients,omitempty"`
    OccupancyLimit  *int                      `json:"occupancy_limit,omitempty"`
    OrgId           *uuid.UUID                `json:"org_id,omitempty"`
    SdkclientsWaits *StatsZoneSdkclientsWaits `json:"sdkclients_waits,omitempty"`
    SiteId          *uuid.UUID                `json:"site_id,omitempty"`
    Vertices        []ZoneVertex              `json:"vertices,omitempty"`
    VerticesM       []ZoneVertexM             `json:"vertices_m,omitempty"`
}

func (s *tempStatsZone) validate() error {
    var errs []string
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `stats_zone`")
    }
    if s.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `stats_zone`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}