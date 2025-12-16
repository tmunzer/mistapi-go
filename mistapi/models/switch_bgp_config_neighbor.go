// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// SwitchBgpConfigNeighbor represents a SwitchBgpConfigNeighbor struct.
type SwitchBgpConfigNeighbor struct {
	// Autonomous System (AS) number of the BGP neighbor. For internal BGP, this must match `local_as`. For external BGP, this must differ from `local_as`.
	NeighborAs SwitchBgpConfigNeighborNeighborAs `json:"neighbor_as"`
	// Hold time is three times the interval at which keepalive messages are sent. It indicates to the peer the length of time that it should consider the sender valid. Must be 0 or a number in the range 3-65535.
	HoldTime *int `json:"hold_time,omitempty"`
	// Export policy must match one of the policy names defined in the `routing_policies` property.
	ExportPolicy *string `json:"export_policy,omitempty"`
	// Import policy must match one of the policy names defined in the `routing_policies` property.
	ImportPolicy         *string                `json:"import_policy,omitempty"`
	MultihopTtl          *int                   `json:"multihop_ttl,omitempty"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SwitchBgpConfigNeighbor,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SwitchBgpConfigNeighbor) String() string {
	return fmt.Sprintf(
		"SwitchBgpConfigNeighbor[NeighborAs=%v, HoldTime=%v, ExportPolicy=%v, ImportPolicy=%v, MultihopTtl=%v, AdditionalProperties=%v]",
		s.NeighborAs, s.HoldTime, s.ExportPolicy, s.ImportPolicy, s.MultihopTtl, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SwitchBgpConfigNeighbor.
// It customizes the JSON marshaling process for SwitchBgpConfigNeighbor objects.
func (s SwitchBgpConfigNeighbor) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(s.AdditionalProperties,
		"neighbor_as", "hold_time", "export_policy", "import_policy", "multihop_ttl"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(s.toMap())
}

// toMap converts the SwitchBgpConfigNeighbor object to a map representation for JSON marshaling.
func (s SwitchBgpConfigNeighbor) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["neighbor_as"] = s.NeighborAs.toMap()
	if s.HoldTime != nil {
		structMap["hold_time"] = s.HoldTime
	}
	if s.ExportPolicy != nil {
		structMap["export_policy"] = s.ExportPolicy
	}
	if s.ImportPolicy != nil {
		structMap["import_policy"] = s.ImportPolicy
	}
	if s.MultihopTtl != nil {
		structMap["multihop_ttl"] = s.MultihopTtl
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SwitchBgpConfigNeighbor.
// It customizes the JSON unmarshaling process for SwitchBgpConfigNeighbor objects.
func (s *SwitchBgpConfigNeighbor) UnmarshalJSON(input []byte) error {
	var temp tempSwitchBgpConfigNeighbor
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "neighbor_as", "hold_time", "export_policy", "import_policy", "multihop_ttl")
	if err != nil {
		return err
	}
	s.AdditionalProperties = additionalProperties

	s.NeighborAs = *temp.NeighborAs
	s.HoldTime = temp.HoldTime
	s.ExportPolicy = temp.ExportPolicy
	s.ImportPolicy = temp.ImportPolicy
	s.MultihopTtl = temp.MultihopTtl
	return nil
}

// tempSwitchBgpConfigNeighbor is a temporary struct used for validating the fields of SwitchBgpConfigNeighbor.
type tempSwitchBgpConfigNeighbor struct {
	NeighborAs   *SwitchBgpConfigNeighborNeighborAs `json:"neighbor_as"`
	HoldTime     *int                               `json:"hold_time,omitempty"`
	ExportPolicy *string                            `json:"export_policy,omitempty"`
	ImportPolicy *string                            `json:"import_policy,omitempty"`
	MultihopTtl  *int                               `json:"multihop_ttl,omitempty"`
}

func (s *tempSwitchBgpConfigNeighbor) validate() error {
	var errs []string
	if s.NeighborAs == nil {
		errs = append(errs, "required field `neighbor_as` is missing for type `switch_bgp_config_neighbor`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
