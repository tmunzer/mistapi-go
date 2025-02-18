package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// WebhookClientJoinEvent represents a WebhookClientJoinEvent struct.
type WebhookClientJoinEvent struct {
    // MAC address of the AP the client connected to
    Ap                   string                 `json:"ap"`
    // user-friendly name of the AP the client connected to.
    ApName               string                 `json:"ap_name"`
    // 5GHz or 2.4GHz band
    Band                 string                 `json:"band"`
    Bssid                string                 `json:"bssid"`
    // Time when the user connects
    Connect              int                    `json:"connect"`
    // floating point connect timestamp with millisecond precision
    ConnectFloat         float64                `json:"connect_float"`
    // Client's MAC Address
    Mac                  string                 `json:"mac"`
    OrgId                uuid.UUID              `json:"org_id"`
    // RSSI when the client associated
    Rssi                 float64                `json:"rssi"`
    SiteId               uuid.UUID              `json:"site_id"`
    SiteName             string                 `json:"site_name"`
    // ESSID
    Ssid                 string                 `json:"ssid"`
    Timestamp            float64                `json:"timestamp"`
    // schema version of this message
    Version              float64                `json:"version"`
    WlanId               uuid.UUID              `json:"wlan_id"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookClientJoinEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookClientJoinEvent) String() string {
    return fmt.Sprintf(
    	"WebhookClientJoinEvent[Ap=%v, ApName=%v, Band=%v, Bssid=%v, Connect=%v, ConnectFloat=%v, Mac=%v, OrgId=%v, Rssi=%v, SiteId=%v, SiteName=%v, Ssid=%v, Timestamp=%v, Version=%v, WlanId=%v, AdditionalProperties=%v]",
    	w.Ap, w.ApName, w.Band, w.Bssid, w.Connect, w.ConnectFloat, w.Mac, w.OrgId, w.Rssi, w.SiteId, w.SiteName, w.Ssid, w.Timestamp, w.Version, w.WlanId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookClientJoinEvent.
// It customizes the JSON marshaling process for WebhookClientJoinEvent objects.
func (w WebhookClientJoinEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "ap", "ap_name", "band", "bssid", "connect", "connect_float", "mac", "org_id", "rssi", "site_id", "site_name", "ssid", "timestamp", "version", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookClientJoinEvent object to a map representation for JSON marshaling.
func (w WebhookClientJoinEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["ap"] = w.Ap
    structMap["ap_name"] = w.ApName
    structMap["band"] = w.Band
    structMap["bssid"] = w.Bssid
    structMap["connect"] = w.Connect
    structMap["connect_float"] = w.ConnectFloat
    structMap["mac"] = w.Mac
    structMap["org_id"] = w.OrgId
    structMap["rssi"] = w.Rssi
    structMap["site_id"] = w.SiteId
    structMap["site_name"] = w.SiteName
    structMap["ssid"] = w.Ssid
    structMap["timestamp"] = w.Timestamp
    structMap["version"] = w.Version
    structMap["wlan_id"] = w.WlanId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookClientJoinEvent.
// It customizes the JSON unmarshaling process for WebhookClientJoinEvent objects.
func (w *WebhookClientJoinEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookClientJoinEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "ap_name", "band", "bssid", "connect", "connect_float", "mac", "org_id", "rssi", "site_id", "site_name", "ssid", "timestamp", "version", "wlan_id")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.Ap = *temp.Ap
    w.ApName = *temp.ApName
    w.Band = *temp.Band
    w.Bssid = *temp.Bssid
    w.Connect = *temp.Connect
    w.ConnectFloat = *temp.ConnectFloat
    w.Mac = *temp.Mac
    w.OrgId = *temp.OrgId
    w.Rssi = *temp.Rssi
    w.SiteId = *temp.SiteId
    w.SiteName = *temp.SiteName
    w.Ssid = *temp.Ssid
    w.Timestamp = *temp.Timestamp
    w.Version = *temp.Version
    w.WlanId = *temp.WlanId
    return nil
}

// tempWebhookClientJoinEvent is a temporary struct used for validating the fields of WebhookClientJoinEvent.
type tempWebhookClientJoinEvent  struct {
    Ap           *string    `json:"ap"`
    ApName       *string    `json:"ap_name"`
    Band         *string    `json:"band"`
    Bssid        *string    `json:"bssid"`
    Connect      *int       `json:"connect"`
    ConnectFloat *float64   `json:"connect_float"`
    Mac          *string    `json:"mac"`
    OrgId        *uuid.UUID `json:"org_id"`
    Rssi         *float64   `json:"rssi"`
    SiteId       *uuid.UUID `json:"site_id"`
    SiteName     *string    `json:"site_name"`
    Ssid         *string    `json:"ssid"`
    Timestamp    *float64   `json:"timestamp"`
    Version      *float64   `json:"version"`
    WlanId       *uuid.UUID `json:"wlan_id"`
}

func (w *tempWebhookClientJoinEvent) validate() error {
    var errs []string
    if w.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `webhook_client_join_event`")
    }
    if w.ApName == nil {
        errs = append(errs, "required field `ap_name` is missing for type `webhook_client_join_event`")
    }
    if w.Band == nil {
        errs = append(errs, "required field `band` is missing for type `webhook_client_join_event`")
    }
    if w.Bssid == nil {
        errs = append(errs, "required field `bssid` is missing for type `webhook_client_join_event`")
    }
    if w.Connect == nil {
        errs = append(errs, "required field `connect` is missing for type `webhook_client_join_event`")
    }
    if w.ConnectFloat == nil {
        errs = append(errs, "required field `connect_float` is missing for type `webhook_client_join_event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `webhook_client_join_event`")
    }
    if w.OrgId == nil {
        errs = append(errs, "required field `org_id` is missing for type `webhook_client_join_event`")
    }
    if w.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `webhook_client_join_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_client_join_event`")
    }
    if w.SiteName == nil {
        errs = append(errs, "required field `site_name` is missing for type `webhook_client_join_event`")
    }
    if w.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `webhook_client_join_event`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `webhook_client_join_event`")
    }
    if w.Version == nil {
        errs = append(errs, "required field `version` is missing for type `webhook_client_join_event`")
    }
    if w.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `webhook_client_join_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
