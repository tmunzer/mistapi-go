package models

import (
    "encoding/json"
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
    AdditionalProperties   map[string]any                `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookLocationCentrakEvent.
// It customizes the JSON marshaling process for WebhookLocationCentrakEvent objects.
func (w WebhookLocationCentrakEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookLocationCentrakEvent object to a map representation for JSON marshaling.
func (w WebhookLocationCentrakEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp webhookLocationCentrakEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "map_id", "mfg_company_id", "mfg_data", "timestamp", "wifi_beacon_extended_info", "x", "y")
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

// webhookLocationCentrakEvent is a temporary struct used for validating the fields of WebhookLocationCentrakEvent.
type webhookLocationCentrakEvent  struct {
    MapId                  *string                       `json:"map_id,omitempty"`
    MfgCompanyId           *int                          `json:"mfg_company_id,omitempty"`
    MfgData                *string                       `json:"mfg_data,omitempty"`
    Timestamp              *int                          `json:"timestamp,omitempty"`
    WifiBeaconExtendedInfo []WifiBeaconExtendedInfoItems `json:"wifi_beacon_extended_info,omitempty"`
    X                      *float64                      `json:"x,omitempty"`
    Y                      *float64                      `json:"y,omitempty"`
}
