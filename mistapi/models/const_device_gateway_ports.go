package models

import (
    "encoding/json"
)

// ConstDeviceGatewayPorts represents a ConstDeviceGatewayPorts struct.
// Object Key is the interface name (e.g. "ge-0/0/1", ...)
type ConstDeviceGatewayPorts struct {
    Display              *string                `json:"display,omitempty"`
    PciAddress           *string                `json:"pci_address,omitempty"`
    Speed                *int                   `json:"speed,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ConstDeviceGatewayPorts.
// It customizes the JSON marshaling process for ConstDeviceGatewayPorts objects.
func (c ConstDeviceGatewayPorts) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(c.AdditionalProperties,
        "display", "pci_address", "speed"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(c.toMap())
}

// toMap converts the ConstDeviceGatewayPorts object to a map representation for JSON marshaling.
func (c ConstDeviceGatewayPorts) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, c.AdditionalProperties)
    if c.Display != nil {
        structMap["display"] = c.Display
    }
    if c.PciAddress != nil {
        structMap["pci_address"] = c.PciAddress
    }
    if c.Speed != nil {
        structMap["speed"] = c.Speed
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ConstDeviceGatewayPorts.
// It customizes the JSON unmarshaling process for ConstDeviceGatewayPorts objects.
func (c *ConstDeviceGatewayPorts) UnmarshalJSON(input []byte) error {
    var temp tempConstDeviceGatewayPorts
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "display", "pci_address", "speed")
    if err != nil {
    	return err
    }
    c.AdditionalProperties = additionalProperties
    
    c.Display = temp.Display
    c.PciAddress = temp.PciAddress
    c.Speed = temp.Speed
    return nil
}

// tempConstDeviceGatewayPorts is a temporary struct used for validating the fields of ConstDeviceGatewayPorts.
type tempConstDeviceGatewayPorts  struct {
    Display    *string `json:"display,omitempty"`
    PciAddress *string `json:"pci_address,omitempty"`
    Speed      *int    `json:"speed,omitempty"`
}
