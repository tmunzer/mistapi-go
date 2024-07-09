package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// ResponseInventoryInventoryDuplicatedItems represents a ResponseInventoryInventoryDuplicatedItems struct.
type ResponseInventoryInventoryDuplicatedItems struct {
    Mac                  string         `json:"mac"`
    Magic                string         `json:"magic"`
    Model                string         `json:"model"`
    Serial               string         `json:"serial"`
    Type                 string         `json:"type"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ResponseInventoryInventoryDuplicatedItems.
// It customizes the JSON marshaling process for ResponseInventoryInventoryDuplicatedItems objects.
func (r ResponseInventoryInventoryDuplicatedItems) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(r.toMap())
}

// toMap converts the ResponseInventoryInventoryDuplicatedItems object to a map representation for JSON marshaling.
func (r ResponseInventoryInventoryDuplicatedItems) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, r.AdditionalProperties)
    structMap["mac"] = r.Mac
    structMap["magic"] = r.Magic
    structMap["model"] = r.Model
    structMap["serial"] = r.Serial
    structMap["type"] = r.Type
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ResponseInventoryInventoryDuplicatedItems.
// It customizes the JSON unmarshaling process for ResponseInventoryInventoryDuplicatedItems objects.
func (r *ResponseInventoryInventoryDuplicatedItems) UnmarshalJSON(input []byte) error {
    var temp responseInventoryInventoryDuplicatedItems
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

// responseInventoryInventoryDuplicatedItems is a temporary struct used for validating the fields of ResponseInventoryInventoryDuplicatedItems.
type responseInventoryInventoryDuplicatedItems  struct {
    Mac    *string `json:"mac"`
    Magic  *string `json:"magic"`
    Model  *string `json:"model"`
    Serial *string `json:"serial"`
    Type   *string `json:"type"`
}

func (r *responseInventoryInventoryDuplicatedItems) validate() error {
    var errs []string
    if r.Mac == nil {
        errs = append(errs, "required field `mac` is missing for type `Response_Inventory_Inventory_Duplicated_Items`")
    }
    if r.Magic == nil {
        errs = append(errs, "required field `magic` is missing for type `Response_Inventory_Inventory_Duplicated_Items`")
    }
    if r.Model == nil {
        errs = append(errs, "required field `model` is missing for type `Response_Inventory_Inventory_Duplicated_Items`")
    }
    if r.Serial == nil {
        errs = append(errs, "required field `serial` is missing for type `Response_Inventory_Inventory_Duplicated_Items`")
    }
    if r.Type == nil {
        errs = append(errs, "required field `type` is missing for type `Response_Inventory_Inventory_Duplicated_Items`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
