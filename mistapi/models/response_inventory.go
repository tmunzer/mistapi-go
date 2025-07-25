// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// ResponseInventory represents a ResponseInventory struct.
type ResponseInventory struct {
    Added                []string                                    `json:"added,omitempty"`
    Duplicated           []string                                    `json:"duplicated,omitempty"`
    Error                []string                                    `json:"error,omitempty"`
    InventoryAdded       []ResponseInventoryInventoryAddedItems      `json:"inventory_added,omitempty"`
    InventoryDuplicated  []ResponseInventoryInventoryDuplicatedItems `json:"inventory_duplicated,omitempty"`
    Reason               []string                                    `json:"reason,omitempty"`
    AdditionalProperties map[string]interface{}                      `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseInventory,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseInventory) String() string {
    return fmt.Sprintf(
    	"ResponseInventory[Added=%v, Duplicated=%v, Error=%v, InventoryAdded=%v, InventoryDuplicated=%v, Reason=%v, AdditionalProperties=%v]",
    	r.Added, r.Duplicated, r.Error, r.InventoryAdded, r.InventoryDuplicated, r.Reason, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseInventory.
// It customizes the JSON marshaling process for ResponseInventory objects.
func (r ResponseInventory) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "added", "duplicated", "error", "inventory_added", "inventory_duplicated", "reason"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInventory object to a map representation for JSON marshaling.
func (r ResponseInventory) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Added != nil {
        structMap["added"] = r.Added
    }
    if r.Duplicated != nil {
        structMap["duplicated"] = r.Duplicated
    }
    if r.Error != nil {
        structMap["error"] = r.Error
    }
    if r.InventoryAdded != nil {
        structMap["inventory_added"] = r.InventoryAdded
    }
    if r.InventoryDuplicated != nil {
        structMap["inventory_duplicated"] = r.InventoryDuplicated
    }
    if r.Reason != nil {
        structMap["reason"] = r.Reason
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInventory.
// It customizes the JSON unmarshaling process for ResponseInventory objects.
func (r *ResponseInventory) UnmarshalJSON(input []byte) error {
    var temp tempResponseInventory
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "added", "duplicated", "error", "inventory_added", "inventory_duplicated", "reason")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Added = temp.Added
    r.Duplicated = temp.Duplicated
    r.Error = temp.Error
    r.InventoryAdded = temp.InventoryAdded
    r.InventoryDuplicated = temp.InventoryDuplicated
    r.Reason = temp.Reason
    return nil
}

// tempResponseInventory is a temporary struct used for validating the fields of ResponseInventory.
type tempResponseInventory  struct {
    Added               []string                                    `json:"added,omitempty"`
    Duplicated          []string                                    `json:"duplicated,omitempty"`
    Error               []string                                    `json:"error,omitempty"`
    InventoryAdded      []ResponseInventoryInventoryAddedItems      `json:"inventory_added,omitempty"`
    InventoryDuplicated []ResponseInventoryInventoryDuplicatedItems `json:"inventory_duplicated,omitempty"`
    Reason              []string                                    `json:"reason,omitempty"`
}
