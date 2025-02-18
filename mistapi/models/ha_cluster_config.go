package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// HaClusterConfig represents a HaClusterConfig struct.
type HaClusterConfig struct {
    // If the device is claimed
    DisableAutoConfig    *bool                  `json:"disable_auto_config,omitempty"`
    // If the device is adopted
    Managed              *bool                  `json:"managed,omitempty"`
    Nodes                []HaClusterConfigNode  `json:"nodes,omitempty"`
    SiteId               *uuid.UUID             `json:"site_id,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for HaClusterConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (h HaClusterConfig) String() string {
    return fmt.Sprintf(
    	"HaClusterConfig[DisableAutoConfig=%v, Managed=%v, Nodes=%v, SiteId=%v, AdditionalProperties=%v]",
    	h.DisableAutoConfig, h.Managed, h.Nodes, h.SiteId, h.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for HaClusterConfig.
// It customizes the JSON marshaling process for HaClusterConfig objects.
func (h HaClusterConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(h.AdditionalProperties,
        "disable_auto_config", "managed", "nodes", "site_id"); err != nil {
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disable_auto_config", "managed", "nodes", "site_id")
    if err != nil {
    	return err
    }
    h.AdditionalProperties = additionalProperties
    
    h.DisableAutoConfig = temp.DisableAutoConfig
    h.Managed = temp.Managed
    h.Nodes = temp.Nodes
    h.SiteId = temp.SiteId
    return nil
}

// tempHaClusterConfig is a temporary struct used for validating the fields of HaClusterConfig.
type tempHaClusterConfig  struct {
    DisableAutoConfig *bool                 `json:"disable_auto_config,omitempty"`
    Managed           *bool                 `json:"managed,omitempty"`
    Nodes             []HaClusterConfigNode `json:"nodes,omitempty"`
    SiteId            *uuid.UUID            `json:"site_id,omitempty"`
}
