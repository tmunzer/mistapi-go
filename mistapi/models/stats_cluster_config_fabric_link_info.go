package models

import (
	"encoding/json"
)

// StatsClusterConfigFabricLinkInfo represents a StatsClusterConfigFabricLinkInfo struct.
type StatsClusterConfigFabricLinkInfo struct {
	DataPlaneNotifiedStatus *string        `json:"DataPlaneNotifiedStatus,omitempty"`
	Interface               []string       `json:"Interface,omitempty"`
	InternalStatus          *string        `json:"InternalStatus,omitempty"`
	State                   *string        `json:"State,omitempty"`
	Status                  *string        `json:"Status,omitempty"`
	AdditionalProperties    map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for StatsClusterConfigFabricLinkInfo.
// It customizes the JSON marshaling process for StatsClusterConfigFabricLinkInfo objects.
func (s StatsClusterConfigFabricLinkInfo) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the StatsClusterConfigFabricLinkInfo object to a map representation for JSON marshaling.
func (s StatsClusterConfigFabricLinkInfo) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	if s.DataPlaneNotifiedStatus != nil {
		structMap["DataPlaneNotifiedStatus"] = s.DataPlaneNotifiedStatus
	}
	if s.Interface != nil {
		structMap["Interface"] = s.Interface
	}
	if s.InternalStatus != nil {
		structMap["InternalStatus"] = s.InternalStatus
	}
	if s.State != nil {
		structMap["State"] = s.State
	}
	if s.Status != nil {
		structMap["Status"] = s.Status
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for StatsClusterConfigFabricLinkInfo.
// It customizes the JSON unmarshaling process for StatsClusterConfigFabricLinkInfo objects.
func (s *StatsClusterConfigFabricLinkInfo) UnmarshalJSON(input []byte) error {
	var temp tempStatsClusterConfigFabricLinkInfo
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "DataPlaneNotifiedStatus", "Interface", "InternalStatus", "State", "Status")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.DataPlaneNotifiedStatus = temp.DataPlaneNotifiedStatus
	s.Interface = temp.Interface
	s.InternalStatus = temp.InternalStatus
	s.State = temp.State
	s.Status = temp.Status
	return nil
}

// tempStatsClusterConfigFabricLinkInfo is a temporary struct used for validating the fields of StatsClusterConfigFabricLinkInfo.
type tempStatsClusterConfigFabricLinkInfo struct {
	DataPlaneNotifiedStatus *string  `json:"DataPlaneNotifiedStatus,omitempty"`
	Interface               []string `json:"Interface,omitempty"`
	InternalStatus          *string  `json:"InternalStatus,omitempty"`
	State                   *string  `json:"State,omitempty"`
	Status                  *string  `json:"Status,omitempty"`
}
