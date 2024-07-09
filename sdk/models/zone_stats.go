package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ZoneStats represents a ZoneStats struct.
// Zone statistics
type ZoneStats struct {
    // ble asset wait time right now
    AssetsWaits          *ZoneStatsAssetsWaits     `json:"assets_waits,omitempty"`
    // client wait time right now
    ClientsWaits         *ZoneStatsClientsWaits    `json:"clients_waits,omitempty"`
    CreatedTime          *int                      `json:"created_time,omitempty"`
    // id of the zone
    Id                   uuid.UUID                 `json:"id"`
    // map_id of the zone
    MapId                uuid.UUID                 `json:"map_id"`
    ModifiedTime         *int                      `json:"modified_time,omitempty"`
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
    SdkclientsWaits      *ZoneStatsSdkclientsWaits `json:"sdkclients_waits,omitempty"`
    SiteId               *uuid.UUID                `json:"site_id,omitempty"`
    // vertices used to define an area. It’s assumed that the last point connects to the first point and forms an closed area
    Vertices             []ZoneVertex              `json:"vertices,omitempty"`
    VerticesM            []ZoneVertexM             `json:"vertices_m,omitempty"`
    AdditionalProperties map[string]any            `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ZoneStats.
// It customizes the JSON marshaling process for ZoneStats objects.
func (z ZoneStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(z.toMap())
}

// toMap converts the ZoneStats object to a map representation for JSON marshaling.
func (z ZoneStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, z.AdditionalProperties)
    if z.AssetsWaits != nil {
        structMap["assets_waits"] = z.AssetsWaits.toMap()
    }
    if z.ClientsWaits != nil {
        structMap["clients_waits"] = z.ClientsWaits.toMap()
    }
    if z.CreatedTime != nil {
        structMap["created_time"] = z.CreatedTime
    }
    structMap["id"] = z.Id
    structMap["map_id"] = z.MapId
    if z.ModifiedTime != nil {
        structMap["modified_time"] = z.ModifiedTime
    }
    structMap["name"] = z.Name
    if z.NumAssets != nil {
        structMap["num_assets"] = z.NumAssets
    }
    if z.NumClients != nil {
        structMap["num_clients"] = z.NumClients
    }
    if z.NumSdkclients != nil {
        structMap["num_sdkclients"] = z.NumSdkclients
    }
    if z.OccupancyLimit != nil {
        structMap["occupancy_limit"] = z.OccupancyLimit
    }
    if z.OrgId != nil {
        structMap["org_id"] = z.OrgId
    }
    if z.SdkclientsWaits != nil {
        structMap["sdkclients_waits"] = z.SdkclientsWaits.toMap()
    }
    if z.SiteId != nil {
        structMap["site_id"] = z.SiteId
    }
    if z.Vertices != nil {
        structMap["vertices"] = z.Vertices
    }
    if z.VerticesM != nil {
        structMap["vertices_m"] = z.VerticesM
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ZoneStats.
// It customizes the JSON unmarshaling process for ZoneStats objects.
func (z *ZoneStats) UnmarshalJSON(input []byte) error {
    var temp zoneStats
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
    
    z.AdditionalProperties = additionalProperties
    z.AssetsWaits = temp.AssetsWaits
    z.ClientsWaits = temp.ClientsWaits
    z.CreatedTime = temp.CreatedTime
    z.Id = *temp.Id
    z.MapId = *temp.MapId
    z.ModifiedTime = temp.ModifiedTime
    z.Name = *temp.Name
    z.NumAssets = temp.NumAssets
    z.NumClients = temp.NumClients
    z.NumSdkclients = temp.NumSdkclients
    z.OccupancyLimit = temp.OccupancyLimit
    z.OrgId = temp.OrgId
    z.SdkclientsWaits = temp.SdkclientsWaits
    z.SiteId = temp.SiteId
    z.Vertices = temp.Vertices
    z.VerticesM = temp.VerticesM
    return nil
}

// zoneStats is a temporary struct used for validating the fields of ZoneStats.
type zoneStats  struct {
    AssetsWaits     *ZoneStatsAssetsWaits     `json:"assets_waits,omitempty"`
    ClientsWaits    *ZoneStatsClientsWaits    `json:"clients_waits,omitempty"`
    CreatedTime     *int                      `json:"created_time,omitempty"`
    Id              *uuid.UUID                `json:"id"`
    MapId           *uuid.UUID                `json:"map_id"`
    ModifiedTime    *int                      `json:"modified_time,omitempty"`
    Name            *string                   `json:"name"`
    NumAssets       *int                      `json:"num_assets,omitempty"`
    NumClients      *int                      `json:"num_clients,omitempty"`
    NumSdkclients   *int                      `json:"num_sdkclients,omitempty"`
    OccupancyLimit  *int                      `json:"occupancy_limit,omitempty"`
    OrgId           *uuid.UUID                `json:"org_id,omitempty"`
    SdkclientsWaits *ZoneStatsSdkclientsWaits `json:"sdkclients_waits,omitempty"`
    SiteId          *uuid.UUID                `json:"site_id,omitempty"`
    Vertices        []ZoneVertex              `json:"vertices,omitempty"`
    VerticesM       []ZoneVertexM             `json:"vertices_m,omitempty"`
}

func (z *zoneStats) validate() error {
    var errs []string
    if z.Id == nil {
        errs = append(errs, "required field `id` is missing for type `Zone_Stats`")
    }
    if z.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Zone_Stats`")
    }
    if z.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Zone_Stats`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
