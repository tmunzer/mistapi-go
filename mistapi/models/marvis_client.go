package models

import (
    "encoding/json"
    "github.com/google/uuid"
)

// MarvisClient represents a MarvisClient struct.
type MarvisClient struct {
    Diabled              *bool          `json:"diabled,omitempty"`
    Id                   *uuid.UUID     `json:"id,omitempty"`
    Name                 *string        `json:"name,omitempty"`
    // in MDM, add `--provision_url <provision_url>` to the instlal command
    ProvisionUrl         *string        `json:"provision_url,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for MarvisClient.
// It customizes the JSON marshaling process for MarvisClient objects.
func (m MarvisClient) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(m.toMap())
}

// toMap converts the MarvisClient object to a map representation for JSON marshaling.
func (m MarvisClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Diabled != nil {
        structMap["diabled"] = m.Diabled
    }
    if m.Id != nil {
        structMap["id"] = m.Id
    }
    if m.Name != nil {
        structMap["name"] = m.Name
    }
    if m.ProvisionUrl != nil {
        structMap["provision_url"] = m.ProvisionUrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for MarvisClient.
// It customizes the JSON unmarshaling process for MarvisClient objects.
func (m *MarvisClient) UnmarshalJSON(input []byte) error {
    var temp marvisClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "diabled", "id", "name", "provision_url")
    if err != nil {
    	return err
    }
    
    m.AdditionalProperties = additionalProperties
    m.Diabled = temp.Diabled
    m.Id = temp.Id
    m.Name = temp.Name
    m.ProvisionUrl = temp.ProvisionUrl
    return nil
}

// marvisClient is a temporary struct used for validating the fields of MarvisClient.
type marvisClient  struct {
    Diabled      *bool      `json:"diabled,omitempty"`
    Id           *uuid.UUID `json:"id,omitempty"`
    Name         *string    `json:"name,omitempty"`
    ProvisionUrl *string    `json:"provision_url,omitempty"`
}
