package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SdkstatsWirelessClient represents a SdkstatsWirelessClient struct.
// SDK Client Details statistics
type SdkstatsWirelessClient struct {
    Id                   uuid.UUID                        `json:"id"`
    // last seen timestamp
    LastSeen             float64                          `json:"last_seen"`
    // map_id of the sdk client (if known), or null
    MapId                Optional[uuid.UUID]              `json:"map_id"`
    // name of the sdk client (if provided)
    Name                 *string                          `json:"name,omitempty"`
    // various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as
    NetworkConnection    *StatsSdkclientNetworkConnection `json:"network_connection,omitempty"`
    // uuid of the sdk client
    Uuid                 uuid.UUID                        `json:"uuid"`
    // list of beacon_id’s of the sdk client is in and since when (if known)
    Vbeacons             []SdkstatsWirelessClientVbeacon  `json:"vbeacons,omitempty"`
    // x (in pixels) of user location on the map (if known)
    X                    *float64                         `json:"x,omitempty"`
    // y (in pixels) of user location on the map (if known)
    Y                    *float64                         `json:"y,omitempty"`
    // list of zone_id’s of the sdk client is in and since when (if known)
    Zones                []SdkstatsWirelessClientZone     `json:"zones,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SdkstatsWirelessClient.
// It customizes the JSON marshaling process for SdkstatsWirelessClient objects.
func (s SdkstatsWirelessClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SdkstatsWirelessClient object to a map representation for JSON marshaling.
func (s SdkstatsWirelessClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["id"] = s.Id
    structMap["last_seen"] = s.LastSeen
    if s.MapId.IsValueSet() {
        if s.MapId.Value() != nil {
            structMap["map_id"] = s.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.NetworkConnection != nil {
        structMap["network_connection"] = s.NetworkConnection.toMap()
    }
    structMap["uuid"] = s.Uuid
    if s.Vbeacons != nil {
        structMap["vbeacons"] = s.Vbeacons
    }
    if s.X != nil {
        structMap["x"] = s.X
    }
    if s.Y != nil {
        structMap["y"] = s.Y
    }
    if s.Zones != nil {
        structMap["zones"] = s.Zones
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SdkstatsWirelessClient.
// It customizes the JSON unmarshaling process for SdkstatsWirelessClient objects.
func (s *SdkstatsWirelessClient) UnmarshalJSON(input []byte) error {
    var temp tempSdkstatsWirelessClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "last_seen", "map_id", "name", "network_connection", "uuid", "vbeacons", "x", "y", "zones")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Id = *temp.Id
    s.LastSeen = *temp.LastSeen
    s.MapId = temp.MapId
    s.Name = temp.Name
    s.NetworkConnection = temp.NetworkConnection
    s.Uuid = *temp.Uuid
    s.Vbeacons = temp.Vbeacons
    s.X = temp.X
    s.Y = temp.Y
    s.Zones = temp.Zones
    return nil
}

// tempSdkstatsWirelessClient is a temporary struct used for validating the fields of SdkstatsWirelessClient.
type tempSdkstatsWirelessClient  struct {
    Id                *uuid.UUID                       `json:"id"`
    LastSeen          *float64                         `json:"last_seen"`
    MapId             Optional[uuid.UUID]              `json:"map_id"`
    Name              *string                          `json:"name,omitempty"`
    NetworkConnection *StatsSdkclientNetworkConnection `json:"network_connection,omitempty"`
    Uuid              *uuid.UUID                       `json:"uuid"`
    Vbeacons          []SdkstatsWirelessClientVbeacon  `json:"vbeacons,omitempty"`
    X                 *float64                         `json:"x,omitempty"`
    Y                 *float64                         `json:"y,omitempty"`
    Zones             []SdkstatsWirelessClientZone     `json:"zones,omitempty"`
}

func (s *tempSdkstatsWirelessClient) validate() error {
    var errs []string
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `sdkstats_wireless_client`")
    }
    if s.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `sdkstats_wireless_client`")
    }
    if s.Uuid == nil {
        errs = append(errs, "required field `uuid` is missing for type `sdkstats_wireless_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
