// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// UtilsServicePing represents a UtilsServicePing struct.
type UtilsServicePing struct {
	Count *int   `json:"count,omitempty"`
	Host  string `json:"host"`
	// only for HA. enum: `node0`, `node1`
	Node *HaClusterNodeEnum `json:"node,omitempty"`
	// Ping packet takes the same path as the service
	Service string `json:"service"`
	Size    *int   `json:"size,omitempty"`
	// Tenant context in which the packet is sent
	Tenant               *string                `json:"tenant,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for UtilsServicePing,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (u UtilsServicePing) String() string {
	return fmt.Sprintf(
		"UtilsServicePing[Count=%v, Host=%v, Node=%v, Service=%v, Size=%v, Tenant=%v, AdditionalProperties=%v]",
		u.Count, u.Host, u.Node, u.Service, u.Size, u.Tenant, u.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for UtilsServicePing.
// It customizes the JSON marshaling process for UtilsServicePing objects.
func (u UtilsServicePing) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(u.AdditionalProperties,
		"count", "host", "node", "service", "size", "tenant"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(u.toMap())
}

// toMap converts the UtilsServicePing object to a map representation for JSON marshaling.
func (u UtilsServicePing) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, u.AdditionalProperties)
	if u.Count != nil {
		structMap["count"] = u.Count
	}
	structMap["host"] = u.Host
	if u.Node != nil {
		structMap["node"] = u.Node
	}
	structMap["service"] = u.Service
	if u.Size != nil {
		structMap["size"] = u.Size
	}
	if u.Tenant != nil {
		structMap["tenant"] = u.Tenant
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for UtilsServicePing.
// It customizes the JSON unmarshaling process for UtilsServicePing objects.
func (u *UtilsServicePing) UnmarshalJSON(input []byte) error {
	var temp tempUtilsServicePing
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "count", "host", "node", "service", "size", "tenant")
	if err != nil {
		return err
	}
	u.AdditionalProperties = additionalProperties

	u.Count = temp.Count
	u.Host = *temp.Host
	u.Node = temp.Node
	u.Service = *temp.Service
	u.Size = temp.Size
	u.Tenant = temp.Tenant
	return nil
}

// tempUtilsServicePing is a temporary struct used for validating the fields of UtilsServicePing.
type tempUtilsServicePing struct {
	Count   *int               `json:"count,omitempty"`
	Host    *string            `json:"host"`
	Node    *HaClusterNodeEnum `json:"node,omitempty"`
	Service *string            `json:"service"`
	Size    *int               `json:"size,omitempty"`
	Tenant  *string            `json:"tenant,omitempty"`
}

func (u *tempUtilsServicePing) validate() error {
	var errs []string
	if u.Host == nil {
		errs = append(errs, "required field `host` is missing for type `utils_service_ping`")
	}
	if u.Service == nil {
		errs = append(errs, "required field `service` is missing for type `utils_service_ping`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
