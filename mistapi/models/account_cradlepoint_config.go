package models

import (
	"encoding/json"
)

// AccountCradlepointConfig represents a AccountCradlepointConfig struct.
type AccountCradlepointConfig struct {
	CpApiId              *string        `json:"cp_api_id,omitempty"`
	CpApiKey             *string        `json:"cp_api_key,omitempty"`
	EcmApiId             *string        `json:"ecm_api_id,omitempty"`
	EcmApiKey            *string        `json:"ecm_api_key,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for AccountCradlepointConfig.
// It customizes the JSON marshaling process for AccountCradlepointConfig objects.
func (a AccountCradlepointConfig) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(a.toMap())
}

// toMap converts the AccountCradlepointConfig object to a map representation for JSON marshaling.
func (a AccountCradlepointConfig) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, a.AdditionalProperties)
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
	additionalProperties, err := UnmarshalAdditionalProperties(input, "cp_api_id", "cp_api_key", "ecm_api_id", "ecm_api_key")
	if err != nil {
		return err
	}

	a.AdditionalProperties = additionalProperties
	a.CpApiId = temp.CpApiId
	a.CpApiKey = temp.CpApiKey
	a.EcmApiId = temp.EcmApiId
	a.EcmApiKey = temp.EcmApiKey
	return nil
}

// tempAccountCradlepointConfig is a temporary struct used for validating the fields of AccountCradlepointConfig.
type tempAccountCradlepointConfig struct {
	CpApiId   *string `json:"cp_api_id,omitempty"`
	CpApiKey  *string `json:"cp_api_key,omitempty"`
	EcmApiId  *string `json:"ecm_api_id,omitempty"`
	EcmApiKey *string `json:"ecm_api_key,omitempty"`
}
