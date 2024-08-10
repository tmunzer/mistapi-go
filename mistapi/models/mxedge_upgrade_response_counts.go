package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// MxedgeUpgradeResponseCounts represents a MxedgeUpgradeResponseCounts struct.
type MxedgeUpgradeResponseCounts struct {
    Failed               int            `json:"failed"`
    Queued               int            `json:"queued"`
    Success              int            `json:"success"`
    Upgrading            int            `json:"upgrading"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeUpgradeResponseCounts.
// It customizes the JSON marshaling process for MxedgeUpgradeResponseCounts objects.
func (m MxedgeUpgradeResponseCounts) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeUpgradeResponseCounts object to a map representation for JSON marshaling.
func (m MxedgeUpgradeResponseCounts) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    structMap["failed"] = m.Failed
    structMap["queued"] = m.Queued
    structMap["success"] = m.Success
    structMap["upgrading"] = m.Upgrading
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeUpgradeResponseCounts.
// It customizes the JSON unmarshaling process for MxedgeUpgradeResponseCounts objects.
func (m *MxedgeUpgradeResponseCounts) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeUpgradeResponseCounts
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "failed", "queued", "success", "upgrading")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Failed = *temp.Failed
    m.Queued = *temp.Queued
    m.Success = *temp.Success
    m.Upgrading = *temp.Upgrading
    return nil
}

// tempMxedgeUpgradeResponseCounts is a temporary struct used for validating the fields of MxedgeUpgradeResponseCounts.
type tempMxedgeUpgradeResponseCounts  struct {
    Failed    *int `json:"failed"`
    Queued    *int `json:"queued"`
    Success   *int `json:"success"`
    Upgrading *int `json:"upgrading"`
}

func (m *tempMxedgeUpgradeResponseCounts) validate() error {
    var errs []string
    if m.Failed == nil {
        errs = append(errs, "required field `failed` is missing for type `mxedge_upgrade_response_counts`")
    }
    if m.Queued == nil {
        errs = append(errs, "required field `queued` is missing for type `mxedge_upgrade_response_counts`")
    }
    if m.Success == nil {
        errs = append(errs, "required field `success` is missing for type `mxedge_upgrade_response_counts`")
    }
    if m.Upgrading == nil {
        errs = append(errs, "required field `upgrading` is missing for type `mxedge_upgrade_response_counts`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
