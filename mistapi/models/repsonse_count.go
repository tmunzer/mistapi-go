package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// RepsonseCount represents a RepsonseCount struct.
type RepsonseCount struct {
    Distinct             string                 `json:"distinct"`
    End                  int                    `json:"end"`
    Limit                int                    `json:"limit"`
    Results              []CountResult          `json:"results"`
    Start                int                    `json:"start"`
    Total                int                    `json:"total"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RepsonseCount,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RepsonseCount) String() string {
    return fmt.Sprintf(
    	"RepsonseCount[Distinct=%v, End=%v, Limit=%v, Results=%v, Start=%v, Total=%v, AdditionalProperties=%v]",
    	r.Distinct, r.End, r.Limit, r.Results, r.Start, r.Total, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RepsonseCount.
// It customizes the JSON marshaling process for RepsonseCount objects.
func (r RepsonseCount) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "distinct", "end", "limit", "results", "start", "total"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the RepsonseCount object to a map representation for JSON marshaling.
func (r RepsonseCount) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["distinct"] = r.Distinct
    structMap["end"] = r.End
    structMap["limit"] = r.Limit
    structMap["results"] = r.Results
    structMap["start"] = r.Start
    structMap["total"] = r.Total
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RepsonseCount.
// It customizes the JSON unmarshaling process for RepsonseCount objects.
func (r *RepsonseCount) UnmarshalJSON(input []byte) error {
    var temp tempRepsonseCount
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "distinct", "end", "limit", "results", "start", "total")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Distinct = *temp.Distinct
    r.End = *temp.End
    r.Limit = *temp.Limit
    r.Results = *temp.Results
    r.Start = *temp.Start
    r.Total = *temp.Total
    return nil
}

// tempRepsonseCount is a temporary struct used for validating the fields of RepsonseCount.
type tempRepsonseCount  struct {
    Distinct *string        `json:"distinct"`
    End      *int           `json:"end"`
    Limit    *int           `json:"limit"`
    Results  *[]CountResult `json:"results"`
    Start    *int           `json:"start"`
    Total    *int           `json:"total"`
}

func (r *tempRepsonseCount) validate() error {
    var errs []string
    if r.Distinct == nil {
        errs = append(errs, "required field `distinct` is missing for type `repsonse_count`")
    }
    if r.End == nil {
        errs = append(errs, "required field `end` is missing for type `repsonse_count`")
    }
    if r.Limit == nil {
        errs = append(errs, "required field `limit` is missing for type `repsonse_count`")
    }
    if r.Results == nil {
        errs = append(errs, "required field `results` is missing for type `repsonse_count`")
    }
    if r.Start == nil {
        errs = append(errs, "required field `start` is missing for type `repsonse_count`")
    }
    if r.Total == nil {
        errs = append(errs, "required field `total` is missing for type `repsonse_count`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
