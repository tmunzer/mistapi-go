package models

import (
    "encoding/json"
    "fmt"
)

// ModuleStatItemTemperaturesItem represents a ModuleStatItemTemperaturesItem struct.
type ModuleStatItemTemperaturesItem struct {
    Celsius              *float64               `json:"celsius,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    Status               *string                `json:"status,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ModuleStatItemTemperaturesItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m ModuleStatItemTemperaturesItem) String() string {
    return fmt.Sprintf(
    	"ModuleStatItemTemperaturesItem[Celsius=%v, Name=%v, Status=%v, AdditionalProperties=%v]",
    	m.Celsius, m.Name, m.Status, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ModuleStatItemTemperaturesItem.
// It customizes the JSON marshaling process for ModuleStatItemTemperaturesItem objects.
func (m ModuleStatItemTemperaturesItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "celsius", "name", "status"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the ModuleStatItemTemperaturesItem object to a map representation for JSON marshaling.
func (m ModuleStatItemTemperaturesItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Celsius != nil {
        structMap["celsius"] = m.Celsius
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.Status != nil {
        structMap["status"] = m.Status
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ModuleStatItemTemperaturesItem.
// It customizes the JSON unmarshaling process for ModuleStatItemTemperaturesItem objects.
func (m *ModuleStatItemTemperaturesItem) UnmarshalJSON(input []byte) error {
    var temp tempModuleStatItemTemperaturesItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "celsius", "name", "status")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Celsius = temp.Celsius
    m.Name = temp.Name
    m.Status = temp.Status
    return nil
}

// tempModuleStatItemTemperaturesItem is a temporary struct used for validating the fields of ModuleStatItemTemperaturesItem.
type tempModuleStatItemTemperaturesItem  struct {
    Celsius *float64 `json:"celsius,omitempty"`
    Name    *string  `json:"name,omitempty"`
    Status  *string  `json:"status,omitempty"`
}
