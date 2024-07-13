package models

import (
    "encoding/json"
)

// ServiceStat represents a ServiceStat struct.
type ServiceStat struct {
    AshVersion           *string        `json:"ash_version,omitempty"`
    CiaVersion           *string        `json:"cia_version,omitempty"`
    EmberVersion         *string        `json:"ember_version,omitempty"`
    IpsecClientVersion   *string        `json:"ipsec_client_version,omitempty"`
    MistAgentVersion     *string        `json:"mist_agent_version,omitempty"`
    PackageVersion       *string        `json:"package_version,omitempty"`
    TestingToolsVersion  *string        `json:"testing_tools_version,omitempty"`
    WheeljackVersion     *string        `json:"wheeljack_version,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for ServiceStat.
// It customizes the JSON marshaling process for ServiceStat objects.
func (s ServiceStat) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the ServiceStat object to a map representation for JSON marshaling.
func (s ServiceStat) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AshVersion != nil {
        structMap["ash_version"] = s.AshVersion
    }
    if s.CiaVersion != nil {
        structMap["cia_version"] = s.CiaVersion
    }
    if s.EmberVersion != nil {
        structMap["ember_version"] = s.EmberVersion
    }
    if s.IpsecClientVersion != nil {
        structMap["ipsec_client_version"] = s.IpsecClientVersion
    }
    if s.MistAgentVersion != nil {
        structMap["mist_agent_version"] = s.MistAgentVersion
    }
    if s.PackageVersion != nil {
        structMap["package_version"] = s.PackageVersion
    }
    if s.TestingToolsVersion != nil {
        structMap["testing_tools_version"] = s.TestingToolsVersion
    }
    if s.WheeljackVersion != nil {
        structMap["wheeljack_version"] = s.WheeljackVersion
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for ServiceStat.
// It customizes the JSON unmarshaling process for ServiceStat objects.
func (s *ServiceStat) UnmarshalJSON(input []byte) error {
    var temp serviceStat
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "ash_version", "cia_version", "ember_version", "ipsec_client_version", "mist_agent_version", "package_version", "testing_tools_version", "wheeljack_version")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.AshVersion = temp.AshVersion
    s.CiaVersion = temp.CiaVersion
    s.EmberVersion = temp.EmberVersion
    s.IpsecClientVersion = temp.IpsecClientVersion
    s.MistAgentVersion = temp.MistAgentVersion
    s.PackageVersion = temp.PackageVersion
    s.TestingToolsVersion = temp.TestingToolsVersion
    s.WheeljackVersion = temp.WheeljackVersion
    return nil
}

// serviceStat is a temporary struct used for validating the fields of ServiceStat.
type serviceStat  struct {
    AshVersion          *string `json:"ash_version,omitempty"`
    CiaVersion          *string `json:"cia_version,omitempty"`
    EmberVersion        *string `json:"ember_version,omitempty"`
    IpsecClientVersion  *string `json:"ipsec_client_version,omitempty"`
    MistAgentVersion    *string `json:"mist_agent_version,omitempty"`
    PackageVersion      *string `json:"package_version,omitempty"`
    TestingToolsVersion *string `json:"testing_tools_version,omitempty"`
    WheeljackVersion    *string `json:"wheeljack_version,omitempty"`
}
