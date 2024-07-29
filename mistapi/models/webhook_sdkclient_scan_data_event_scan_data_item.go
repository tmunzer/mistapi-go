package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// WebhookSdkclientScanDataEventScanDataItem represents a WebhookSdkclientScanDataEventScanDataItem struct.
type WebhookSdkclientScanDataEventScanDataItem struct {
    // mac address of the AP associated with the BSSID scanned
    Ap                   string               `json:"ap"`
    // 5GHz or 2.4GHz band, associated with the BSSID scanned. enum: `2.4`, `5`
    Band                 ScanDataItemBandEnum `json:"band"`
    // BSSID found during client’s background scan for wifi
    Bssid                string               `json:"bssid"`
    // channel of the band found in the scan
    Channel              int                  `json:"channel"`
    // client’s RSSI relative to the BSSID scanned
    Rssi                 float64              `json:"rssi"`
    // ESSID containing the BSSID scanned
    Ssid                 string               `json:"ssid"`
    // time the scan of the particular BSSID occurred
    Timestamp            float64              `json:"timestamp"`
    AdditionalProperties map[string]any       `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookSdkclientScanDataEventScanDataItem.
// It customizes the JSON marshaling process for WebhookSdkclientScanDataEventScanDataItem objects.
func (w WebhookSdkclientScanDataEventScanDataItem) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookSdkclientScanDataEventScanDataItem object to a map representation for JSON marshaling.
func (w WebhookSdkclientScanDataEventScanDataItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    structMap["ap"] = w.Ap
    structMap["band"] = w.Band
    structMap["bssid"] = w.Bssid
    structMap["channel"] = w.Channel
    structMap["rssi"] = w.Rssi
    structMap["ssid"] = w.Ssid
    structMap["timestamp"] = w.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookSdkclientScanDataEventScanDataItem.
// It customizes the JSON unmarshaling process for WebhookSdkclientScanDataEventScanDataItem objects.
func (w *WebhookSdkclientScanDataEventScanDataItem) UnmarshalJSON(input []byte) error {
    var temp webhookSdkclientScanDataEventScanDataItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "band", "bssid", "channel", "rssi", "ssid", "timestamp")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Ap = *temp.Ap
    w.Band = *temp.Band
    w.Bssid = *temp.Bssid
    w.Channel = *temp.Channel
    w.Rssi = *temp.Rssi
    w.Ssid = *temp.Ssid
    w.Timestamp = *temp.Timestamp
    return nil
}

// webhookSdkclientScanDataEventScanDataItem is a temporary struct used for validating the fields of WebhookSdkclientScanDataEventScanDataItem.
type webhookSdkclientScanDataEventScanDataItem  struct {
    Ap        *string               `json:"ap"`
    Band      *ScanDataItemBandEnum `json:"band"`
    Bssid     *string               `json:"bssid"`
    Channel   *int                  `json:"channel"`
    Rssi      *float64              `json:"rssi"`
    Ssid      *string               `json:"ssid"`
    Timestamp *float64              `json:"timestamp"`
}

func (w *webhookSdkclientScanDataEventScanDataItem) validate() error {
    var errs []string
    if w.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if w.Band == nil {
        errs = append(errs, "required field `band` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if w.Bssid == nil {
        errs = append(errs, "required field `bssid` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if w.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if w.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if w.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if w.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Webhook_Sdkclient_Scan_Data_Event_Scan_Data_Item`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
