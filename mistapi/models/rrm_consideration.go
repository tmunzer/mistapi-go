package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// RrmConsideration represents a RrmConsideration struct.
type RrmConsideration struct {
    Channel              int            `json:"channel"`
    Noise                float64        `json:"noise"`
    // the avg RSSI heard from other APs (that does NOT belongs to the same site)
    OtherRssi            *float64       `json:"other_rssi,omitempty"`
    // SSID from other AP that we heard from with the max RSSI
    OtherSsid            *string        `json:"other_ssid,omitempty"`
    // utilization score, 0-1, lower means less utilization (cleaner RF)
    UtilScore            float64        `json:"util_score"`
    // non-wifi utilization score, 0-1, lower means less utilization (cleaner RF)
    UtilScoreNonWifi     float64        `json:"util_score_non_wifi"`
    // other utilization score, 0-1, lower means less utilization (cleaner RF)
    UtilScoreOther       float64        `json:"util_score_other"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for RrmConsideration.
// It customizes the JSON marshaling process for RrmConsideration objects.
func (r RrmConsideration) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the RrmConsideration object to a map representation for JSON marshaling.
func (r RrmConsideration) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["channel"] = r.Channel
    structMap["noise"] = r.Noise
    if r.OtherRssi != nil {
        structMap["other_rssi"] = r.OtherRssi
    }
    if r.OtherSsid != nil {
        structMap["other_ssid"] = r.OtherSsid
    }
    structMap["util_score"] = r.UtilScore
    structMap["util_score_non_wifi"] = r.UtilScoreNonWifi
    structMap["util_score_other"] = r.UtilScoreOther
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmConsideration.
// It customizes the JSON unmarshaling process for RrmConsideration objects.
func (r *RrmConsideration) UnmarshalJSON(input []byte) error {
    var temp rrmConsideration
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "noise", "other_rssi", "other_ssid", "util_score", "util_score_non_wifi", "util_score_other")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Channel = *temp.Channel
    r.Noise = *temp.Noise
    r.OtherRssi = temp.OtherRssi
    r.OtherSsid = temp.OtherSsid
    r.UtilScore = *temp.UtilScore
    r.UtilScoreNonWifi = *temp.UtilScoreNonWifi
    r.UtilScoreOther = *temp.UtilScoreOther
    return nil
}

// rrmConsideration is a temporary struct used for validating the fields of RrmConsideration.
type rrmConsideration  struct {
    Channel          *int     `json:"channel"`
    Noise            *float64 `json:"noise"`
    OtherRssi        *float64 `json:"other_rssi,omitempty"`
    OtherSsid        *string  `json:"other_ssid,omitempty"`
    UtilScore        *float64 `json:"util_score"`
    UtilScoreNonWifi *float64 `json:"util_score_non_wifi"`
    UtilScoreOther   *float64 `json:"util_score_other"`
}

func (r *rrmConsideration) validate() error {
    var errs []string
    if r.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `Rrm_Consideration`")
    }
    if r.Noise == nil {
        errs = append(errs, "required field `noise` is missing for type `Rrm_Consideration`")
    }
    if r.UtilScore == nil {
        errs = append(errs, "required field `util_score` is missing for type `Rrm_Consideration`")
    }
    if r.UtilScoreNonWifi == nil {
        errs = append(errs, "required field `util_score_non_wifi` is missing for type `Rrm_Consideration`")
    }
    if r.UtilScoreOther == nil {
        errs = append(errs, "required field `util_score_other` is missing for type `Rrm_Consideration`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}