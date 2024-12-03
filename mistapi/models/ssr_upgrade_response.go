package models

import (
    "encoding/json"
    "errors"
    "github.com/google/uuid"
    "strings"
)

// SsrUpgradeResponse represents a SsrUpgradeResponse struct.
type SsrUpgradeResponse struct {
    Channel              string                   `json:"channel"`
    Counts               SsrUpgradeResponseCounts `json:"counts"`
    DeviceType           string                   `json:"device_type"`
    // Unique ID of the object instance in the Mist Organnization
    Id                   uuid.UUID                `json:"id"`
    Status               string                   `json:"status"`
    Strategy             string                   `json:"strategy"`
    Versions             map[string]string        `json:"versions"`
    AdditionalProperties map[string]interface{}   `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsrUpgradeResponse.
// It customizes the JSON marshaling process for SsrUpgradeResponse objects.
func (s SsrUpgradeResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "channel", "counts", "device_type", "id", "status", "strategy", "versions"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SsrUpgradeResponse object to a map representation for JSON marshaling.
func (s SsrUpgradeResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["channel"] = s.Channel
    structMap["counts"] = s.Counts.toMap()
    structMap["device_type"] = s.DeviceType
    structMap["id"] = s.Id
    structMap["status"] = s.Status
    structMap["strategy"] = s.Strategy
    structMap["versions"] = s.Versions
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrUpgradeResponse.
// It customizes the JSON unmarshaling process for SsrUpgradeResponse objects.
func (s *SsrUpgradeResponse) UnmarshalJSON(input []byte) error {
    var temp tempSsrUpgradeResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "channel", "counts", "device_type", "id", "status", "strategy", "versions")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Channel = *temp.Channel
    s.Counts = *temp.Counts
    s.DeviceType = *temp.DeviceType
    s.Id = *temp.Id
    s.Status = *temp.Status
    s.Strategy = *temp.Strategy
    s.Versions = *temp.Versions
    return nil
}

// tempSsrUpgradeResponse is a temporary struct used for validating the fields of SsrUpgradeResponse.
type tempSsrUpgradeResponse  struct {
    Channel    *string                   `json:"channel"`
    Counts     *SsrUpgradeResponseCounts `json:"counts"`
    DeviceType *string                   `json:"device_type"`
    Id         *uuid.UUID                `json:"id"`
    Status     *string                   `json:"status"`
    Strategy   *string                   `json:"strategy"`
    Versions   *map[string]string        `json:"versions"`
}

func (s *tempSsrUpgradeResponse) validate() error {
    var errs []string
    if s.Channel == nil {
        errs = append(errs, "required field `channel` is missing for type `ssr_upgrade_response`")
    }
    if s.Counts == nil {
        errs = append(errs, "required field `counts` is missing for type `ssr_upgrade_response`")
    }
    if s.DeviceType == nil {
        errs = append(errs, "required field `device_type` is missing for type `ssr_upgrade_response`")
    }
    if s.Id == nil {
        errs = append(errs, "required field `id` is missing for type `ssr_upgrade_response`")
    }
    if s.Status == nil {
        errs = append(errs, "required field `status` is missing for type `ssr_upgrade_response`")
    }
    if s.Strategy == nil {
        errs = append(errs, "required field `strategy` is missing for type `ssr_upgrade_response`")
    }
    if s.Versions == nil {
        errs = append(errs, "required field `versions` is missing for type `ssr_upgrade_response`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
