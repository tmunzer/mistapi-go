package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// WebhookNacAccountingEvent represents a WebhookNacAccountingEvent struct.
type WebhookNacAccountingEvent struct {
    // mac address of the AP the client roamed or disconnected from
    Ap                   *string        `json:"ap,omitempty"`
    // radius authentication type
    AuthType             *string        `json:"auth_type,omitempty"`
    // it’s the MAC physical address of the access point
    Bssid                *string        `json:"bssid,omitempty"`
    // IP Address of client
    ClientIp             *string        `json:"client_ip,omitempty"`
    // client type E.g. “wired”, “wireless”, “vty”
    ClientType           *string        `json:"client_type,omitempty"`
    // the client’s mac
    Mac                  *string        `json:"mac,omitempty"`
    // NAS Device vendor name E.g. “Juniper”, “Cisco”
    NasVendor            *string        `json:"nas_vendor,omitempty"`
    OrgId                *uuid.UUID     `json:"org_id,omitempty"`
    // number of packets received
    RxPkts               *int           `json:"rx_pkts,omitempty"`
    SiteId               *uuid.UUID     `json:"site_id,omitempty"`
    // ESSID
    Ssid                 *string        `json:"ssid,omitempty"`
    // sampling time (in epoch)
    Timestamp            *float64       `json:"timestamp,omitempty"`
    // number of packets sent
    TxPkts               *int           `json:"tx_pkts,omitempty"`
    // type of event. E.g. “ACCOUNTING_START”, “ACCOUNTING_UPDATE”, “ACCOUNTING_STOP”
    Type                 *string        `json:"type,omitempty"`
    // username authenticated with
    Username             *string        `json:"username,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacAccountingEvent.
// It customizes the JSON marshaling process for WebhookNacAccountingEvent objects.
func (w WebhookNacAccountingEvent) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacAccountingEvent object to a map representation for JSON marshaling.
func (w WebhookNacAccountingEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.Ap != nil {
        structMap["ap"] = w.Ap
    }
    if w.AuthType != nil {
        structMap["auth_type"] = w.AuthType
    }
    if w.Bssid != nil {
        structMap["bssid"] = w.Bssid
    }
    if w.ClientIp != nil {
        structMap["client_ip"] = w.ClientIp
    }
    if w.ClientType != nil {
        structMap["client_type"] = w.ClientType
    }
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.NasVendor != nil {
        structMap["nas_vendor"] = w.NasVendor
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.RxPkts != nil {
        structMap["rx_pkts"] = w.RxPkts
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Ssid != nil {
        structMap["ssid"] = w.Ssid
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.TxPkts != nil {
        structMap["tx_pkts"] = w.TxPkts
    }
    if w.Type != nil {
        structMap["type"] = w.Type
    }
    if w.Username != nil {
        structMap["username"] = w.Username
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WebhookNacAccountingEvent.
// It customizes the JSON unmarshaling process for WebhookNacAccountingEvent objects.
func (w *WebhookNacAccountingEvent) UnmarshalJSON(input []byte) error {
    var temp webhookNacAccountingEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ap", "auth_type", "bssid", "client_ip", "client_type", "mac", "nas_vendor", "org_id", "rx_pkts", "site_id", "ssid", "timestamp", "tx_pkts", "type", "username")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.Ap = temp.Ap
    w.AuthType = temp.AuthType
    w.Bssid = temp.Bssid
    w.ClientIp = temp.ClientIp
    w.ClientType = temp.ClientType
    w.Mac = temp.Mac
    w.NasVendor = temp.NasVendor
    w.OrgId = temp.OrgId
    w.RxPkts = temp.RxPkts
    w.SiteId = temp.SiteId
    w.Ssid = temp.Ssid
    w.Timestamp = temp.Timestamp
    w.TxPkts = temp.TxPkts
    w.Type = temp.Type
    w.Username = temp.Username
    return nil
}

// webhookNacAccountingEvent is a temporary struct used for validating the fields of WebhookNacAccountingEvent.
type webhookNacAccountingEvent  struct {
    Ap         *string    `json:"ap,omitempty"`
    AuthType   *string    `json:"auth_type,omitempty"`
    Bssid      *string    `json:"bssid,omitempty"`
    ClientIp   *string    `json:"client_ip,omitempty"`
    ClientType *string    `json:"client_type,omitempty"`
    Mac        *string    `json:"mac,omitempty"`
    NasVendor  *string    `json:"nas_vendor,omitempty"`
    OrgId      *uuid.UUID `json:"org_id,omitempty"`
    RxPkts     *int       `json:"rx_pkts,omitempty"`
    SiteId     *uuid.UUID `json:"site_id,omitempty"`
    Ssid       *string    `json:"ssid,omitempty"`
    Timestamp  *float64   `json:"timestamp,omitempty"`
    TxPkts     *int       `json:"tx_pkts,omitempty"`
    Type       *string    `json:"type,omitempty"`
    Username   *string    `json:"username,omitempty"`
}
