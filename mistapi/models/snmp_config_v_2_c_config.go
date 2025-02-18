package models

import (
    "encoding/json"
    "fmt"
)

// SnmpConfigV2cConfig represents a SnmpConfigV2cConfig struct.
type SnmpConfigV2cConfig struct {
    Authorization        *string                `json:"authorization,omitempty"`
    // Client_list_name here should refer to client_list above
    ClientListName       *string                `json:"client_list_name,omitempty"`
    CommunityName        *string                `json:"community_name,omitempty"`
    // View name here should be defined in views above
    View                 *string                `json:"view,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SnmpConfigV2cConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SnmpConfigV2cConfig) String() string {
    return fmt.Sprintf(
    	"SnmpConfigV2cConfig[Authorization=%v, ClientListName=%v, CommunityName=%v, View=%v, AdditionalProperties=%v]",
    	s.Authorization, s.ClientListName, s.CommunityName, s.View, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SnmpConfigV2cConfig.
// It customizes the JSON marshaling process for SnmpConfigV2cConfig objects.
func (s SnmpConfigV2cConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "authorization", "client_list_name", "community_name", "view"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SnmpConfigV2cConfig object to a map representation for JSON marshaling.
func (s SnmpConfigV2cConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
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

// UnmarshalJSON implements the json.Unmarshaler interface for SnmpConfigV2cConfig.
// It customizes the JSON unmarshaling process for SnmpConfigV2cConfig objects.
func (s *SnmpConfigV2cConfig) UnmarshalJSON(input []byte) error {
    var temp tempSnmpConfigV2cConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "authorization", "client_list_name", "community_name", "view")
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

// tempSnmpConfigV2cConfig is a temporary struct used for validating the fields of SnmpConfigV2cConfig.
type tempSnmpConfigV2cConfig  struct {
    Authorization  *string `json:"authorization,omitempty"`
    ClientListName *string `json:"client_list_name,omitempty"`
    CommunityName  *string `json:"community_name,omitempty"`
    View           *string `json:"view,omitempty"`
}
