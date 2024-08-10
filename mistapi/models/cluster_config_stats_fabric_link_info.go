package models

import (
    "encoding/json"
)

// ClusterConfigStatsFabricLinkInfo represents a ClusterConfigStatsFabricLinkInfo struct.
type ClusterConfigStatsFabricLinkInfo struct {
    DataPlaneNotifiedStatus *string        `json:"DataPlaneNotifiedStatus,omitempty"`
    Interface               []string       `json:"Interface,omitempty"`
    InternalStatus          *string        `json:"InternalStatus,omitempty"`
    State                   *string        `json:"State,omitempty"`
    Status                  *string        `json:"Status,omitempty"`
    AdditionalProperties    map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ClusterConfigStatsFabricLinkInfo.
// It customizes the JSON marshaling process for ClusterConfigStatsFabricLinkInfo objects.
func (c ClusterConfigStatsFabricLinkInfo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ClusterConfigStatsFabricLinkInfo object to a map representation for JSON marshaling.
func (c ClusterConfigStatsFabricLinkInfo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    if c.DataPlaneNotifiedStatus != nil {
        structMap["DataPlaneNotifiedStatus"] = c.DataPlaneNotifiedStatus
    }
    if c.Interface != nil {
        structMap["Interface"] = c.Interface
    }
    if c.InternalStatus != nil {
        structMap["InternalStatus"] = c.InternalStatus
    }
    if c.State != nil {
        structMap["State"] = c.State
    }
    if c.Status != nil {
        structMap["Status"] = c.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ClusterConfigStatsFabricLinkInfo.
// It customizes the JSON unmarshaling process for ClusterConfigStatsFabricLinkInfo objects.
func (c *ClusterConfigStatsFabricLinkInfo) UnmarshalJSON(input []byte) error {
    var temp tempClusterConfigStatsFabricLinkInfo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "DataPlaneNotifiedStatus", "Interface", "InternalStatus", "State", "Status")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.DataPlaneNotifiedStatus = temp.DataPlaneNotifiedStatus
    c.Interface = temp.Interface
    c.InternalStatus = temp.InternalStatus
    c.State = temp.State
    c.Status = temp.Status
    return nil
}

// tempClusterConfigStatsFabricLinkInfo is a temporary struct used for validating the fields of ClusterConfigStatsFabricLinkInfo.
type tempClusterConfigStatsFabricLinkInfo  struct {
    DataPlaneNotifiedStatus *string  `json:"DataPlaneNotifiedStatus,omitempty"`
    Interface               []string `json:"Interface,omitempty"`
    InternalStatus          *string  `json:"InternalStatus,omitempty"`
    State                   *string  `json:"State,omitempty"`
    Status                  *string  `json:"Status,omitempty"`
}
