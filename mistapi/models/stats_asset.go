package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsAsset represents a StatsAsset struct.
// Asset statistics
type StatsAsset struct {
    // Battery voltage, in mV
    BatteryVoltage        *float64               `json:"battery_voltage,omitempty"`
    Beam                  *int                   `json:"beam,omitempty"`
    DeviceName            *string                `json:"device_name,omitempty"`
    Duration              *int                   `json:"duration,omitempty"`
    EddystoneUidInstance  *string                `json:"eddystone_uid_instance,omitempty"`
    EddystoneUidNamespace *string                `json:"eddystone_uid_namespace,omitempty"`
    EddystoneUrlUrl       *string                `json:"eddystone_url_url,omitempty"`
    IbeaconMajor          *int                   `json:"ibeacon_major,omitempty"`
    IbeaconMinor          *int                   `json:"ibeacon_minor,omitempty"`
    IbeaconUuid           *uuid.UUID             `json:"ibeacon_uuid,omitempty"`
    // Last seen timestamp
    LastSeen              *float64               `json:"last_seen,omitempty"`
    // Bluetooth MAC
    Mac                   string                 `json:"mac"`
    // Map where the device belongs to
    MapId                 *uuid.UUID             `json:"map_id,omitempty"`
    // Name / label of the device
    Name                  *string                `json:"name,omitempty"`
    Rssi                  *int                   `json:"rssi,omitempty"`
    // Only send this for individual asset stat
    Rssizones             []AssetRssiZone        `json:"rssizones,omitempty"`
    Temperature           *float64               `json:"temperature,omitempty"`
    // X in pixel
    X                     *float64               `json:"x,omitempty"`
    // Y in pixel
    Y                     *float64               `json:"y,omitempty"`
    // Only send this for individual asset stat
    Zones                 []AssetZone            `json:"zones,omitempty"`
    AdditionalProperties  map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsAsset,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsAsset) String() string {
    return fmt.Sprintf(
    	"StatsAsset[BatteryVoltage=%v, Beam=%v, DeviceName=%v, Duration=%v, EddystoneUidInstance=%v, EddystoneUidNamespace=%v, EddystoneUrlUrl=%v, IbeaconMajor=%v, IbeaconMinor=%v, IbeaconUuid=%v, LastSeen=%v, Mac=%v, MapId=%v, Name=%v, Rssi=%v, Rssizones=%v, Temperature=%v, X=%v, Y=%v, Zones=%v, AdditionalProperties=%v]",
    	s.BatteryVoltage, s.Beam, s.DeviceName, s.Duration, s.EddystoneUidInstance, s.EddystoneUidNamespace, s.EddystoneUrlUrl, s.IbeaconMajor, s.IbeaconMinor, s.IbeaconUuid, s.LastSeen, s.Mac, s.MapId, s.Name, s.Rssi, s.Rssizones, s.Temperature, s.X, s.Y, s.Zones, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsAsset.
// It customizes the JSON marshaling process for StatsAsset objects.
func (s StatsAsset) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "battery_voltage", "beam", "device_name", "duration", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "last_seen", "mac", "map_id", "name", "rssi", "rssizones", "temperature", "x", "y", "zones"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsAsset object to a map representation for JSON marshaling.
func (s StatsAsset) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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
    if s.Temperature != nil {
        structMap["temperature"] = s.Temperature
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
    var temp tempStatsAsset
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "battery_voltage", "beam", "device_name", "duration", "eddystone_uid_instance", "eddystone_uid_namespace", "eddystone_url_url", "ibeacon_major", "ibeacon_minor", "ibeacon_uuid", "last_seen", "mac", "map_id", "name", "rssi", "rssizones", "temperature", "x", "y", "zones")
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
    s.Temperature = temp.Temperature
    s.X = temp.X
    s.Y = temp.Y
    s.Zones = temp.Zones
    return nil
}

// tempStatsAsset is a temporary struct used for validating the fields of StatsAsset.
type tempStatsAsset  struct {
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
    Temperature           *float64        `json:"temperature,omitempty"`
    X                     *float64        `json:"x,omitempty"`
    Y                     *float64        `json:"y,omitempty"`
    Zones                 []AssetZone     `json:"zones,omitempty"`
}

func (s *tempStatsAsset) validate() error {
    var errs []string
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `stats_asset`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
