package models

import (
    "encoding/json"
    "fmt"
)

// WebhookLocationCentrakEvent represents a WebhookLocationCentrakEvent struct.
type WebhookLocationCentrakEvent struct {
    // map id
    MapId                  *string                       `json:"map_id,omitempty"`
    // optional, BLE manufacturing company ID
    MfgCompanyId           *int                          `json:"mfg_company_id,omitempty"`
    // optional, BLE manufacturing data in hex byte-string format (i.e. “112233AABBCC”)
    MfgData                *string                       `json:"mfg_data,omitempty"`
    // timestamp of the event, epoch
    Timestamp              *int                          `json:"timestamp,omitempty"`
    // optional, list of extended beacon info packets heard from the client, frame and sequence control included with the payload
    WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
    // x, in meter
    X                      *float64                      `json:"x,omitempty"`
    // y, in meter
    Y                      *float64                      `json:"y,omitempty"`
    AdditionalProperties   map[string]interface{}        `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookLocationCentrakEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookLocationCentrakEvent) String() string {
    return fmt.Sprintf(
    	"WebhookLocationCentrakEvent[MapId=%v, MfgCompanyId=%v, MfgData=%v, Timestamp=%v, WifiBeaconExtendedInfo=%v, X=%v, Y=%v, AdditionalProperties=%v]",
    	w.MapId, w.MfgCompanyId, w.MfgData, w.Timestamp, w.WifiBeaconExtendedInfo, w.X, w.Y, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationCentrakEvent.
// It customizes the JSON marshaling process for WebhookLocationCentrakEvent objects.
func (w WebhookLocationCentrakEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "map_id", "mfg_company_id", "mfg_data", "timestamp", "wifi_beacon_extended_info", "x", "y"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationCentrakEvent object to a map representation for JSON marshaling.
func (w WebhookLocationCentrakEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.MapId != nil {
        structMap["map_id"] = w.MapId
    }
    if w.MfgCompanyId != nil {
        structMap["mfg_company_id"] = w.MfgCompanyId
    }
    if w.MfgData != nil {
        structMap["mfg_data"] = w.MfgData
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.WifiBeaconExtendedInfo != nil {
        structMap["wifi_beacon_extended_info"] = w.WifiBeaconExtendedInfo
    }
    if w.X != nil {
        structMap["x"] = w.X
    }
    if w.Y != nil {
        structMap["y"] = w.Y
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookLocationCentrakEvent.
// It customizes the JSON unmarshaling process for WebhookLocationCentrakEvent objects.
func (w *WebhookLocationCentrakEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookLocationCentrakEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "map_id", "mfg_company_id", "mfg_data", "timestamp", "wifi_beacon_extended_info", "x", "y")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.MapId = temp.MapId
    w.MfgCompanyId = temp.MfgCompanyId
    w.MfgData = temp.MfgData
    w.Timestamp = temp.Timestamp
    w.WifiBeaconExtendedInfo = temp.WifiBeaconExtendedInfo
    w.X = temp.X
    w.Y = temp.Y
    return nil
}

// tempWebhookLocationCentrakEvent is a temporary struct used for validating the fields of WebhookLocationCentrakEvent.
type tempWebhookLocationCentrakEvent  struct {
    MapId                  *string                       `json:"map_id,omitempty"`
    MfgCompanyId           *int                          `json:"mfg_company_id,omitempty"`
    MfgData                *string                       `json:"mfg_data,omitempty"`
    Timestamp              *int                          `json:"timestamp,omitempty"`
    WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
    X                      *float64                      `json:"x,omitempty"`
    Y                      *float64                      `json:"y,omitempty"`
}
