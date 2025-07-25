// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseDeviceUpgrade represents a ResponseDeviceUpgrade struct.
type ResponseDeviceUpgrade struct {
    // enum: `error`, `inprogress`, `scheduled`, `starting`, `success`
    Status               UpgradeInfoStatusEnum  `json:"status"`
    // Epoch (seconds)
    Timestamp            float64                `json:"timestamp"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseDeviceUpgrade,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseDeviceUpgrade) String() string {
    return fmt.Sprintf(
    	"ResponseDeviceUpgrade[Status=%v, Timestamp=%v, AdditionalProperties=%v]",
    	r.Status, r.Timestamp, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseDeviceUpgrade.
// It customizes the JSON marshaling process for ResponseDeviceUpgrade objects.
func (r ResponseDeviceUpgrade) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "status", "timestamp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseDeviceUpgrade object to a map representation for JSON marshaling.
func (r ResponseDeviceUpgrade) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["status"] = r.Status
    structMap["timestamp"] = r.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseDeviceUpgrade.
// It customizes the JSON unmarshaling process for ResponseDeviceUpgrade objects.
func (r *ResponseDeviceUpgrade) UnmarshalJSON(input []byte) error {
    var temp tempResponseDeviceUpgrade
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "status", "timestamp")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Status = *temp.Status
    r.Timestamp = *temp.Timestamp
    return nil
}

// tempResponseDeviceUpgrade is a temporary struct used for validating the fields of ResponseDeviceUpgrade.
type tempResponseDeviceUpgrade  struct {
    Status    *UpgradeInfoStatusEnum `json:"status"`
    Timestamp *float64               `json:"timestamp"`
}

func (r *tempResponseDeviceUpgrade) validate() error {
    var errs []string
    if r.Status == nil {
        errs = append(errs, "required field `status` is missing for type `response_device_upgrade`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `response_device_upgrade`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
