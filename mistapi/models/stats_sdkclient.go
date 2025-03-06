package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsSdkclient represents a StatsSdkclient struct.
// SDK Client statistics
type StatsSdkclient struct {
    // Unique ID of the object instance in the Mist Organization
    Id                   uuid.UUID                       `json:"id"`
    // Last seen timestamp
    LastSeen             *float64                        `json:"last_seen"`
    // Map_id of the sdk client (if known), or null
    MapId                Optional[uuid.UUID]             `json:"map_id"`
    // Name of the sdk client (if provided)
    Name                 *string                         `json:"name,omitempty"`
    // Various network connection info for the SDK client (if known, else omitted), with RSSI in dBm, and signal level as
    NetworkConnection    StatsSdkclientNetworkConnection `json:"network_connection"`
    // UUID of the sdk client
    Uuid                 uuid.UUID                       `json:"uuid"`
    // X (in pixels) of user location on the map (if known)
    X                    *float64                        `json:"x,omitempty"`
    // Y (in pixels) of user location on the map (if known)
    Y                    *float64                        `json:"y,omitempty"`
    AdditionalProperties map[string]interface{}          `json:"_"`
}

// String implements the fmt.Stringer interface for StatsSdkclient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsSdkclient) String() string {
    return fmt.Sprintf(
    	"StatsSdkclient[Id=%v, LastSeen=%v, MapId=%v, Name=%v, NetworkConnection=%v, Uuid=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	s.Id, s.LastSeen, s.MapId, s.Name, s.NetworkConnection, s.Uuid, s.X, s.Y, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsSdkclient.
// It customizes the JSON marshaling process for StatsSdkclient objects.
func (s StatsSdkclient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "id", "last_seen", "map_id", "name", "network_connection", "uuid", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsSdkclient object to a map representation for JSON marshaling.
func (s StatsSdkclient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["id"] = s.Id
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    } else {
        structMap["last_seen"] = nil
    }
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "id", "last_seen", "map_id", "name", "network_connection", "uuid", "x", "y")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Id = *temp.Id
    s.LastSeen = temp.LastSeen
    s.MapId = temp.MapId
    s.Name = temp.Name
    s.NetworkConnection = *temp.NetworkConnection
    s.Uuid = *temp.Uuid
    s.X = temp.X
    s.Y = temp.Y
    return nil
}

// tempStatsSdkclient is a temporary struct used for validating the fields of StatsSdkclient.
type tempStatsSdkclient  struct {
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
    if s.NetworkConnection == nil {
        errs = append(errs, "required field `network_connection` is missing for type `stats_sdkclient`")
    }
    if s.Uuid == nil {
        errs = append(errs, "required field `uuid` is missing for type `stats_sdkclient`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
