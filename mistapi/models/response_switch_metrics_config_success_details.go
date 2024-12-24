package models

import (
    "encoding/json"
    "fmt"
)

// ResponseSwitchMetricsConfigSuccessDetails represents a ResponseSwitchMetricsConfigSuccessDetails struct.
type ResponseSwitchMetricsConfigSuccessDetails struct {
    ConfigSuccessCount   *int                   `json:"config_success_count,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSwitchMetricsConfigSuccessDetails,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSwitchMetricsConfigSuccessDetails) String() string {
    return fmt.Sprintf(
    	"ResponseSwitchMetricsConfigSuccessDetails[ConfigSuccessCount=%v, AdditionalProperties=%v]",
    	r.ConfigSuccessCount, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSwitchMetricsConfigSuccessDetails.
// It customizes the JSON marshaling process for ResponseSwitchMetricsConfigSuccessDetails objects.
func (r ResponseSwitchMetricsConfigSuccessDetails) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "config_success_count"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSwitchMetricsConfigSuccessDetails object to a map representation for JSON marshaling.
func (r ResponseSwitchMetricsConfigSuccessDetails) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "config_success_count")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.ConfigSuccessCount = temp.ConfigSuccessCount
    return nil
}

// tempResponseSwitchMetricsConfigSuccessDetails is a temporary struct used for validating the fields of ResponseSwitchMetricsConfigSuccessDetails.
type tempResponseSwitchMetricsConfigSuccessDetails  struct {
    ConfigSuccessCount *int `json:"config_success_count,omitempty"`
}
