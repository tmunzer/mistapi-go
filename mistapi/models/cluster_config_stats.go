package models

import (
    "encoding/json"
)

// ClusterConfigStats represents a ClusterConfigStats struct.
type ClusterConfigStats struct {
    Configuration              *string                                     `json:"configuration,omitempty"`
    ControlLinkInfo            *ClusterConfigStatsControlLinkInfo          `json:"control_link_info,omitempty"`
    EthernetConnection         []ClusterConfigStatsEthernetConnectionItem  `json:"ethernet_connection,omitempty"`
    FabricLinkInfo             *ClusterConfigStatsFabricLinkInfo           `json:"fabric_link_info,omitempty"`
    LastStatusChangeReason     *string                                     `json:"last_status_change_reason,omitempty"`
    Operational                *string                                     `json:"operational,omitempty"`
    PrimaryNodeHealth          *string                                     `json:"primary_node_health,omitempty"`
    RedundancyGroupInformation []ClusterConfigStatsRedundancyGroupInfoItem `json:"redundancy_group_information,omitempty"`
    SecondaryNodeHealth        *string                                     `json:"secondary_node_health,omitempty"`
    Status                     *string                                     `json:"status,omitempty"`
    AdditionalProperties       map[string]any                              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClusterConfigStats.
// It customizes the JSON marshaling process for ClusterConfigStats objects.
func (c ClusterConfigStats) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClusterConfigStats object to a map representation for JSON marshaling.
func (c ClusterConfigStats) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Configuration != nil {
        structMap["configuration"] = c.Configuration
    }
    if c.ControlLinkInfo != nil {
        structMap["control_link_info"] = c.ControlLinkInfo.toMap()
    }
    if c.EthernetConnection != nil {
        structMap["ethernet_connection"] = c.EthernetConnection
    }
    if c.FabricLinkInfo != nil {
        structMap["fabric_link_info"] = c.FabricLinkInfo.toMap()
    }
    if c.LastStatusChangeReason != nil {
        structMap["last_status_change_reason"] = c.LastStatusChangeReason
    }
    if c.Operational != nil {
        structMap["operational"] = c.Operational
    }
    if c.PrimaryNodeHealth != nil {
        structMap["primary_node_health"] = c.PrimaryNodeHealth
    }
    if c.RedundancyGroupInformation != nil {
        structMap["redundancy_group_information"] = c.RedundancyGroupInformation
    }
    if c.SecondaryNodeHealth != nil {
        structMap["secondary_node_health"] = c.SecondaryNodeHealth
    }
    if c.Status != nil {
        structMap["status"] = c.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClusterConfigStats.
// It customizes the JSON unmarshaling process for ClusterConfigStats objects.
func (c *ClusterConfigStats) UnmarshalJSON(input []byte) error {
    var temp tempClusterConfigStats
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "configuration", "control_link_info", "ethernet_connection", "fabric_link_info", "last_status_change_reason", "operational", "primary_node_health", "redundancy_group_information", "secondary_node_health", "status")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Configuration = temp.Configuration
    c.ControlLinkInfo = temp.ControlLinkInfo
    c.EthernetConnection = temp.EthernetConnection
    c.FabricLinkInfo = temp.FabricLinkInfo
    c.LastStatusChangeReason = temp.LastStatusChangeReason
    c.Operational = temp.Operational
    c.PrimaryNodeHealth = temp.PrimaryNodeHealth
    c.RedundancyGroupInformation = temp.RedundancyGroupInformation
    c.SecondaryNodeHealth = temp.SecondaryNodeHealth
    c.Status = temp.Status
    return nil
}

// tempClusterConfigStats is a temporary struct used for validating the fields of ClusterConfigStats.
type tempClusterConfigStats  struct {
    Configuration              *string                                     `json:"configuration,omitempty"`
    ControlLinkInfo            *ClusterConfigStatsControlLinkInfo          `json:"control_link_info,omitempty"`
    EthernetConnection         []ClusterConfigStatsEthernetConnectionItem  `json:"ethernet_connection,omitempty"`
    FabricLinkInfo             *ClusterConfigStatsFabricLinkInfo           `json:"fabric_link_info,omitempty"`
    LastStatusChangeReason     *string                                     `json:"last_status_change_reason,omitempty"`
    Operational                *string                                     `json:"operational,omitempty"`
    PrimaryNodeHealth          *string                                     `json:"primary_node_health,omitempty"`
    RedundancyGroupInformation []ClusterConfigStatsRedundancyGroupInfoItem `json:"redundancy_group_information,omitempty"`
    SecondaryNodeHealth        *string                                     `json:"secondary_node_health,omitempty"`
    Status                     *string                                     `json:"status,omitempty"`
}
