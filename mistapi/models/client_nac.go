package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// ClientNac represents a ClientNac struct.
type ClientNac struct {
    Ap                   []string               `json:"ap,omitempty"`
    // enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap`
    AuthType             *NacAuthTypeEnum       `json:"auth_type,omitempty"`
    // When certificate based authentication is used, the CN from the certificates used for the specified duration
    CertCn               []string               `json:"cert_cn,omitempty"`
    // When certificate based authentication is used, the Issuer from the certificates used for the specified duration
    CertIssuer           []string               `json:"cert_issuer,omitempty"`
    // When certificate based authentication is used, the Serial from the certificates used for the specified duration
    CertSerial           []string               `json:"cert_serial,omitempty"`
    // When certificate based authentication is used, the Subject from the certificates used for the specified duration
    CertSubject          []string               `json:"cert_subject,omitempty"`
    // The known IP Addresses used by the client for the specified duration
    ClientIp             []string               `json:"client_ip,omitempty"`
    // MAC Address of the device (AP, Switch) the client is connected to
    DeviceMac            *string                `json:"device_mac,omitempty"`
    Group                *string                `json:"group,omitempty"`
    IdpId                *string                `json:"idp_id,omitempty"`
    IdpRole              []string               `json:"idp_role,omitempty"`
    // Latest AP where the client is/was connected to
    LastAp               *string                `json:"last_ap,omitempty"`
    // When certificate based authentication is used, the CN from the latest certificate used
    LastCertCn           *string                `json:"last_cert_cn,omitempty"`
    // When certificate based authentication is used, the expiration date from the latest certificate used
    LastCertExpiry       *float64               `json:"last_cert_expiry,omitempty"`
    // When certificate based authentication is used, the Issuer from the latest certificate used
    LastCertIssuer       *string                `json:"last_cert_issuer,omitempty"`
    // When certificate based authentication is used, the Serial from the latest certificate used
    LastCertSerial       *string                `json:"last_cert_serial,omitempty"`
    // When certificate based authentication is used, the Subject from the latest certificate used
    LastCertSubject      *string                `json:"last_cert_subject,omitempty"`
    // The last known IP Address for the client
    LastClientIp         *string                `json:"last_client_ip,omitempty"`
    // ID of the latest NAC Rule used to authenticate the client
    LastNacruleId        *string                `json:"last_nacrule_id,omitempty"`
    // Name of the latest NAC Rule used to authenticate the client
    LastNacruleName      *string                `json:"last_nacrule_name,omitempty"`
    // Vendor name of the NAS for the latest authentication
    LastNasVendor        *string                `json:"last_nas_vendor,omitempty"`
    // If Wired authentication, the latest Port-id the client was connected to
    LastPortId           *string                `json:"last_port_id,omitempty"`
    // If Wireless authentication, the latest SSID the client was connected to
    LastSsid             *string                `json:"last_ssid,omitempty"`
    // Latest Authentication status of the client. enum: `denied`, `permitted`, `session_started`, `session_stopped`
    LastStatus           *LastStatusEnum        `json:"last_status,omitempty"`
    // If dot1x authentication, the username used during the latest authentication. Otherwise, the MAC address of the client
    LastUsername         *string                `json:"last_username,omitempty"`
    // Latest VLAN ID assigned to the client
    LastVlan             *int                   `json:"last_vlan,omitempty"`
    // Client MAC address
    Mac                  *string                `json:"mac,omitempty"`
    // IDs of the NAC Rules used to authenticate the client for the specified duration
    NacruleId            []string               `json:"nacrule_id,omitempty"`
    NacruleMatched       *bool                  `json:"nacrule_matched,omitempty"`
    // Name of the NAC Rules used to authenticate the client for the specified duration
    NacruleName          []string               `json:"nacrule_name,omitempty"`
    // Vendor name of the NAS for the specified duration
    NasVendor            []string               `json:"nas_vendor,omitempty"`
    OrgId                *uuid.UUID             `json:"org_id,omitempty"`
    // Port-ids the client was connected to  for the specified duration
    PortId               []string               `json:"port_id,omitempty"`
    // Whether the client is using randomized MAC Address or not
    RandomMac            *bool                  `json:"random_mac,omitempty"`
    // List of Radius AVP returned by the Authentication Server
    RespAttrs            []string               `json:"resp_attrs,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    // SSIDs the client was connected to  for the specified duration
    Ssid                 []string               `json:"ssid,omitempty"`
    // Epoch (seconds)
    Timestamp            *float64               `json:"timestamp,omitempty"`
    // Type of network access. enum: `wireless`, `wired`
    Type                 *NacAccessTypeEnum     `json:"type,omitempty"`
    UsermacLabel         []string               `json:"usermac_label,omitempty"`
    // List of usernames that have been assigned to the client
    Username             []string               `json:"username,omitempty"`
    // List of vlans that have been assigned to the client
    Vlan                 []string               `json:"vlan,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ClientNac,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ClientNac) String() string {
    return fmt.Sprintf(
    	"ClientNac[Ap=%v, AuthType=%v, CertCn=%v, CertIssuer=%v, CertSerial=%v, CertSubject=%v, ClientIp=%v, DeviceMac=%v, Group=%v, IdpId=%v, IdpRole=%v, LastAp=%v, LastCertCn=%v, LastCertExpiry=%v, LastCertIssuer=%v, LastCertSerial=%v, LastCertSubject=%v, LastClientIp=%v, LastNacruleId=%v, LastNacruleName=%v, LastNasVendor=%v, LastPortId=%v, LastSsid=%v, LastStatus=%v, LastUsername=%v, LastVlan=%v, Mac=%v, NacruleId=%v, NacruleMatched=%v, NacruleName=%v, NasVendor=%v, OrgId=%v, PortId=%v, RandomMac=%v, RespAttrs=%v, SiteId=%v, Ssid=%v, Timestamp=%v, Type=%v, UsermacLabel=%v, Username=%v, Vlan=%v, AdditionalProperties=%v]",
    	c.Ap, c.AuthType, c.CertCn, c.CertIssuer, c.CertSerial, c.CertSubject, c.ClientIp, c.DeviceMac, c.Group, c.IdpId, c.IdpRole, c.LastAp, c.LastCertCn, c.LastCertExpiry, c.LastCertIssuer, c.LastCertSerial, c.LastCertSubject, c.LastClientIp, c.LastNacruleId, c.LastNacruleName, c.LastNasVendor, c.LastPortId, c.LastSsid, c.LastStatus, c.LastUsername, c.LastVlan, c.Mac, c.NacruleId, c.NacruleMatched, c.NacruleName, c.NasVendor, c.OrgId, c.PortId, c.RandomMac, c.RespAttrs, c.SiteId, c.Ssid, c.Timestamp, c.Type, c.UsermacLabel, c.Username, c.Vlan, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ClientNac.
// It customizes the JSON marshaling process for ClientNac objects.
func (c ClientNac) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "ap", "auth_type", "cert_cn", "cert_issuer", "cert_serial", "cert_subject", "client_ip", "device_mac", "group", "idp_id", "idp_role", "last_ap", "last_cert_cn", "last_cert_expiry", "last_cert_issuer", "last_cert_serial", "last_cert_subject", "last_client_ip", "last_nacrule_id", "last_nacrule_name", "last_nas_vendor", "last_port_id", "last_ssid", "last_status", "last_username", "last_vlan", "mac", "nacrule_id", "nacrule_matched", "nacrule_name", "nas_vendor", "org_id", "port_id", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "usermac_label", "username", "vlan"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ClientNac object to a map representation for JSON marshaling.
func (c ClientNac) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Ap != nil {
        structMap["ap"] = c.Ap
    }
    if c.AuthType != nil {
        structMap["auth_type"] = c.AuthType
    }
    if c.CertCn != nil {
        structMap["cert_cn"] = c.CertCn
    }
    if c.CertIssuer != nil {
        structMap["cert_issuer"] = c.CertIssuer
    }
    if c.CertSerial != nil {
        structMap["cert_serial"] = c.CertSerial
    }
    if c.CertSubject != nil {
        structMap["cert_subject"] = c.CertSubject
    }
    if c.ClientIp != nil {
        structMap["client_ip"] = c.ClientIp
    }
    if c.DeviceMac != nil {
        structMap["device_mac"] = c.DeviceMac
    }
    if c.Group != nil {
        structMap["group"] = c.Group
    }
    if c.IdpId != nil {
        structMap["idp_id"] = c.IdpId
    }
    if c.IdpRole != nil {
        structMap["idp_role"] = c.IdpRole
    }
    if c.LastAp != nil {
        structMap["last_ap"] = c.LastAp
    }
    if c.LastCertCn != nil {
        structMap["last_cert_cn"] = c.LastCertCn
    }
    if c.LastCertExpiry != nil {
        structMap["last_cert_expiry"] = c.LastCertExpiry
    }
    if c.LastCertIssuer != nil {
        structMap["last_cert_issuer"] = c.LastCertIssuer
    }
    if c.LastCertSerial != nil {
        structMap["last_cert_serial"] = c.LastCertSerial
    }
    if c.LastCertSubject != nil {
        structMap["last_cert_subject"] = c.LastCertSubject
    }
    if c.LastClientIp != nil {
        structMap["last_client_ip"] = c.LastClientIp
    }
    if c.LastNacruleId != nil {
        structMap["last_nacrule_id"] = c.LastNacruleId
    }
    if c.LastNacruleName != nil {
        structMap["last_nacrule_name"] = c.LastNacruleName
    }
    if c.LastNasVendor != nil {
        structMap["last_nas_vendor"] = c.LastNasVendor
    }
    if c.LastPortId != nil {
        structMap["last_port_id"] = c.LastPortId
    }
    if c.LastSsid != nil {
        structMap["last_ssid"] = c.LastSsid
    }
    if c.LastStatus != nil {
        structMap["last_status"] = c.LastStatus
    }
    if c.LastUsername != nil {
        structMap["last_username"] = c.LastUsername
    }
    if c.LastVlan != nil {
        structMap["last_vlan"] = c.LastVlan
    }
    if c.Mac != nil {
        structMap["mac"] = c.Mac
    }
    if c.NacruleId != nil {
        structMap["nacrule_id"] = c.NacruleId
    }
    if c.NacruleMatched != nil {
        structMap["nacrule_matched"] = c.NacruleMatched
    }
    if c.NacruleName != nil {
        structMap["nacrule_name"] = c.NacruleName
    }
    if c.NasVendor != nil {
        structMap["nas_vendor"] = c.NasVendor
    }
    if c.OrgId != nil {
        structMap["org_id"] = c.OrgId
    }
    if c.PortId != nil {
        structMap["port_id"] = c.PortId
    }
    if c.RandomMac != nil {
        structMap["random_mac"] = c.RandomMac
    }
    if c.RespAttrs != nil {
        structMap["resp_attrs"] = c.RespAttrs
    }
    if c.SiteId != nil {
        structMap["site_id"] = c.SiteId
    }
    if c.Ssid != nil {
        structMap["ssid"] = c.Ssid
    }
    if c.Timestamp != nil {
        structMap["timestamp"] = c.Timestamp
    }
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    if c.UsermacLabel != nil {
        structMap["usermac_label"] = c.UsermacLabel
    }
    if c.Username != nil {
        structMap["username"] = c.Username
    }
    if c.Vlan != nil {
        structMap["vlan"] = c.Vlan
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClientNac.
// It customizes the JSON unmarshaling process for ClientNac objects.
func (c *ClientNac) UnmarshalJSON(input []byte) error {
    var temp tempClientNac
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "auth_type", "cert_cn", "cert_issuer", "cert_serial", "cert_subject", "client_ip", "device_mac", "group", "idp_id", "idp_role", "last_ap", "last_cert_cn", "last_cert_expiry", "last_cert_issuer", "last_cert_serial", "last_cert_subject", "last_client_ip", "last_nacrule_id", "last_nacrule_name", "last_nas_vendor", "last_port_id", "last_ssid", "last_status", "last_username", "last_vlan", "mac", "nacrule_id", "nacrule_matched", "nacrule_name", "nas_vendor", "org_id", "port_id", "random_mac", "resp_attrs", "site_id", "ssid", "timestamp", "type", "usermac_label", "username", "vlan")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Ap = temp.Ap
    c.AuthType = temp.AuthType
    c.CertCn = temp.CertCn
    c.CertIssuer = temp.CertIssuer
    c.CertSerial = temp.CertSerial
    c.CertSubject = temp.CertSubject
    c.ClientIp = temp.ClientIp
    c.DeviceMac = temp.DeviceMac
    c.Group = temp.Group
    c.IdpId = temp.IdpId
    c.IdpRole = temp.IdpRole
    c.LastAp = temp.LastAp
    c.LastCertCn = temp.LastCertCn
    c.LastCertExpiry = temp.LastCertExpiry
    c.LastCertIssuer = temp.LastCertIssuer
    c.LastCertSerial = temp.LastCertSerial
    c.LastCertSubject = temp.LastCertSubject
    c.LastClientIp = temp.LastClientIp
    c.LastNacruleId = temp.LastNacruleId
    c.LastNacruleName = temp.LastNacruleName
    c.LastNasVendor = temp.LastNasVendor
    c.LastPortId = temp.LastPortId
    c.LastSsid = temp.LastSsid
    c.LastStatus = temp.LastStatus
    c.LastUsername = temp.LastUsername
    c.LastVlan = temp.LastVlan
    c.Mac = temp.Mac
    c.NacruleId = temp.NacruleId
    c.NacruleMatched = temp.NacruleMatched
    c.NacruleName = temp.NacruleName
    c.NasVendor = temp.NasVendor
    c.OrgId = temp.OrgId
    c.PortId = temp.PortId
    c.RandomMac = temp.RandomMac
    c.RespAttrs = temp.RespAttrs
    c.SiteId = temp.SiteId
    c.Ssid = temp.Ssid
    c.Timestamp = temp.Timestamp
    c.Type = temp.Type
    c.UsermacLabel = temp.UsermacLabel
    c.Username = temp.Username
    c.Vlan = temp.Vlan
    return nil
}

// tempClientNac is a temporary struct used for validating the fields of ClientNac.
type tempClientNac  struct {
    Ap              []string           `json:"ap,omitempty"`
    AuthType        *NacAuthTypeEnum   `json:"auth_type,omitempty"`
    CertCn          []string           `json:"cert_cn,omitempty"`
    CertIssuer      []string           `json:"cert_issuer,omitempty"`
    CertSerial      []string           `json:"cert_serial,omitempty"`
    CertSubject     []string           `json:"cert_subject,omitempty"`
    ClientIp        []string           `json:"client_ip,omitempty"`
    DeviceMac       *string            `json:"device_mac,omitempty"`
    Group           *string            `json:"group,omitempty"`
    IdpId           *string            `json:"idp_id,omitempty"`
    IdpRole         []string           `json:"idp_role,omitempty"`
    LastAp          *string            `json:"last_ap,omitempty"`
    LastCertCn      *string            `json:"last_cert_cn,omitempty"`
    LastCertExpiry  *float64           `json:"last_cert_expiry,omitempty"`
    LastCertIssuer  *string            `json:"last_cert_issuer,omitempty"`
    LastCertSerial  *string            `json:"last_cert_serial,omitempty"`
    LastCertSubject *string            `json:"last_cert_subject,omitempty"`
    LastClientIp    *string            `json:"last_client_ip,omitempty"`
    LastNacruleId   *string            `json:"last_nacrule_id,omitempty"`
    LastNacruleName *string            `json:"last_nacrule_name,omitempty"`
    LastNasVendor   *string            `json:"last_nas_vendor,omitempty"`
    LastPortId      *string            `json:"last_port_id,omitempty"`
    LastSsid        *string            `json:"last_ssid,omitempty"`
    LastStatus      *LastStatusEnum    `json:"last_status,omitempty"`
    LastUsername    *string            `json:"last_username,omitempty"`
    LastVlan        *int               `json:"last_vlan,omitempty"`
    Mac             *string            `json:"mac,omitempty"`
    NacruleId       []string           `json:"nacrule_id,omitempty"`
    NacruleMatched  *bool              `json:"nacrule_matched,omitempty"`
    NacruleName     []string           `json:"nacrule_name,omitempty"`
    NasVendor       []string           `json:"nas_vendor,omitempty"`
    OrgId           *uuid.UUID         `json:"org_id,omitempty"`
    PortId          []string           `json:"port_id,omitempty"`
    RandomMac       *bool              `json:"random_mac,omitempty"`
    RespAttrs       []string           `json:"resp_attrs,omitempty"`
    SiteId          *uuid.UUID         `json:"site_id,omitempty"`
    Ssid            []string           `json:"ssid,omitempty"`
    Timestamp       *float64           `json:"timestamp,omitempty"`
    Type            *NacAccessTypeEnum `json:"type,omitempty"`
    UsermacLabel    []string           `json:"usermac_label,omitempty"`
    Username        []string           `json:"username,omitempty"`
    Vlan            []string           `json:"vlan,omitempty"`
}
