package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// WiredClientResponse represents a WiredClientResponse struct.
type WiredClientResponse struct {
    AuthMethod                *string                                `json:"auth_method,omitempty"`
    AuthState                 *string                                `json:"auth_state,omitempty"`
    // MAC Address of the switch the client is connected to
    DeviceMac                 []string                               `json:"device_mac,omitempty"`
    DeviceMacPort             []WiredClientResponseDeviceMacPortItem `json:"device_mac_port,omitempty"`
    DhcpClientIdentifier      *string                                `json:"dhcp_client_identifier,omitempty"`
    DhcpClientOptions         []DhcpClientOption                     `json:"dhcp_client_options,omitempty"`
    DhcpFqdn                  *string                                `json:"dhcp_fqdn,omitempty"`
    DhcpHostname              *string                                `json:"dhcp_hostname,omitempty"`
    DhcpRequestParams         *string                                `json:"dhcp_request_params,omitempty"`
    DhcpVendorClassIdentifier *string                                `json:"dhcp_vendor_class_identifier,omitempty"`
    Ip                        []string                               `json:"ip,omitempty"`
    Mac                       *string                                `json:"mac,omitempty"`
    OrgId                     *uuid.UUID                             `json:"org_id,omitempty"`
    PortId                    []string                               `json:"port_id,omitempty"`
    SiteId                    *uuid.UUID                             `json:"site_id,omitempty"`
    // Epoch (seconds)
    Timestamp                 *float64                               `json:"timestamp,omitempty"`
    Vlan                      []int                                  `json:"vlan,omitempty"`
    AdditionalProperties      map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for WiredClientResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WiredClientResponse) String() string {
    return fmt.Sprintf(
    	"WiredClientResponse[AuthMethod=%v, AuthState=%v, DeviceMac=%v, DeviceMacPort=%v, DhcpClientIdentifier=%v, DhcpClientOptions=%v, DhcpFqdn=%v, DhcpHostname=%v, DhcpRequestParams=%v, DhcpVendorClassIdentifier=%v, Ip=%v, Mac=%v, OrgId=%v, PortId=%v, SiteId=%v, Timestamp=%v, Vlan=%v, AdditionalProperties=%v]",
    	w.AuthMethod, w.AuthState, w.DeviceMac, w.DeviceMacPort, w.DhcpClientIdentifier, w.DhcpClientOptions, w.DhcpFqdn, w.DhcpHostname, w.DhcpRequestParams, w.DhcpVendorClassIdentifier, w.Ip, w.Mac, w.OrgId, w.PortId, w.SiteId, w.Timestamp, w.Vlan, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WiredClientResponse.
// It customizes the JSON marshaling process for WiredClientResponse objects.
func (w WiredClientResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "auth_method", "auth_state", "device_mac", "device_mac_port", "dhcp_client_identifier", "dhcp_client_options", "dhcp_fqdn", "dhcp_hostname", "dhcp_request_params", "dhcp_vendor_class_identifier", "ip", "mac", "org_id", "port_id", "site_id", "timestamp", "vlan"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WiredClientResponse object to a map representation for JSON marshaling.
func (w WiredClientResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AuthMethod != nil {
        structMap["auth_method"] = w.AuthMethod
    }
    if w.AuthState != nil {
        structMap["auth_state"] = w.AuthState
    }
    if w.DeviceMac != nil {
        structMap["device_mac"] = w.DeviceMac
    }
    if w.DeviceMacPort != nil {
        structMap["device_mac_port"] = w.DeviceMacPort
    }
    if w.DhcpClientIdentifier != nil {
        structMap["dhcp_client_identifier"] = w.DhcpClientIdentifier
    }
    if w.DhcpClientOptions != nil {
        structMap["dhcp_client_options"] = w.DhcpClientOptions
    }
    if w.DhcpFqdn != nil {
        structMap["dhcp_fqdn"] = w.DhcpFqdn
    }
    if w.DhcpHostname != nil {
        structMap["dhcp_hostname"] = w.DhcpHostname
    }
    if w.DhcpRequestParams != nil {
        structMap["dhcp_request_params"] = w.DhcpRequestParams
    }
    if w.DhcpVendorClassIdentifier != nil {
        structMap["dhcp_vendor_class_identifier"] = w.DhcpVendorClassIdentifier
    }
    if w.Ip != nil {
        structMap["ip"] = w.Ip
    }
    if w.Mac != nil {
        structMap["mac"] = w.Mac
    }
    if w.OrgId != nil {
        structMap["org_id"] = w.OrgId
    }
    if w.PortId != nil {
        structMap["port_id"] = w.PortId
    }
    if w.SiteId != nil {
        structMap["site_id"] = w.SiteId
    }
    if w.Timestamp != nil {
        structMap["timestamp"] = w.Timestamp
    }
    if w.Vlan != nil {
        structMap["vlan"] = w.Vlan
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WiredClientResponse.
// It customizes the JSON unmarshaling process for WiredClientResponse objects.
func (w *WiredClientResponse) UnmarshalJSON(input []byte) error {
    var temp tempWiredClientResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "auth_method", "auth_state", "device_mac", "device_mac_port", "dhcp_client_identifier", "dhcp_client_options", "dhcp_fqdn", "dhcp_hostname", "dhcp_request_params", "dhcp_vendor_class_identifier", "ip", "mac", "org_id", "port_id", "site_id", "timestamp", "vlan")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.AuthMethod = temp.AuthMethod
    w.AuthState = temp.AuthState
    w.DeviceMac = temp.DeviceMac
    w.DeviceMacPort = temp.DeviceMacPort
    w.DhcpClientIdentifier = temp.DhcpClientIdentifier
    w.DhcpClientOptions = temp.DhcpClientOptions
    w.DhcpFqdn = temp.DhcpFqdn
    w.DhcpHostname = temp.DhcpHostname
    w.DhcpRequestParams = temp.DhcpRequestParams
    w.DhcpVendorClassIdentifier = temp.DhcpVendorClassIdentifier
    w.Ip = temp.Ip
    w.Mac = temp.Mac
    w.OrgId = temp.OrgId
    w.PortId = temp.PortId
    w.SiteId = temp.SiteId
    w.Timestamp = temp.Timestamp
    w.Vlan = temp.Vlan
    return nil
}

// tempWiredClientResponse is a temporary struct used for validating the fields of WiredClientResponse.
type tempWiredClientResponse  struct {
    AuthMethod                *string                                `json:"auth_method,omitempty"`
    AuthState                 *string                                `json:"auth_state,omitempty"`
    DeviceMac                 []string                               `json:"device_mac,omitempty"`
    DeviceMacPort             []WiredClientResponseDeviceMacPortItem `json:"device_mac_port,omitempty"`
    DhcpClientIdentifier      *string                                `json:"dhcp_client_identifier,omitempty"`
    DhcpClientOptions         []DhcpClientOption                     `json:"dhcp_client_options,omitempty"`
    DhcpFqdn                  *string                                `json:"dhcp_fqdn,omitempty"`
    DhcpHostname              *string                                `json:"dhcp_hostname,omitempty"`
    DhcpRequestParams         *string                                `json:"dhcp_request_params,omitempty"`
    DhcpVendorClassIdentifier *string                                `json:"dhcp_vendor_class_identifier,omitempty"`
    Ip                        []string                               `json:"ip,omitempty"`
    Mac                       *string                                `json:"mac,omitempty"`
    OrgId                     *uuid.UUID                             `json:"org_id,omitempty"`
    PortId                    []string                               `json:"port_id,omitempty"`
    SiteId                    *uuid.UUID                             `json:"site_id,omitempty"`
    Timestamp                 *float64                               `json:"timestamp,omitempty"`
    Vlan                      []int                                  `json:"vlan,omitempty"`
}
