// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
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
	AssetsWait *StatsZoneAssetsWaits `json:"assets_wait,omitempty"`
	// Client wait time right now
	ClientsWait *StatsZoneClientsWaits `json:"clients_wait,omitempty"`
	// When the object has been created, in epoch
	CreatedTime *float64               `json:"created_time,omitempty"`
	Devices     []StatsRssiZonesDevice `json:"devices"`
	// Discovered asset wait time right now
	DiscoveredAssetsWait *StatsZoneDiscoveredAssetsWaits `json:"discovered_assets_wait,omitempty"`
	// Unique ID of the object instance in the Mist Organization
	Id uuid.UUID `json:"id"`
	// When the object has been modified for the last time, in epoch
	ModifiedTime *float64 `json:"modified_time,omitempty"`
	// Name of the zone
	Name string `json:"name"`
	// Number of assets
	NumAssets *int `json:"num_assets,omitempty"`
	// Number of Wi-Fi clients (unconnected + connected)
	NumClients *int `json:"num_clients,omitempty"`
	// Number of discoveredassets
	NumDiscoveredAssets *int `json:"num_discovered_assets,omitempty"`
	// Number of sdk clients
	NumSdkclients *int `json:"num_sdkclients,omitempty"`
	// Number of unconnected Wi-Fi clients
	NumUnconnectedClients *int       `json:"num_unconnected_clients,omitempty"`
	OccupancyLimit        *int       `json:"occupancy_limit,omitempty"`
	OrgId                 *uuid.UUID `json:"org_id,omitempty"`
	// SDK client wait time right now
	SdkclientsWait *StatsZoneSdkclientsWaits `json:"sdkclients_wait,omitempty"`
	SiteId         *uuid.UUID                `json:"site_id,omitempty"`
	// Unconnected Wi-Fi client wait time right now
	UnconnectedClientsWait *StatsZoneUnconnectedClientsWaits `json:"unconnected_clients_wait,omitempty"`
	AdditionalProperties   map[string]interface{}            `json:"_"`
}

