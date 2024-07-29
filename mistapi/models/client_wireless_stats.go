package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ClientWirelessStats represents a ClientWirelessStats struct.
type ClientWirelessStats struct {
    // TTL of the validity of the stat
    Ttl                  float64                          `json:"_ttl"`
    // estimated client location accuracy, in meter
    Accuracy             *int                             `json:"accuracy,omitempty"`
    AirespaceIfname      *string                          `json:"airespace_ifname,omitempty"`
    // information if airwatch enabled
    Airwatch             *ClientWirelessStatsAirwatch     `json:"airwatch,omitempty"`
    // AP ID the client is connected to
    ApId                 uuid.UUID                        `json:"ap_id"`
    // AP the client is connected to
    ApMac                string                           `json:"ap_mac"`
    // enum: `24`, `5`, `6`
    Band                 Dot11BandEnum                    `json:"band"`
    // current channel
    Channel              int                              `json:"channel"`
    // whether the client is dual_band capable (determined by whether we’ve seen probe requests from both bands)
    DualBand             bool                             `json:"dual_band"`
    // device family, through fingerprinting. iPod / Nexus Galaxy / Windows Mobile or CE …
    Family               string                           `json:"family"`
    // information about this portal
    Guest                *ClientWirelessStatsGuest        `json:"guest,omitempty"`
    // hostname that we learned from sniffing DHCP
    Hostname             string                           `json:"hostname"`
    // how long, in seconds, has the client been idle (since the last RX packet)
    IdleTime             float64                          `json:"idle_time"`
    Ip                   string                           `json:"ip"`
    // whether this is a guest
    IsGuest              bool                             `json:"is_guest"`
    // e.g. WPA2-PSK/CCMP
    KeyMgmt              string                           `json:"key_mgmt"`
    // last seen timestamp
    LastSeen             float64                          `json:"last_seen"`
    // client mac
    Mac                  string                           `json:"mac"`
    // device manufacture, through fingerprinting or OUI
    Manufacture          string                           `json:"manufacture"`
    // estimated client location - map_id
    MapId                *uuid.UUID                       `json:"map_id,omitempty"`
    // device model, may be available if we can identify them
    Model                string                           `json:"model"`
    // number of APs used to locate this client
    NumLocatingAps       *int                             `json:"num_locating_aps,omitempty"`
    // device os, through fingerprinting
    Os                   string                           `json:"os"`
    // if it’s currently in power-save mode
    PowerSaving          bool                             `json:"power_saving"`
    // enum: `a`, `ac`, `ax`, `b`, `g`, `n`
    Proto                Dot11ProtoEnum                   `json:"proto"`
    // PSK id (if multi-psk is used)
    PskId                *uuid.UUID                       `json:"psk_id,omitempty"`
    // signal strength
    Rssi                 float64                          `json:"rssi"`
    // list of rssizone_id’s where client is in and since when (if known)
    Rssizones            []ClientWirelessStatsRssiZone    `json:"rssizones,omitempty"`
    // rate of receiving traffic from the clients, bits/seconds, last known
    RxBps                float64                          `json:"rx_bps"`
    // amount of traffic received from client since client connects
    RxBytes              float64                          `json:"rx_bytes"`
    // amount of traffic received from client since client connects
    RxPackets            float64                          `json:"rx_packets"`
    // RX Rate, Mbps
    RxRate               float64                          `json:"rx_rate"`
    // amount of rx retries
    RxRetries            float64                          `json:"rx_retries"`
    // signal over noise
    Snr                  float64                          `json:"snr"`
    // SSID the client is connected to
    Ssid                 string                           `json:"ssid"`
    // rate of transmitting traffic to the clients, bits/seconds, last known
    TxBps                float64                          `json:"tx_bps"`
    // amount of traffic sent to client since client connects
    TxBytes              float64                          `json:"tx_bytes"`
    // amount of traffic sent to client since client connects
    TxPackets            float64                          `json:"tx_packets"`
    // TX Rate, Mbps
    TxRate               float64                          `json:"tx_rate"`
    // amount of tx retries
    TxRetries            float64                          `json:"tx_retries"`
    // client’s type, regular / vip / resource / blocked (if client object is created)
    Type                 *string                          `json:"type,omitempty"`
    // how long, in seconds, has the client been connected
    Uptime               float64                          `json:"uptime"`
    // username that we learned from 802.1X exchange or Per_user PSK or User Portal
    Username             string                           `json:"username"`
    // list of beacon_id’s where the client is in and since when (if known)
    Vbeacons             []ClientWirelessStatsVbeacon     `json:"vbeacons,omitempty"`
    // vlan id, could be empty (from older AP)
    VlanId               *int                             `json:"vlan_id,omitempty"`
    // WLAN ID the client is connected to
    WlanId               uuid.UUID                        `json:"wlan_id"`
    // current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched.
    WxruleId             *uuid.UUID                       `json:"wxrule_id,omitempty"`
    // current WxlanRule usage per tag_id
    WxruleUsage          []ClientWirelessStatsWxruleUsage `json:"wxrule_usage,omitempty"`
    // estimated clinet location in pixels
    X                    *float64                         `json:"x,omitempty"`
    // estimated client location in meter
    XM                   *float64                         `json:"x_m,omitempty"`
    // estimated clinet location in pixels
    Y                    *float64                         `json:"y,omitempty"`
    // estimated client location in meter
    YM                   *float64                         `json:"y_m,omitempty"`
    // list of zone_id’s where client is in and since when (if known)
    Zones                []ClientWirelessStatsZone        `json:"zones,omitempty"`
    AdditionalProperties map[string]any                   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWirelessStats.
// It customizes the JSON marshaling process for ClientWirelessStats objects.
func (c ClientWirelessStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWirelessStats object to a map representation for JSON marshaling.
func (c ClientWirelessStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["_ttl"] = c.Ttl
    if c.Accuracy != nil {
        structMap["accuracy"] = c.Accuracy
    }
    if c.AirespaceIfname != nil {
        structMap["airespace_ifname"] = c.AirespaceIfname
    }
    if c.Airwatch != nil {
        structMap["airwatch"] = c.Airwatch.toMap()
    }
    structMap["ap_id"] = c.ApId
    structMap["ap_mac"] = c.ApMac
    structMap["band"] = c.Band
    structMap["channel"] = c.Channel
    structMap["dual_band"] = c.DualBand
    structMap["family"] = c.Family
    if c.Guest != nil {
        structMap["guest"] = c.Guest.toMap()
    }
    structMap["hostname"] = c.Hostname
    structMap["idle_time"] = c.IdleTime
    structMap["ip"] = c.Ip
    structMap["is_guest"] = c.IsGuest
    structMap["key_mgmt"] = c.KeyMgmt
    structMap["last_seen"] = c.LastSeen
    structMap["mac"] = c.Mac
    structMap["manufacture"] = c.Manufacture
    if c.MapId != nil {
        structMap["map_id"] = c.MapId
    }
    structMap["model"] = c.Model
    if c.NumLocatingAps != nil {
        structMap["num_locating_aps"] = c.NumLocatingAps
    }
    structMap["os"] = c.Os
    structMap["power_saving"] = c.PowerSaving
    structMap["proto"] = c.Proto
    if c.PskId != nil {
        structMap["psk_id"] = c.PskId
    }
    structMap["rssi"] = c.Rssi
    if c.Rssizones != nil {
        structMap["rssizones"] = c.Rssizones
    }
    structMap["rx_bps"] = c.RxBps
    structMap["rx_bytes"] = c.RxBytes
    structMap["rx_packets"] = c.RxPackets
    structMap["rx_rate"] = c.RxRate
    structMap["rx_retries"] = c.RxRetries
    structMap["snr"] = c.Snr
    structMap["ssid"] = c.Ssid
    structMap["tx_bps"] = c.TxBps
    structMap["tx_bytes"] = c.TxBytes
    structMap["tx_packets"] = c.TxPackets
    structMap["tx_rate"] = c.TxRate
    structMap["tx_retries"] = c.TxRetries
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    structMap["uptime"] = c.Uptime
    structMap["username"] = c.Username
    if c.Vbeacons != nil {
        structMap["vbeacons"] = c.Vbeacons
    }
    if c.VlanId != nil {
        structMap["vlan_id"] = c.VlanId
    }
    structMap["wlan_id"] = c.WlanId
    if c.WxruleId != nil {
        structMap["wxrule_id"] = c.WxruleId
    }
    if c.WxruleUsage != nil {
        structMap["wxrule_usage"] = c.WxruleUsage
    }
    if c.X != nil {
        structMap["x"] = c.X
    }
    if c.XM != nil {
        structMap["x_m"] = c.XM
    }
    if c.Y != nil {
        structMap["y"] = c.Y
    }
    if c.YM != nil {
        structMap["y_m"] = c.YM
    }
    if c.Zones != nil {
        structMap["zones"] = c.Zones
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWirelessStats.
// It customizes the JSON unmarshaling process for ClientWirelessStats objects.
func (c *ClientWirelessStats) UnmarshalJSON(input []byte) error {
    var temp clientWirelessStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "_ttl", "accuracy", "airespace_ifname", "airwatch", "ap_id", "ap_mac", "band", "channel", "dual_band", "family", "guest", "hostname", "idle_time", "ip", "is_guest", "key_mgmt", "last_seen", "mac", "manufacture", "map_id", "model", "num_locating_aps", "os", "power_saving", "proto", "psk_id", "rssi", "rssizones", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "snr", "ssid", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries", "type", "uptime", "username", "vbeacons", "vlan_id", "wlan_id", "wxrule_id", "wxrule_usage", "x", "x_m", "y", "y_m", "zones")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Ttl = *temp.Ttl
    c.Accuracy = temp.Accuracy
    c.AirespaceIfname = temp.AirespaceIfname
    c.Airwatch = temp.Airwatch
    c.ApId = *temp.ApId
    c.ApMac = *temp.ApMac
    c.Band = *temp.Band
    c.Channel = *temp.Channel
    c.DualBand = *temp.DualBand
    c.Family = *temp.Family
    c.Guest = temp.Guest
    c.Hostname = *temp.Hostname
    c.IdleTime = *temp.IdleTime
    c.Ip = *temp.Ip
    c.IsGuest = *temp.IsGuest
    c.KeyMgmt = *temp.KeyMgmt
    c.LastSeen = *temp.LastSeen
    c.Mac = *temp.Mac
    c.Manufacture = *temp.Manufacture
    c.MapId = temp.MapId
    c.Model = *temp.Model
    c.NumLocatingAps = temp.NumLocatingAps
    c.Os = *temp.Os
    c.PowerSaving = *temp.PowerSaving
    c.Proto = *temp.Proto
    c.PskId = temp.PskId
    c.Rssi = *temp.Rssi
    c.Rssizones = temp.Rssizones
    c.RxBps = *temp.RxBps
    c.RxBytes = *temp.RxBytes
    c.RxPackets = *temp.RxPackets
    c.RxRate = *temp.RxRate
    c.RxRetries = *temp.RxRetries
    c.Snr = *temp.Snr
    c.Ssid = *temp.Ssid
    c.TxBps = *temp.TxBps
    c.TxBytes = *temp.TxBytes
    c.TxPackets = *temp.TxPackets
    c.TxRate = *temp.TxRate
    c.TxRetries = *temp.TxRetries
    c.Type = temp.Type
    c.Uptime = *temp.Uptime
    c.Username = *temp.Username
    c.Vbeacons = temp.Vbeacons
    c.VlanId = temp.VlanId
    c.WlanId = *temp.WlanId
    c.WxruleId = temp.WxruleId
    c.WxruleUsage = temp.WxruleUsage
    c.X = temp.X
    c.XM = temp.XM
    c.Y = temp.Y
    c.YM = temp.YM
    c.Zones = temp.Zones
    return nil
}

// clientWirelessStats is a temporary struct used for validating the fields of ClientWirelessStats.
type clientWirelessStats  struct {
    Ttl             *float64                         `json:"_ttl"`
    Accuracy        *int                             `json:"accuracy,omitempty"`
    AirespaceIfname *string                          `json:"airespace_ifname,omitempty"`
    Airwatch        *ClientWirelessStatsAirwatch     `json:"airwatch,omitempty"`
    ApId            *uuid.UUID                       `json:"ap_id"`
    ApMac           *string                          `json:"ap_mac"`
    Band            *Dot11BandEnum                   `json:"band"`
    Channel         *int                             `json:"channel"`
    DualBand        *bool                            `json:"dual_band"`
    Family          *string                          `json:"family"`
    Guest           *ClientWirelessStatsGuest        `json:"guest,omitempty"`
    Hostname        *string                          `json:"hostname"`
    IdleTime        *float64                         `json:"idle_time"`
    Ip              *string                          `json:"ip"`
    IsGuest         *bool                            `json:"is_guest"`
    KeyMgmt         *string                          `json:"key_mgmt"`
    LastSeen        *float64                         `json:"last_seen"`
    Mac             *string                          `json:"mac"`
    Manufacture     *string                          `json:"manufacture"`
    MapId           *uuid.UUID                       `json:"map_id,omitempty"`
    Model           *string                          `json:"model"`
    NumLocatingAps  *int                             `json:"num_locating_aps,omitempty"`
    Os              *string                          `json:"os"`
    PowerSaving     *bool                            `json:"power_saving"`
    Proto           *Dot11ProtoEnum                  `json:"proto"`
    PskId           *uuid.UUID                       `json:"psk_id,omitempty"`
    Rssi            *float64                         `json:"rssi"`
    Rssizones       []ClientWirelessStatsRssiZone    `json:"rssizones,omitempty"`
    RxBps           *float64                         `json:"rx_bps"`
    RxBytes         *float64                         `json:"rx_bytes"`
    RxPackets       *float64                         `json:"rx_packets"`
    RxRate          *float64                         `json:"rx_rate"`
    RxRetries       *float64                         `json:"rx_retries"`
    Snr             *float64                         `json:"snr"`
    Ssid            *string                          `json:"ssid"`
    TxBps           *float64                         `json:"tx_bps"`
    TxBytes         *float64                         `json:"tx_bytes"`
    TxPackets       *float64                         `json:"tx_packets"`
    TxRate          *float64                         `json:"tx_rate"`
    TxRetries       *float64                         `json:"tx_retries"`
    Type            *string                          `json:"type,omitempty"`
    Uptime          *float64                         `json:"uptime"`
    Username        *string                          `json:"username"`
    Vbeacons        []ClientWirelessStatsVbeacon     `json:"vbeacons,omitempty"`
    VlanId          *int                             `json:"vlan_id,omitempty"`
    WlanId          *uuid.UUID                       `json:"wlan_id"`
    WxruleId        *uuid.UUID                       `json:"wxrule_id,omitempty"`
    WxruleUsage     []ClientWirelessStatsWxruleUsage `json:"wxrule_usage,omitempty"`
    X               *float64                         `json:"x,omitempty"`
    XM              *float64                         `json:"x_m,omitempty"`
    Y               *float64                         `json:"y,omitempty"`
    YM              *float64                         `json:"y_m,omitempty"`
    Zones           []ClientWirelessStatsZone        `json:"zones,omitempty"`
}

func (c *clientWirelessStats) validate() error {
    var errs []string
    if c.Ttl == nil {
        errs = append(errs, "required field `_ttl` is missing for type `Client_Wireless_Stats`")
    }
    if c.ApId == nil {
        errs = append(errs, "required field `ap_id` is missing for type `Client_Wireless_Stats`")
    }
    if c.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `Client_Wireless_Stats`")
    }
    if c.Band == nil {
        errs = append(errs, "required field `band` is missing for type `Client_Wireless_Stats`")
    }
    if c.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `Client_Wireless_Stats`")
    }
    if c.DualBand == nil {
        errs = append(errs, "required field `dual_band` is missing for type `Client_Wireless_Stats`")
    }
    if c.Family == nil {
        errs = append(errs, "required field `family` is missing for type `Client_Wireless_Stats`")
    }
    if c.Hostname == nil {
        errs = append(errs, "required field `hostname` is missing for type `Client_Wireless_Stats`")
    }
    if c.IdleTime == nil {
        errs = append(errs, "required field `idle_time` is missing for type `Client_Wireless_Stats`")
    }
    if c.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `Client_Wireless_Stats`")
    }
    if c.IsGuest == nil {
        errs = append(errs, "required field `is_guest` is missing for type `Client_Wireless_Stats`")
    }
    if c.KeyMgmt == nil {
        errs = append(errs, "required field `key_mgmt` is missing for type `Client_Wireless_Stats`")
    }
    if c.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `Client_Wireless_Stats`")
    }
    if c.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Client_Wireless_Stats`")
    }
    if c.Manufacture == nil {
        errs = append(errs, "required field `manufacture` is missing for type `Client_Wireless_Stats`")
    }
    if c.Model == nil {
        errs = append(errs, "required field `model` is missing for type `Client_Wireless_Stats`")
    }
    if c.Os == nil {
        errs = append(errs, "required field `os` is missing for type `Client_Wireless_Stats`")
    }
    if c.PowerSaving == nil {
        errs = append(errs, "required field `power_saving` is missing for type `Client_Wireless_Stats`")
    }
    if c.Proto == nil {
        errs = append(errs, "required field `proto` is missing for type `Client_Wireless_Stats`")
    }
    if c.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `Client_Wireless_Stats`")
    }
    if c.RxBps == nil {
        errs = append(errs, "required field `rx_bps` is missing for type `Client_Wireless_Stats`")
    }
    if c.RxBytes == nil {
        errs = append(errs, "required field `rx_bytes` is missing for type `Client_Wireless_Stats`")
    }
    if c.RxPackets == nil {
        errs = append(errs, "required field `rx_packets` is missing for type `Client_Wireless_Stats`")
    }
    if c.RxRate == nil {
        errs = append(errs, "required field `rx_rate` is missing for type `Client_Wireless_Stats`")
    }
    if c.RxRetries == nil {
        errs = append(errs, "required field `rx_retries` is missing for type `Client_Wireless_Stats`")
    }
    if c.Snr == nil {
        errs = append(errs, "required field `snr` is missing for type `Client_Wireless_Stats`")
    }
    if c.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `Client_Wireless_Stats`")
    }
    if c.TxBps == nil {
        errs = append(errs, "required field `tx_bps` is missing for type `Client_Wireless_Stats`")
    }
    if c.TxBytes == nil {
        errs = append(errs, "required field `tx_bytes` is missing for type `Client_Wireless_Stats`")
    }
    if c.TxPackets == nil {
        errs = append(errs, "required field `tx_packets` is missing for type `Client_Wireless_Stats`")
    }
    if c.TxRate == nil {
        errs = append(errs, "required field `tx_rate` is missing for type `Client_Wireless_Stats`")
    }
    if c.TxRetries == nil {
        errs = append(errs, "required field `tx_retries` is missing for type `Client_Wireless_Stats`")
    }
    if c.Uptime == nil {
        errs = append(errs, "required field `uptime` is missing for type `Client_Wireless_Stats`")
    }
    if c.Username == nil {
        errs = append(errs, "required field `username` is missing for type `Client_Wireless_Stats`")
    }
    if c.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `Client_Wireless_Stats`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
