package models

import (
    "encoding/json"
)

// MspLogo represents a MspLogo struct.
type MspLogo struct {
    LogoUrl              *string        `json:"logo_url,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MspLogo.
// It customizes the JSON marshaling process for MspLogo objects.
func (m MspLogo) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MspLogo object to a map representation for JSON marshaling.
func (m MspLogo) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.LogoUrl != nil {
        structMap["logo_url"] = m.LogoUrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MspLogo.
// It customizes the JSON unmarshaling process for MspLogo objects.
func (m *MspLogo) UnmarshalJSON(input []byte) error {
    var temp tempMspLogo
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "logo_url")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.LogoUrl = temp.LogoUrl
    return nil
}

// tempMspLogo is a temporary struct used for validating the fields of MspLogo.
type tempMspLogo  struct {
    LogoUrl *string `json:"logo_url,omitempty"`
}
