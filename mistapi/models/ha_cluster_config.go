// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
)

// HaClusterConfig represents a HaClusterConfig struct.
// Request to create an HA cluster from unassigned gateway inventory nodes
type HaClusterConfig struct {
	// This disables the default behavior of a cloud-ready switch/gateway being managed/configured by Mist. Setting this to `true` means you want to disable the default behavior and do not want the device to be Mist-managed.
	DisableAutoConfig *bool `json:"disable_auto_config,omitempty"` // Deprecated
	// An adopted switch/gateway will not be managed/configured by Mist by default. Setting this parameter to `true` enables the adopted switch/gateway to be managed/configured by Mist.
	Managed *bool `json:"managed,omitempty"` // Deprecated
	// whether the device can be configured by Mist or not. This deprecates `managed` (for adopted device) and `disable_auto_config` for claimed device)
	MistConfigured *bool `json:"mist_configured,omitempty"`
	// Gateway inventory nodes used to create an HA cluster
	Nodes []HaClusterConfigNode `json:"nodes,omitempty"`
	// Site where the HA cluster should be created
	SiteId               *uuid.UUID             `json:"site_id,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for HaClusterConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (h HaClusterConfig) String() string {
	return fmt.Sprintf(
		"HaClusterConfig[DisableAutoConfig=%v, Managed=%v, MistConfigured=%v, Nodes=%v, SiteId=%v, AdditionalProperties=%v]",
		h.DisableAutoConfig, h.Managed, h.MistConfigured, h.Nodes, h.SiteId, h.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for HaClusterConfig.
// It customizes the JSON marshaling process for HaClusterConfig objects.
func (h HaClusterConfig) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(h.AdditionalProperties,
		"disable_auto_config", "managed", "mist_configured", "nodes", "site_id"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(h.toMap())
}

// toMap converts the HaClusterConfig object to a map representation for JSON marshaling.
func (h HaClusterConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, h.AdditionalProperties)
	if h.DisableAutoConfig != nil {
		structMap["disable_auto_config"] = h.DisableAutoConfig
	}
	if h.Managed != nil {
		structMap["managed"] = h.Managed
	}
	if h.MistConfigured != nil {
		structMap["mist_configured"] = h.MistConfigured
	}
	if h.Nodes != nil {
		structMap["nodes"] = h.Nodes
	}
	if h.SiteId != nil {
		structMap["site_id"] = h.SiteId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for HaClusterConfig.
// It customizes the JSON unmarshaling process for HaClusterConfig objects.
func (h *HaClusterConfig) UnmarshalJSON(input []byte) error {
	var temp tempHaClusterConfig
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_auto_config", "managed", "mist_configured", "nodes", "site_id")
	if err != nil {
		return err
	}
	h.AdditionalProperties = additionalProperties

	h.DisableAutoConfig = temp.DisableAutoConfig
	h.Managed = temp.Managed
	h.MistConfigured = temp.MistConfigured
	h.Nodes = temp.Nodes
	h.SiteId = temp.SiteId
	return nil
}

// tempHaClusterConfig is a temporary struct used for validating the fields of HaClusterConfig.
type tempHaClusterConfig struct {
	DisableAutoConfig *bool                 `json:"disable_auto_config,omitempty"`
	Managed           *bool                 `json:"managed,omitempty"`
	MistConfigured    *bool                 `json:"mist_configured,omitempty"`
	Nodes             []HaClusterConfigNode `json:"nodes,omitempty"`
	SiteId            *uuid.UUID            `json:"site_id,omitempty"`
}
