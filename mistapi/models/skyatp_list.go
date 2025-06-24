package models

import (
    "encoding/json"
    "fmt"
)

// SkyatpList represents a SkyatpList struct.
type SkyatpList struct {
    Domains              []SkyatpListDomain     `json:"domains,omitempty"`
    Ips                  []SkyatpListIp         `json:"ips,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SkyatpList,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SkyatpList) String() string {
    return fmt.Sprintf(
    	"SkyatpList[Domains=%v, Ips=%v, AdditionalProperties=%v]",
    	s.Domains, s.Ips, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SkyatpList.
// It customizes the JSON marshaling process for SkyatpList objects.
func (s SkyatpList) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "domains", "ips"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SkyatpList object to a map representation for JSON marshaling.
func (s SkyatpList) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Domains != nil {
        structMap["domains"] = s.Domains
    }
    if s.Ips != nil {
        structMap["ips"] = s.Ips
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SkyatpList.
// It customizes the JSON unmarshaling process for SkyatpList objects.
func (s *SkyatpList) UnmarshalJSON(input []byte) error {
    var temp tempSkyatpList
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "domains", "ips")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.Domains = temp.Domains
    s.Ips = temp.Ips
    return nil
}

// tempSkyatpList is a temporary struct used for validating the fields of SkyatpList.
type tempSkyatpList  struct {
    Domains []SkyatpListDomain `json:"domains,omitempty"`
    Ips     []SkyatpListIp     `json:"ips,omitempty"`
}
