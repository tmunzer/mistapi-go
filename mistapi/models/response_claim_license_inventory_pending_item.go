// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseClaimLicenseInventoryPendingItem represents a ResponseClaimLicenseInventoryPendingItem struct.
type ResponseClaimLicenseInventoryPendingItem struct {
    Mac                  *string                `json:"mac,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseClaimLicenseInventoryPendingItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseClaimLicenseInventoryPendingItem) String() string {
    return fmt.Sprintf(
    	"ResponseClaimLicenseInventoryPendingItem[Mac=%v, AdditionalProperties=%v]",
    	r.Mac, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseClaimLicenseInventoryPendingItem.
// It customizes the JSON marshaling process for ResponseClaimLicenseInventoryPendingItem objects.
func (r ResponseClaimLicenseInventoryPendingItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "mac"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseClaimLicenseInventoryPendingItem object to a map representation for JSON marshaling.
func (r ResponseClaimLicenseInventoryPendingItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseClaimLicenseInventoryPendingItem.
// It customizes the JSON unmarshaling process for ResponseClaimLicenseInventoryPendingItem objects.
func (r *ResponseClaimLicenseInventoryPendingItem) UnmarshalJSON(input []byte) error {
    var temp tempResponseClaimLicenseInventoryPendingItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Mac = temp.Mac
    return nil
}

// tempResponseClaimLicenseInventoryPendingItem is a temporary struct used for validating the fields of ResponseClaimLicenseInventoryPendingItem.
type tempResponseClaimLicenseInventoryPendingItem  struct {
    Mac *string `json:"mac,omitempty"`
}
