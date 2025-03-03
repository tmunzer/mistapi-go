package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsRssiZone represents a StatsRssiZone struct.
// Zone statistics
type StatsRssiZone struct {
    // BLE asset wait time right now
    AssetsWaits          *StatsZoneAssetsWaits     `json:"assets_waits,omitempty"`
    // Client wait time right now
    ClientsWaits         *StatsZoneClientsWaits    `json:"clients_waits,omitempty"`
    // When the object has been created, in epoch
    CreatedTime          *float64                  `json:"created_time,omitempty"`
    Devices              []StatsRssiZonesDevice    `json:"devices"`
    // Unique ID of the object instance in the Mist Organization
    Id                   uuid.UUID                 `json:"id"`
    // When the object has been modified for the last time, in epoch
    ModifiedTime         *float64                  `json:"modified_time,omitempty"`
    // Name of the zone
    Name                 string                    `json:"name"`
    // Number of assets
    NumAssets            *int                      `json:"num_assets,omitempty"`
    // Number of Wi-Fi clients (unconnected + connected)
    NumClients           *int                      `json:"num_clients,omitempty"`
    // Number of sdk clients
    NumSdkclients        *int                      `json:"num_sdkclients,omitempty"`
    OccupancyLimit       *int                      `json:"occupancy_limit,omitempty"`
    OrgId                *uuid.UUID                `json:"org_id,omitempty"`
    // SDK Clients wait time right now
    SdkclientsWaits      *StatsZoneSdkclientsWaits `json:"sdkclients_waits,omitempty"`
    SiteId               *uuid.UUID                `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{}    `json:"_"`
}

// String implements the fmt.Stringer interface for StatsRssiZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsRssiZone) String() string {
    return fmt.Sprintf(
    	"StatsRssiZone[AssetsWaits=%v, ClientsWaits=%v, CreatedTime=%v, Devices=%v, Id=%v, ModifiedTime=%v, Name=%v, NumAssets=%v, NumClients=%v, NumSdkclients=%v, OccupancyLimit=%v, OrgId=%v, SdkclientsWaits=%v, SiteId=%v, AdditionalProperties=%v]",
    	s.AssetsWaits, s.ClientsWaits, s.CreatedTime, s.Devices, s.Id, s.ModifiedTime, s.Name, s.NumAssets, s.NumClients, s.NumSdkclients, s.OccupancyLimit, s.OrgId, s.SdkclientsWaits, s.SiteId, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsRssiZone.
// It customizes the JSON marshaling process for StatsRssiZone objects.
func (s StatsRssiZone) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "assets_waits", "clients_waits", "created_time", "devices", "id", "modified_time", "name", "num_assets", "num_clients", "num_sdkclients", "occupancy_limit", "org_id", "sdkclients_waits", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsRssiZone object to a map representation for JSON marshaling.
func (s StatsRssiZone) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AssetsWaits != nil {
        structMap["assets_waits"] = s.AssetsWaits.toMap()
    }
    if s.ClientsWaits != nil {
        structMap["clients_waits"] = s.ClientsWaits.toMap()
    }
    if s.CreatedTime != nil {
        structMap["created_time"] = s.CreatedTime
    }
    structMap["devices"] = s.Devices
    structMap["id"] = s.Id
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsRssiZone.
// It customizes the JSON unmarshaling process for StatsRssiZone objects.
func (s *StatsRssiZone) UnmarshalJSON(input []byte) error {
    var temp tempStatsRssiZone
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "assets_waits", "clients_waits", "created_time", "devices", "id", "modified_time", "name", "num_assets", "num_clients", "num_sdkclients", "occupancy_limit", "org_id", "sdkclients_waits", "site_id")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AssetsWaits = temp.AssetsWaits
    s.ClientsWaits = temp.ClientsWaits
    s.CreatedTime = temp.CreatedTime
    s.Devices = *temp.Devices
    s.Id = *temp.Id
    s.ModifiedTime = temp.ModifiedTime
    s.Name = *temp.Name
    s.NumAssets = temp.NumAssets
    s.NumClients = temp.NumClients
    s.NumSdkclients = temp.NumSdkclients
    s.OccupancyLimit = temp.OccupancyLimit
    s.OrgId = temp.OrgId
    s.SdkclientsWaits = temp.SdkclientsWaits
    s.SiteId = temp.SiteId
    return nil
}

// tempStatsRssiZone is a temporary struct used for validating the fields of StatsRssiZone.
type tempStatsRssiZone  struct {
    AssetsWaits     *StatsZoneAssetsWaits     `json:"assets_waits,omitempty"`
    ClientsWaits    *StatsZoneClientsWaits    `json:"clients_waits,omitempty"`
    CreatedTime     *float64                  `json:"created_time,omitempty"`
    Devices         *[]StatsRssiZonesDevice   `json:"devices"`
    Id              *uuid.UUID                `json:"id"`
    ModifiedTime    *float64                  `json:"modified_time,omitempty"`
    Name            *string                   `json:"name"`
    NumAssets       *int                      `json:"num_assets,omitempty"`
    NumClients      *int                      `json:"num_clients,omitempty"`
    NumSdkclients   *int                      `json:"num_sdkclients,omitempty"`
    OccupancyLimit  *int                      `json:"occupancy_limit,omitempty"`
    OrgId           *uuid.UUID                `json:"org_id,omitempty"`
    SdkclientsWaits *StatsZoneSdkclientsWaits `json:"sdkclients_waits,omitempty"`
    SiteId          *uuid.UUID                `json:"site_id,omitempty"`
}

func (s *tempStatsRssiZone) validate() error {
    var errs []string
    if s.Devices == nil {
        errs = append(errs, "required field `devices` is missing for type `stats_rssi_zone`")
    }
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `stats_rssi_zone`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_rssi_zone`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
