package models

import (
    "encoding/json"
    "fmt"
)

// JsInventoryItem represents a JsInventoryItem struct.
type JsInventoryItem struct {
    // End of life time
    EolTime              *int                   `json:"eol_time,omitempty"`
    // End of support time
    EosTime              *int                   `json:"eos_time,omitempty"`
    // Model of the install base inventory
    Model                *string                `json:"model,omitempty"`
    // Serial Number of the inventory
    Serial               *string                `json:"serial,omitempty"`
    // Serviceable device stock
    Sku                  *string                `json:"sku,omitempty"`
    // enum: `ap`, `gateway`, `switch`
    Type                 *DeviceTypeEnum        `json:"type,omitempty"`
    WarrantyType         *string                `json:"warranty_type,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for JsInventoryItem,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (j JsInventoryItem) String() string {
    return fmt.Sprintf(
    	"JsInventoryItem[EolTime=%v, EosTime=%v, Model=%v, Serial=%v, Sku=%v, Type=%v, WarrantyType=%v, AdditionalProperties=%v]",
    	j.EolTime, j.EosTime, j.Model, j.Serial, j.Sku, j.Type, j.WarrantyType, j.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for JsInventoryItem.
// It customizes the JSON marshaling process for JsInventoryItem objects.
func (j JsInventoryItem) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(j.AdditionalProperties,
        "eol_time", "eos_time", "model", "serial", "sku", "type", "warranty_type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(j.toMap())
}

// toMap converts the JsInventoryItem object to a map representation for JSON marshaling.
func (j JsInventoryItem) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, j.AdditionalProperties)
    if j.EolTime != nil {
        structMap["eol_time"] = j.EolTime
    }
    if j.EosTime != nil {
        structMap["eos_time"] = j.EosTime
    }
    if j.Model != nil {
        structMap["model"] = j.Model
    }
    if j.Serial != nil {
        structMap["serial"] = j.Serial
    }
    if j.Sku != nil {
        structMap["sku"] = j.Sku
    }
    if j.Type != nil {
        structMap["type"] = j.Type
    }
    if j.WarrantyType != nil {
        structMap["warranty_type"] = j.WarrantyType
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for JsInventoryItem.
// It customizes the JSON unmarshaling process for JsInventoryItem objects.
func (j *JsInventoryItem) UnmarshalJSON(input []byte) error {
    var temp tempJsInventoryItem
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "eol_time", "eos_time", "model", "serial", "sku", "type", "warranty_type")
    if err != nil {
    	return err
    }
    j.AdditionalProperties = additionalProperties
    
    j.EolTime = temp.EolTime
    j.EosTime = temp.EosTime
    j.Model = temp.Model
    j.Serial = temp.Serial
    j.Sku = temp.Sku
    j.Type = temp.Type
    j.WarrantyType = temp.WarrantyType
    return nil
}

// tempJsInventoryItem is a temporary struct used for validating the fields of JsInventoryItem.
type tempJsInventoryItem  struct {
    EolTime      *int            `json:"eol_time,omitempty"`
    EosTime      *int            `json:"eos_time,omitempty"`
    Model        *string         `json:"model,omitempty"`
    Serial       *string         `json:"serial,omitempty"`
    Sku          *string         `json:"sku,omitempty"`
    Type         *DeviceTypeEnum `json:"type,omitempty"`
    WarrantyType *string         `json:"warranty_type,omitempty"`
}
