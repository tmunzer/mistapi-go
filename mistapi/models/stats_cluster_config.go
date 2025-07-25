// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// StatsClusterConfig represents a StatsClusterConfig struct.
type StatsClusterConfig struct {
    Configuration              *string                                     `json:"configuration,omitempty"`
    ControlLinkInfo            *StatsClusterConfigControlLinkInfo          `json:"control_link_info,omitempty"`
    EthernetConnection         []StatsClusterConfigEthernetConnectionItem  `json:"ethernet_connection,omitempty"`
    FabricLinkInfo             *StatsClusterConfigFabricLinkInfo           `json:"fabric_link_info,omitempty"`
    LastStatusChangeReason     *string                                     `json:"last_status_change_reason,omitempty"`
    Operational                *string                                     `json:"operational,omitempty"`
    PrimaryNodeHealth          *string                                     `json:"primary_node_health,omitempty"`
    RedundancyGroupInformation []StatsClusterConfigRedundancyGroupInfoItem `json:"redundancy_group_information,omitempty"`
    SecondaryNodeHealth        *string                                     `json:"secondary_node_health,omitempty"`
    Status                     *string                                     `json:"status,omitempty"`
    AdditionalProperties       map[string]interface{}                      `json:"_"`
}

// String implements the fmt.Stringer interface for StatsClusterConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s StatsClusterConfig) String() string {
    return fmt.Sprintf(
    	"StatsClusterConfig[Configuration=%v, ControlLinkInfo=%v, EthernetConnection=%v, FabricLinkInfo=%v, LastStatusChangeReason=%v, Operational=%v, PrimaryNodeHealth=%v, RedundancyGroupInformation=%v, SecondaryNodeHealth=%v, Status=%v, AdditionalProperties=%v]",
    	s.Configuration, s.ControlLinkInfo, s.EthernetConnection, s.FabricLinkInfo, s.LastStatusChangeReason, s.Operational, s.PrimaryNodeHealth, s.RedundancyGroupInformation, s.SecondaryNodeHealth, s.Status, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for StatsClusterConfig.
// It customizes the JSON marshaling process for StatsClusterConfig objects.
func (s StatsClusterConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "configuration", "control_link_info", "ethernet_connection", "fabric_link_info", "last_status_change_reason", "operational", "primary_node_health", "redundancy_group_information", "secondary_node_health", "status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the StatsClusterConfig object to a map representation for JSON marshaling.
func (s StatsClusterConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Configuration != nil {
        structMap["configuration"] = s.Configuration
    }
    if s.ControlLinkInfo != nil {
        structMap["control_link_info"] = s.ControlLinkInfo.toMap()
    }
    if s.EthernetConnection != nil {
        structMap["ethernet_connection"] = s.EthernetConnection
    }
    if s.FabricLinkInfo != nil {
        structMap["fabric_link_info"] = s.FabricLinkInfo.toMap()
    }
    if s.LastStatusChangeReason != nil {
        structMap["last_status_change_reason"] = s.LastStatusChangeReason
    }
    if s.Operational != nil {
        structMap["operational"] = s.Operational
    }
    if s.PrimaryNodeHealth != nil {
        structMap["primary_node_health"] = s.PrimaryNodeHealth
    }
    if s.RedundancyGroupInformation != nil {
        structMap["redundancy_group_information"] = s.RedundancyGroupInformation
    }
    if s.SecondaryNodeHealth != nil {
        structMap["secondary_node_health"] = s.SecondaryNodeHealth
    }
    if s.Status != nil {
        structMap["status"] = s.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClusterConfig.
// It customizes the JSON unmarshaling process for StatsClusterConfig objects.
func (s *StatsClusterConfig) UnmarshalJSON(input []byte) error {
    var temp tempStatsClusterConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "configuration", "control_link_info", "ethernet_connection", "fabric_link_info", "last_status_change_reason", "operational", "primary_node_health", "redundancy_group_information", "secondary_node_health", "status")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Configuration = temp.Configuration
    s.ControlLinkInfo = temp.ControlLinkInfo
    s.EthernetConnection = temp.EthernetConnection
    s.FabricLinkInfo = temp.FabricLinkInfo
    s.LastStatusChangeReason = temp.LastStatusChangeReason
    s.Operational = temp.Operational
    s.PrimaryNodeHealth = temp.PrimaryNodeHealth
    s.RedundancyGroupInformation = temp.RedundancyGroupInformation
    s.SecondaryNodeHealth = temp.SecondaryNodeHealth
    s.Status = temp.Status
    return nil
}

// tempStatsClusterConfig is a temporary struct used for validating the fields of StatsClusterConfig.
type tempStatsClusterConfig  struct {
    Configuration              *string                                     `json:"configuration,omitempty"`
    ControlLinkInfo            *StatsClusterConfigControlLinkInfo          `json:"control_link_info,omitempty"`
    EthernetConnection         []StatsClusterConfigEthernetConnectionItem  `json:"ethernet_connection,omitempty"`
    FabricLinkInfo             *StatsClusterConfigFabricLinkInfo           `json:"fabric_link_info,omitempty"`
    LastStatusChangeReason     *string                                     `json:"last_status_change_reason,omitempty"`
    Operational                *string                                     `json:"operational,omitempty"`
    PrimaryNodeHealth          *string                                     `json:"primary_node_health,omitempty"`
    RedundancyGroupInformation []StatsClusterConfigRedundancyGroupInfoItem `json:"redundancy_group_information,omitempty"`
    SecondaryNodeHealth        *string                                     `json:"secondary_node_health,omitempty"`
    Status                     *string                                     `json:"status,omitempty"`
}
