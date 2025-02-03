package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RrmBandMetric represents a RrmBandMetric struct.
type RrmBandMetric struct {
    // Average number of co-channel neighbors
    CochannelNeighbors   float64                              `json:"cochannel_neighbors"`
    // defined by how APs can hear from one and another, 0 - 1 (everyone can hear everyone)
    Density              float64                              `json:"density"`
    // Property key is the channel number
    Interferences        map[string]RrmBandMetricInterference `json:"interferences,omitempty"`
    // Average number of neighbors
    Neighbors            float64                              `json:"neighbors"`
    // Average noise in dBm
    Noise                float64                              `json:"noise"`
    AdditionalProperties map[string]interface{}               `json:"_"`
}

// String implements the fmt.Stringer interface for RrmBandMetric,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RrmBandMetric) String() string {
    return fmt.Sprintf(
    	"RrmBandMetric[CochannelNeighbors=%v, Density=%v, Interferences=%v, Neighbors=%v, Noise=%v, AdditionalProperties=%v]",
    	r.CochannelNeighbors, r.Density, r.Interferences, r.Neighbors, r.Noise, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RrmBandMetric.
// It customizes the JSON marshaling process for RrmBandMetric objects.
func (r RrmBandMetric) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "cochannel_neighbors", "density", "interferences", "neighbors", "noise"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RrmBandMetric object to a map representation for JSON marshaling.
func (r RrmBandMetric) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cochannel_neighbors", "density", "interferences", "neighbors", "noise")
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
type tempRrmBandMetric  struct {
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
    return errors.New(strings.Join (errs, "\n"))
}
