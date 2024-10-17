package models

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"strings"
)

// StatsSdkclient represents a StatsSdkclient struct.
// SDK Client statistics
type StatsSdkclient struct {
	Id uuid.UUID `json:"id"`
	// last seen timestamp
	LastSeen float64 `json:"last_seen"`
	// map_id of the sdk client (if known), or null
	MapId Optional[uuid.UUID] `json:"map_id"`
	// name of the sdk client (if provided)
	Name *string `json:"name,omitempty"`
	// various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as
	NetworkConnection StatsSdkclientNetworkConnection `json:"network_connection"`
	// uuid of the sdk client
	Uuid uuid.UUID `json:"uuid"`
	// x (in pixels) of user location on the map (if known)
	X *float64 `json:"x,omitempty"`
	// y (in pixels) of user location on the map (if known)
	Y                    *float64       `json:"y,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsSdkclient.
// It customizes the JSON marshaling process for StatsSdkclient objects.
func (s StatsSdkclient) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsSdkclient object to a map representation for JSON marshaling.
func (s StatsSdkclient) toMap() map[string]any {
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
	structMap["network_connection"] = s.NetworkConnection.toMap()
	structMap["uuid"] = s.Uuid
	if s.X != nil {
		structMap["x"] = s.X
	}
	if s.Y != nil {
		structMap["y"] = s.Y
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsSdkclient.
// It customizes the JSON unmarshaling process for StatsSdkclient objects.
func (s *StatsSdkclient) UnmarshalJSON(input []byte) error {
	var temp tempStatsSdkclient
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "id", "last_seen", "map_id", "name", "network_connection", "uuid", "x", "y")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.Id = *temp.Id
	s.LastSeen = *temp.LastSeen
	s.MapId = temp.MapId
	s.Name = temp.Name
	s.NetworkConnection = *temp.NetworkConnection
	s.Uuid = *temp.Uuid
	s.X = temp.X
	s.Y = temp.Y
	return nil
}

// tempStatsSdkclient is a temporary struct used for validating the fields of StatsSdkclient.
type tempStatsSdkclient struct {
	Id                *uuid.UUID                       `json:"id"`
	LastSeen          *float64                         `json:"last_seen"`
	MapId             Optional[uuid.UUID]              `json:"map_id"`
	Name              *string                          `json:"name,omitempty"`
	NetworkConnection *StatsSdkclientNetworkConnection `json:"network_connection"`
	Uuid              *uuid.UUID                       `json:"uuid"`
	X                 *float64                         `json:"x,omitempty"`
	Y                 *float64                         `json:"y,omitempty"`
}

func (s *tempStatsSdkclient) validate() error {
	var errs []string
	if s.Id == nil {
		errs = append(errs, "required field `id` is missing for type `stats_sdkclient`")
	}
	if s.LastSeen == nil {
		errs = append(errs, "required field `last_seen` is missing for type `stats_sdkclient`")
	}
	if s.NetworkConnection == nil {
		errs = append(errs, "required field `network_connection` is missing for type `stats_sdkclient`")
	}
	if s.Uuid == nil {
		errs = append(errs, "required field `uuid` is missing for type `stats_sdkclient`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
