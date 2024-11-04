package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsUnconnectedClient represents a StatsUnconnectedClient struct.
// Unconnected clients statistics
type StatsUnconnectedClient struct {
    // mac address of the AP that heard the client
    ApMac                string              `json:"ap_mac"`
    // last seen timestamp
    LastSeen             float64             `json:"last_seen"`
    // mac address of the (unconnected) client
    Mac                  string              `json:"mac"`
    // device manufacture, through fingerprinting or OUI
    Manufacture          string              `json:"manufacture"`
    // map_id of the client (if known), or null
    MapId                Optional[uuid.UUID] `json:"map_id"`
    // client RSSI observered by the AP that heard the client (in dBm)
    Rssi                 int                 `json:"rssi"`
    // x (in pixels) of user location on the map (if known)
    X                    *float64            `json:"x,omitempty"`
    // y (in pixels) of user location on the map (if known)
    Y                    float64             `json:"y"`
    AdditionalProperties map[string]any      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsUnconnectedClient.
// It customizes the JSON marshaling process for StatsUnconnectedClient objects.
func (s StatsUnconnectedClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the StatsUnconnectedClient object to a map representation for JSON marshaling.
func (s StatsUnconnectedClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["ap_mac"] = s.ApMac
    structMap["last_seen"] = s.LastSeen
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
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap_mac", "last_seen", "mac", "manufacture", "map_id", "rssi", "x", "y")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.ApMac = *temp.ApMac
    s.LastSeen = *temp.LastSeen
    s.Mac = *temp.Mac
    s.Manufacture = *temp.Manufacture
    s.MapId = temp.MapId
    s.Rssi = *temp.Rssi
    s.X = temp.X
    s.Y = *temp.Y
    return nil
}

// tempStatsUnconnectedClient is a temporary struct used for validating the fields of StatsUnconnectedClient.
type tempStatsUnconnectedClient  struct {
    ApMac       *string             `json:"ap_mac"`
    LastSeen    *float64            `json:"last_seen"`
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
    if s.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `stats_unconnected_client`")
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
    return errors.New(strings.Join (errs, "\n"))
}
