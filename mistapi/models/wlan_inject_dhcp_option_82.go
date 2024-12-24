package models

import (
    "encoding/json"
    "fmt"
)

// WlanInjectDhcpOption82 represents a WlanInjectDhcpOption82 struct.
type WlanInjectDhcpOption82 struct {
    // information to set in the `circuit_id` field of the DHCP Option 82. It is possible to use static string or the following variables (e.g. `{{SSID}}:{{AP_MAC}}`):
    // * {{AP_MAC}}
    // * {{AP_MAC_DASHED}}
    // * {{AP_MODEL}}
    // * {{AP_NAME}}
    // * {{SITE_NAME}}
    // * {{SSID}}
    CircuitId            *string                `json:"circuit_id,omitempty"`
    // whether to inject option 82 when forwarding DHCP packets
    Enabled              *bool                  `json:"enabled,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for WlanInjectDhcpOption82,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (w WlanInjectDhcpOption82) String() string {
    return fmt.Sprintf(
    	"WlanInjectDhcpOption82[CircuitId=%v, Enabled=%v, AdditionalProperties=%v]",
    	w.CircuitId, w.Enabled, w.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for WlanInjectDhcpOption82.
// It customizes the JSON marshaling process for WlanInjectDhcpOption82 objects.
func (w WlanInjectDhcpOption82) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(w.AdditionalProperties,
        "circuit_id", "enabled"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(w.toMap())
}

// toMap converts the WlanInjectDhcpOption82 object to a map representation for JSON marshaling.
func (w WlanInjectDhcpOption82) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, w.AdditionalProperties)
    if w.CircuitId != nil {
        structMap["circuit_id"] = w.CircuitId
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanInjectDhcpOption82.
// It customizes the JSON unmarshaling process for WlanInjectDhcpOption82 objects.
func (w *WlanInjectDhcpOption82) UnmarshalJSON(input []byte) error {
    var temp tempWlanInjectDhcpOption82
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "circuit_id", "enabled")
    if err != nil {
    	return err
    }
    w.AdditionalProperties = additionalProperties
    
    w.CircuitId = temp.CircuitId
    w.Enabled = temp.Enabled
    return nil
}

// tempWlanInjectDhcpOption82 is a temporary struct used for validating the fields of WlanInjectDhcpOption82.
type tempWlanInjectDhcpOption82  struct {
    CircuitId *string `json:"circuit_id,omitempty"`
    Enabled   *bool   `json:"enabled,omitempty"`
}
