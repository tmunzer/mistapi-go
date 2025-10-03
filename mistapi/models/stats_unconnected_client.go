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

// StatsUnconnectedClient represents a StatsUnconnectedClient struct.
// Unconnected clients statistics
type StatsUnconnectedClient struct {
	// MAC address of the AP that heard the client
	ApMac string `json:"ap_mac"`
	// Last seen timestamp
	LastSeen Optional[float64] `json:"last_seen"`
	// MAC address of the (unconnected) client
	Mac string `json:"mac"`
	// Device manufacture, through fingerprinting or OUI
	Manufacture string `json:"manufacture"`
	// Map_id of the client (if known), or null
	MapId Optional[uuid.UUID] `json:"map_id"`
	// Client RSSI observed by the AP that heard the client (in dBm)
	Rssi int `json:"rssi"`
	// X (in pixels) of user location on the map (if known)
	X *float64 `json:"x,omitempty"`
	// Y (in pixels) of user location on the map (if known)
	Y                    float64                `json:"y"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for StatsUnconnectedClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsUnconnectedClient) String() string {
	return fmt.Sprintf(
		"StatsUnconnectedClient[ApMac=%v, LastSeen=%v, Mac=%v, Manufacture=%v, MapId=%v, Rssi=%v, X=%v, Y=%v, AdditionalProperties=%v]",
		s.ApMac, s.LastSeen, s.Mac, s.Manufacture, s.MapId, s.Rssi, s.X, s.Y, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsUnconnectedClient.
// It customizes the JSON marshaling process for StatsUnconnectedClient objects.
func (s StatsUnconnectedClient) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"ap_mac", "last_seen", "mac", "manufacture", "map_id", "rssi", "x", "y"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the StatsUnconnectedClient object to a map representation for JSON marshaling.
func (s StatsUnconnectedClient) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["ap_mac"] = s.ApMac
	if s.LastSeen.IsValueSet() {
		if s.LastSeen.Value() != nil {
			structMap["last_seen"] = s.LastSeen.Value()
		} else {
			structMap["last_seen"] = nil
		}
	}
	structMap["mac"] = s.Mac
	structMap["manufacture"] = s.Manufacture
	if s.MapId.IsValueSet() {
		if s.MapId.Value() != nil {
			structMap["map_id"] = s.MapId.Value()
		} else {
			structMap["map_id"] = nil
		}
	}
	structMap["rssi"] = s.Rssi
	if s.X != nil {
		structMap["x"] = s.X
	}
	structMap["y"] = s.Y
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsUnconnectedClient.
// It customizes the JSON unmarshaling process for StatsUnconnectedClient objects.
func (s *StatsUnconnectedClient) UnmarshalJSON(input []byte) error {
	var temp tempStatsUnconnectedClient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap_mac", "last_seen", "mac", "manufacture", "map_id", "rssi", "x", "y")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.ApMac = *temp.ApMac
	s.LastSeen = temp.LastSeen
	s.Mac = *temp.Mac
	s.Manufacture = *temp.Manufacture
	s.MapId = temp.MapId
	s.Rssi = *temp.Rssi
	s.X = temp.X
	s.Y = *temp.Y
	return nil
}

// tempStatsUnconnectedClient is a temporary struct used for validating the fields of StatsUnconnectedClient.
type tempStatsUnconnectedClient struct {
	ApMac       *string             `json:"ap_mac"`
	LastSeen    Optional[float64]   `json:"last_seen"`
	Mac         *string             `json:"mac"`
	Manufacture *string             `json:"manufacture"`
	MapId       Optional[uuid.UUID] `json:"map_id"`
	Rssi        *int                `json:"rssi"`
	X           *float64            `json:"x,omitempty"`
	Y           *float64            `json:"y"`
}

func (s *tempStatsUnconnectedClient) validate() error {
	var errs []string
	if s.ApMac == nil {
		errs = append(errs, "required field `ap_mac` is missing for type `stats_unconnected_client`")
	}
	if s.Mac == nil {
		errs = append(errs, "required field `mac` is missing for type `stats_unconnected_client`")
	}
	if s.Manufacture == nil {
		errs = append(errs, "required field `manufacture` is missing for type `stats_unconnected_client`")
	}
	if s.Rssi == nil {
		errs = append(errs, "required field `rssi` is missing for type `stats_unconnected_client`")
	}
	if s.Y == nil {
		errs = append(errs, "required field `y` is missing for type `stats_unconnected_client`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
