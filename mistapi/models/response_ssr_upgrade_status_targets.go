package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseSsrUpgradeStatusTargets represents a ResponseSsrUpgradeStatusTargets struct.
type ResponseSsrUpgradeStatusTargets struct {
    Failed               []string               `json:"failed"`
    Queued               []string               `json:"queued"`
    Success              []string               `json:"success"`
    Upgrading            []string               `json:"upgrading"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseSsrUpgradeStatusTargets,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseSsrUpgradeStatusTargets) String() string {
    return fmt.Sprintf(
    	"ResponseSsrUpgradeStatusTargets[Failed=%v, Queued=%v, Success=%v, Upgrading=%v, AdditionalProperties=%v]",
    	r.Failed, r.Queued, r.Success, r.Upgrading, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseSsrUpgradeStatusTargets.
// It customizes the JSON marshaling process for ResponseSsrUpgradeStatusTargets objects.
func (r ResponseSsrUpgradeStatusTargets) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "failed", "queued", "success", "upgrading"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseSsrUpgradeStatusTargets object to a map representation for JSON marshaling.
func (r ResponseSsrUpgradeStatusTargets) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["failed"] = r.Failed
    structMap["queued"] = r.Queued
    structMap["success"] = r.Success
    structMap["upgrading"] = r.Upgrading
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseSsrUpgradeStatusTargets.
// It customizes the JSON unmarshaling process for ResponseSsrUpgradeStatusTargets objects.
func (r *ResponseSsrUpgradeStatusTargets) UnmarshalJSON(input []byte) error {
    var temp tempResponseSsrUpgradeStatusTargets
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
    r.AdditionalProperties = additionalProperties
    
    r.Failed = *temp.Failed
    r.Queued = *temp.Queued
    r.Success = *temp.Success
    r.Upgrading = *temp.Upgrading
    return nil
}

// tempResponseSsrUpgradeStatusTargets is a temporary struct used for validating the fields of ResponseSsrUpgradeStatusTargets.
type tempResponseSsrUpgradeStatusTargets  struct {
    Failed    *[]string `json:"failed"`
    Queued    *[]string `json:"queued"`
    Success   *[]string `json:"success"`
    Upgrading *[]string `json:"upgrading"`
}

func (r *tempResponseSsrUpgradeStatusTargets) validate() error {
    var errs []string
    if r.Failed == nil {
        errs = append(errs, "required field `failed` is missing for type `response_ssr_upgrade_status_targets`")
    }
    if r.Queued == nil {
        errs = append(errs, "required field `queued` is missing for type `response_ssr_upgrade_status_targets`")
    }
    if r.Success == nil {
        errs = append(errs, "required field `success` is missing for type `response_ssr_upgrade_status_targets`")
    }
    if r.Upgrading == nil {
        errs = append(errs, "required field `upgrading` is missing for type `response_ssr_upgrade_status_targets`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
