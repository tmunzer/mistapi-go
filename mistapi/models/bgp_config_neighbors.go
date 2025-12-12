// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// BgpConfigNeighbors represents a BgpConfigNeighbors struct.
type BgpConfigNeighbors struct {
	// If true, the BGP session to this neighbor will be administratively disabled/shutdown
	Disabled     *bool   `json:"disabled,omitempty"`
	ExportPolicy *string `json:"export_policy,omitempty"`
	HoldTime     *int    `json:"hold_time,omitempty"`
	ImportPolicy *string `json:"import_policy,omitempty"`
	// Assuming BGP neighbor is directly connected
	MultihopTtl *int `json:"multihop_ttl,omitempty"`
	// BGP AS, value in range 1-4294967294
	NeighborAs           BgpAs                  `json:"neighbor_as"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for BgpConfigNeighbors,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (b BgpConfigNeighbors) String() string {
	return fmt.Sprintf(
		"BgpConfigNeighbors[Disabled=%v, ExportPolicy=%v, HoldTime=%v, ImportPolicy=%v, MultihopTtl=%v, NeighborAs=%v, AdditionalProperties=%v]",
		b.Disabled, b.ExportPolicy, b.HoldTime, b.ImportPolicy, b.MultihopTtl, b.NeighborAs, b.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for BgpConfigNeighbors.
// It customizes the JSON marshaling process for BgpConfigNeighbors objects.
func (b BgpConfigNeighbors) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(b.AdditionalProperties,
		"disabled", "export_policy", "hold_time", "import_policy", "multihop_ttl", "neighbor_as"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(b.toMap())
}

// toMap converts the BgpConfigNeighbors object to a map representation for JSON marshaling.
func (b BgpConfigNeighbors) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, b.AdditionalProperties)
	if b.Disabled != nil {
		structMap["disabled"] = b.Disabled
	}
	if b.ExportPolicy != nil {
		structMap["export_policy"] = b.ExportPolicy
	}
	if b.HoldTime != nil {
		structMap["hold_time"] = b.HoldTime
	}
	if b.ImportPolicy != nil {
		structMap["import_policy"] = b.ImportPolicy
	}
	if b.MultihopTtl != nil {
		structMap["multihop_ttl"] = b.MultihopTtl
	}
	structMap["neighbor_as"] = b.NeighborAs.toMap()
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for BgpConfigNeighbors.
// It customizes the JSON unmarshaling process for BgpConfigNeighbors objects.
func (b *BgpConfigNeighbors) UnmarshalJSON(input []byte) error {
	var temp tempBgpConfigNeighbors
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "export_policy", "hold_time", "import_policy", "multihop_ttl", "neighbor_as")
	if err != nil {
		return err
	}
	b.AdditionalProperties = additionalProperties

	b.Disabled = temp.Disabled
	b.ExportPolicy = temp.ExportPolicy
	b.HoldTime = temp.HoldTime
	b.ImportPolicy = temp.ImportPolicy
	b.MultihopTtl = temp.MultihopTtl
	b.NeighborAs = *temp.NeighborAs
	return nil
}

// tempBgpConfigNeighbors is a temporary struct used for validating the fields of BgpConfigNeighbors.
type tempBgpConfigNeighbors struct {
	Disabled     *bool   `json:"disabled,omitempty"`
	ExportPolicy *string `json:"export_policy,omitempty"`
	HoldTime     *int    `json:"hold_time,omitempty"`
	ImportPolicy *string `json:"import_policy,omitempty"`
	MultihopTtl  *int    `json:"multihop_ttl,omitempty"`
	NeighborAs   *BgpAs  `json:"neighbor_as"`
}

func (b *tempBgpConfigNeighbors) validate() error {
	var errs []string
	if b.NeighborAs == nil {
		errs = append(errs, "required field `neighbor_as` is missing for type `bgp_config_neighbors`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
