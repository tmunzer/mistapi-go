package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsAsset represents a StatsAsset struct.
// Asset statistics
type StatsAsset struct {
    // battery voltage, in mV
    BatteryVoltage        *float64        `json:"battery_voltage,omitempty"`
    Beam                  *int            `json:"beam,omitempty"`
    DeviceName            *string         `json:"device_name,omitempty"`
    Duration              *int            `json:"duration,omitempty"`
    EddystoneUidInstance  *string         `json:"eddystone_uid_instance,omitempty"`
    EddystoneUidNamespace *string         `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrlUrl       *string         `json:"eddystone_url_url,omitempty"`
    IbeaconMajor          *int            `json:"ibeacon_major,omitempty"`
    IbeaconMinor          *int            `json:"ibeacon_minor,omitempty"`
    IbeaconUuid           *uuid.UUID      `json:"ibeacon_uuid,omitempty"`
    // last seen timestamp
    LastSeen              *float64        `json:"last_seen,omitempty"`
    // bluetooth MAC
    Mac                   string          `json:"mac"`
    // map where the device belongs to
    MapId                 *uuid.UUID      `json:"map_id,omitempty"`
    // name / label of the device
    Name                  *string         `json:"name,omitempty"`
    Rssi                  *int            `json:"rssi,omitempty"`
    // only send this for individual asset stat
    Rssizones             []AssetRssiZone `json:"rssizones,omitempty"`
    Temperatur            *float64        `json:"temperatur,omitempty"`
    // x in pixel
    X                     *float64        `json:"x,omitempty"`
    // y in pixel
    Y                     *float64        `json:"y,omitempty"`
    // only send this for individual asset stat
    Zones                 []AssetZone     `json:"zones,omitempty"`
    AdditionalProperties  map[string]any  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsAsset.
// It customizes the JSON marshaling process for StatsAsset objects.
func (s StatsAsset) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsAsset object to a map representation for JSON marshaling.
func (s StatsAsset) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.BatteryVoltage != nil {
        structMap["battery_voltage"] = s.BatteryVoltage
    }
    if s.Beam != nil {
        structMap["beam"] = s.Beam
    }
    if s.DeviceName != nil {
        structMap["device_name"] = s.DeviceName
    }
    if s.Duration != nil {
        structMap["duration"] = s.Duration
    }
    if s.EddystoneUidInstance != nil {
        structMap["eddystone_uid_instance"] = s.EddystoneUidInstance
    }
    if s.EddystoneUidNamespace != nil {
        structMap["eddystone_uid_namespace"] = s.EddystoneUidNamespace
    }
    if s.EddystoneUrlUrl != nil {
        structMap["eddystone_url_url"] = s.EddystoneUrlUrl
    }
    if s.IbeaconMajor != nil {
        structMap["ibeacon_major"] = s.IbeaconMajor
    }
    if s.IbeaconMinor != nil {
        structMap["ibeacon_minor"] = s.IbeaconMinor
    }
    if s.IbeaconUuid != nil {
        structMap["ibeacon_uuid"] = s.IbeaconUuid
    }
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    }
    structMap["mac"] = s.Mac
    if s.MapId != nil {
        structMap["map_id"] = s.MapId
    }
    if s.Name != nil {
        structMap["name"] = s.Name
    }
    if s.Rssi != nil {
        structMap["rssi"] = s.Rssi
    }
    if s.Rssizones != nil {
        structMap["rssizones"] = s.Rssizones
    }
    if s.Temperatur != nil {
        structMap["temperatur"] = s.Temperatur
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

// UnmarshalJSON implements the json.Unmarshaler interface for StatsAsset.
// It customizes the JSON unmarshaling process for StatsAsset objects.
func (s *StatsAsset) UnmarshalJSON(input []byte) error {
    var temp statsAsset
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "battery_voltage", "beam", "device_name", "duration", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "last_seen", "mac", "map_id", "name", "rssi", "rssizones", "temperatur", "x", "y", "zones")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.BatteryVoltage = temp.BatteryVoltage
    s.Beam = temp.Beam
    s.DeviceName = temp.DeviceName
    s.Duration = temp.Duration
    s.EddystoneUidInstance = temp.EddystoneUidInstance
    s.EddystoneUidNamespace = temp.EddystoneUidNamespace
    s.EddystoneUrlUrl = temp.EddystoneUrlUrl
    s.IbeaconMajor = temp.IbeaconMajor
    s.IbeaconMinor = temp.IbeaconMinor
    s.IbeaconUuid = temp.IbeaconUuid
    s.LastSeen = temp.LastSeen
    s.Mac = *temp.Mac
    s.MapId = temp.MapId
    s.Name = temp.Name
    s.Rssi = temp.Rssi
    s.Rssizones = temp.Rssizones
    s.Temperatur = temp.Temperatur
    s.X = temp.X
    s.Y = temp.Y
    s.Zones = temp.Zones
    return nil
}

// statsAsset is a temporary struct used for validating the fields of StatsAsset.
type statsAsset  struct {
    BatteryVoltage        *float64        `json:"battery_voltage,omitempty"`
    Beam                  *int            `json:"beam,omitempty"`
    DeviceName            *string         `json:"device_name,omitempty"`
    Duration              *int            `json:"duration,omitempty"`
    EddystoneUidInstance  *string         `json:"eddystone_uid_instance,omitempty"`
    EddystoneUidNamespace *string         `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrlUrl       *string         `json:"eddystone_url_url,omitempty"`
    IbeaconMajor          *int            `json:"ibeacon_major,omitempty"`
    IbeaconMinor          *int            `json:"ibeacon_minor,omitempty"`
    IbeaconUuid           *uuid.UUID      `json:"ibeacon_uuid,omitempty"`
    LastSeen              *float64        `json:"last_seen,omitempty"`
    Mac                   *string         `json:"mac"`
    MapId                 *uuid.UUID      `json:"map_id,omitempty"`
    Name                  *string         `json:"name,omitempty"`
    Rssi                  *int            `json:"rssi,omitempty"`
    Rssizones             []AssetRssiZone `json:"rssizones,omitempty"`
    Temperatur            *float64        `json:"temperatur,omitempty"`
    X                     *float64        `json:"x,omitempty"`
    Y                     *float64        `json:"y,omitempty"`
    Zones                 []AssetZone     `json:"zones,omitempty"`
}

func (s *statsAsset) validate() error {
    var errs []string
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Stats_Asset`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
