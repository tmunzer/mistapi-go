package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseLocationCoverage represents a ResponseLocationCoverage struct.
type ResponseLocationCoverage struct {
    // List of [x, y, mean]s, x/y are in meters (UI would need to use map.ppm to calulate the pixel location from top-left).
    BeamsMeans           [][]float64            `json:"beams_means"`
    End                  int                    `json:"end"`
    // Size of grid, in meter
    Gridsize             float64                `json:"gridsize"`
    // List of names annotating the fields in results
    ResultDef            []string               `json:"result_def"`
    // List of results, see result_def.
    Results              [][]float64            `json:"results"`
    Start                int                    `json:"start"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseLocationCoverage,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseLocationCoverage) String() string {
    return fmt.Sprintf(
    	"ResponseLocationCoverage[BeamsMeans=%v, End=%v, Gridsize=%v, ResultDef=%v, Results=%v, Start=%v, AdditionalProperties=%v]",
    	r.BeamsMeans, r.End, r.Gridsize, r.ResultDef, r.Results, r.Start, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseLocationCoverage.
// It customizes the JSON marshaling process for ResponseLocationCoverage objects.
func (r ResponseLocationCoverage) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "beams_means", "end", "gridsize", "result_def", "results", "start"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseLocationCoverage object to a map representation for JSON marshaling.
func (r ResponseLocationCoverage) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["beams_means"] = r.BeamsMeans
    structMap["end"] = r.End
    structMap["gridsize"] = r.Gridsize
    structMap["result_def"] = r.ResultDef
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseLocationCoverage.
// It customizes the JSON unmarshaling process for ResponseLocationCoverage objects.
func (r *ResponseLocationCoverage) UnmarshalJSON(input []byte) error {
    var temp tempResponseLocationCoverage
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "beams_means", "end", "gridsize", "result_def", "results", "start")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.BeamsMeans = *temp.BeamsMeans
    r.End = *temp.End
    r.Gridsize = *temp.Gridsize
    r.ResultDef = *temp.ResultDef
    r.Results = *temp.Results
    r.Start = *temp.Start
    return nil
}

// tempResponseLocationCoverage is a temporary struct used for validating the fields of ResponseLocationCoverage.
type tempResponseLocationCoverage  struct {
    BeamsMeans *[][]float64 `json:"beams_means"`
    End        *int         `json:"end"`
    Gridsize   *float64     `json:"gridsize"`
    ResultDef  *[]string    `json:"result_def"`
    Results    *[][]float64 `json:"results"`
    Start      *int         `json:"start"`
}

func (r *tempResponseLocationCoverage) validate() error {
    var errs []string
    if r.BeamsMeans == nil {
        errs = append(errs, "required field `beams_means` is missing for type `response_location_coverage`")
    }
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `response_location_coverage`")
    }
    if r.Gridsize == nil {
        errs = append(errs, "required field `gridsize` is missing for type `response_location_coverage`")
    }
    if r.ResultDef == nil {
        errs = append(errs, "required field `result_def` is missing for type `response_location_coverage`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `response_location_coverage`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `response_location_coverage`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
