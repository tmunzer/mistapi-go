package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// SsrUpgradeResponseCounts represents a SsrUpgradeResponseCounts struct.
type SsrUpgradeResponseCounts struct {
    Failed               int                    `json:"failed"`
    Queued               int                    `json:"queued"`
    Success              int                    `json:"success"`
    Upgrading            int                    `json:"upgrading"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsrUpgradeResponseCounts.
// It customizes the JSON marshaling process for SsrUpgradeResponseCounts objects.
func (s SsrUpgradeResponseCounts) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "failed", "queued", "success", "upgrading"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SsrUpgradeResponseCounts object to a map representation for JSON marshaling.
func (s SsrUpgradeResponseCounts) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    structMap["failed"] = s.Failed
    structMap["queued"] = s.Queued
    structMap["success"] = s.Success
    structMap["upgrading"] = s.Upgrading
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsrUpgradeResponseCounts.
// It customizes the JSON unmarshaling process for SsrUpgradeResponseCounts objects.
func (s *SsrUpgradeResponseCounts) UnmarshalJSON(input []byte) error {
    var temp tempSsrUpgradeResponseCounts
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "failed", "queued", "success", "upgrading")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Failed = *temp.Failed
    s.Queued = *temp.Queued
    s.Success = *temp.Success
    s.Upgrading = *temp.Upgrading
    return nil
}

// tempSsrUpgradeResponseCounts is a temporary struct used for validating the fields of SsrUpgradeResponseCounts.
type tempSsrUpgradeResponseCounts  struct {
    Failed    *int `json:"failed"`
    Queued    *int `json:"queued"`
    Success   *int `json:"success"`
    Upgrading *int `json:"upgrading"`
}

func (s *tempSsrUpgradeResponseCounts) validate() error {
    var errs []string
    if s.Failed == nil {
        errs = append(errs, "required field `failed` is missing for type `ssr_upgrade_response_counts`")
    }
    if s.Queued == nil {
        errs = append(errs, "required field `queued` is missing for type `ssr_upgrade_response_counts`")
    }
    if s.Success == nil {
        errs = append(errs, "required field `success` is missing for type `ssr_upgrade_response_counts`")
    }
    if s.Upgrading == nil {
        errs = append(errs, "required field `upgrading` is missing for type `ssr_upgrade_response_counts`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
