package models

import (
    "encoding/json"
    "fmt"
)

// StatsApGpsStat represents a StatsApGpsStat struct.
type StatsApGpsStat struct {
    // The estimated accuracy or accuracy of the GPS coordinates, measured in meters.
    Accuracy             *float64               `json:"accuracy,omitempty"`
    // The elevation of the AP above sea level, measured in meters.
    Altitude             *float64               `json:"altitude,omitempty"`
    // The geographic latitude of the AP, measured in degrees.
    Latitude             *float64               `json:"latitude,omitempty"`
    // The geographic longitude of the AP, measured in degrees.
    Longitude            *float64               `json:"longitude,omitempty"`
    // The origin of the GPS data. enum:
    // * `gps`: from this device’s GPS estimates
    // * `other_ap` from neighboring device GPS estimates
    Src                  *StatsApGpsStatSrcEnum `json:"src,omitempty"`
    // The unix timestamp when the GPS data was recorded.
    Timestamp            *float64               `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsApGpsStat,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsApGpsStat) String() string {
    return fmt.Sprintf(
    	"StatsApGpsStat[Accuracy=%v, Altitude=%v, Latitude=%v, Longitude=%v, Src=%v, Timestamp=%v, AdditionalProperties=%v]",
    	s.Accuracy, s.Altitude, s.Latitude, s.Longitude, s.Src, s.Timestamp, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsApGpsStat.
// It customizes the JSON marshaling process for StatsApGpsStat objects.
func (s StatsApGpsStat) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "accuracy", "altitude", "latitude", "longitude", "src", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsApGpsStat object to a map representation for JSON marshaling.
func (s StatsApGpsStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Accuracy != nil {
        structMap["accuracy"] = s.Accuracy
    }
    if s.Altitude != nil {
        structMap["altitude"] = s.Altitude
    }
    if s.Latitude != nil {
        structMap["latitude"] = s.Latitude
    }
    if s.Longitude != nil {
        structMap["longitude"] = s.Longitude
    }
    if s.Src != nil {
        structMap["src"] = s.Src
    }
    if s.Timestamp != nil {
        structMap["timestamp"] = s.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsApGpsStat.
// It customizes the JSON unmarshaling process for StatsApGpsStat objects.
func (s *StatsApGpsStat) UnmarshalJSON(input []byte) error {
    var temp tempStatsApGpsStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accuracy", "altitude", "latitude", "longitude", "src", "timestamp")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Accuracy = temp.Accuracy
    s.Altitude = temp.Altitude
    s.Latitude = temp.Latitude
    s.Longitude = temp.Longitude
    s.Src = temp.Src
    s.Timestamp = temp.Timestamp
    return nil
}

// tempStatsApGpsStat is a temporary struct used for validating the fields of StatsApGpsStat.
type tempStatsApGpsStat  struct {
    Accuracy  *float64               `json:"accuracy,omitempty"`
    Altitude  *float64               `json:"altitude,omitempty"`
    Latitude  *float64               `json:"latitude,omitempty"`
    Longitude *float64               `json:"longitude,omitempty"`
    Src       *StatsApGpsStatSrcEnum `json:"src,omitempty"`
    Timestamp *float64               `json:"timestamp,omitempty"`
}
