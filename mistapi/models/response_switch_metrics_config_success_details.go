package models

import (
	"encoding/json"
)

// ResponseSwitchMetricsConfigSuccessDetails represents a ResponseSwitchMetricsConfigSuccessDetails struct.
type ResponseSwitchMetricsConfigSuccessDetails struct {
	ConfigSuccessCount   *int           `json:"config_success_count,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsConfigSuccessDetails.
// It customizes the JSON marshaling process for ResponseSwitchMetricsConfigSuccessDetails objects.
func (r ResponseSwitchMetricsConfigSuccessDetails) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsConfigSuccessDetails object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsConfigSuccessDetails) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, r.AdditionalProperties)
	if r.ConfigSuccessCount != nil {
		structMap["config_success_count"] = r.ConfigSuccessCount
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSwitchMetricsConfigSuccessDetails.
// It customizes the JSON unmarshaling process for ResponseSwitchMetricsConfigSuccessDetails objects.
func (r *ResponseSwitchMetricsConfigSuccessDetails) UnmarshalJSON(input []byte) error {
	var temp tempResponseSwitchMetricsConfigSuccessDetails
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "config_success_count")
	if err != nil {
		return err
	}

	r.AdditionalProperties = additionalProperties
	r.ConfigSuccessCount = temp.ConfigSuccessCount
	return nil
}

// tempResponseSwitchMetricsConfigSuccessDetails is a temporary struct used for validating the fields of ResponseSwitchMetricsConfigSuccessDetails.
type tempResponseSwitchMetricsConfigSuccessDetails struct {
	ConfigSuccessCount *int `json:"config_success_count,omitempty"`
}
