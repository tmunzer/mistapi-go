package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ConstAppSubcategoryDefinition represents a ConstAppSubcategoryDefinition struct.
type ConstAppSubcategoryDefinition struct {
    // Description of the app subcategory
    Display              string         `json:"display"`
    // Key name of the app subcategory
    Key                  string         `json:"key"`
    // Type of traffic (QoS) of the app subcategory
    TrafficType          string         `json:"traffic_type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstAppSubcategoryDefinition.
// It customizes the JSON marshaling process for ConstAppSubcategoryDefinition objects.
func (c ConstAppSubcategoryDefinition) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(c.toMap())
}

// toMap converts the ConstAppSubcategoryDefinition object to a map representation for JSON marshaling.
func (c ConstAppSubcategoryDefinition) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, c.AdditionalProperties)
    structMap["display"] = c.Display
    structMap["key"] = c.Key
    structMap["traffic_type"] = c.TrafficType
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstAppSubcategoryDefinition.
// It customizes the JSON unmarshaling process for ConstAppSubcategoryDefinition objects.
func (c *ConstAppSubcategoryDefinition) UnmarshalJSON(input []byte) error {
    var temp constAppSubcategoryDefinition
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "display", "key", "traffic_type")
    if err != nil {
    	return err
    }
    
    c.AdditionalProperties = additionalProperties
    c.Display = *temp.Display
    c.Key = *temp.Key
    c.TrafficType = *temp.TrafficType
    return nil
}

// constAppSubcategoryDefinition is a temporary struct used for validating the fields of ConstAppSubcategoryDefinition.
type constAppSubcategoryDefinition  struct {
    Display     *string `json:"display"`
    Key         *string `json:"key"`
    TrafficType *string `json:"traffic_type"`
}

func (c *constAppSubcategoryDefinition) validate() error {
    var errs []string
    if c.Display == nil {
        errs = append(errs, "required field `display` is missing for type `Const_App_Subcategory_Definition`")
    }
    if c.Key == nil {
        errs = append(errs, "required field `key` is missing for type `Const_App_Subcategory_Definition`")
    }
    if c.TrafficType == nil {
        errs = append(errs, "required field `traffic_type` is missing for type `Const_App_Subcategory_Definition`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
