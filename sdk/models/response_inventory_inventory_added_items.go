package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseInventoryInventoryAddedItems represents a ResponseInventoryInventoryAddedItems struct.
type ResponseInventoryInventoryAddedItems struct {
    Mac                  string         `json:"mac"`
    Magic                string         `json:"magic"`
    Model                string         `json:"model"`
    Serial               string         `json:"serial"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseInventoryInventoryAddedItems.
// It customizes the JSON marshaling process for ResponseInventoryInventoryAddedItems objects.
func (r ResponseInventoryInventoryAddedItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInventoryInventoryAddedItems object to a map representation for JSON marshaling.
func (r ResponseInventoryInventoryAddedItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
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
    var temp responseInventoryInventoryAddedItems
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "mac", "magic", "model", "serial", "type")
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

// responseInventoryInventoryAddedItems is a temporary struct used for validating the fields of ResponseInventoryInventoryAddedItems.
type responseInventoryInventoryAddedItems  struct {
    Mac    *string `json:"mac"`
    Magic  *string `json:"magic"`
    Model  *string `json:"model"`
    Serial *string `json:"serial"`
    Type   *string `json:"type"`
}

func (r *responseInventoryInventoryAddedItems) validate() error {
    var errs []string
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Response_Inventory_Inventory_Added_Items`")
    }
    if r.Magic == nil {
        errs = append(errs, "required field `magic` is missing for type `Response_Inventory_Inventory_Added_Items`")
    }
    if r.Model == nil {
        errs = append(errs, "required field `model` is missing for type `Response_Inventory_Inventory_Added_Items`")
    }
    if r.Serial == nil {
        errs = append(errs, "required field `serial` is missing for type `Response_Inventory_Inventory_Added_Items`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Response_Inventory_Inventory_Added_Items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