// String implements the fmt.Stringer interface for StatsRssiZone,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsRssiZone) String() string {
	return fmt.Sprintf(
		"StatsRssiZone[AssetsWait=%v, ClientsWait=%v, CreatedTime=%v, Devices=%v, DiscoveredAssetsWait=%v, Id=%v, ModifiedTime=%v, Name=%v, NumAssets=%v, NumClients=%v, NumDiscoveredAssets=%v, NumSdkclients=%v, NumUnconnectedClients=%v, OccupancyLimit=%v, OrgId=%v, SdkclientsWait=%v, SiteId=%v, UnconnectedClientsWait=%v, AdditionalProperties=%v]",
		s.AssetsWait, s.ClientsWait, s.CreatedTime, s.Devices, s.DiscoveredAssetsWait, s.Id, s.ModifiedTime, s.Name, s.NumAssets, s.NumClients, s.NumDiscoveredAssets, s.NumSdkclients, s.NumUnconnectedClients, s.OccupancyLimit, s.OrgId, s.SdkclientsWait, s.SiteId, s.UnconnectedClientsWait, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsRssiZone.
// It customizes the JSON marshaling process for StatsRssiZone objects.
func (s StatsRssiZone) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"assets_wait", "clients_wait", "created_time", "devices", "discovered_assets_wait", "id", "modified_time", "name", "num_assets", "num_clients", "num_discovered_assets", "num_sdkclients", "num_unconnected_clients", "occupancy_limit", "org_id", "sdkclients_wait", "site_id", "unconnected_clients_wait"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsRssiZone object to a map representation for JSON marshaling.
func (s StatsRssiZone) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.AssetsWait != nil {
		structMap["assets_wait"] = s.AssetsWait.toMap()
	}
	if s.ClientsWait != nil {
		structMap["clients_wait"] = s.ClientsWait.toMap()
	}
	if s.CreatedTime != nil {
		structMap["created_time"] = s.CreatedTime
	}
	structMap["devices"] = s.Devices
	if s.DiscoveredAssetsWait != nil {
		structMap["discovered_assets_wait"] = s.DiscoveredAssetsWait.toMap()
	}
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
	if s.NumDiscoveredAssets != nil {
		structMap["num_discovered_assets"] = s.NumDiscoveredAssets
	}
	if s.NumSdkclients != nil {
		structMap["num_sdkclients"] = s.NumSdkclients
	}
	if s.NumUnconnectedClients != nil {
		structMap["num_unconnected_clients"] = s.NumUnconnectedClients
	}
	if s.OccupancyLimit != nil {
		structMap["occupancy_limit"] = s.OccupancyLimit
	}
	if s.OrgId != nil {
		structMap["org_id"] = s.OrgId
	}
	if s.SdkclientsWait != nil {
		structMap["sdkclients_wait"] = s.SdkclientsWait.toMap()
	}
	if s.SiteId != nil {
		structMap["site_id"] = s.SiteId
	}
	if s.UnconnectedClientsWait != nil {
		structMap["unconnected_clients_wait"] = s.UnconnectedClientsWait.toMap()
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
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "assets_wait", "clients_wait", "created_time", "devices", "discovered_assets_wait", "id", "modified_time", "name", "num_assets", "num_clients", "num_discovered_assets", "num_sdkclients", "num_unconnected_clients", "occupancy_limit", "org_id", "sdkclients_wait", "site_id", "unconnected_clients_wait")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.AssetsWait = temp.AssetsWait
	s.ClientsWait = temp.ClientsWait
	s.CreatedTime = temp.CreatedTime
	s.Devices = *temp.Devices
	s.DiscoveredAssetsWait = temp.DiscoveredAssetsWait
	s.Id = *temp.Id
	s.ModifiedTime = temp.ModifiedTime
	s.Name = *temp.Name
	s.NumAssets = temp.NumAssets
	s.NumClients = temp.NumClients
	s.NumDiscoveredAssets = temp.NumDiscoveredAssets
	s.NumSdkclients = temp.NumSdkclients
	s.NumUnconnectedClients = temp.NumUnconnectedClients
	s.OccupancyLimit = temp.OccupancyLimit
	s.OrgId = temp.OrgId
	s.SdkclientsWait = temp.SdkclientsWait
	s.SiteId = temp.SiteId
	s.UnconnectedClientsWait = temp.UnconnectedClientsWait
	return nil
}

// tempStatsRssiZone is a temporary struct used for validating the fields of StatsRssiZone.
type tempStatsRssiZone struct {
	AssetsWait             *StatsZoneAssetsWaits             `json:"assets_wait,omitempty"`
	ClientsWait            *StatsZoneClientsWaits            `json:"clients_wait,omitempty"`
	CreatedTime            *float64                          `json:"created_time,omitempty"`
	Devices                *[]StatsRssiZonesDevice           `json:"devices"`
	DiscoveredAssetsWait   *StatsZoneDiscoveredAssetsWaits   `json:"discovered_assets_wait,omitempty"`
	Id                     *uuid.UUID                        `json:"id"`
	ModifiedTime           *float64                          `json:"modified_time,omitempty"`
	Name                   *string                           `json:"name"`
	NumAssets              *int                              `json:"num_assets,omitempty"`
	NumClients             *int                              `json:"num_clients,omitempty"`
	NumDiscoveredAssets    *int                              `json:"num_discovered_assets,omitempty"`
	NumSdkclients          *int                              `json:"num_sdkclients,omitempty"`
	NumUnconnectedClients  *int                              `json:"num_unconnected_clients,omitempty"`
	OccupancyLimit         *int                              `json:"occupancy_limit,omitempty"`
	OrgId                  *uuid.UUID                        `json:"org_id,omitempty"`
	SdkclientsWait         *StatsZoneSdkclientsWaits         `json:"sdkclients_wait,omitempty"`
	SiteId                 *uuid.UUID                        `json:"site_id,omitempty"`
	UnconnectedClientsWait *StatsZoneUnconnectedClientsWaits `json:"unconnected_clients_wait,omitempty"`
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
	return errors.New(strings.Join(errs, "\n"))
}
