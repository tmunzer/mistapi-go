package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// MarvisClient represents a MarvisClient struct.
type MarvisClient struct {
    Disabled             *bool                  `json:"disabled,omitempty"`
    // Unique ID of the object instance in the Mist Organization
    Id                   *uuid.UUID             `json:"id,omitempty"`
    Name                 *string                `json:"name,omitempty"`
    // In MDM, add `--provision_url <provision_url>` to the install command
    ProvisionUrl         *string                `json:"provision_url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for MarvisClient,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (m MarvisClient) String() string {
    return fmt.Sprintf(
    	"MarvisClient[Disabled=%v, Id=%v, Name=%v, ProvisionUrl=%v, AdditionalProperties=%v]",
    	m.Disabled, m.Id, m.Name, m.ProvisionUrl, m.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for MarvisClient.
// It customizes the JSON marshaling process for MarvisClient objects.
func (m MarvisClient) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(m.AdditionalProperties,
        "disabled", "id", "name", "provision_url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(m.toMap())
}

// toMap converts the MarvisClient object to a map representation for JSON marshaling.
func (m MarvisClient) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, m.AdditionalProperties)
    if m.Disabled != nil {
        structMap["disabled"] = m.Disabled
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
    var temp tempMarvisClient
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "disabled", "id", "name", "provision_url")
    if err != nil {
    	return err
    }
    m.AdditionalProperties = additionalProperties
    
    m.Disabled = temp.Disabled
    m.Id = temp.Id
    m.Name = temp.Name
    m.ProvisionUrl = temp.ProvisionUrl
    return nil
}

// tempMarvisClient is a temporary struct used for validating the fields of MarvisClient.
type tempMarvisClient  struct {
    Disabled     *bool      `json:"disabled,omitempty"`
    Id           *uuid.UUID `json:"id,omitempty"`
    Name         *string    `json:"name,omitempty"`
    ProvisionUrl *string    `json:"provision_url,omitempty"`
}
