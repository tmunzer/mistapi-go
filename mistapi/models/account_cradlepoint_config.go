package models

import (
    "encoding/json"
)

// AccountCradlepointConfig represents a AccountCradlepointConfig struct.
type AccountCradlepointConfig struct {
    CpApiId              *string                `json:"cp_api_id,omitempty"`
    CpApiKey             *string                `json:"cp_api_key,omitempty"`
    EcmApiId             *string                `json:"ecm_api_id,omitempty"`
    EcmApiKey            *string                `json:"ecm_api_key,omitempty"`
    EnableLldp           *bool                  `json:"enable_lldp,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountCradlepointConfig.
// It customizes the JSON marshaling process for AccountCradlepointConfig objects.
func (a AccountCradlepointConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(a.AdditionalProperties,
        "cp_api_id", "cp_api_key", "ecm_api_id", "ecm_api_key", "enable_lldp"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(a.toMap())
}

// toMap converts the AccountCradlepointConfig object to a map representation for JSON marshaling.
func (a AccountCradlepointConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, a.AdditionalProperties)
    if a.CpApiId != nil {
        structMap["cp_api_id"] = a.CpApiId
    }
    if a.CpApiKey != nil {
        structMap["cp_api_key"] = a.CpApiKey
    }
    if a.EcmApiId != nil {
        structMap["ecm_api_id"] = a.EcmApiId
    }
    if a.EcmApiKey != nil {
        structMap["ecm_api_key"] = a.EcmApiKey
    }
    if a.EnableLldp != nil {
        structMap["enable_lldp"] = a.EnableLldp
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for AccountCradlepointConfig.
// It customizes the JSON unmarshaling process for AccountCradlepointConfig objects.
func (a *AccountCradlepointConfig) UnmarshalJSON(input []byte) error {
    var temp tempAccountCradlepointConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cp_api_id", "cp_api_key", "ecm_api_id", "ecm_api_key", "enable_lldp")
    if err != nil {
    	return err
    }
    a.AdditionalProperties = additionalProperties
    
    a.CpApiId = temp.CpApiId
    a.CpApiKey = temp.CpApiKey
    a.EcmApiId = temp.EcmApiId
    a.EcmApiKey = temp.EcmApiKey
    a.EnableLldp = temp.EnableLldp
    return nil
}

// tempAccountCradlepointConfig is a temporary struct used for validating the fields of AccountCradlepointConfig.
type tempAccountCradlepointConfig  struct {
    CpApiId    *string `json:"cp_api_id,omitempty"`
    CpApiKey   *string `json:"cp_api_key,omitempty"`
    EcmApiId   *string `json:"ecm_api_id,omitempty"`
    EcmApiKey  *string `json:"ecm_api_key,omitempty"`
    EnableLldp *bool   `json:"enable_lldp,omitempty"`
}
