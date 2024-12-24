package models

import (
    "encoding/json"
    "errors"
    "fmt"
    "strings"
)

// ResponseInventoryInventoryAddedItems represents a ResponseInventoryInventoryAddedItems struct.
type ResponseInventoryInventoryAddedItems struct {
    Mac                  string                 `json:"mac"`
    Magic                string                 `json:"magic"`
    Model                string                 `json:"model"`
    Serial               string                 `json:"serial"`
    Type                 string                 `json:"type"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ResponseInventoryInventoryAddedItems,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ResponseInventoryInventoryAddedItems) String() string {
    return fmt.Sprintf(
    	"ResponseInventoryInventoryAddedItems[Mac=%v, Magic=%v, Model=%v, Serial=%v, Type=%v, AdditionalProperties=%v]",
    	r.Mac, r.Magic, r.Model, r.Serial, r.Type, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ResponseInventoryInventoryAddedItems.
// It customizes the JSON marshaling process for ResponseInventoryInventoryAddedItems objects.
func (r ResponseInventoryInventoryAddedItems) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "mac", "magic", "model", "serial", "type"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInventoryInventoryAddedItems object to a map representation for JSON marshaling.
func (r ResponseInventoryInventoryAddedItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["mac"] = r.Mac
    structMap["magic"] = r.Magic
    structMap["model"] = r.Model
    structMap["serial"] = r.Serial
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInventoryInventoryAddedItems.
// It customizes the JSON unmarshaling process for ResponseInventoryInventoryAddedItems objects.
func (r *ResponseInventoryInventoryAddedItems) UnmarshalJSON(input []byte) error {
    var temp tempResponseInventoryInventoryAddedItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mac", "magic", "model", "serial", "type")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Mac = *temp.Mac
    r.Magic = *temp.Magic
    r.Model = *temp.Model
    r.Serial = *temp.Serial
    r.Type = *temp.Type
    return nil
}

// tempResponseInventoryInventoryAddedItems is a temporary struct used for validating the fields of ResponseInventoryInventoryAddedItems.
type tempResponseInventoryInventoryAddedItems  struct {
    Mac    *string `json:"mac"`
    Magic  *string `json:"magic"`
    Model  *string `json:"model"`
    Serial *string `json:"serial"`
    Type   *string `json:"type"`
}

func (r *tempResponseInventoryInventoryAddedItems) validate() error {
    var errs []string
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `response_inventory_inventory_added_items`")
    }
    if r.Magic == nil {
        errs = append(errs, "required field `magic` is missing for type `response_inventory_inventory_added_items`")
    }
    if r.Model == nil {
        errs = append(errs, "required field `model` is missing for type `response_inventory_inventory_added_items`")
    }
    if r.Serial == nil {
        errs = append(errs, "required field `serial` is missing for type `response_inventory_inventory_added_items`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `response_inventory_inventory_added_items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join (errs, "\n"))
}
