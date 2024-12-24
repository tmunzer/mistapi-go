package models

import (
    "encoding/json"
    "fmt"
)

// InventorySearch represents a InventorySearch struct.
type InventorySearch struct {
    Limit                *int                    `json:"limit,omitempty"`
    Page                 *int                    `json:"page,omitempty"`
    Results              []InventorySearchResult `json:"results,omitempty"`
    AdditionalProperties map[string]interface{}  `json:"_"`
}

// String implements the fmt.Stringer interface for InventorySearch,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i InventorySearch) String() string {
    return fmt.Sprintf(
    	"InventorySearch[Limit=%v, Page=%v, Results=%v, AdditionalProperties=%v]",
    	i.Limit, i.Page, i.Results, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for InventorySearch.
// It customizes the JSON marshaling process for InventorySearch objects.
func (i InventorySearch) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "limit", "page", "results"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the InventorySearch object to a map representation for JSON marshaling.
func (i InventorySearch) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
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
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "limit", "page", "results")
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
