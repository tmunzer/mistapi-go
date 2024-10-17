package models

import (
	"encoding/json"
)

// BgpConfigNeighbors represents a BgpConfigNeighbors struct.
type BgpConfigNeighbors struct {
	// If true, the BGP session to this neighbor will be administratively disabled/shutdown
	Disabled     *bool   `json:"disabled,omitempty"`
	ExportPolicy *string `json:"export_policy,omitempty"`
	HoldTime     *int    `json:"hold_time,omitempty"`
	ImportPolicy *string `json:"import_policy,omitempty"`
	// assuming BGP neighbor is directly connected
	MultihopTtl          *int           `json:"multihop_ttl,omitempty"`
	NeighborAs           *int           `json:"neighbor_as,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for BgpConfigNeighbors.
// It customizes the JSON marshaling process for BgpConfigNeighbors objects.
func (b BgpConfigNeighbors) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(b.toMap())
}

// toMap converts the BgpConfigNeighbors object to a map representation for JSON marshaling.
func (b BgpConfigNeighbors) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, b.AdditionalProperties)
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
	if b.NeighborAs != nil {
		structMap["neighbor_as"] = b.NeighborAs
	}
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "disabled", "export_policy", "hold_time", "import_policy", "multihop_ttl", "neighbor_as")
	if err != nil {
		return err
	}

	b.AdditionalProperties = additionalProperties
	b.Disabled = temp.Disabled
	b.ExportPolicy = temp.ExportPolicy
	b.HoldTime = temp.HoldTime
	b.ImportPolicy = temp.ImportPolicy
	b.MultihopTtl = temp.MultihopTtl
	b.NeighborAs = temp.NeighborAs
	return nil
}

// tempBgpConfigNeighbors is a temporary struct used for validating the fields of BgpConfigNeighbors.
type tempBgpConfigNeighbors struct {
	Disabled     *bool   `json:"disabled,omitempty"`
	ExportPolicy *string `json:"export_policy,omitempty"`
	HoldTime     *int    `json:"hold_time,omitempty"`
	ImportPolicy *string `json:"import_policy,omitempty"`
	MultihopTtl  *int    `json:"multihop_ttl,omitempty"`
	NeighborAs   *int    `json:"neighbor_as,omitempty"`
}
