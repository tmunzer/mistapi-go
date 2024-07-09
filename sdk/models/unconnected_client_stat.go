package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// UnconnectedClientStat represents a UnconnectedClientStat struct.
// Unconnected clients statistics
type UnconnectedClientStat struct {
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

// MarshalJSON implements the json.Marshaler interface for UnconnectedClientStat.
// It customizes the JSON marshaling process for UnconnectedClientStat objects.
func (u UnconnectedClientStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(u.toMap())
}

// toMap converts the UnconnectedClientStat object to a map representation for JSON marshaling.
func (u UnconnectedClientStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, u.AdditionalProperties)
    structMap["ap_mac"] = u.ApMac
    structMap["last_seen"] = u.LastSeen
    structMap["mac"] = u.Mac
    structMap["manufacture"] = u.Manufacture
    if u.MapId.IsValueSet() {
        if u.MapId.Value() != nil {
            structMap["map_id"] = u.MapId.Value()
        } else {
            structMap["map_id"] = nil
        }
    }
    structMap["rssi"] = u.Rssi
    if u.X != nil {
        structMap["x"] = u.X
    }
    structMap["y"] = u.Y
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UnconnectedClientStat.
// It customizes the JSON unmarshaling process for UnconnectedClientStat objects.
func (u *UnconnectedClientStat) UnmarshalJSON(input []byte) error {
    var temp unconnectedClientStat
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
    
    u.AdditionalProperties = additionalProperties
    u.ApMac = *temp.ApMac
    u.LastSeen = *temp.LastSeen
    u.Mac = *temp.Mac
    u.Manufacture = *temp.Manufacture
    u.MapId = temp.MapId
    u.Rssi = *temp.Rssi
    u.X = temp.X
    u.Y = *temp.Y
    return nil
}

// unconnectedClientStat is a temporary struct used for validating the fields of UnconnectedClientStat.
type unconnectedClientStat  struct {
    ApMac       *string             `json:"ap_mac"`
    LastSeen    *float64            `json:"last_seen"`
    Mac         *string             `json:"mac"`
    Manufacture *string             `json:"manufacture"`
    MapId       Optional[uuid.UUID] `json:"map_id"`
    Rssi        *int                `json:"rssi"`
    X           *float64            `json:"x,omitempty"`
    Y           *float64            `json:"y"`
}

func (u *unconnectedClientStat) validate() error {
    var errs []string
    if u.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `Unconnected_Client_Stat`")
    }
    if u.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `Unconnected_Client_Stat`")
    }
    if u.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Unconnected_Client_Stat`")
    }
    if u.Manufacture == nil {
        errs = append(errs, "required field `manufacture` is missing for type `Unconnected_Client_Stat`")
    }
    if u.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `Unconnected_Client_Stat`")
    }
    if u.Y == nil {
        errs = append(errs, "required field `y` is missing for type `Unconnected_Client_Stat`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
