package models

import (
    "encoding/json"
    "fmt"
)

// SynthetictestConfigLanNetwork represents a SynthetictestConfigLanNetwork struct.
// configure minis probes to be tested on lan networks of gateways
type SynthetictestConfigLanNetwork struct {
    // List of networks to be used for synthetic tests
    Networks             []string               `json:"networks,omitempty"`
    // app name comes from `custom_probes` above or /const/synthetic_test_probes
    Probes               []string               `json:"probes,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SynthetictestConfigLanNetwork,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SynthetictestConfigLanNetwork) String() string {
    return fmt.Sprintf(
    	"SynthetictestConfigLanNetwork[Networks=%v, Probes=%v, AdditionalProperties=%v]",
    	s.Networks, s.Probes, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SynthetictestConfigLanNetwork.
// It customizes the JSON marshaling process for SynthetictestConfigLanNetwork objects.
func (s SynthetictestConfigLanNetwork) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "networks", "probes"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SynthetictestConfigLanNetwork object to a map representation for JSON marshaling.
func (s SynthetictestConfigLanNetwork) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Networks != nil {
        structMap["networks"] = s.Networks
    }
    if s.Probes != nil {
        structMap["probes"] = s.Probes
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SynthetictestConfigLanNetwork.
// It customizes the JSON unmarshaling process for SynthetictestConfigLanNetwork objects.
func (s *SynthetictestConfigLanNetwork) UnmarshalJSON(input []byte) error {
    var temp tempSynthetictestConfigLanNetwork
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "networks", "probes")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Networks = temp.Networks
    s.Probes = temp.Probes
    return nil
}

// tempSynthetictestConfigLanNetwork is a temporary struct used for validating the fields of SynthetictestConfigLanNetwork.
type tempSynthetictestConfigLanNetwork  struct {
    Networks []string `json:"networks,omitempty"`
    Probes   []string `json:"probes,omitempty"`
}
