package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseSsrUpgradeStatusTargets represents a ResponseSsrUpgradeStatusTargets struct.
type ResponseSsrUpgradeStatusTargets struct {
    Failed               []string       `json:"failed"`
    Queued               []string       `json:"queued"`
    Success              []string       `json:"success"`
    Upgrading            []string       `json:"upgrading"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrUpgradeStatusTargets.
// It customizes the JSON marshaling process for ResponseSsrUpgradeStatusTargets objects.
func (r ResponseSsrUpgradeStatusTargets) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrUpgradeStatusTargets object to a map representation for JSON marshaling.
func (r ResponseSsrUpgradeStatusTargets) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["failed"] = r.Failed
    structMap["queued"] = r.Queued
    structMap["success"] = r.Success
    structMap["upgrading"] = r.Upgrading
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrUpgradeStatusTargets.
// It customizes the JSON unmarshaling process for ResponseSsrUpgradeStatusTargets objects.
func (r *ResponseSsrUpgradeStatusTargets) UnmarshalJSON(input []byte) error {
    var temp responseSsrUpgradeStatusTargets
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
    
    r.AdditionalProperties = additionalProperties
    r.Failed = *temp.Failed
    r.Queued = *temp.Queued
    r.Success = *temp.Success
    r.Upgrading = *temp.Upgrading
    return nil
}

// responseSsrUpgradeStatusTargets is a temporary struct used for validating the fields of ResponseSsrUpgradeStatusTargets.
type responseSsrUpgradeStatusTargets  struct {
    Failed    *[]string `json:"failed"`
    Queued    *[]string `json:"queued"`
    Success   *[]string `json:"success"`
    Upgrading *[]string `json:"upgrading"`
}

func (r *responseSsrUpgradeStatusTargets) validate() error {
    var errs []string
    if r.Failed == nil {
        errs = append(errs, "required field `failed` is missing for type `Response_Ssr_Upgrade_Status_Targets`")
    }
    if r.Queued == nil {
        errs = append(errs, "required field `queued` is missing for type `Response_Ssr_Upgrade_Status_Targets`")
    }
    if r.Success == nil {
        errs = append(errs, "required field `success` is missing for type `Response_Ssr_Upgrade_Status_Targets`")
    }
    if r.Upgrading == nil {
        errs = append(errs, "required field `upgrading` is missing for type `Response_Ssr_Upgrade_Status_Targets`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
