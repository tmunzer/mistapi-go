package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// BeaconStatsItems represents a BeaconStatsItems struct.
type BeaconStatsItems struct {
    // battery voltage, in mV
    BatteryVoltage       *float64       `json:"battery_voltage,omitempty"`
    EddystoneInstance    *string        `json:"eddystone_instance,omitempty"`
    EddystoneNamespace   *string        `json:"eddystone_namespace,omitempty"`
    LastSeen             float64        `json:"last_seen"`
    Mac                  string         `json:"mac"`
    MapId                uuid.UUID      `json:"map_id"`
    Name                 string         `json:"name"`
    Power                int            `json:"power"`
    Type                 string         `json:"type"`
    X                    float64        `json:"x"`
    Y                    float64        `json:"y"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BeaconStatsItems.
// It customizes the JSON marshaling process for BeaconStatsItems objects.
func (b BeaconStatsItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(b.toMap())
}

// toMap converts the BeaconStatsItems object to a map representation for JSON marshaling.
func (b BeaconStatsItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, b.AdditionalProperties)
    if b.BatteryVoltage != nil {
        structMap["battery_voltage"] = b.BatteryVoltage
    }
    if b.EddystoneInstance != nil {
        structMap["eddystone_instance"] = b.EddystoneInstance
    }
    if b.EddystoneNamespace != nil {
        structMap["eddystone_namespace"] = b.EddystoneNamespace
    }
    structMap["last_seen"] = b.LastSeen
    structMap["mac"] = b.Mac
    structMap["map_id"] = b.MapId
    structMap["name"] = b.Name
    structMap["power"] = b.Power
    structMap["type"] = b.Type
    structMap["x"] = b.X
    structMap["y"] = b.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BeaconStatsItems.
// It customizes the JSON unmarshaling process for BeaconStatsItems objects.
func (b *BeaconStatsItems) UnmarshalJSON(input []byte) error {
    var temp beaconStatsItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "battery_voltage", "eddystone_instance", "eddystone_namespace", "last_seen", "mac", "map_id", "name", "power", "type", "x", "y")
    if err != nil {
    	return err
    }
    
    b.AdditionalProperties = additionalProperties
    b.BatteryVoltage = temp.BatteryVoltage
    b.EddystoneInstance = temp.EddystoneInstance
    b.EddystoneNamespace = temp.EddystoneNamespace
    b.LastSeen = *temp.LastSeen
    b.Mac = *temp.Mac
    b.MapId = *temp.MapId
    b.Name = *temp.Name
    b.Power = *temp.Power
    b.Type = *temp.Type
    b.X = *temp.X
    b.Y = *temp.Y
    return nil
}

// beaconStatsItems is a temporary struct used for validating the fields of BeaconStatsItems.
type beaconStatsItems  struct {
    BatteryVoltage     *float64   `json:"battery_voltage,omitempty"`
    EddystoneInstance  *string    `json:"eddystone_instance,omitempty"`
    EddystoneNamespace *string    `json:"eddystone_namespace,omitempty"`
    LastSeen           *float64   `json:"last_seen"`
    Mac                *string    `json:"mac"`
    MapId              *uuid.UUID `json:"map_id"`
    Name               *string    `json:"name"`
    Power              *int       `json:"power"`
    Type               *string    `json:"type"`
    X                  *float64   `json:"x"`
    Y                  *float64   `json:"y"`
}

func (b *beaconStatsItems) validate() error {
    var errs []string
    if b.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `Beacon_Stats_Items`")
    }
    if b.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Beacon_Stats_Items`")
    }
    if b.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `Beacon_Stats_Items`")
    }
    if b.Name == nil {
        errs = append(errs, "required field `name` is missing for type `Beacon_Stats_Items`")
    }
    if b.Power == nil {
        errs = append(errs, "required field `power` is missing for type `Beacon_Stats_Items`")
    }
    if b.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Beacon_Stats_Items`")
    }
    if b.X == nil {
        errs = append(errs, "required field `x` is missing for type `Beacon_Stats_Items`")
    }
    if b.Y == nil {
        errs = append(errs, "required field `y` is missing for type `Beacon_Stats_Items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
