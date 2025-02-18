package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// WebhookSdkclientScanDataEvent represents a WebhookSdkclientScanDataEvent struct.
type WebhookSdkclientScanDataEvent struct {
    // MAC address of the AP the client is connected to
    ConnectionAp         string                                      `json:"connection_ap"`
    // 5GHz or 2.4GHz band, of the BSSID the client is connected to
    ConnectionBand       string                                      `json:"connection_band"`
    // BSSID of the AP the client is connected to
    ConnectionBssid      string                                      `json:"connection_bssid"`
    // Channel of the band the client is connected to
    ConnectionChannel    int                                         `json:"connection_channel"`
    // RSSI of the client’s connection to the AP/BSSID
    ConnectionRssi       float64                                     `json:"connection_rssi"`
    // Time client last seen with scan data
    LastSeen             float64                                     `json:"last_seen"`
    // Client's MAC Address
    Mac                  string                                      `json:"mac"`
    ScanData             []WebhookSdkclientScanDataEventScanDataItem `json:"scan_data,omitempty"`
    SiteId               uuid.UUID                                   `json:"site_id"`
    AdditionalProperties map[string]interface{}                      `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookSdkclientScanDataEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookSdkclientScanDataEvent) String() string {
    return fmt.Sprintf(
    	"WebhookSdkclientScanDataEvent[ConnectionAp=%v, ConnectionBand=%v, ConnectionBssid=%v, ConnectionChannel=%v, ConnectionRssi=%v, LastSeen=%v, Mac=%v, ScanData=%v, SiteId=%v, AdditionalProperties=%v]",
    	w.ConnectionAp, w.ConnectionBand, w.ConnectionBssid, w.ConnectionChannel, w.ConnectionRssi, w.LastSeen, w.Mac, w.ScanData, w.SiteId, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookSdkclientScanDataEvent.
// It customizes the JSON marshaling process for WebhookSdkclientScanDataEvent objects.
func (w WebhookSdkclientScanDataEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "connection_ap", "connection_band", "connection_bssid", "connection_channel", "connection_rssi", "last_seen", "mac", "scan_data", "site_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSdkclientScanDataEvent object to a map representation for JSON marshaling.
func (w WebhookSdkclientScanDataEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["connection_ap"] = w.ConnectionAp
    structMap["connection_band"] = w.ConnectionBand
    structMap["connection_bssid"] = w.ConnectionBssid
    structMap["connection_channel"] = w.ConnectionChannel
    structMap["connection_rssi"] = w.ConnectionRssi
    structMap["last_seen"] = w.LastSeen
    structMap["mac"] = w.Mac
    if w.ScanData != nil {
        structMap["scan_data"] = w.ScanData
    }
    structMap["site_id"] = w.SiteId
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSdkclientScanDataEvent.
// It customizes the JSON unmarshaling process for WebhookSdkclientScanDataEvent objects.
func (w *WebhookSdkclientScanDataEvent) UnmarshalJSON(input []byte) error {
    var temp tempWebhookSdkclientScanDataEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "connection_ap", "connection_band", "connection_bssid", "connection_channel", "connection_rssi", "last_seen", "mac", "scan_data", "site_id")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.ConnectionAp = *temp.ConnectionAp
    w.ConnectionBand = *temp.ConnectionBand
    w.ConnectionBssid = *temp.ConnectionBssid
    w.ConnectionChannel = *temp.ConnectionChannel
    w.ConnectionRssi = *temp.ConnectionRssi
    w.LastSeen = *temp.LastSeen
    w.Mac = *temp.Mac
    w.ScanData = temp.ScanData
    w.SiteId = *temp.SiteId
    return nil
}

// tempWebhookSdkclientScanDataEvent is a temporary struct used for validating the fields of WebhookSdkclientScanDataEvent.
type tempWebhookSdkclientScanDataEvent  struct {
    ConnectionAp      *string                                     `json:"connection_ap"`
    ConnectionBand    *string                                     `json:"connection_band"`
    ConnectionBssid   *string                                     `json:"connection_bssid"`
    ConnectionChannel *int                                        `json:"connection_channel"`
    ConnectionRssi    *float64                                    `json:"connection_rssi"`
    LastSeen          *float64                                    `json:"last_seen"`
    Mac               *string                                     `json:"mac"`
    ScanData          []WebhookSdkclientScanDataEventScanDataItem `json:"scan_data,omitempty"`
    SiteId            *uuid.UUID                                  `json:"site_id"`
}

func (w *tempWebhookSdkclientScanDataEvent) validate() error {
    var errs []string
    if w.ConnectionAp == nil {
        errs = append(errs, "required field `connection_ap` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.ConnectionBand == nil {
        errs = append(errs, "required field `connection_band` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.ConnectionBssid == nil {
        errs = append(errs, "required field `connection_bssid` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.ConnectionChannel == nil {
        errs = append(errs, "required field `connection_channel` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.ConnectionRssi == nil {
        errs = append(errs, "required field `connection_rssi` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if w.SiteId == nil {
        errs = append(errs, "required field `site_id` is missing for type `webhook_sdkclient_scan_data_event`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
