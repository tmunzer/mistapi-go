package models

import (
    "encoding/json"
)

// ResponseInventory represents a ResponseInventory struct.
type ResponseInventory struct {
    Added                []string                                    `json:"added,omitempty"`
    Duplicated           []string                                    `json:"duplicated,omitempty"`
    Error                []string                                    `json:"error,omitempty"`
    InventoryAdded       []ResponseInventoryInventoryAddedItems      `json:"inventory_added,omitempty"`
    InventoryDuplicated  []ResponseInventoryInventoryDuplicatedItems `json:"inventory_duplicated,omitempty"`
    AdditionalProperties map[string]any                              `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseInventory.
// It customizes the JSON marshaling process for ResponseInventory objects.
func (r ResponseInventory) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInventory object to a map representation for JSON marshaling.
func (r ResponseInventory) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInventory.
// It customizes the JSON unmarshaling process for ResponseInventory objects.
func (r *ResponseInventory) UnmarshalJSON(input []byte) error {
    var temp responseInventory
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "added", "duplicated", "error", "inventory_added", "inventory_duplicated")
    if err != nil {
    	return err
    }
    
    r.AdditionalProperties = additionalProperties
    r.Added = temp.Added
    r.Duplicated = temp.Duplicated
    r.Error = temp.Error
    r.InventoryAdded = temp.InventoryAdded
    r.InventoryDuplicated = temp.InventoryDuplicated
    return nil
}

// responseInventory is a temporary struct used for validating the fields of ResponseInventory.
type responseInventory  struct {
    Added               []string                                    `json:"added,omitempty"`
    Duplicated          []string                                    `json:"duplicated,omitempty"`
    Error               []string                                    `json:"error,omitempty"`
    InventoryAdded      []ResponseInventoryInventoryAddedItems      `json:"inventory_added,omitempty"`
    InventoryDuplicated []ResponseInventoryInventoryDuplicatedItems `json:"inventory_duplicated,omitempty"`
}
