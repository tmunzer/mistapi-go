package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConstAppCategoryDefinition represents a ConstAppCategoryDefinition struct.
type ConstAppCategoryDefinition struct {
    // Description of the app category
    Display              string                             `json:"display"`
    Filters              *ConstAppCategoryDefinitionFilters `json:"filters,omitempty"`
    // List of other App Categories contained by this one
    Includes             []string                           `json:"includes,omitempty"`
    // Key name of the app category
    Key                  string                             `json:"key"`
    AdditionalProperties map[string]interface{}             `json:"_"`
}

// String implements the fmt.Stringer interface for ConstAppCategoryDefinition,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstAppCategoryDefinition) String() string {
    return fmt.Sprintf(
    	"ConstAppCategoryDefinition[Display=%v, Filters=%v, Includes=%v, Key=%v, AdditionalProperties=%v]",
    	c.Display, c.Filters, c.Includes, c.Key, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstAppCategoryDefinition.
// It customizes the JSON marshaling process for ConstAppCategoryDefinition objects.
func (c ConstAppCategoryDefinition) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "display", "filters", "includes", "key"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstAppCategoryDefinition object to a map representation for JSON marshaling.
func (c ConstAppCategoryDefinition) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["display"] = c.Display
    if c.Filters != nil {
        structMap["filters"] = c.Filters.toMap()
    }
    if c.Includes != nil {
        structMap["includes"] = c.Includes
    }
    structMap["key"] = c.Key
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstAppCategoryDefinition.
// It customizes the JSON unmarshaling process for ConstAppCategoryDefinition objects.
func (c *ConstAppCategoryDefinition) UnmarshalJSON(input []byte) error {
    var temp tempConstAppCategoryDefinition
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "filters", "includes", "key")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Display = *temp.Display
    c.Filters = temp.Filters
    c.Includes = temp.Includes
    c.Key = *temp.Key
    return nil
}

// tempConstAppCategoryDefinition is a temporary struct used for validating the fields of ConstAppCategoryDefinition.
type tempConstAppCategoryDefinition  struct {
    Display  *string                            `json:"display"`
    Filters  *ConstAppCategoryDefinitionFilters `json:"filters,omitempty"`
    Includes []string                           `json:"includes,omitempty"`
    Key      *string                            `json:"key"`
}

func (c *tempConstAppCategoryDefinition) validate() error {
    var errs []string
    if c.Display == nil {
        errs = append(errs, "required field `display` is missing for type `const_app_category_definition`")
    }
    if c.Key == nil {
        errs = append(errs, "required field `key` is missing for type `const_app_category_definition`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
