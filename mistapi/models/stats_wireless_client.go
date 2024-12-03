package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// StatsWirelessClient represents a StatsWirelessClient struct.
type StatsWirelessClient struct {
    // estimated client location accuracy, in meter
    Accuracy             *int                             `json:"accuracy,omitempty"`
    AirespaceIfname      *string                          `json:"airespace_ifname,omitempty"`
    // information if airwatch enabled
    Airwatch             *StatsWirelessClientAirwatch     `json:"airwatch,omitempty"`
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
    Guest                *StatsWirelessClientGuest        `json:"guest,omitempty"`
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
    Rssizones            []StatsWirelessClientRssiZone    `json:"rssizones,omitempty"`
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
    Vbeacons             []StatsWirelessClientVbeacon     `json:"vbeacons,omitempty"`
    // vlan id, could be empty (from older AP)
    VlanId               *int                             `json:"vlan_id,omitempty"`
    // WLAN ID the client is connected to
    WlanId               uuid.UUID                        `json:"wlan_id"`
    // current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched.
    WxruleId             *uuid.UUID                       `json:"wxrule_id,omitempty"`
    // current WxlanRule usage per tag_id
    WxruleUsage          []StatsWirelessClientWxruleUsage `json:"wxrule_usage,omitempty"`
    // estimated clinet location in pixels
    X                    *float64                         `json:"x,omitempty"`
    // estimated client location in meter
    XM                   *float64                         `json:"x_m,omitempty"`
    // estimated clinet location in pixels
    Y                    *float64                         `json:"y,omitempty"`
    // estimated client location in meter
    YM                   *float64                         `json:"y_m,omitempty"`
    // list of zone_id’s where client is in and since when (if known)
    Zones                []StatsWirelessClientZone        `json:"zones,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsWirelessClient.
// It customizes the JSON marshaling process for StatsWirelessClient objects.
func (s StatsWirelessClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "accuracy", "airespace_ifname", "airwatch", "ap_id", "ap_mac", "band", "channel", "dual_band", "family", "guest", "hostname", "idle_time", "ip", "is_guest", "key_mgmt", "last_seen", "mac", "manufacture", "map_id", "model", "num_locating_aps", "os", "power_saving", "proto", "psk_id", "rssi", "rssizones", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "snr", "ssid", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries", "type", "uptime", "username", "vbeacons", "vlan_id", "wlan_id", "wxrule_id", "wxrule_usage", "x", "x_m", "y", "y_m", "zones"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsWirelessClient object to a map representation for JSON marshaling.
func (s StatsWirelessClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Accuracy != nil {
        structMap["accuracy"] = s.Accuracy
    }
    if s.AirespaceIfname != nil {
        structMap["airespace_ifname"] = s.AirespaceIfname
    }
    if s.Airwatch != nil {
        structMap["airwatch"] = s.Airwatch.toMap()
    }
    structMap["ap_id"] = s.ApId
    structMap["ap_mac"] = s.ApMac
    structMap["band"] = s.Band
    structMap["channel"] = s.Channel
    structMap["dual_band"] = s.DualBand
    structMap["family"] = s.Family
    if s.Guest != nil {
        structMap["guest"] = s.Guest.toMap()
    }
    structMap["hostname"] = s.Hostname
    structMap["idle_time"] = s.IdleTime
    structMap["ip"] = s.Ip
    structMap["is_guest"] = s.IsGuest
    structMap["key_mgmt"] = s.KeyMgmt
    structMap["last_seen"] = s.LastSeen
    structMap["mac"] = s.Mac
    structMap["manufacture"] = s.Manufacture
    if s.MapId != nil {
        structMap["map_id"] = s.MapId
    }
    structMap["model"] = s.Model
    if s.NumLocatingAps != nil {
        structMap["num_locating_aps"] = s.NumLocatingAps
    }
    structMap["os"] = s.Os
    structMap["power_saving"] = s.PowerSaving
    structMap["proto"] = s.Proto
    if s.PskId != nil {
        structMap["psk_id"] = s.PskId
    }
    structMap["rssi"] = s.Rssi
    if s.Rssizones != nil {
        structMap["rssizones"] = s.Rssizones
    }
    structMap["rx_bps"] = s.RxBps
    structMap["rx_bytes"] = s.RxBytes
    structMap["rx_packets"] = s.RxPackets
    structMap["rx_rate"] = s.RxRate
    structMap["rx_retries"] = s.RxRetries
    structMap["snr"] = s.Snr
    structMap["ssid"] = s.Ssid
    structMap["tx_bps"] = s.TxBps
    structMap["tx_bytes"] = s.TxBytes
    structMap["tx_packets"] = s.TxPackets
    structMap["tx_rate"] = s.TxRate
    structMap["tx_retries"] = s.TxRetries
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    structMap["uptime"] = s.Uptime
    structMap["username"] = s.Username
    if s.Vbeacons != nil {
        structMap["vbeacons"] = s.Vbeacons
    }
    if s.VlanId != nil {
        structMap["vlan_id"] = s.VlanId
    }
    structMap["wlan_id"] = s.WlanId
    if s.WxruleId != nil {
        structMap["wxrule_id"] = s.WxruleId
    }
    if s.WxruleUsage != nil {
        structMap["wxrule_usage"] = s.WxruleUsage
    }
    if s.X != nil {
        structMap["x"] = s.X
    }
    if s.XM != nil {
        structMap["x_m"] = s.XM
    }
    if s.Y != nil {
        structMap["y"] = s.Y
    }
    if s.YM != nil {
        structMap["y_m"] = s.YM
    }
    if s.Zones != nil {
        structMap["zones"] = s.Zones
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsWirelessClient.
// It customizes the JSON unmarshaling process for StatsWirelessClient objects.
func (s *StatsWirelessClient) UnmarshalJSON(input []byte) error {
    var temp tempStatsWirelessClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "accuracy", "airespace_ifname", "airwatch", "ap_id", "ap_mac", "band", "channel", "dual_band", "family", "guest", "hostname", "idle_time", "ip", "is_guest", "key_mgmt", "last_seen", "mac", "manufacture", "map_id", "model", "num_locating_aps", "os", "power_saving", "proto", "psk_id", "rssi", "rssizones", "rx_bps", "rx_bytes", "rx_packets", "rx_rate", "rx_retries", "snr", "ssid", "tx_bps", "tx_bytes", "tx_packets", "tx_rate", "tx_retries", "type", "uptime", "username", "vbeacons", "vlan_id", "wlan_id", "wxrule_id", "wxrule_usage", "x", "x_m", "y", "y_m", "zones")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Accuracy = temp.Accuracy
    s.AirespaceIfname = temp.AirespaceIfname
    s.Airwatch = temp.Airwatch
    s.ApId = *temp.ApId
    s.ApMac = *temp.ApMac
    s.Band = *temp.Band
    s.Channel = *temp.Channel
    s.DualBand = *temp.DualBand
    s.Family = *temp.Family
    s.Guest = temp.Guest
    s.Hostname = *temp.Hostname
    s.IdleTime = *temp.IdleTime
    s.Ip = *temp.Ip
    s.IsGuest = *temp.IsGuest
    s.KeyMgmt = *temp.KeyMgmt
    s.LastSeen = *temp.LastSeen
    s.Mac = *temp.Mac
    s.Manufacture = *temp.Manufacture
    s.MapId = temp.MapId
    s.Model = *temp.Model
    s.NumLocatingAps = temp.NumLocatingAps
    s.Os = *temp.Os
    s.PowerSaving = *temp.PowerSaving
    s.Proto = *temp.Proto
    s.PskId = temp.PskId
    s.Rssi = *temp.Rssi
    s.Rssizones = temp.Rssizones
    s.RxBps = *temp.RxBps
    s.RxBytes = *temp.RxBytes
    s.RxPackets = *temp.RxPackets
    s.RxRate = *temp.RxRate
    s.RxRetries = *temp.RxRetries
    s.Snr = *temp.Snr
    s.Ssid = *temp.Ssid
    s.TxBps = *temp.TxBps
    s.TxBytes = *temp.TxBytes
    s.TxPackets = *temp.TxPackets
    s.TxRate = *temp.TxRate
    s.TxRetries = *temp.TxRetries
    s.Type = temp.Type
    s.Uptime = *temp.Uptime
    s.Username = *temp.Username
    s.Vbeacons = temp.Vbeacons
    s.VlanId = temp.VlanId
    s.WlanId = *temp.WlanId
    s.WxruleId = temp.WxruleId
    s.WxruleUsage = temp.WxruleUsage
    s.X = temp.X
    s.XM = temp.XM
    s.Y = temp.Y
    s.YM = temp.YM
    s.Zones = temp.Zones
    return nil
}

// tempStatsWirelessClient is a temporary struct used for validating the fields of StatsWirelessClient.
type tempStatsWirelessClient  struct {
    Accuracy        *int                             `json:"accuracy,omitempty"`
    AirespaceIfname *string                          `json:"airespace_ifname,omitempty"`
    Airwatch        *StatsWirelessClientAirwatch     `json:"airwatch,omitempty"`
    ApId            *uuid.UUID                       `json:"ap_id"`
    ApMac           *string                          `json:"ap_mac"`
    Band            *Dot11BandEnum                   `json:"band"`
    Channel         *int                             `json:"channel"`
    DualBand        *bool                            `json:"dual_band"`
    Family          *string                          `json:"family"`
    Guest           *StatsWirelessClientGuest        `json:"guest,omitempty"`
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
    Rssizones       []StatsWirelessClientRssiZone    `json:"rssizones,omitempty"`
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
    Vbeacons        []StatsWirelessClientVbeacon     `json:"vbeacons,omitempty"`
    VlanId          *int                             `json:"vlan_id,omitempty"`
    WlanId          *uuid.UUID                       `json:"wlan_id"`
    WxruleId        *uuid.UUID                       `json:"wxrule_id,omitempty"`
    WxruleUsage     []StatsWirelessClientWxruleUsage `json:"wxrule_usage,omitempty"`
    X               *float64                         `json:"x,omitempty"`
    XM              *float64                         `json:"x_m,omitempty"`
    Y               *float64                         `json:"y,omitempty"`
    YM              *float64                         `json:"y_m,omitempty"`
    Zones           []StatsWirelessClientZone        `json:"zones,omitempty"`
}

func (s *tempStatsWirelessClient) validate() error {
    var errs []string
    if s.ApId == nil {
        errs = append(errs, "required field `ap_id` is missing for type `stats_wireless_client`")
    }
    if s.ApMac == nil {
        errs = append(errs, "required field `ap_mac` is missing for type `stats_wireless_client`")
    }
    if s.Band == nil {
        errs = append(errs, "required field `band` is missing for type `stats_wireless_client`")
    }
    if s.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `stats_wireless_client`")
    }
    if s.DualBand == nil {
        errs = append(errs, "required field `dual_band` is missing for type `stats_wireless_client`")
    }
    if s.Family == nil {
        errs = append(errs, "required field `family` is missing for type `stats_wireless_client`")
    }
    if s.Hostname == nil {
        errs = append(errs, "required field `hostname` is missing for type `stats_wireless_client`")
    }
    if s.IdleTime == nil {
        errs = append(errs, "required field `idle_time` is missing for type `stats_wireless_client`")
    }
    if s.Ip == nil {
        errs = append(errs, "required field `ip` is missing for type `stats_wireless_client`")
    }
    if s.IsGuest == nil {
        errs = append(errs, "required field `is_guest` is missing for type `stats_wireless_client`")
    }
    if s.KeyMgmt == nil {
        errs = append(errs, "required field `key_mgmt` is missing for type `stats_wireless_client`")
    }
    if s.LastSeen == nil {
        errs = append(errs, "required field `last_seen` is missing for type `stats_wireless_client`")
    }
    if s.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `stats_wireless_client`")
    }
    if s.Manufacture == nil {
        errs = append(errs, "required field `manufacture` is missing for type `stats_wireless_client`")
    }
    if s.Model == nil {
        errs = append(errs, "required field `model` is missing for type `stats_wireless_client`")
    }
    if s.Os == nil {
        errs = append(errs, "required field `os` is missing for type `stats_wireless_client`")
    }
    if s.PowerSaving == nil {
        errs = append(errs, "required field `power_saving` is missing for type `stats_wireless_client`")
    }
    if s.Proto == nil {
        errs = append(errs, "required field `proto` is missing for type `stats_wireless_client`")
    }
    if s.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `stats_wireless_client`")
    }
    if s.RxBps == nil {
        errs = append(errs, "required field `rx_bps` is missing for type `stats_wireless_client`")
    }
    if s.RxBytes == nil {
        errs = append(errs, "required field `rx_bytes` is missing for type `stats_wireless_client`")
    }
    if s.RxPackets == nil {
        errs = append(errs, "required field `rx_packets` is missing for type `stats_wireless_client`")
    }
    if s.RxRate == nil {
        errs = append(errs, "required field `rx_rate` is missing for type `stats_wireless_client`")
    }
    if s.RxRetries == nil {
        errs = append(errs, "required field `rx_retries` is missing for type `stats_wireless_client`")
    }
    if s.Snr == nil {
        errs = append(errs, "required field `snr` is missing for type `stats_wireless_client`")
    }
    if s.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `stats_wireless_client`")
    }
    if s.TxBps == nil {
        errs = append(errs, "required field `tx_bps` is missing for type `stats_wireless_client`")
    }
    if s.TxBytes == nil {
        errs = append(errs, "required field `tx_bytes` is missing for type `stats_wireless_client`")
    }
    if s.TxPackets == nil {
        errs = append(errs, "required field `tx_packets` is missing for type `stats_wireless_client`")
    }
    if s.TxRate == nil {
        errs = append(errs, "required field `tx_rate` is missing for type `stats_wireless_client`")
    }
    if s.TxRetries == nil {
        errs = append(errs, "required field `tx_retries` is missing for type `stats_wireless_client`")
    }
    if s.Uptime == nil {
        errs = append(errs, "required field `uptime` is missing for type `stats_wireless_client`")
    }
    if s.Username == nil {
        errs = append(errs, "required field `username` is missing for type `stats_wireless_client`")
    }
    if s.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `stats_wireless_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
