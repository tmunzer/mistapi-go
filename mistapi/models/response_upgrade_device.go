package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseUpgradeDevice represents a ResponseUpgradeDevice struct.
type ResponseUpgradeDevice struct {
    // enum: `error`, `inprogress`, `scheduled`, `starting`, `success`
    Status               UpgradeInfoStatusEnum `json:"status"`
    // timestamp
    Timestamp            float64               `json:"timestamp"`
    AdditionalProperties map[string]any        `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseUpgradeDevice.
// It customizes the JSON marshaling process for ResponseUpgradeDevice objects.
func (r ResponseUpgradeDevice) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseUpgradeDevice object to a map representation for JSON marshaling.
func (r ResponseUpgradeDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["status"] = r.Status
    structMap["timestamp"] = r.Timestamp
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseUpgradeDevice.
// It customizes the JSON unmarshaling process for ResponseUpgradeDevice objects.
func (r *ResponseUpgradeDevice) UnmarshalJSON(input []byte) error {
    var temp tempResponseUpgradeDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "status", "timestamp")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Status = *temp.Status
    r.Timestamp = *temp.Timestamp
    return nil
}

// tempResponseUpgradeDevice is a temporary struct used for validating the fields of ResponseUpgradeDevice.
type tempResponseUpgradeDevice  struct {
    Status    *UpgradeInfoStatusEnum `json:"status"`
    Timestamp *float64               `json:"timestamp"`
}

func (r *tempResponseUpgradeDevice) validate() error {
    var errs []string
    if r.Status == nil {
        errs = append(errs, "required field `status` is missing for type `response_upgrade_device`")
    }
    if r.Timestamp == nil {
        errs = append(errs, "required field `timestamp` is missing for type `response_upgrade_device`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
