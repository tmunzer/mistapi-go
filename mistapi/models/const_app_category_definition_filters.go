package models

import (
    "encoding/json"
)

// ConstAppCategoryDefinitionFilters represents a ConstAppCategoryDefinitionFilters struct.
type ConstAppCategoryDefinitionFilters struct {
    Srx                  []string               `json:"srx,omitempty"`
    Ssr                  []string               `json:"ssr,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstAppCategoryDefinitionFilters.
// It customizes the JSON marshaling process for ConstAppCategoryDefinitionFilters objects.
func (c ConstAppCategoryDefinitionFilters) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "srx", "ssr"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstAppCategoryDefinitionFilters object to a map representation for JSON marshaling.
func (c ConstAppCategoryDefinitionFilters) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Srx != nil {
        structMap["srx"] = c.Srx
    }
    if c.Ssr != nil {
        structMap["ssr"] = c.Ssr
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstAppCategoryDefinitionFilters.
// It customizes the JSON unmarshaling process for ConstAppCategoryDefinitionFilters objects.
func (c *ConstAppCategoryDefinitionFilters) UnmarshalJSON(input []byte) error {
    var temp tempConstAppCategoryDefinitionFilters
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "srx", "ssr")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Srx = temp.Srx
    c.Ssr = temp.Ssr
    return nil
}

// tempConstAppCategoryDefinitionFilters is a temporary struct used for validating the fields of ConstAppCategoryDefinitionFilters.
type tempConstAppCategoryDefinitionFilters  struct {
    Srx []string `json:"srx,omitempty"`
    Ssr []string `json:"ssr,omitempty"`
}
