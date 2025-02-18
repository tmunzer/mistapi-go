package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ConstAppSubcategoryDefinition represents a ConstAppSubcategoryDefinition struct.
type ConstAppSubcategoryDefinition struct {
    // Description of the app subcategory
    Display              string                 `json:"display"`
    // Key name of the app subcategory
    Key                  string                 `json:"key"`
    // Type of traffic (QoS) of the app subcategory
    TrafficType          string                 `json:"traffic_type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ConstAppSubcategoryDefinition,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (c ConstAppSubcategoryDefinition) String() string {
    return fmt.Sprintf(
    	"ConstAppSubcategoryDefinition[Display=%v, Key=%v, TrafficType=%v, AdditionalProperties=%v]",
    	c.Display, c.Key, c.TrafficType, c.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ConstAppSubcategoryDefinition.
// It customizes the JSON marshaling process for ConstAppSubcategoryDefinition objects.
func (c ConstAppSubcategoryDefinition) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "display", "key", "traffic_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstAppSubcategoryDefinition object to a map representation for JSON marshaling.
func (c ConstAppSubcategoryDefinition) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["display"] = c.Display
    structMap["key"] = c.Key
    structMap["traffic_type"] = c.TrafficType
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstAppSubcategoryDefinition.
// It customizes the JSON unmarshaling process for ConstAppSubcategoryDefinition objects.
func (c *ConstAppSubcategoryDefinition) UnmarshalJSON(input []byte) error {
    var temp tempConstAppSubcategoryDefinition
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "key", "traffic_type")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Display = *temp.Display
    c.Key = *temp.Key
    c.TrafficType = *temp.TrafficType
    return nil
}

// tempConstAppSubcategoryDefinition is a temporary struct used for validating the fields of ConstAppSubcategoryDefinition.
type tempConstAppSubcategoryDefinition  struct {
    Display     *string `json:"display"`
    Key         *string `json:"key"`
    TrafficType *string `json:"traffic_type"`
}

func (c *tempConstAppSubcategoryDefinition) validate() error {
    var errs []string
    if c.Display == nil {
        errs = append(errs, "required field `display` is missing for type `const_app_subcategory_definition`")
    }
    if c.Key == nil {
        errs = append(errs, "required field `key` is missing for type `const_app_subcategory_definition`")
    }
    if c.TrafficType == nil {
        errs = append(errs, "required field `traffic_type` is missing for type `const_app_subcategory_definition`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
