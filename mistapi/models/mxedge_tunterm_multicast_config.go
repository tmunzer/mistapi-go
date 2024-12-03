package models

import (
    "encoding/json"
)

// MxedgeTuntermMulticastConfig represents a MxedgeTuntermMulticastConfig struct.
type MxedgeTuntermMulticastConfig struct {
    Mdns                 *MxedgeTuntermMulticastMdns `json:"mdns,omitempty"`
    Ssdp                 *MxedgeTuntermMulticastSsdp `json:"ssdp,omitempty"`
    AdditionalProperties map[string]interface{}      `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MxedgeTuntermMulticastConfig.
// It customizes the JSON marshaling process for MxedgeTuntermMulticastConfig objects.
func (m MxedgeTuntermMulticastConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "mdns", "ssdp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MxedgeTuntermMulticastConfig object to a map representation for JSON marshaling.
func (m MxedgeTuntermMulticastConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Mdns != nil {
        structMap["mdns"] = m.Mdns.toMap()
    }
    if m.Ssdp != nil {
        structMap["ssdp"] = m.Ssdp.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MxedgeTuntermMulticastConfig.
// It customizes the JSON unmarshaling process for MxedgeTuntermMulticastConfig objects.
func (m *MxedgeTuntermMulticastConfig) UnmarshalJSON(input []byte) error {
    var temp tempMxedgeTuntermMulticastConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mdns", "ssdp")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Mdns = temp.Mdns
    m.Ssdp = temp.Ssdp
    return nil
}

// tempMxedgeTuntermMulticastConfig is a temporary struct used for validating the fields of MxedgeTuntermMulticastConfig.
type tempMxedgeTuntermMulticastConfig  struct {
    Mdns *MxedgeTuntermMulticastMdns `json:"mdns,omitempty"`
    Ssdp *MxedgeTuntermMulticastSsdp `json:"ssdp,omitempty"`
}
