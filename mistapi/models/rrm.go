package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// Rrm represents a Rrm struct.
// RRM
type Rrm struct {
    // proposal on band 2.4G, key is ap_id, value is the proposal
    Band24               map[string]RrmBand `json:"band_24"`
    Band24Metric         RrmBandMetric      `json:"band_24_metric"`
    // proposal on band 5G, key is ap_id, value is the proposal
    Band5                map[string]RrmBand `json:"band_5"`
    Band5Metric          RrmBandMetric      `json:"band_5_metric"`
    // proposal on band 6G, key is ap_id, value is the proposal
    Band6                map[string]RrmBand `json:"band_6,omitempty"`
    Band6Metric          *RrmBandMetric     `json:"band_6_metric,omitempty"`
    // RF Template
    Rftemplate           RfTemplate         `json:"rftemplate"`
    RftemplateId         uuid.UUID          `json:"rftemplate_id"`
    RftemplateName       string             `json:"rftemplate_name"`
    // enum: `ready`, `unknown`, `updating`
    Status               RrmStatusEnum      `json:"status"`
    // time where the status was updated
    Timestamp            float64            `json:"timestamp"`
    AdditionalProperties map[string]any     `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Rrm.
// It customizes the JSON marshaling process for Rrm objects.
func (r Rrm) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the Rrm object to a map representation for JSON marshaling.
func (r Rrm) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["band_24"] = r.Band24
    structMap["band_24_metric"] = r.Band24Metric.toMap()
    structMap["band_5"] = r.Band5
    structMap["band_5_metric"] = r.Band5Metric.toMap()
    if r.Band6 != nil {
        structMap["band_6"] = r.Band6
    }
    if r.Band6Metric != nil {
        structMap["band_6_metric"] = r.Band6Metric.toMap()
    }
    structMap["rftemplate"] = r.Rftemplate.toMap()
    structMap["rftemplate_id"] = r.RftemplateId
    structMap["rftemplate_name"] = r.RftemplateName
    structMap["status"] = r.Status
    structMap["timestamp"] = r.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Rrm.
// It customizes the JSON unmarshaling process for Rrm objects.
func (r *Rrm) UnmarshalJSON(input []byte) error {
    var temp rrm
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "band_24", "band_24_metric", "band_5", "band_5_metric", "band_6", "band_6_metric", "rftemplate", "rftemplate_id", "rftemplate_name", "status", "timestamp")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Band24 = *temp.Band24
    r.Band24Metric = *temp.Band24Metric
    r.Band5 = *temp.Band5
    r.Band5Metric = *temp.Band5Metric
    r.Band6 = temp.Band6
    r.Band6Metric = temp.Band6Metric
    r.Rftemplate = *temp.Rftemplate
    r.RftemplateId = *temp.RftemplateId
    r.RftemplateName = *temp.RftemplateName
    r.Status = *temp.Status
    r.Timestamp = *temp.Timestamp
    return nil
}

// rrm is a temporary struct used for validating the fields of Rrm.
type rrm  struct {
    Band24         *map[string]RrmBand `json:"band_24"`
    Band24Metric   *RrmBandMetric      `json:"band_24_metric"`
    Band5          *map[string]RrmBand `json:"band_5"`
    Band5Metric    *RrmBandMetric      `json:"band_5_metric"`
    Band6          map[string]RrmBand  `json:"band_6,omitempty"`
    Band6Metric    *RrmBandMetric      `json:"band_6_metric,omitempty"`
    Rftemplate     *RfTemplate         `json:"rftemplate"`
    RftemplateId   *uuid.UUID          `json:"rftemplate_id"`
    RftemplateName *string             `json:"rftemplate_name"`
    Status         *RrmStatusEnum      `json:"status"`
    Timestamp      *float64            `json:"timestamp"`
}

func (r *rrm) validate() error {
    var errs []string
    if r.Band24 == nil {
        errs = append(errs, "required field `band_24` is missing for type `Rrm`")
    }
    if r.Band24Metric == nil {
        errs = append(errs, "required field `band_24_metric` is missing for type `Rrm`")
    }
    if r.Band5 == nil {
        errs = append(errs, "required field `band_5` is missing for type `Rrm`")
    }
    if r.Band5Metric == nil {
        errs = append(errs, "required field `band_5_metric` is missing for type `Rrm`")
    }
    if r.Rftemplate == nil {
        errs = append(errs, "required field `rftemplate` is missing for type `Rrm`")
    }
    if r.RftemplateId == nil {
        errs = append(errs, "required field `rftemplate_id` is missing for type `Rrm`")
    }
    if r.RftemplateName == nil {
        errs = append(errs, "required field `rftemplate_name` is missing for type `Rrm`")
    }
    if r.Status == nil {
        errs = append(errs, "required field `status` is missing for type `Rrm`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `Rrm`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
