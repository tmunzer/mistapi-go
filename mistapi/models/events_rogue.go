package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// EventsRogue represents a EventsRogue struct.
// rogue events
type EventsRogue struct {
    Ap                   string                 `json:"ap"`
    Bssid                string                 `json:"bssid"`
    Channel              int                    `json:"channel"`
    Rssi                 int                    `json:"rssi"`
    Ssid                 string                 `json:"ssid"`
    Timestamp            float64                `json:"timestamp"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for EventsRogue.
// It customizes the JSON marshaling process for EventsRogue objects.
func (e EventsRogue) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "ap", "bssid", "channel", "rssi", "ssid", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EventsRogue object to a map representation for JSON marshaling.
func (e EventsRogue) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["ap"] = e.Ap
    structMap["bssid"] = e.Bssid
    structMap["channel"] = e.Channel
    structMap["rssi"] = e.Rssi
    structMap["ssid"] = e.Ssid
    structMap["timestamp"] = e.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventsRogue.
// It customizes the JSON unmarshaling process for EventsRogue objects.
func (e *EventsRogue) UnmarshalJSON(input []byte) error {
    var temp tempEventsRogue
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "bssid", "channel", "rssi", "ssid", "timestamp")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.Ap = *temp.Ap
    e.Bssid = *temp.Bssid
    e.Channel = *temp.Channel
    e.Rssi = *temp.Rssi
    e.Ssid = *temp.Ssid
    e.Timestamp = *temp.Timestamp
    return nil
}

// tempEventsRogue is a temporary struct used for validating the fields of EventsRogue.
type tempEventsRogue  struct {
    Ap        *string  `json:"ap"`
    Bssid     *string  `json:"bssid"`
    Channel   *int     `json:"channel"`
    Rssi      *int     `json:"rssi"`
    Ssid      *string  `json:"ssid"`
    Timestamp *float64 `json:"timestamp"`
}

func (e *tempEventsRogue) validate() error {
    var errs []string
    if e.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `events_rogue`")
    }
    if e.Bssid == nil {
        errs = append(errs, "required field `bssid` is missing for type `events_rogue`")
    }
    if e.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `events_rogue`")
    }
    if e.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `events_rogue`")
    }
    if e.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `events_rogue`")
    }
    if e.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `events_rogue`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
