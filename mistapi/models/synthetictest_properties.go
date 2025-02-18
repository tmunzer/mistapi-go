package models

import (
    "encoding/json"
    "fmt"
)

// SynthetictestProperties represents a SynthetictestProperties struct.
type SynthetictestProperties struct {
    CustomTestUrls       []string               `json:"custom_test_urls,omitempty"`
    // For some vlans where we don't want this to run
    Disabled             *bool                  `json:"disabled,omitempty"`
    VlanIds              []VlanIdWithVariable   `json:"vlan_ids,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestProperties,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestProperties) String() string {
    return fmt.Sprintf(
    	"SynthetictestProperties[CustomTestUrls=%v, Disabled=%v, VlanIds=%v, AdditionalProperties=%v]",
    	s.CustomTestUrls, s.Disabled, s.VlanIds, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestProperties.
// It customizes the JSON marshaling process for SynthetictestProperties objects.
func (s SynthetictestProperties) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "custom_test_urls", "disabled", "vlan_ids"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestProperties object to a map representation for JSON marshaling.
func (s SynthetictestProperties) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.CustomTestUrls != nil {
        structMap["custom_test_urls"] = s.CustomTestUrls
    }
    if s.Disabled != nil {
        structMap["disabled"] = s.Disabled
    }
    if s.VlanIds != nil {
        structMap["vlan_ids"] = s.VlanIds
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestProperties.
// It customizes the JSON unmarshaling process for SynthetictestProperties objects.
func (s *SynthetictestProperties) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestProperties
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "custom_test_urls", "disabled", "vlan_ids")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.CustomTestUrls = temp.CustomTestUrls
    s.Disabled = temp.Disabled
    s.VlanIds = temp.VlanIds
    return nil
}

// tempSynthetictestProperties is a temporary struct used for validating the fields of SynthetictestProperties.
type tempSynthetictestProperties  struct {
    CustomTestUrls []string             `json:"custom_test_urls,omitempty"`
    Disabled       *bool                `json:"disabled,omitempty"`
    VlanIds        []VlanIdWithVariable `json:"vlan_ids,omitempty"`
}
