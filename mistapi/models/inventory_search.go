package models

import (
    "encoding/json"
)

// InventorySearch represents a InventorySearch struct.
type InventorySearch struct {
    Limit                *int                    `json:"limit,omitempty"`
    Page                 *int                    `json:"page,omitempty"`
    Results              []InventorySearchResult `json:"results,omitempty"`
    AdditionalProperties map[string]any          `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for InventorySearch.
// It customizes the JSON marshaling process for InventorySearch objects.
func (i InventorySearch) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(i.toMap())
}

// toMap converts the InventorySearch object to a map representation for JSON marshaling.
func (i InventorySearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, i.AdditionalProperties)
    if i.Limit != nil {
        structMap["limit"] = i.Limit
    }
    if i.Page != nil {
        structMap["page"] = i.Page
    }
    if i.Results != nil {
        structMap["results"] = i.Results
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for InventorySearch.
// It customizes the JSON unmarshaling process for InventorySearch objects.
func (i *InventorySearch) UnmarshalJSON(input []byte) error {
    var temp tempInventorySearch
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "limit", "page", "results")
    if err != nil {
    	return err
    }
    
    i.AdditionalProperties = additionalProperties
    i.Limit = temp.Limit
    i.Page = temp.Page
    i.Results = temp.Results
    return nil
}

// tempInventorySearch is a temporary struct used for validating the fields of InventorySearch.
type tempInventorySearch  struct {
    Limit   *int                    `json:"limit,omitempty"`
    Page    *int                    `json:"page,omitempty"`
    Results []InventorySearchResult `json:"results,omitempty"`
}
