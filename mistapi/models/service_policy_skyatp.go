// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
)

// ServicePolicySkyatp represents a ServicePolicySkyatp struct.
// SRX Sky ATP threat inspection settings for a service policy
type ServicePolicySkyatp struct {
	// Sky ATP DNS DGA detection settings
	DnsDgaDetection *ServicePolicySkyatpDnsDgaDetection `json:"dns_dga_detection,omitempty"`
	// Sky ATP DNS tunneling detection settings
	DnsTunnelDetection *ServicePolicySkyatpDnsTunnelDetection `json:"dns_tunnel_detection,omitempty"`
	// Sky ATP HTTP inspection settings
	HttpInspection *ServicePolicySkyatpHttpInspection `json:"http_inspection,omitempty"`
	// Sky ATP IoT device policy settings
	IotDevicePolicy      *ServicePolicySkyatpIotDevicePolicy `json:"iot_device_policy,omitempty"`
	AdditionalProperties map[string]interface{}              `json:"_"`
}

// String implements the fmt.Stringer interface for ServicePolicySkyatp,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s ServicePolicySkyatp) String() string {
	return fmt.Sprintf(
		"ServicePolicySkyatp[DnsDgaDetection=%v, DnsTunnelDetection=%v, HttpInspection=%v, IotDevicePolicy=%v, AdditionalProperties=%v]",
		s.DnsDgaDetection, s.DnsTunnelDetection, s.HttpInspection, s.IotDevicePolicy, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ServicePolicySkyatp.
// It customizes the JSON marshaling process for ServicePolicySkyatp objects.
func (s ServicePolicySkyatp) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"dns_dga_detection", "dns_tunnel_detection", "http_inspection", "iot_device_policy"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the ServicePolicySkyatp object to a map representation for JSON marshaling.
func (s ServicePolicySkyatp) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	if s.DnsDgaDetection != nil {
		structMap["dns_dga_detection"] = s.DnsDgaDetection.toMap()
	}
	if s.DnsTunnelDetection != nil {
		structMap["dns_tunnel_detection"] = s.DnsTunnelDetection.toMap()
	}
	if s.HttpInspection != nil {
		structMap["http_inspection"] = s.HttpInspection.toMap()
	}
	if s.IotDevicePolicy != nil {
		structMap["iot_device_policy"] = s.IotDevicePolicy.toMap()
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServicePolicySkyatp.
// It customizes the JSON unmarshaling process for ServicePolicySkyatp objects.
func (s *ServicePolicySkyatp) UnmarshalJSON(input []byte) error {
	var temp tempServicePolicySkyatp
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "dns_dga_detection", "dns_tunnel_detection", "http_inspection", "iot_device_policy")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.DnsDgaDetection = temp.DnsDgaDetection
	s.DnsTunnelDetection = temp.DnsTunnelDetection
	s.HttpInspection = temp.HttpInspection
	s.IotDevicePolicy = temp.IotDevicePolicy
	return nil
}

// tempServicePolicySkyatp is a temporary struct used for validating the fields of ServicePolicySkyatp.
type tempServicePolicySkyatp struct {
	DnsDgaDetection    *ServicePolicySkyatpDnsDgaDetection    `json:"dns_dga_detection,omitempty"`
	DnsTunnelDetection *ServicePolicySkyatpDnsTunnelDetection `json:"dns_tunnel_detection,omitempty"`
	HttpInspection     *ServicePolicySkyatpHttpInspection     `json:"http_inspection,omitempty"`
	IotDevicePolicy    *ServicePolicySkyatpIotDevicePolicy    `json:"iot_device_policy,omitempty"`
}
