package models

import (
    "encoding/json"
)

// ConstOtherDeviceModel represents a ConstOtherDeviceModel struct.
type ConstOtherDeviceModel struct {
    VendorModelId        *string                `json:"_vendor_model_id,omitempty"`
    Display              *string                `json:"display,omitempty"`
    Model                *string                `json:"model,omitempty"`
    Type                 *string                `json:"type,omitempty"`
    Vendor               *string                `json:"vendor,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstOtherDeviceModel.
// It customizes the JSON marshaling process for ConstOtherDeviceModel objects.
func (c ConstOtherDeviceModel) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "_vendor_model_id", "display", "model", "type", "vendor"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstOtherDeviceModel object to a map representation for JSON marshaling.
func (c ConstOtherDeviceModel) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.VendorModelId != nil {
        structMap["_vendor_model_id"] = c.VendorModelId
    }
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.Model != nil {
        structMap["model"] = c.Model
    }
    if c.Type != nil {
        structMap["type"] = c.Type
    }
    if c.Vendor != nil {
        structMap["vendor"] = c.Vendor
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstOtherDeviceModel.
// It customizes the JSON unmarshaling process for ConstOtherDeviceModel objects.
func (c *ConstOtherDeviceModel) UnmarshalJSON(input []byte) error {
    var temp tempConstOtherDeviceModel
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "_vendor_model_id", "display", "model", "type", "vendor")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.VendorModelId = temp.VendorModelId
    c.Display = temp.Display
    c.Model = temp.Model
    c.Type = temp.Type
    c.Vendor = temp.Vendor
    return nil
}

// tempConstOtherDeviceModel is a temporary struct used for validating the fields of ConstOtherDeviceModel.
type tempConstOtherDeviceModel  struct {
    VendorModelId *string `json:"_vendor_model_id,omitempty"`
    Display       *string `json:"display,omitempty"`
    Model         *string `json:"model,omitempty"`
    Type          *string `json:"type,omitempty"`
    Vendor        *string `json:"vendor,omitempty"`
}
