package models

import (
    "encoding/json"
    "fmt"
)

// ResponseAutoplacement represents a ResponseAutoplacement struct.
type ResponseAutoplacement struct {
    // Property key is the AP MAC Address. Contains the validation status of each device.
    Devices              map[string]ResponseAutoplacementDevice `json:"devices,omitempty"`
    // Estimated runtime for the process in seconds.
    EstimatedRuntime     *int                                   `json:"estimated_runtime,omitempty"`
    // Provides the reason for the status.
    Reason               *string                                `json:"reason,omitempty"`
    // Indicates whether the autoplacement process has started.
    Started              *bool                                  `json:"started,omitempty"`
    // Indicates whether the autoplacement request is valid.
    Valid                *bool                                  `json:"valid,omitempty"`
    // Indicates whether the auto placement process will interrupt WiFi traffic.
    WifiInterrupting     *bool                                  `json:"wifi_interrupting,omitempty"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseAutoplacement,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseAutoplacement) String() string {
    return fmt.Sprintf(
    	"ResponseAutoplacement[Devices=%v, EstimatedRuntime=%v, Reason=%v, Started=%v, Valid=%v, WifiInterrupting=%v, AdditionalProperties=%v]",
    	r.Devices, r.EstimatedRuntime, r.Reason, r.Started, r.Valid, r.WifiInterrupting, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseAutoplacement.
// It customizes the JSON marshaling process for ResponseAutoplacement objects.
func (r ResponseAutoplacement) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "devices", "estimated_runtime", "reason", "started", "valid", "wifi_interrupting"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseAutoplacement object to a map representation for JSON marshaling.
func (r ResponseAutoplacement) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Devices != nil {
        structMap["devices"] = r.Devices
    }
    if r.EstimatedRuntime != nil {
        structMap["estimated_runtime"] = r.EstimatedRuntime
    }
    if r.Reason != nil {
        structMap["reason"] = r.Reason
    }
    if r.Started != nil {
        structMap["started"] = r.Started
    }
    if r.Valid != nil {
        structMap["valid"] = r.Valid
    }
    if r.WifiInterrupting != nil {
        structMap["wifi_interrupting"] = r.WifiInterrupting
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseAutoplacement.
// It customizes the JSON unmarshaling process for ResponseAutoplacement objects.
func (r *ResponseAutoplacement) UnmarshalJSON(input []byte) error {
    var temp tempResponseAutoplacement
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "devices", "estimated_runtime", "reason", "started", "valid", "wifi_interrupting")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Devices = temp.Devices
    r.EstimatedRuntime = temp.EstimatedRuntime
    r.Reason = temp.Reason
    r.Started = temp.Started
    r.Valid = temp.Valid
    r.WifiInterrupting = temp.WifiInterrupting
    return nil
}

// tempResponseAutoplacement is a temporary struct used for validating the fields of ResponseAutoplacement.
type tempResponseAutoplacement  struct {
    Devices          map[string]ResponseAutoplacementDevice `json:"devices,omitempty"`
    EstimatedRuntime *int                                   `json:"estimated_runtime,omitempty"`
    Reason           *string                                `json:"reason,omitempty"`
    Started          *bool                                  `json:"started,omitempty"`
    Valid            *bool                                  `json:"valid,omitempty"`
    WifiInterrupting *bool                                  `json:"wifi_interrupting,omitempty"`
}
