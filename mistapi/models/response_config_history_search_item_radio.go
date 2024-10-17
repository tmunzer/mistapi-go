package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// ResponseConfigHistorySearchItemRadio represents a ResponseConfigHistorySearchItemRadio struct.
type ResponseConfigHistorySearchItemRadio struct {
	Band                 string         `json:"band"`
	Channel              int            `json:"channel"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseConfigHistorySearchItemRadio.
// It customizes the JSON marshaling process for ResponseConfigHistorySearchItemRadio objects.
func (r ResponseConfigHistorySearchItemRadio) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseConfigHistorySearchItemRadio object to a map representation for JSON marshaling.
func (r ResponseConfigHistorySearchItemRadio) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	structMap["band"] = r.Band
	structMap["channel"] = r.Channel
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseConfigHistorySearchItemRadio.
// It customizes the JSON unmarshaling process for ResponseConfigHistorySearchItemRadio objects.
func (r *ResponseConfigHistorySearchItemRadio) UnmarshalJSON(input []byte) error {
	var temp tempResponseConfigHistorySearchItemRadio
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "band", "channel")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.Band = *temp.Band
	r.Channel = *temp.Channel
	return nil
}

// tempResponseConfigHistorySearchItemRadio is a temporary struct used for validating the fields of ResponseConfigHistorySearchItemRadio.
type tempResponseConfigHistorySearchItemRadio struct {
	Band    *string `json:"band"`
	Channel *int    `json:"channel"`
}

func (r *tempResponseConfigHistorySearchItemRadio) validate() error {
	var errs []string
	if r.Band == nil {
		errs = append(errs, "required field `band` is missing for type `response_config_history_search_item_radio`")
	}
	if r.Channel == nil {
		errs = append(errs, "required field `channel` is missing for type `response_config_history_search_item_radio`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
