package models

import (
    "encoding/json"
    "fmt"
)

// SiteSettingTuntermMulticastConfig represents a SiteSettingTuntermMulticastConfig struct.
type SiteSettingTuntermMulticastConfig struct {
    Mdns                 *SiteSettingTuntermMulticastConfigMdns `json:"mdns,omitempty"`
    MulticastAll         *bool                                  `json:"multicast_all,omitempty"`
    Ssdp                 *SiteSettingTuntermMulticastConfigSsdp `json:"ssdp,omitempty"`
    AdditionalProperties map[string]interface{}                 `json:"_"`
}

// String implements the fmt.Stringer interface for SiteSettingTuntermMulticastConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SiteSettingTuntermMulticastConfig) String() string {
    return fmt.Sprintf(
    	"SiteSettingTuntermMulticastConfig[Mdns=%v, MulticastAll=%v, Ssdp=%v, AdditionalProperties=%v]",
    	s.Mdns, s.MulticastAll, s.Ssdp, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SiteSettingTuntermMulticastConfig.
// It customizes the JSON marshaling process for SiteSettingTuntermMulticastConfig objects.
func (s SiteSettingTuntermMulticastConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "mdns", "multicast_all", "ssdp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SiteSettingTuntermMulticastConfig object to a map representation for JSON marshaling.
func (s SiteSettingTuntermMulticastConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Mdns != nil {
        structMap["mdns"] = s.Mdns.toMap()
    }
    if s.MulticastAll != nil {
        structMap["multicast_all"] = s.MulticastAll
    }
    if s.Ssdp != nil {
        structMap["ssdp"] = s.Ssdp.toMap()
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SiteSettingTuntermMulticastConfig.
// It customizes the JSON unmarshaling process for SiteSettingTuntermMulticastConfig objects.
func (s *SiteSettingTuntermMulticastConfig) UnmarshalJSON(input []byte) error {
    var temp tempSiteSettingTuntermMulticastConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "mdns", "multicast_all", "ssdp")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Mdns = temp.Mdns
    s.MulticastAll = temp.MulticastAll
    s.Ssdp = temp.Ssdp
    return nil
}

// tempSiteSettingTuntermMulticastConfig is a temporary struct used for validating the fields of SiteSettingTuntermMulticastConfig.
type tempSiteSettingTuntermMulticastConfig  struct {
    Mdns         *SiteSettingTuntermMulticastConfigMdns `json:"mdns,omitempty"`
    MulticastAll *bool                                  `json:"multicast_all,omitempty"`
    Ssdp         *SiteSettingTuntermMulticastConfigSsdp `json:"ssdp,omitempty"`
}
