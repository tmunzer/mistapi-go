package models

import (
    "encoding/json"
    "fmt"
)

// SynthetictestConfigVlan represents a SynthetictestConfigVlan struct.
type SynthetictestConfigVlan struct {
    CustomTestUrls       []string               `json:"custom_test_urls,omitempty"` // Deprecated
    // For some vlans where we don't want this to run
    Disabled             *bool                  `json:"disabled,omitempty"`
    // app name comes from `custom_probes` above or /const/synthetic_test_probes
    Probes               []string               `json:"probes,omitempty"`
    VlanIds              []VlanIdWithVariable   `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestConfigVlan,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestConfigVlan) String() string {
    return fmt.Sprintf(
    	"SynthetictestConfigVlan[CustomTestUrls=%v, Disabled=%v, Probes=%v, VlanIds=%v, AdditionalProperties=%v]",
    	s.CustomTestUrls, s.Disabled, s.Probes, s.VlanIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfigVlan.
// It customizes the JSON marshaling process for SynthetictestConfigVlan objects.
func (s SynthetictestConfigVlan) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "custom_test_urls", "disabled", "probes", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfigVlan object to a map representation for JSON marshaling.
func (s SynthetictestConfigVlan) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CustomTestUrls != nil {
        structMap["custom_test_urls"] = s.CustomTestUrls
    }
    if s.Disabled != nil {
        structMap["disabled"] = s.Disabled
    }
    if s.Probes != nil {
        structMap["probes"] = s.Probes
    }
    if s.VlanIds != nil {
        structMap["vlan_ids"] = s.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfigVlan.
// It customizes the JSON unmarshaling process for SynthetictestConfigVlan objects.
func (s *SynthetictestConfigVlan) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestConfigVlan
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "custom_test_urls", "disabled", "probes", "vlan_ids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CustomTestUrls = temp.CustomTestUrls
    s.Disabled = temp.Disabled
    s.Probes = temp.Probes
    s.VlanIds = temp.VlanIds
    return nil
}

// tempSynthetictestConfigVlan is a temporary struct used for validating the fields of SynthetictestConfigVlan.
type tempSynthetictestConfigVlan  struct {
    CustomTestUrls []string             `json:"custom_test_urls,omitempty"`
    Disabled       *bool                `json:"disabled,omitempty"`
    Probes         []string             `json:"probes,omitempty"`
    VlanIds        []VlanIdWithVariable `json:"vlan_ids,omitempty"`
}
