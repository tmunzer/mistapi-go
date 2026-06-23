// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// RrmChannelScore represents a RrmChannelScore struct.
// RRM utilization score for a channel
type RrmChannelScore struct {
	// RF channel number represented by this score record
	Channel int `json:"channel"`
	// Utilization score for the channel, 0-1, lower means cleaner RF
	UtilScore float64 `json:"util_score"`
	// Score contribution from noise, 0-1, lower means cleaner RF
	UtilScoreNoiseFloor float64 `json:"util_score_noise_floor"`
	// Score contribution from non-wifi utilization, 0-1, lower means cleaner RF
	UtilScoreNonWifi float64 `json:"util_score_non_wifi"`
	// Score contribution from RxOtherBss utilization (wifi packets destined for other radios), 0-1, lower means cleaner RF
	UtilScoreOther float64 `json:"util_score_other"`
	// Score contribution from radar detections, 0-1, lower means cleaner RF
	UtilScoreRadar float64 `json:"util_score_radar"`
	// Score contribution from undecodable wifi utilization (wifi packets which can't be decoded), 0-1, lower means cleaner RF
	UtilScoreUndecodableWifi float64 `json:"util_score_undecodable_wifi"`
	// Score contribution from unknown wifi utilization (wifi packets of unknown type), 0-1, lower means cleaner RF
	UtilScoreUnknownWifi float64                `json:"util_score_unknown_wifi"`
	AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for RrmChannelScore,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r RrmChannelScore) String() string {
	return fmt.Sprintf(
		"RrmChannelScore[Channel=%v, UtilScore=%v, UtilScoreNoiseFloor=%v, UtilScoreNonWifi=%v, UtilScoreOther=%v, UtilScoreRadar=%v, UtilScoreUndecodableWifi=%v, UtilScoreUnknownWifi=%v, AdditionalProperties=%v]",
		r.Channel, r.UtilScore, r.UtilScoreNoiseFloor, r.UtilScoreNonWifi, r.UtilScoreOther, r.UtilScoreRadar, r.UtilScoreUndecodableWifi, r.UtilScoreUnknownWifi, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for RrmChannelScore.
// It customizes the JSON marshaling process for RrmChannelScore objects.
func (r RrmChannelScore) MarshalJSON() (
	[]byte,
	error) {
	if err := DetectConflictingProperties(r.AdditionalProperties,
		"channel", "util_score", "util_score_noise_floor", "util_score_non_wifi", "util_score_other", "util_score_radar", "util_score_undecodable_wifi", "util_score_unknown_wifi"); err != nil {
		return []byte{}, err
	}
	return json.Marshal(r.toMap())
}

// toMap converts the RrmChannelScore object to a map representation for JSON marshaling.
func (r RrmChannelScore) toMap() map[string]any {
	structMap := make(map[string]any)
	MergeAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["channel"] = r.Channel
	structMap["util_score"] = r.UtilScore
	structMap["util_score_noise_floor"] = r.UtilScoreNoiseFloor
	structMap["util_score_non_wifi"] = r.UtilScoreNonWifi
	structMap["util_score_other"] = r.UtilScoreOther
	structMap["util_score_radar"] = r.UtilScoreRadar
	structMap["util_score_undecodable_wifi"] = r.UtilScoreUndecodableWifi
	structMap["util_score_unknown_wifi"] = r.UtilScoreUnknownWifi
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for RrmChannelScore.
// It customizes the JSON unmarshaling process for RrmChannelScore objects.
func (r *RrmChannelScore) UnmarshalJSON(input []byte) error {
	var temp tempRrmChannelScore
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "util_score", "util_score_noise_floor", "util_score_non_wifi", "util_score_other", "util_score_radar", "util_score_undecodable_wifi", "util_score_unknown_wifi")
	if err != nil {
		return err
	}
	r.AdditionalProperties = additionalProperties

	r.Channel = *temp.Channel
	r.UtilScore = *temp.UtilScore
	r.UtilScoreNoiseFloor = *temp.UtilScoreNoiseFloor
	r.UtilScoreNonWifi = *temp.UtilScoreNonWifi
	r.UtilScoreOther = *temp.UtilScoreOther
	r.UtilScoreRadar = *temp.UtilScoreRadar
	r.UtilScoreUndecodableWifi = *temp.UtilScoreUndecodableWifi
	r.UtilScoreUnknownWifi = *temp.UtilScoreUnknownWifi
	return nil
}

// tempRrmChannelScore is a temporary struct used for validating the fields of RrmChannelScore.
type tempRrmChannelScore struct {
	Channel                  *int     `json:"channel"`
	UtilScore                *float64 `json:"util_score"`
	UtilScoreNoiseFloor      *float64 `json:"util_score_noise_floor"`
	UtilScoreNonWifi         *float64 `json:"util_score_non_wifi"`
	UtilScoreOther           *float64 `json:"util_score_other"`
	UtilScoreRadar           *float64 `json:"util_score_radar"`
	UtilScoreUndecodableWifi *float64 `json:"util_score_undecodable_wifi"`
	UtilScoreUnknownWifi     *float64 `json:"util_score_unknown_wifi"`
}

func (r *tempRrmChannelScore) validate() error {
	var errs []string
	if r.Channel == nil {
		errs = append(errs, "required field `channel` is missing for type `rrm_channel_score`")
	}
	if r.UtilScore == nil {
		errs = append(errs, "required field `util_score` is missing for type `rrm_channel_score`")
	}
	if r.UtilScoreNoiseFloor == nil {
		errs = append(errs, "required field `util_score_noise_floor` is missing for type `rrm_channel_score`")
	}
	if r.UtilScoreNonWifi == nil {
		errs = append(errs, "required field `util_score_non_wifi` is missing for type `rrm_channel_score`")
	}
	if r.UtilScoreOther == nil {
		errs = append(errs, "required field `util_score_other` is missing for type `rrm_channel_score`")
	}
	if r.UtilScoreRadar == nil {
		errs = append(errs, "required field `util_score_radar` is missing for type `rrm_channel_score`")
	}
	if r.UtilScoreUndecodableWifi == nil {
		errs = append(errs, "required field `util_score_undecodable_wifi` is missing for type `rrm_channel_score`")
	}
	if r.UtilScoreUnknownWifi == nil {
		errs = append(errs, "required field `util_score_unknown_wifi` is missing for type `rrm_channel_score`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
