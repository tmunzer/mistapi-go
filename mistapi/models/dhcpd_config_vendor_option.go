package models

import (
    "encoding/json"
    "fmt"
)

// DhcpdConfigVendorOption represents a DhcpdConfigVendorOption struct.
type DhcpdConfigVendorOption struct {
    // enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32`
    Type                 *DhcpdConfigVendorOptionTypeEnum `json:"type,omitempty"`
    Value                *string                          `json:"value,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for DhcpdConfigVendorOption,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (d DhcpdConfigVendorOption) String() string {
    return fmt.Sprintf(
    	"DhcpdConfigVendorOption[Type=%v, Value=%v, AdditionalProperties=%v]",
    	d.Type, d.Value, d.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for DhcpdConfigVendorOption.
// It customizes the JSON marshaling process for DhcpdConfigVendorOption objects.
func (d DhcpdConfigVendorOption) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(d.AdditionalProperties,
        "type", "value"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(d.toMap())
}

// toMap converts the DhcpdConfigVendorOption object to a map representation for JSON marshaling.
func (d DhcpdConfigVendorOption) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, d.AdditionalProperties)
    if d.Type != nil {
        structMap["type"] = d.Type
    }
    if d.Value != nil {
        structMap["value"] = d.Value
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for DhcpdConfigVendorOption.
// It customizes the JSON unmarshaling process for DhcpdConfigVendorOption objects.
func (d *DhcpdConfigVendorOption) UnmarshalJSON(input []byte) error {
    var temp tempDhcpdConfigVendorOption
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "type", "value")
    if err != nil {
    	return err
    }
    d.AdditionalProperties = additionalProperties
    
    d.Type = temp.Type
    d.Value = temp.Value
    return nil
}

// tempDhcpdConfigVendorOption is a temporary struct used for validating the fields of DhcpdConfigVendorOption.
type tempDhcpdConfigVendorOption  struct {
    Type  *DhcpdConfigVendorOptionTypeEnum `json:"type,omitempty"`
    Value *string                          `json:"value,omitempty"`
}
