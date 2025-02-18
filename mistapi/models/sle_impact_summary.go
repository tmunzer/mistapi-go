package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// SleImpactSummary represents a SleImpactSummary struct.
type SleImpactSummary struct {
    Ap                   []SleImpactSummaryApItem         `json:"ap"`
    Band                 []SleImpactSummaryBandItem       `json:"band"`
    Classifier           string                           `json:"classifier"`
    DeviceOs             []SleImpactSummaryDeviceOsItem   `json:"device_os"`
    DeviceType           []SleImpactSummaryDeviceTypeItem `json:"device_type"`
    End                  float64                          `json:"end"`
    Failure              string                           `json:"failure"`
    Metric               string                           `json:"metric"`
    Start                float64                          `json:"start"`
    Wlan                 []SleImpactSummaryWlanItem       `json:"wlan"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for SleImpactSummary,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SleImpactSummary) String() string {
    return fmt.Sprintf(
    	"SleImpactSummary[Ap=%v, Band=%v, Classifier=%v, DeviceOs=%v, DeviceType=%v, End=%v, Failure=%v, Metric=%v, Start=%v, Wlan=%v, AdditionalProperties=%v]",
    	s.Ap, s.Band, s.Classifier, s.DeviceOs, s.DeviceType, s.End, s.Failure, s.Metric, s.Start, s.Wlan, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SleImpactSummary.
// It customizes the JSON marshaling process for SleImpactSummary objects.
func (s SleImpactSummary) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "ap", "band", "classifier", "device_os", "device_type", "end", "failure", "metric", "start", "wlan"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SleImpactSummary object to a map representation for JSON marshaling.
func (s SleImpactSummary) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["ap"] = s.Ap
    structMap["band"] = s.Band
    structMap["classifier"] = s.Classifier
    structMap["device_os"] = s.DeviceOs
    structMap["device_type"] = s.DeviceType
    structMap["end"] = s.End
    structMap["failure"] = s.Failure
    structMap["metric"] = s.Metric
    structMap["start"] = s.Start
    structMap["wlan"] = s.Wlan
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SleImpactSummary.
// It customizes the JSON unmarshaling process for SleImpactSummary objects.
func (s *SleImpactSummary) UnmarshalJSON(input []byte) error {
    var temp tempSleImpactSummary
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "ap", "band", "classifier", "device_os", "device_type", "end", "failure", "metric", "start", "wlan")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Ap = *temp.Ap
    s.Band = *temp.Band
    s.Classifier = *temp.Classifier
    s.DeviceOs = *temp.DeviceOs
    s.DeviceType = *temp.DeviceType
    s.End = *temp.End
    s.Failure = *temp.Failure
    s.Metric = *temp.Metric
    s.Start = *temp.Start
    s.Wlan = *temp.Wlan
    return nil
}

// tempSleImpactSummary is a temporary struct used for validating the fields of SleImpactSummary.
type tempSleImpactSummary  struct {
    Ap         *[]SleImpactSummaryApItem         `json:"ap"`
    Band       *[]SleImpactSummaryBandItem       `json:"band"`
    Classifier *string                           `json:"classifier"`
    DeviceOs   *[]SleImpactSummaryDeviceOsItem   `json:"device_os"`
    DeviceType *[]SleImpactSummaryDeviceTypeItem `json:"device_type"`
    End        *float64                          `json:"end"`
    Failure    *string                           `json:"failure"`
    Metric     *string                           `json:"metric"`
    Start      *float64                          `json:"start"`
    Wlan       *[]SleImpactSummaryWlanItem       `json:"wlan"`
}

func (s *tempSleImpactSummary) validate() error {
    var errs []string
    if s.Ap == nil {
        errs = append(errs, "required field `ap` is missing for type `sle_impact_summary`")
    }
    if s.Band == nil {
        errs = append(errs, "required field `band` is missing for type `sle_impact_summary`")
    }
    if s.Classifier == nil {
        errs = append(errs, "required field `classifier` is missing for type `sle_impact_summary`")
    }
    if s.DeviceOs == nil {
        errs = append(errs, "required field `device_os` is missing for type `sle_impact_summary`")
    }
    if s.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `sle_impact_summary`")
    }
    if s.End == nil {
        errs = append(errs, "required field `end` is missing for type `sle_impact_summary`")
    }
    if s.Failure == nil {
        errs = append(errs, "required field `failure` is missing for type `sle_impact_summary`")
    }
    if s.Metric == nil {
        errs = append(errs, "required field `metric` is missing for type `sle_impact_summary`")
    }
    if s.Start == nil {
        errs = append(errs, "required field `start` is missing for type `sle_impact_summary`")
    }
    if s.Wlan == nil {
        errs = append(errs, "required field `wlan` is missing for type `sle_impact_summary`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
