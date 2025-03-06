package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// WebhookNacAccountingEvent represents a WebhookNacAccountingEvent struct.
type WebhookNacAccountingEvent struct {
    // MAC address of the AP the client roamed or disconnected from
    Ap                   *string                `json:"ap,omitempty"`
    // enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `peap-tls`, `psk`
    AuthType             *NacAuthTypeEnum       `json:"auth_type,omitempty"`
    // MAC physical address of the access point
    Bssid                *string                `json:"bssid,omitempty"`
    // IP Address of client
    ClientIp             *string                `json:"client_ip,omitempty"`
    // Client type E.g. "wired", "wireless", "vty"
    ClientType           *string                `json:"client_type,omitempty"`
    // Client's MAC Address
    Mac                  *string                `json:"mac,omitempty"`
    // NAS Device vendor name E.g. "Juniper", "Cisco"
    NasVendor            *string                `json:"nas_vendor,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // Number of packets received
    RxPkts               *int                   `json:"rx_pkts,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // ESSID
    Ssid                 *string                `json:"ssid,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    // Number of packets sent
    TxPkts               *int                   `json:"tx_pkts,omitempty"`
    // Type of event. E.g. "ACCOUNTING_START", "ACCOUNTING_UPDATE", "ACCOUNTING_STOP"
    Type                 *string                `json:"type,omitempty"`
    // Username authenticated with
    Username             *string                `json:"username,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WebhookNacAccountingEvent,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WebhookNacAccountingEvent) String() string {
    return fmt.Sprintf(
    	"WebhookNacAccountingEvent[Ap=%v, AuthType=%v, Bssid=%v, ClientIp=%v, ClientType=%v, Mac=%v, NasVendor=%v, OrgId=%v, RxPkts=%v, SiteId=%v, Ssid=%v, Timestamp=%v, TxPkts=%v, Type=%v, Username=%v, AdditionalProperties=%v]",
    	w.Ap, w.AuthType, w.Bssid, w.ClientIp, w.ClientType, w.Mac, w.NasVendor, w.OrgId, w.RxPkts, w.SiteId, w.Ssid, w.Timestamp, w.TxPkts, w.Type, w.Username, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WebhookNacAccountingEvent.
// It customizes the JSON marshaling process for WebhookNacAccountingEvent objects.
func (w WebhookNacAccountingEvent) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "ap", "auth_type", "bssid", "client_ip", "client_type", "mac", "nas_vendor", "org_id", "rx_pkts", "site_id", "ssid", "timestamp", "tx_pkts", "type", "username"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WebhookNacAccountingEvent object to a map representation for JSON marshaling.
func (w WebhookNacAccountingEvent) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
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
    var temp tempWebhookNacAccountingEvent
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "auth_type", "bssid", "client_ip", "client_type", "mac", "nas_vendor", "org_id", "rx_pkts", "site_id", "ssid", "timestamp", "tx_pkts", "type", "username")
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

// tempWebhookNacAccountingEvent is a temporary struct used for validating the fields of WebhookNacAccountingEvent.
type tempWebhookNacAccountingEvent  struct {
    Ap         *string          `json:"ap,omitempty"`
    AuthType   *NacAuthTypeEnum `json:"auth_type,omitempty"`
    Bssid      *string          `json:"bssid,omitempty"`
    ClientIp   *string          `json:"client_ip,omitempty"`
    ClientType *string          `json:"client_type,omitempty"`
    Mac        *string          `json:"mac,omitempty"`
    NasVendor  *string          `json:"nas_vendor,omitempty"`
    OrgId      *uuid.UUID       `json:"org_id,omitempty"`
    RxPkts     *int             `json:"rx_pkts,omitempty"`
    SiteId     *uuid.UUID       `json:"site_id,omitempty"`
    Ssid       *string          `json:"ssid,omitempty"`
    Timestamp  *float64         `json:"timestamp,omitempty"`
    TxPkts     *int             `json:"tx_pkts,omitempty"`
    Type       *string          `json:"type,omitempty"`
    Username   *string          `json:"username,omitempty"`
}
