package models

import (
    "encoding/json"
)

// SnmpConfigV2CConfig represents a SnmpConfigV2CConfig struct.
type SnmpConfigV2CConfig struct {
    Authorization        *string        `json:"authorization,omitempty"`
    // client_list_name here should refer to client_list above
    ClientListName       *string        `json:"client_list_name,omitempty"`
    CommunityName        *string        `json:"community_name,omitempty"`
    // view name here should be defined in views above
    View                 *string        `json:"view,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigV2CConfig.
// It customizes the JSON marshaling process for SnmpConfigV2CConfig objects.
func (s SnmpConfigV2CConfig) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigV2CConfig object to a map representation for JSON marshaling.
func (s SnmpConfigV2CConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, s.AdditionalProperties)
    if s.Authorization != nil {
        structMap["authorization"] = s.Authorization
    }
    if s.ClientListName != nil {
        structMap["client_list_name"] = s.ClientListName
    }
    if s.CommunityName != nil {
        structMap["community_name"] = s.CommunityName
    }
    if s.View != nil {
        structMap["view"] = s.View
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpConfigV2CConfig.
// It customizes the JSON unmarshaling process for SnmpConfigV2CConfig objects.
func (s *SnmpConfigV2CConfig) UnmarshalJSON(input []byte) error {
    var temp snmpConfigV2CConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "authorization", "client_list_name", "community_name", "view")
    if err != nil {
    	return err
    }
    
    s.AdditionalProperties = additionalProperties
    s.Authorization = temp.Authorization
    s.ClientListName = temp.ClientListName
    s.CommunityName = temp.CommunityName
    s.View = temp.View
    return nil
}

// snmpConfigV2CConfig is a temporary struct used for validating the fields of SnmpConfigV2CConfig.
type snmpConfigV2CConfig  struct {
    Authorization  *string `json:"authorization,omitempty"`
    ClientListName *string `json:"client_list_name,omitempty"`
    CommunityName  *string `json:"community_name,omitempty"`
    View           *string `json:"view,omitempty"`
}
