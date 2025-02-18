package models

import (
    "encoding/json"
    "fmt"
)

// ResponseDeviceBiosUpgrade represents a ResponseDeviceBiosUpgrade struct.
type ResponseDeviceBiosUpgrade struct {
    Status               *string                `json:"status,omitempty"`
    Timestamp            *int                   `json:"timestamp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDeviceBiosUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceBiosUpgrade) String() string {
    return fmt.Sprintf(
    	"ResponseDeviceBiosUpgrade[Status=%v, Timestamp=%v, AdditionalProperties=%v]",
    	r.Status, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceBiosUpgrade.
// It customizes the JSON marshaling process for ResponseDeviceBiosUpgrade objects.
func (r ResponseDeviceBiosUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "status", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceBiosUpgrade object to a map representation for JSON marshaling.
func (r ResponseDeviceBiosUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Status != nil {
        structMap["status"] = r.Status
    }
    if r.Timestamp != nil {
        structMap["timestamp"] = r.Timestamp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceBiosUpgrade.
// It customizes the JSON unmarshaling process for ResponseDeviceBiosUpgrade objects.
func (r *ResponseDeviceBiosUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceBiosUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "status", "timestamp")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Status = temp.Status
    r.Timestamp = temp.Timestamp
    return nil
}

// tempResponseDeviceBiosUpgrade is a temporary struct used for validating the fields of ResponseDeviceBiosUpgrade.
type tempResponseDeviceBiosUpgrade  struct {
    Status    *string `json:"status,omitempty"`
    Timestamp *int    `json:"timestamp,omitempty"`
}
