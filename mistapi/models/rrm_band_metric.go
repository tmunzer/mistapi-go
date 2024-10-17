package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// RrmBandMetric represents a RrmBandMetric struct.
type RrmBandMetric struct {
	// average number of co-channel neighbors
	CochannelNeighbors float64 `json:"cochannel_neighbors"`
	// defined by how APs can hear from one and another, 0 - 1 (everyone can hear everyone)
	Density float64 `json:"density"`
	// Property key is the channel number
	Interferences map[string]RrmBandMetricInterference `json:"interferences,omitempty"`
	// average number of neighbors
	Neighbors float64 `json:"neighbors"`
	// average noise in dBm
	Noise                float64        `json:"noise"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmBandMetric.
// It customizes the JSON marshaling process for RrmBandMetric objects.
func (r RrmBandMetric) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the RrmBandMetric object to a map representation for JSON marshaling.
func (r RrmBandMetric) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["cochannel_neighbors"] = r.CochannelNeighbors
	structMap["density"] = r.Density
	if r.Interferences != nil {
		structMap["interferences"] = r.Interferences
	}
	structMap["neighbors"] = r.Neighbors
	structMap["noise"] = r.Noise
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmBandMetric.
// It customizes the JSON unmarshaling process for RrmBandMetric objects.
func (r *RrmBandMetric) UnmarshalJSON(input []byte) error {
	var temp tempRrmBandMetric
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cochannel_neighbors", "density", "interferences", "neighbors", "noise")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.CochannelNeighbors = *temp.CochannelNeighbors
	r.Density = *temp.Density
	r.Interferences = temp.Interferences
	r.Neighbors = *temp.Neighbors
	r.Noise = *temp.Noise
	return nil
}

// tempRrmBandMetric is a temporary struct used for validating the fields of RrmBandMetric.
type tempRrmBandMetric struct {
	CochannelNeighbors *float64                             `json:"cochannel_neighbors"`
	Density            *float64                             `json:"density"`
	Interferences      map[string]RrmBandMetricInterference `json:"interferences,omitempty"`
	Neighbors          *float64                             `json:"neighbors"`
	Noise              *float64                             `json:"noise"`
}

func (r *tempRrmBandMetric) validate() error {
	var errs []string
	if r.CochannelNeighbors == nil {
		errs = append(errs, "required field `cochannel_neighbors` is missing for type `rrm_band_metric`")
	}
	if r.Density == nil {
		errs = append(errs, "required field `density` is missing for type `rrm_band_metric`")
	}
	if r.Neighbors == nil {
		errs = append(errs, "required field `neighbors` is missing for type `rrm_band_metric`")
	}
	if r.Noise == nil {
		errs = append(errs, "required field `noise` is missing for type `rrm_band_metric`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
