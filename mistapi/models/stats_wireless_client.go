package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "github.com/google/uuid"
    "strings"
)

// StatsWirelessClient represents a StatsWirelessClient struct.
type StatsWirelessClient struct {
    // Estimated client location accuracy, in meter
    Accuracy             *int                             `json:"accuracy,omitempty"`
    AirespaceIfname      *string                          `json:"airespace_ifname,omitempty"`
    // Information if airwatch enabled
    Airwatch             *StatsWirelessClientAirwatch     `json:"airwatch,omitempty"`
    // AP ID the client is connected to
    ApId                 uuid.UUID                        `json:"ap_id"`
    // AP the client is connected to
    ApMac                string                           `json:"ap_mac"`
    // enum: `24`, `5`, `6`
    Band                 Dot11BandEnum                    `json:"band"`
    // Current channel
    Channel              int                              `json:"channel"`
    // Whether the client is dual_band capable (determined by whether we’ve seen probe requests from both bands)
    DualBand             *bool                            `json:"dual_band,omitempty"`
    // Device family, through fingerprinting. iPod / Nexus Galaxy / Windows Mobile or CE …
    Family               string                           `json:"family"`
    // Guest
    Guest                *Guest                           `json:"guest,omitempty"`
    // Hostname that we learned from sniffing DHCP
    Hostname             string                           `json:"hostname"`
    // How long, in seconds, has the client been idle (since the last RX packet)
    IdleTime             *float64                         `json:"idle_time,omitempty"`
    Ip                   string                           `json:"ip"`
    // Whether this is a guest
    IsGuest              bool                             `json:"is_guest"`
    // E.g. WPA2-PSK/CCMP
    KeyMgmt              string                           `json:"key_mgmt"`
    // Last seen timestamp
    LastSeen             *float64                         `json:"last_seen"`
    // Client mac
    Mac                  string                           `json:"mac"`
    // Device manufacture, through fingerprinting or OUI
    Manufacture          string                           `json:"manufacture"`
    // Estimated client location - map_id
    MapId                *uuid.UUID                       `json:"map_id,omitempty"`
    // Device model, may be available if we can identify them
    Model                string                           `json:"model"`
    // Number of APs used to locate this client
    NumLocatingAps       *int                             `json:"num_locating_aps,omitempty"`
    // Device os, through fingerprinting
    Os                   string                           `json:"os"`
    // If it’s currently in power-save mode
    PowerSaving          *bool                            `json:"power_saving,omitempty"`
    // enum: `a`, `ac`, `ax`, `b`, `g`, `n`
    Proto                Dot11ProtoEnum                   `json:"proto"`
    // PSK id (if multi-psk is used)
    PskId                *uuid.UUID                       `json:"psk_id,omitempty"`
    // Signal strength
    Rssi                 float64                          `json:"rssi"`
    // List of rssizone_id’s where client is in and since when (if known)
    Rssizones            []StatsWirelessClientRssiZone    `json:"rssizones,omitempty"`
    // Rate of receiving traffic, bits/seconds, last known
    RxBps                *int64                           `json:"rx_bps"`
    // Amount of traffic received since connection
    RxBytes              *int64                           `json:"rx_bytes"`
    // Amount of packets received since connection
    RxPackets            Optional[int64]                  `json:"rx_packets"`
    // RX Rate, Mbps
    RxRate               *float64                         `json:"rx_rate"`
    // Amount of rx retries
    RxRetries            *int                             `json:"rx_retries"`
    // Signal over noise
    Snr                  float64                          `json:"snr"`
    // SSID the client is connected to
    Ssid                 string                           `json:"ssid"`
    // Rate of transmitting traffic, bits/seconds, last known
    TxBps                *int64                           `json:"tx_bps"`
    // Amount of traffic sent since connection
    TxBytes              *int64                           `json:"tx_bytes"`
    // Amount of packets sent since connection
    TxPackets            Optional[int64]                  `json:"tx_packets"`
    // TX Rate, Mbps
    TxRate               *float64                         `json:"tx_rate"`
    // Amount of tx retries
    TxRetries            *int                             `json:"tx_retries"`
    // Client’s type, regular / vip / resource / blocked (if client object is created)
    Type                 *string                          `json:"type,omitempty"`
    // How long, in seconds, has the client been connected
    Uptime               *float64                         `json:"uptime,omitempty"`
    // Username that we learned from 802.1X exchange or Per_user PSK or User Portal
    Username             *string                          `json:"username,omitempty"`
    // List of beacon_id’s where the client is in and since when (if known)
    Vbeacons             []StatsWirelessClientVbeacon     `json:"vbeacons,omitempty"`
    // VLAN id, could be empty (from older AP)
    VlanId               *string                          `json:"vlan_id,omitempty"`
    // WLAN ID the client is connected to
    WlanId               uuid.UUID                        `json:"wlan_id"`
    // Current WxlanRule using for a Client or an authorized Guest (portal user). null if default rule is matched.
    WxruleId             *uuid.UUID                       `json:"wxrule_id,omitempty"`
    // Current WxlanRule usage per tag_id
    WxruleUsage          []StatsWirelessClientWxruleUsage `json:"wxrule_usage,omitempty"`
    // Estimated client location in pixels
    X                    *float64                         `json:"x,omitempty"`
    // Estimated client location in meter
    XM                   *float64                         `json:"x_m,omitempty"`
    // Estimated client location in pixels
    Y                    *float64                         `json:"y,omitempty"`
    // Estimated client location in meter
    YM                   *float64                         `json:"y_m,omitempty"`
    // List of zone_id’s where client is in and since when (if known)
    Zones                []StatsWirelessClientZone        `json:"zones,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for StatsWirelessClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsWirelessClient) String() string {
    return fmt.Sprintf(
    	"StatsWirelessClient[Accuracy=%v, AirespaceIfname=%v, Airwatch=%v, ApId=%v, ApMac=%v, Band=%v, Channel=%v, DualBand=%v, Family=%v, Guest=%v, Hostname=%v, IdleTime=%v, Ip=%v, IsGuest=%v, KeyMgmt=%v, LastSeen=%v, Mac=%v, Manufacture=%v, MapId=%v, Model=%v, NumLocatingAps=%v, Os=%v, PowerSaving=%v, Proto=%v, PskId=%v, Rssi=%v, Rssizones=%v, RxBps=%v, RxBytes=%v, RxPackets=%v, RxRate=%v, RxRetries=%v, Snr=%v, Ssid=%v, TxBps=%v, TxBytes=%v, TxPackets=%v, TxRate=%v, TxRetries=%v, Type=%v, Uptime=%v, Username=%v, Vbeacons=%v, VlanId=%v, WlanId=%v, WxruleId=%v, WxruleUsage=%v, X=%v, XM=%v, Y=%v, YM=%v, Zones=%v, AdditionalProperties=%v]",
    	s.Accuracy, s.AirespaceIfname, s.Airwatch, s.ApId, s.ApMac, s.Band, s.Channel, s.DualBand, s.Family, s.Guest, s.Hostname, s.IdleTime, s.Ip, s.IsGuest, s.KeyMgmt, s.LastSeen, s.Mac, s.Manufacture, s.MapId, s.Model, s.NumLocatingAps, s.Os, s.PowerSaving, s.Proto, s.PskId, s.Rssi, s.Rssizones, s.RxBps, s.RxBytes, s.RxPackets, s.RxRate, s.RxRetries, s.Snr, s.Ssid, s.TxBps, s.TxBytes, s.TxPackets, s.TxRate, s.TxRetries, s.Type, s.Uptime, s.Username, s.Vbeacons, s.VlanId, s.WlanId, s.WxruleId, s.WxruleUsage, s.X, s.XM, s.Y, s.YM, s.Zones, s.AdditionalProperties)
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
    if s.DualBand != nil {
        structMap["dual_band"] = s.DualBand
    }
    structMap["family"] = s.Family
    if s.Guest != nil {
        structMap["guest"] = s.Guest.toMap()
    }
    structMap["hostname"] = s.Hostname
    if s.IdleTime != nil {
        structMap["idle_time"] = s.IdleTime
    }
    structMap["ip"] = s.Ip
    structMap["is_guest"] = s.IsGuest
    structMap["key_mgmt"] = s.KeyMgmt
    if s.LastSeen != nil {
        structMap["last_seen"] = s.LastSeen
    } else {
        structMap["last_seen"] = nil
    }
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
    if s.PowerSaving != nil {
        structMap["power_saving"] = s.PowerSaving
    }
    structMap["proto"] = s.Proto
    if s.PskId != nil {
        structMap["psk_id"] = s.PskId
    }
    structMap["rssi"] = s.Rssi
    if s.Rssizones != nil {
        structMap["rssizones"] = s.Rssizones
    }
    if s.RxBps != nil {
        structMap["rx_bps"] = s.RxBps
    } else {
        structMap["rx_bps"] = nil
    }
    if s.RxBytes != nil {
        structMap["rx_bytes"] = s.RxBytes
    } else {
        structMap["rx_bytes"] = nil
    }
    if s.RxPackets.IsValueSet() {
        if s.RxPackets.Value() != nil {
            structMap["rx_packets"] = s.RxPackets.Value()
        } else {
            structMap["rx_packets"] = nil
        }
    }
    if s.RxRate != nil {
        structMap["rx_rate"] = s.RxRate
    } else {
        structMap["rx_rate"] = nil
    }
    if s.RxRetries != nil {
        structMap["rx_retries"] = s.RxRetries
    } else {
        structMap["rx_retries"] = nil
    }
    structMap["snr"] = s.Snr
    structMap["ssid"] = s.Ssid
    if s.TxBps != nil {
        structMap["tx_bps"] = s.TxBps
    } else {
        structMap["tx_bps"] = nil
    }
    if s.TxBytes != nil {
        structMap["tx_bytes"] = s.TxBytes
    } else {
        structMap["tx_bytes"] = nil
    }
    if s.TxPackets.IsValueSet() {
        if s.TxPackets.Value() != nil {
            structMap["tx_packets"] = s.TxPackets.Value()
        } else {
            structMap["tx_packets"] = nil
        }
    }
    if s.TxRate != nil {
        structMap["tx_rate"] = s.TxRate
    } else {
        structMap["tx_rate"] = nil
    }
    if s.TxRetries != nil {
        structMap["tx_retries"] = s.TxRetries
    } else {
        structMap["tx_retries"] = nil
    }
    if s.Type != nil {
        structMap["type"] = s.Type
    }
    if s.Uptime != nil {
        structMap["uptime"] = s.Uptime
    }
    if s.Username != nil {
        structMap["username"] = s.Username
    }
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
    s.DualBand = temp.DualBand
    s.Family = *temp.Family
    s.Guest = temp.Guest
    s.Hostname = *temp.Hostname
    s.IdleTime = temp.IdleTime
    s.Ip = *temp.Ip
    s.IsGuest = *temp.IsGuest
    s.KeyMgmt = *temp.KeyMgmt
    s.LastSeen = temp.LastSeen
    s.Mac = *temp.Mac
    s.Manufacture = *temp.Manufacture
    s.MapId = temp.MapId
    s.Model = *temp.Model
    s.NumLocatingAps = temp.NumLocatingAps
    s.Os = *temp.Os
    s.PowerSaving = temp.PowerSaving
    s.Proto = *temp.Proto
    s.PskId = temp.PskId
    s.Rssi = *temp.Rssi
    s.Rssizones = temp.Rssizones
    s.RxBps = temp.RxBps
    s.RxBytes = temp.RxBytes
    s.RxPackets = temp.RxPackets
    s.RxRate = temp.RxRate
    s.RxRetries = temp.RxRetries
    s.Snr = *temp.Snr
    s.Ssid = *temp.Ssid
    s.TxBps = temp.TxBps
    s.TxBytes = temp.TxBytes
    s.TxPackets = temp.TxPackets
    s.TxRate = temp.TxRate
    s.TxRetries = temp.TxRetries
    s.Type = temp.Type
    s.Uptime = temp.Uptime
    s.Username = temp.Username
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
    DualBand        *bool                            `json:"dual_band,omitempty"`
    Family          *string                          `json:"family"`
    Guest           *Guest                           `json:"guest,omitempty"`
    Hostname        *string                          `json:"hostname"`
    IdleTime        *float64                         `json:"idle_time,omitempty"`
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
    PowerSaving     *bool                            `json:"power_saving,omitempty"`
    Proto           *Dot11ProtoEnum                  `json:"proto"`
    PskId           *uuid.UUID                       `json:"psk_id,omitempty"`
    Rssi            *float64                         `json:"rssi"`
    Rssizones       []StatsWirelessClientRssiZone    `json:"rssizones,omitempty"`
    RxBps           *int64                           `json:"rx_bps"`
    RxBytes         *int64                           `json:"rx_bytes"`
    RxPackets       Optional[int64]                  `json:"rx_packets"`
    RxRate          *float64                         `json:"rx_rate"`
    RxRetries       *int                             `json:"rx_retries"`
    Snr             *float64                         `json:"snr"`
    Ssid            *string                          `json:"ssid"`
    TxBps           *int64                           `json:"tx_bps"`
    TxBytes         *int64                           `json:"tx_bytes"`
    TxPackets       Optional[int64]                  `json:"tx_packets"`
    TxRate          *float64                         `json:"tx_rate"`
    TxRetries       *int                             `json:"tx_retries"`
    Type            *string                          `json:"type,omitempty"`
    Uptime          *float64                         `json:"uptime,omitempty"`
    Username        *string                          `json:"username,omitempty"`
    Vbeacons        []StatsWirelessClientVbeacon     `json:"vbeacons,omitempty"`
    VlanId          *string                          `json:"vlan_id,omitempty"`
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
    if s.Family == nil {
        errs = append(errs, "required field `family` is missing for type `stats_wireless_client`")
    }
    if s.Hostname == nil {
        errs = append(errs, "required field `hostname` is missing for type `stats_wireless_client`")
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
    if s.Proto == nil {
        errs = append(errs, "required field `proto` is missing for type `stats_wireless_client`")
    }
    if s.Rssi == nil {
        errs = append(errs, "required field `rssi` is missing for type `stats_wireless_client`")
    }
    if s.Snr == nil {
        errs = append(errs, "required field `snr` is missing for type `stats_wireless_client`")
    }
    if s.Ssid == nil {
        errs = append(errs, "required field `ssid` is missing for type `stats_wireless_client`")
    }
    if s.WlanId == nil {
        errs = append(errs, "required field `wlan_id` is missing for type `stats_wireless_client`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
