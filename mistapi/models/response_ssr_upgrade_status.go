package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// ResponseSsrUpgradeStatus represents a ResponseSsrUpgradeStatus struct.
type ResponseSsrUpgradeStatus struct {
    Channel              string                          `json:"channel"`
    DeviceType           *string                         `json:"device_type,omitempty"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID                       `json:"id"`
    Status               string                          `json:"status"`
    Targets              ResponseSsrUpgradeStatusTargets `json:"targets"`
    Versions             interface{}                     `json:"versions"`
    AdditionalProperties map[string]any                  `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrUpgradeStatus.
// It customizes the JSON marshaling process for ResponseSsrUpgradeStatus objects.
func (r ResponseSsrUpgradeStatus) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrUpgradeStatus object to a map representation for JSON marshaling.
func (r ResponseSsrUpgradeStatus) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["channel"] = r.Channel
    if r.DeviceType != nil {
        structMap["device_type"] = r.DeviceType
    }
    structMap["id"] = r.Id
    structMap["status"] = r.Status
    structMap["targets"] = r.Targets.toMap()
    structMap["versions"] = r.Versions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrUpgradeStatus.
// It customizes the JSON unmarshaling process for ResponseSsrUpgradeStatus objects.
func (r *ResponseSsrUpgradeStatus) UnmarshalJSON(input []byte) error {
    var temp tempResponseSsrUpgradeStatus
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "channel", "device_type", "id", "status", "targets", "versions")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Channel = *temp.Channel
    r.DeviceType = temp.DeviceType
    r.Id = *temp.Id
    r.Status = *temp.Status
    r.Targets = *temp.Targets
    r.Versions = *temp.Versions
    return nil
}

// tempResponseSsrUpgradeStatus is a temporary struct used for validating the fields of ResponseSsrUpgradeStatus.
type tempResponseSsrUpgradeStatus  struct {
    Channel    *string                          `json:"channel"`
    DeviceType *string                          `json:"device_type,omitempty"`
    Id         *uuid.UUID                       `json:"id"`
    Status     *string                          `json:"status"`
    Targets    *ResponseSsrUpgradeStatusTargets `json:"targets"`
    Versions   *interface{}                     `json:"versions"`
}

func (r *tempResponseSsrUpgradeStatus) validate() error {
    var errs []string
    if r.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `response_ssr_upgrade_status`")
    }
    if r.Id == nil {
        errs = append(errs, "required field `id` is missing for type `response_ssr_upgrade_status`")
    }
    if r.Status == nil {
        errs = append(errs, "required field `status` is missing for type `response_ssr_upgrade_status`")
    }
    if r.Targets == nil {
        errs = append(errs, "required field `targets` is missing for type `response_ssr_upgrade_status`")
    }
    if r.Versions == nil {
        errs = append(errs, "required field `versions` is missing for type `response_ssr_upgrade_status`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
