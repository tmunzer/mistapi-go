package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// ClientWireless represents a ClientWireless struct.
type ClientWireless struct {
    // List of AP MAC Addresses the client was connected to
    Ap                   []string               `json:"ap,omitempty"`
    // only when client has the Marvis Client app running. List of the versions of the Marvis Client
    AppVersion           []string               `json:"app_version,omitempty"`
    // Wi-Fi Radio band
    Band                 *string                `json:"band,omitempty"`
    // only when client has the Marvis Client app running. List of the type of device type detected
    Device               []string               `json:"device,omitempty"`
    Ftc                  *bool                  `json:"ftc,omitempty"`
    // only when client has the Marvis Client app running. Type of Wi-Fi adapter
    Hardware             *string                `json:"hardware,omitempty"`
    // list of hostname detected for this client
    Hostname             []string               `json:"hostname,omitempty"`
    // list if the ip addresses detected for this client
    Ip                   []string               `json:"ip,omitempty"`
    // latest AP where the client is/was connected to
    LastAp               *string                `json:"last_ap,omitempty"`
    // latest type of device we identified (e.g. iPhone, Mac, …)
    LastDevice           *string                `json:"last_device,omitempty"`
    // only when client has the Marvis Client app running. Same as "firmware"
    LastFirmware         *string                `json:"last_firmware,omitempty"`
    // latest hostname we detected for the client
    LastHostname         *string                `json:"last_hostname,omitempty"`
    // latest ip address we detected for the client
    LastIp               *string                `json:"last_ip,omitempty"`
    // only when client has the Marvis Client app running. latest client hardware model we detected for the client
    LastModel            *string                `json:"last_model,omitempty"`
    // only when client has the Marvis Client app running. Latest version of OS Type we detected for the client
    LastOs               *string                `json:"last_os,omitempty"`
    // only when client has the Marvis Client app running. Latest version of OS Version we detected for the client
    LastOsVersion        *string                `json:"last_os_version,omitempty"`
    // only for PPSK authentication. Latest PPSK ID used by the client
    LastPskId            *uuid.UUID             `json:"last_psk_id,omitempty"`
    // only for PPSK authentication. Latest PPSK Name used by the client
    LastPskName          *string                `json:"last_psk_name,omitempty"`
    // Name of the latest SSID (WLAN) the client is/was connected to client
    LastSsid             *string                `json:"last_ssid,omitempty"`
    // only for 802.1X authentication. Latest username used by the client
    LastUsername         *string                `json:"last_username,omitempty"`
    // latest VLAN ID assigned to the client
    LastVlan             *int                   `json:"last_vlan,omitempty"`
    // ID of the latest SSID (WLAN) the client is/was connected to
    LastWlanId           *uuid.UUID             `json:"last_wlan_id,omitempty"`
    // Client MAC Address
    Mac                  *string                `json:"mac,omitempty"`
    // manufacturer of the client hardware (MAC OUI based)
    Mfg                  *string                `json:"mfg,omitempty"`
    // only when client has the Marvis Client app running. Client hardware model
    Model                *string                `json:"model,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // only when client is having the Marvis Client app running. List of OS detected for the client
    Os                   []string               `json:"os,omitempty"`
    // only when client is having the Marvis Client app running. List of OS version detected for the client
    OsVersion            []string               `json:"os_version,omitempty"`
    // 802.11 amendment
    Protocol             *string                `json:"protocol,omitempty"`
    // list of IDs of the PPSK used by the client
    PskId                []uuid.UUID            `json:"psk_id,omitempty"`
    // list of names of the PPSK used by the client
    PskName              []string               `json:"psk_name,omitempty"`
    // whether the client is using randomized MAC Address or not
    RandomMac            *bool                  `json:"random_mac,omitempty"`
    // only when client has the Marvis Client app running. List of Marvis Client SDK version detected for the client
    SdkVersion           []string               `json:"sdk_version,omitempty"`
    // Mist Site ID where the client is connected
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // list of Mist Site IDs where the client was connected
    SiteIds              []uuid.UUID            `json:"site_ids,omitempty"`
    // list of the WLAN names the client was connected to
    Ssid                 []string               `json:"ssid,omitempty"`
    // when the data has been updated
    Timestamp            *float64               `json:"timestamp,omitempty"`
    // only for 802.1X authentication. List of usernames used by the client
    Username             []string               `json:"username,omitempty"`
    // list of vlans that have been assigned to the client
    Vlan                 []int                  `json:"vlan,omitempty"`
    // list of IDs of WLANs the client was connected to
    WlanId               []uuid.UUID            `json:"wlan_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClientWireless.
// It customizes the JSON marshaling process for ClientWireless objects.
func (c ClientWireless) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap", "app_version", "band", "device", "ftc", "hardware", "hostname", "ip", "last_ap", "last_device", "last_firmware", "last_hostname", "last_ip", "last_model", "last_os", "last_os_version", "last_psk_id", "last_psk_name", "last_ssid", "last_username", "last_vlan", "last_wlan_id", "mac", "mfg", "model", "org_id", "os", "os_version", "protocol", "psk_id", "psk_name", "random_mac", "sdk_version", "site_id", "site_ids", "ssid", "timestamp", "username", "vlan", "wlan_id"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ClientWireless object to a map representation for JSON marshaling.
func (c ClientWireless) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ap != nil {
        structMap["ap"] = c.Ap
    }
    if c.AppVersion != nil {
        structMap["app_version"] = c.AppVersion
    }
    if c.Band != nil {
        structMap["band"] = c.Band
    }
    if c.Device != nil {
        structMap["device"] = c.Device
    }
    if c.Ftc != nil {
        structMap["ftc"] = c.Ftc
    }
    if c.Hardware != nil {
        structMap["hardware"] = c.Hardware
    }
    if c.Hostname != nil {
        structMap["hostname"] = c.Hostname
    }
    if c.Ip != nil {
        structMap["ip"] = c.Ip
    }
    if c.LastAp != nil {
        structMap["last_ap"] = c.LastAp
    }
    if c.LastDevice != nil {
        structMap["last_device"] = c.LastDevice
    }
    if c.LastFirmware != nil {
        structMap["last_firmware"] = c.LastFirmware
    }
    if c.LastHostname != nil {
        structMap["last_hostname"] = c.LastHostname
    }
    if c.LastIp != nil {
        structMap["last_ip"] = c.LastIp
    }
    if c.LastModel != nil {
        structMap["last_model"] = c.LastModel
    }
    if c.LastOs != nil {
        structMap["last_os"] = c.LastOs
    }
    if c.LastOsVersion != nil {
        structMap["last_os_version"] = c.LastOsVersion
    }
    if c.LastPskId != nil {
        structMap["last_psk_id"] = c.LastPskId
    }
    if c.LastPskName != nil {
        structMap["last_psk_name"] = c.LastPskName
    }
    if c.LastSsid != nil {
        structMap["last_ssid"] = c.LastSsid
    }
    if c.LastUsername != nil {
        structMap["last_username"] = c.LastUsername
    }
    if c.LastVlan != nil {
        structMap["last_vlan"] = c.LastVlan
    }
    if c.LastWlanId != nil {
        structMap["last_wlan_id"] = c.LastWlanId
    }
    if c.Mac != nil {
        structMap["mac"] = c.Mac
    }
    if c.Mfg != nil {
        structMap["mfg"] = c.Mfg
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.OrgId != nil {
        structMap["org_id"] = c.OrgId
    }
    if c.Os != nil {
        structMap["os"] = c.Os
    }
    if c.OsVersion != nil {
        structMap["os_version"] = c.OsVersion
    }
    if c.Protocol != nil {
        structMap["protocol"] = c.Protocol
    }
    if c.PskId != nil {
        structMap["psk_id"] = c.PskId
    }
    if c.PskName != nil {
        structMap["psk_name"] = c.PskName
    }
    if c.RandomMac != nil {
        structMap["random_mac"] = c.RandomMac
    }
    if c.SdkVersion != nil {
        structMap["sdk_version"] = c.SdkVersion
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.SiteIds != nil {
        structMap["site_ids"] = c.SiteIds
    }
    if c.Ssid != nil {
        structMap["ssid"] = c.Ssid
    }
    if c.Timestamp != nil {
        structMap["timestamp"] = c.Timestamp
    }
    if c.Username != nil {
        structMap["username"] = c.Username
    }
    if c.Vlan != nil {
        structMap["vlan"] = c.Vlan
    }
    if c.WlanId != nil {
        structMap["wlan_id"] = c.WlanId
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientWireless.
// It customizes the JSON unmarshaling process for ClientWireless objects.
func (c *ClientWireless) UnmarshalJSON(input []byte) error {
    var temp tempClientWireless
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "app_version", "band", "device", "ftc", "hardware", "hostname", "ip", "last_ap", "last_device", "last_firmware", "last_hostname", "last_ip", "last_model", "last_os", "last_os_version", "last_psk_id", "last_psk_name", "last_ssid", "last_username", "last_vlan", "last_wlan_id", "mac", "mfg", "model", "org_id", "os", "os_version", "protocol", "psk_id", "psk_name", "random_mac", "sdk_version", "site_id", "site_ids", "ssid", "timestamp", "username", "vlan", "wlan_id")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Ap = temp.Ap
    c.AppVersion = temp.AppVersion
    c.Band = temp.Band
    c.Device = temp.Device
    c.Ftc = temp.Ftc
    c.Hardware = temp.Hardware
    c.Hostname = temp.Hostname
    c.Ip = temp.Ip
    c.LastAp = temp.LastAp
    c.LastDevice = temp.LastDevice
    c.LastFirmware = temp.LastFirmware
    c.LastHostname = temp.LastHostname
    c.LastIp = temp.LastIp
    c.LastModel = temp.LastModel
    c.LastOs = temp.LastOs
    c.LastOsVersion = temp.LastOsVersion
    c.LastPskId = temp.LastPskId
    c.LastPskName = temp.LastPskName
    c.LastSsid = temp.LastSsid
    c.LastUsername = temp.LastUsername
    c.LastVlan = temp.LastVlan
    c.LastWlanId = temp.LastWlanId
    c.Mac = temp.Mac
    c.Mfg = temp.Mfg
    c.Model = temp.Model
    c.OrgId = temp.OrgId
    c.Os = temp.Os
    c.OsVersion = temp.OsVersion
    c.Protocol = temp.Protocol
    c.PskId = temp.PskId
    c.PskName = temp.PskName
    c.RandomMac = temp.RandomMac
    c.SdkVersion = temp.SdkVersion
    c.SiteId = temp.SiteId
    c.SiteIds = temp.SiteIds
    c.Ssid = temp.Ssid
    c.Timestamp = temp.Timestamp
    c.Username = temp.Username
    c.Vlan = temp.Vlan
    c.WlanId = temp.WlanId
    return nil
}

// tempClientWireless is a temporary struct used for validating the fields of ClientWireless.
type tempClientWireless  struct {
    Ap            []string    `json:"ap,omitempty"`
    AppVersion    []string    `json:"app_version,omitempty"`
    Band          *string     `json:"band,omitempty"`
    Device        []string    `json:"device,omitempty"`
    Ftc           *bool       `json:"ftc,omitempty"`
    Hardware      *string     `json:"hardware,omitempty"`
    Hostname      []string    `json:"hostname,omitempty"`
    Ip            []string    `json:"ip,omitempty"`
    LastAp        *string     `json:"last_ap,omitempty"`
    LastDevice    *string     `json:"last_device,omitempty"`
    LastFirmware  *string     `json:"last_firmware,omitempty"`
    LastHostname  *string     `json:"last_hostname,omitempty"`
    LastIp        *string     `json:"last_ip,omitempty"`
    LastModel     *string     `json:"last_model,omitempty"`
    LastOs        *string     `json:"last_os,omitempty"`
    LastOsVersion *string     `json:"last_os_version,omitempty"`
    LastPskId     *uuid.UUID  `json:"last_psk_id,omitempty"`
    LastPskName   *string     `json:"last_psk_name,omitempty"`
    LastSsid      *string     `json:"last_ssid,omitempty"`
    LastUsername  *string     `json:"last_username,omitempty"`
    LastVlan      *int        `json:"last_vlan,omitempty"`
    LastWlanId    *uuid.UUID  `json:"last_wlan_id,omitempty"`
    Mac           *string     `json:"mac,omitempty"`
    Mfg           *string     `json:"mfg,omitempty"`
    Model         *string     `json:"model,omitempty"`
    OrgId         *uuid.UUID  `json:"org_id,omitempty"`
    Os            []string    `json:"os,omitempty"`
    OsVersion     []string    `json:"os_version,omitempty"`
    Protocol      *string     `json:"protocol,omitempty"`
    PskId         []uuid.UUID `json:"psk_id,omitempty"`
    PskName       []string    `json:"psk_name,omitempty"`
    RandomMac     *bool       `json:"random_mac,omitempty"`
    SdkVersion    []string    `json:"sdk_version,omitempty"`
    SiteId        *uuid.UUID  `json:"site_id,omitempty"`
    SiteIds       []uuid.UUID `json:"site_ids,omitempty"`
    Ssid          []string    `json:"ssid,omitempty"`
    Timestamp     *float64    `json:"timestamp,omitempty"`
    Username      []string    `json:"username,omitempty"`
    Vlan          []int       `json:"vlan,omitempty"`
    WlanId        []uuid.UUID `json:"wlan_id,omitempty"`
}
