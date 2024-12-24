package models

import (
    "encoding/json"
    "fmt"
)

// SamlMetadata represents a SamlMetadata struct.
type SamlMetadata struct {
    // if `idp_type`==`saml`
    AcsUrl               *string                `json:"acs_url,omitempty"`
    // if `idp_type`==`saml`
    EntityId             *string                `json:"entity_id,omitempty"`
    // if `idp_type`==`saml`
    LogoutUrl            *string                `json:"logout_url,omitempty"`
    // if `idp_type`==`saml`
    Metadata             *string                `json:"metadata,omitempty"`
    // if `idp_type`==`oauth` and `scim_enabled`==`true`
    ScimBaseUrl          *string                `json:"scim_base_url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for SamlMetadata,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (s SamlMetadata) String() string {
    return fmt.Sprintf(
    	"SamlMetadata[AcsUrl=%v, EntityId=%v, LogoutUrl=%v, Metadata=%v, ScimBaseUrl=%v, AdditionalProperties=%v]",
    	s.AcsUrl, s.EntityId, s.LogoutUrl, s.Metadata, s.ScimBaseUrl, s.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for SamlMetadata.
// It customizes the JSON marshaling process for SamlMetadata objects.
func (s SamlMetadata) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(s.AdditionalProperties,
        "acs_url", "entity_id", "logout_url", "metadata", "scim_base_url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(s.toMap())
}

// toMap converts the SamlMetadata object to a map representation for JSON marshaling.
func (s SamlMetadata) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, s.AdditionalProperties)
    if s.AcsUrl != nil {
        structMap["acs_url"] = s.AcsUrl
    }
    if s.EntityId != nil {
        structMap["entity_id"] = s.EntityId
    }
    if s.LogoutUrl != nil {
        structMap["logout_url"] = s.LogoutUrl
    }
    if s.Metadata != nil {
        structMap["metadata"] = s.Metadata
    }
    if s.ScimBaseUrl != nil {
        structMap["scim_base_url"] = s.ScimBaseUrl
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SamlMetadata.
// It customizes the JSON unmarshaling process for SamlMetadata objects.
func (s *SamlMetadata) UnmarshalJSON(input []byte) error {
    var temp tempSamlMetadata
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "acs_url", "entity_id", "logout_url", "metadata", "scim_base_url")
    if err != nil {
    	return err
    }
    s.AdditionalProperties = additionalProperties
    
    s.AcsUrl = temp.AcsUrl
    s.EntityId = temp.EntityId
    s.LogoutUrl = temp.LogoutUrl
    s.Metadata = temp.Metadata
    s.ScimBaseUrl = temp.ScimBaseUrl
    return nil
}

// tempSamlMetadata is a temporary struct used for validating the fields of SamlMetadata.
type tempSamlMetadata  struct {
    AcsUrl      *string `json:"acs_url,omitempty"`
    EntityId    *string `json:"entity_id,omitempty"`
    LogoutUrl   *string `json:"logout_url,omitempty"`
    Metadata    *string `json:"metadata,omitempty"`
    ScimBaseUrl *string `json:"scim_base_url,omitempty"`
}
