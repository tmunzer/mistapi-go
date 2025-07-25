// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseOrgInventoryChange represents a ResponseOrgInventoryChange struct.
type ResponseOrgInventoryChange struct {
    Error                []string                         `json:"error"`
    // enum: `assign`, `delete`, `downgrade_to_jsi`, `unassign`, `upgrade_to_mist`
    Op                   ResponseOrgInventoryChangeOpEnum `json:"op"`
    Reason               []string                         `json:"reason"`
    Success              []string                         `json:"success"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseOrgInventoryChange,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseOrgInventoryChange) String() string {
    return fmt.Sprintf(
    	"ResponseOrgInventoryChange[Error=%v, Op=%v, Reason=%v, Success=%v, AdditionalProperties=%v]",
    	r.Error, r.Op, r.Reason, r.Success, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseOrgInventoryChange.
// It customizes the JSON marshaling process for ResponseOrgInventoryChange objects.
func (r ResponseOrgInventoryChange) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "error", "op", "reason", "success"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseOrgInventoryChange object to a map representation for JSON marshaling.
func (r ResponseOrgInventoryChange) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["error"] = r.Error
    structMap["op"] = r.Op
    structMap["reason"] = r.Reason
    structMap["success"] = r.Success
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseOrgInventoryChange.
// It customizes the JSON unmarshaling process for ResponseOrgInventoryChange objects.
func (r *ResponseOrgInventoryChange) UnmarshalJSON(input []byte) error {
    var temp tempResponseOrgInventoryChange
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "error", "op", "reason", "success")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Error = *temp.Error
    r.Op = *temp.Op
    r.Reason = *temp.Reason
    r.Success = *temp.Success
    return nil
}

// tempResponseOrgInventoryChange is a temporary struct used for validating the fields of ResponseOrgInventoryChange.
type tempResponseOrgInventoryChange  struct {
    Error   *[]string                         `json:"error"`
    Op      *ResponseOrgInventoryChangeOpEnum `json:"op"`
    Reason  *[]string                         `json:"reason"`
    Success *[]string                         `json:"success"`
}

func (r *tempResponseOrgInventoryChange) validate() error {
    var errs []string
    if r.Error == nil {
        errs = append(errs, "required field `error` is missing for type `response_org_inventory_change`")
    }
    if r.Op == nil {
        errs = append(errs, "required field `op` is missing for type `response_org_inventory_change`")
    }
    if r.Reason == nil {
        errs = append(errs, "required field `reason` is missing for type `response_org_inventory_change`")
    }
    if r.Success == nil {
        errs = append(errs, "required field `success` is missing for type `response_org_inventory_change`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
