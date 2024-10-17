package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// WiredClientResponse represents a WiredClientResponse struct.
type WiredClientResponse struct {
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
	AdditionalProperties      map[string]any                         `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WiredClientResponse.
// It customizes the JSON marshaling process for WiredClientResponse objects.
func (w WiredClientResponse) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(w.toMap())
}

// toMap converts the WiredClientResponse object to a map representation for JSON marshaling.
func (w WiredClientResponse) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, w.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "device_mac", "device_mac_port", "dhcp_client_identifier", "dhcp_client_options", "dhcp_fqdn", "dhcp_hostname", "dhcp_request_params", "dhcp_vendor_class_identifier", "ip", "mac", "org_id", "port_id", "site_id", "timestamp", "vlan")
	if err != nil {
		return err
	}

	w.AdditionalProperties = additionalProperties
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
type tempWiredClientResponse struct {
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
