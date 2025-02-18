package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsBeacon represents a StatsBeacon struct.
type StatsBeacon struct {
    // Battery voltage, in mV
    BatteryVoltage       *float64               `json:"battery_voltage,omitempty"`
    EddystoneInstance    *string                `json:"eddystone_instance,omitempty"`
    EddystoneNamespace   *string                `json:"eddystone_namespace,omitempty"`
    LastSeen             float64                `json:"last_seen"`
    Mac                  string                 `json:"mac"`
    MapId                uuid.UUID              `json:"map_id"`
    Name                 string                 `json:"name"`
    Power                int                    `json:"power"`
    Type                 string                 `json:"type"`
    X                    float64                `json:"x"`
    Y                    float64                `json:"y"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsBeacon,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsBeacon) String() string {
    return fmt.Sprintf(
    	"StatsBeacon[BatteryVoltage=%v, EddystoneInstance=%v, EddystoneNamespace=%v, LastSeen=%v, Mac=%v, MapId=%v, Name=%v, Power=%v, Type=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	s.BatteryVoltage, s.EddystoneInstance, s.EddystoneNamespace, s.LastSeen, s.Mac, s.MapId, s.Name, s.Power, s.Type, s.X, s.Y, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsBeacon.
// It customizes the JSON marshaling process for StatsBeacon objects.
func (s StatsBeacon) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "battery_voltage", "eddystone_instance", "eddystone_namespace", "last_seen", "mac", "map_id", "name", "power", "type", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsBeacon object to a map representation for JSON marshaling.
func (s StatsBeacon) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.BatteryVoltage != nil {
        structMap["battery_voltage"] = s.BatteryVoltage
    }
    if s.EddystoneInstance != nil {
        structMap["eddystone_instance"] = s.EddystoneInstance
    }
    if s.EddystoneNamespace != nil {
        structMap["eddystone_namespace"] = s.EddystoneNamespace
    }
    structMap["last_seen"] = s.LastSeen
    structMap["mac"] = s.Mac
    structMap["map_id"] = s.MapId
    structMap["name"] = s.Name
    structMap["power"] = s.Power
    structMap["type"] = s.Type
    structMap["x"] = s.X
    structMap["y"] = s.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsBeacon.
// It customizes the JSON unmarshaling process for StatsBeacon objects.
func (s *StatsBeacon) UnmarshalJSON(input []byte) error {
    var temp tempStatsBeacon
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "battery_voltage", "eddystone_instance", "eddystone_namespace", "last_seen", "mac", "map_id", "name", "power", "type", "x", "y")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.BatteryVoltage = temp.BatteryVoltage
    s.EddystoneInstance = temp.EddystoneInstance
    s.EddystoneNamespace = temp.EddystoneNamespace
    s.LastSeen = *temp.LastSeen
    s.Mac = *temp.Mac
    s.MapId = *temp.MapId
    s.Name = *temp.Name
    s.Power = *temp.Power
    s.Type = *temp.Type
    s.X = *temp.X
    s.Y = *temp.Y
    return nil
}

// tempStatsBeacon is a temporary struct used for validating the fields of StatsBeacon.
type tempStatsBeacon  struct {
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

func (s *tempStatsBeacon) validate() error {
    var errs []string
    if s.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `stats_beacon`")
    }
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `stats_beacon`")
    }
    if s.MapId == nil {
        errs = append(errs, "required field `map_id` is missing for type `stats_beacon`")
    }
    if s.Name == nil {
        errs = append(errs, "required field `name` is missing for type `stats_beacon`")
    }
    if s.Power == nil {
        errs = append(errs, "required field `power` is missing for type `stats_beacon`")
    }
    if s.Type == nil {
        errs = append(errs, "required field `type` is missing for type `stats_beacon`")
    }
    if s.X == nil {
        errs = append(errs, "required field `x` is missing for type `stats_beacon`")
    }
    if s.Y == nil {
        errs = append(errs, "required field `y` is missing for type `stats_beacon`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
