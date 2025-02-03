package models

import (
    "encoding/json"
    "fmt"
)

// MxedgeTuntermIgmpSnoopingQuerier represents a MxedgeTuntermIgmpSnoopingQuerier struct.
type MxedgeTuntermIgmpSnoopingQuerier struct {
    // Querier's query response interval, in tenths-of-seconds
    MaxResponseTime      *int                   `json:"max_response_time,omitempty"`
    // The MTU we use (needed when forming large IGMPv3 Reports)
    Mtu                  *int                   `json:"mtu,omitempty"`
    // Querier's query interval, in seconds
    QueryInterval        *int                   `json:"query_interval,omitempty"`
    // Querier's robustness
    Robustness           *int                   `json:"robustness,omitempty"`
    // Querier's maximum protocol version
    Version              *int                   `json:"version,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MxedgeTuntermIgmpSnoopingQuerier,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MxedgeTuntermIgmpSnoopingQuerier) String() string {
    return fmt.Sprintf(
    	"MxedgeTuntermIgmpSnoopingQuerier[MaxResponseTime=%v, Mtu=%v, QueryInterval=%v, Robustness=%v, Version=%v, AdditionalProperties=%v]",
    	m.MaxResponseTime, m.Mtu, m.QueryInterval, m.Robustness, m.Version, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermIgmpSnoopingQuerier.
// It customizes the JSON marshaling process for MxedgeTuntermIgmpSnoopingQuerier objects.
func (m MxedgeTuntermIgmpSnoopingQuerier) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "max_response_time", "mtu", "query_interval", "robustness", "version"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermIgmpSnoopingQuerier object to a map representation for JSON marshaling.
func (m MxedgeTuntermIgmpSnoopingQuerier) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.MaxResponseTime != nil {
        structMap["max_response_time"] = m.MaxResponseTime
    }
    if m.Mtu != nil {
        structMap["mtu"] = m.Mtu
    }
    if m.QueryInterval != nil {
        structMap["query_interval"] = m.QueryInterval
    }
    if m.Robustness != nil {
        structMap["robustness"] = m.Robustness
    }
    if m.Version != nil {
        structMap["version"] = m.Version
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermIgmpSnoopingQuerier.
// It customizes the JSON unmarshaling process for MxedgeTuntermIgmpSnoopingQuerier objects.
func (m *MxedgeTuntermIgmpSnoopingQuerier) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermIgmpSnoopingQuerier
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "max_response_time", "mtu", "query_interval", "robustness", "version")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.MaxResponseTime = temp.MaxResponseTime
    m.Mtu = temp.Mtu
    m.QueryInterval = temp.QueryInterval
    m.Robustness = temp.Robustness
    m.Version = temp.Version
    return nil
}

// tempMxedgeTuntermIgmpSnoopingQuerier is a temporary struct used for validating the fields of MxedgeTuntermIgmpSnoopingQuerier.
type tempMxedgeTuntermIgmpSnoopingQuerier  struct {
    MaxResponseTime *int `json:"max_response_time,omitempty"`
    Mtu             *int `json:"mtu,omitempty"`
    QueryInterval   *int `json:"query_interval,omitempty"`
    Robustness      *int `json:"robustness,omitempty"`
    Version         *int `json:"version,omitempty"`
}
