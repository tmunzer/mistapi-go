// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// EventFastroam represents a EventFastroam struct.
type EventFastroam struct {
    ApMac                string                 `json:"ap_mac"`
    ClientMac            string                 `json:"client_mac"`
    Fromap               string                 `json:"fromap"`
    Latency              float64                `json:"latency"`
    Ssid                 string                 `json:"ssid"`
    Subtype              *string                `json:"subtype,omitempty"`
    // Epoch (seconds)
    Timestamp            float64                `json:"timestamp"`
    // enum: `fail`, `none`, `pingpong`, `poor`, `slow`, `success`
    Type                 *EventFastroamTypeEnum `json:"type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for EventFastroam,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (e EventFastroam) String() string {
    return fmt.Sprintf(
    	"EventFastroam[ApMac=%v, ClientMac=%v, Fromap=%v, Latency=%v, Ssid=%v, Subtype=%v, Timestamp=%v, Type=%v, AdditionalProperties=%v]",
    	e.ApMac, e.ClientMac, e.Fromap, e.Latency, e.Ssid, e.Subtype, e.Timestamp, e.Type, e.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for EventFastroam.
// It customizes the JSON marshaling process for EventFastroam objects.
func (e EventFastroam) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(e.AdditionalProperties,
        "ap_mac", "client_mac", "fromap", "latency", "ssid", "subtype", "timestamp", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(e.toMap())
}

// toMap converts the EventFastroam object to a map representation for JSON marshaling.
func (e EventFastroam) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, e.AdditionalProperties)
    structMap["ap_mac"] = e.ApMac
    structMap["client_mac"] = e.ClientMac
    structMap["fromap"] = e.Fromap
    structMap["latency"] = e.Latency
    structMap["ssid"] = e.Ssid
    if e.Subtype != nil {
        structMap["subtype"] = e.Subtype
    }
    structMap["timestamp"] = e.Timestamp
    if e.Type != nil {
        structMap["type"] = e.Type
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for EventFastroam.
// It customizes the JSON unmarshaling process for EventFastroam objects.
func (e *EventFastroam) UnmarshalJSON(input []byte) error {
    var temp tempEventFastroam
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "client_mac", "fromap", "latency", "ssid", "subtype", "timestamp", "type")
    if err != nil {
    	return err
    }
    e.AdditionalProperties = additionalProperties
    
    e.ApMac = *temp.ApMac
    e.ClientMac = *temp.ClientMac
    e.Fromap = *temp.Fromap
    e.Latency = *temp.Latency
    e.Ssid = *temp.Ssid
    e.Subtype = temp.Subtype
    e.Timestamp = *temp.Timestamp
    e.Type = temp.Type
    return nil
}

// tempEventFastroam is a temporary struct used for validating the fields of EventFastroam.
type tempEventFastroam  struct {
    ApMac     *string                `json:"ap_mac"`
    ClientMac *string                `json:"client_mac"`
    Fromap    *string                `json:"fromap"`
    Latency   *float64               `json:"latency"`
    Ssid      *string                `json:"ssid"`
    Subtype   *string                `json:"subtype,omitempty"`
    Timestamp *float64               `json:"timestamp"`
    Type      *EventFastroamTypeEnum `json:"type,omitempty"`
}

func (e *tempEventFastroam) validate() error {
    var errs []string
    if e.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `event_fastroam`")
    }
    if e.ClientMac == nil {
        errs = append(errs, "required field `client_mac` is missing for type `event_fastroam`")
    }
    if e.Fromap == nil {
        errs = append(errs, "required field `fromap` is missing for type `event_fastroam`")
    }
    if e.Latency == nil {
        errs = append(errs, "required field `latency` is missing for type `event_fastroam`")
    }
    if e.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `event_fastroam`")
    }
    if e.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `event_fastroam`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
