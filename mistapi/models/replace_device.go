package models

import (
    "encoding/json"
    "fmt"
)

// ReplaceDevice represents a ReplaceDevice struct.
type ReplaceDevice struct {
    // Attributes that you don’t want to copy
    Discard              []string               `json:"discard,omitempty"`
    // MAC Address of the inventory that will be replacing the old one. It has to be claimed and unassigned
    InventoryMac         *string                `json:"inventory_mac,omitempty"`
    // MAC Address of the device to replace
    Mac                  *string                `json:"mac,omitempty"`
    // Site_id of the device to be replaced
    SiteId               *string                `json:"site_id,omitempty"`
    // Ethernet port configurations
    TuntermPortConfig    *TuntermPortConfig     `json:"tunterm_port_config,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for ReplaceDevice,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (r ReplaceDevice) String() string {
    return fmt.Sprintf(
    	"ReplaceDevice[Discard=%v, InventoryMac=%v, Mac=%v, SiteId=%v, TuntermPortConfig=%v, AdditionalProperties=%v]",
    	r.Discard, r.InventoryMac, r.Mac, r.SiteId, r.TuntermPortConfig, r.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for ReplaceDevice.
// It customizes the JSON marshaling process for ReplaceDevice objects.
func (r ReplaceDevice) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(r.AdditionalProperties,
        "discard", "inventory_mac", "mac", "site_id", "tunterm_port_config"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(r.toMap())
}

// toMap converts the ReplaceDevice object to a map representation for JSON marshaling.
func (r ReplaceDevice) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, r.AdditionalProperties)
    if r.Discard != nil {
        structMap["discard"] = r.Discard
    }
    if r.InventoryMac != nil {
        structMap["inventory_mac"] = r.InventoryMac
    }
    if r.Mac != nil {
        structMap["mac"] = r.Mac
    }
    if r.SiteId != nil {
        structMap["site_id"] = r.SiteId
    }
    if r.TuntermPortConfig != nil {
        structMap["tunterm_port_config"] = r.TuntermPortConfig.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ReplaceDevice.
// It customizes the JSON unmarshaling process for ReplaceDevice objects.
func (r *ReplaceDevice) UnmarshalJSON(input []byte) error {
    var temp tempReplaceDevice
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "discard", "inventory_mac", "mac", "site_id", "tunterm_port_config")
    if err != nil {
    	return err
    }
    r.AdditionalProperties = additionalProperties
    
    r.Discard = temp.Discard
    r.InventoryMac = temp.InventoryMac
    r.Mac = temp.Mac
    r.SiteId = temp.SiteId
    r.TuntermPortConfig = temp.TuntermPortConfig
    return nil
}

// tempReplaceDevice is a temporary struct used for validating the fields of ReplaceDevice.
type tempReplaceDevice  struct {
    Discard           []string           `json:"discard,omitempty"`
    InventoryMac      *string            `json:"inventory_mac,omitempty"`
    Mac               *string            `json:"mac,omitempty"`
    SiteId            *string            `json:"site_id,omitempty"`
    TuntermPortConfig *TuntermPortConfig `json:"tunterm_port_config,omitempty"`
}
