package models

import (
	"encoding/json"
	"errors"
	"strings"
)

// SsoSamlMetadata represents a SsoSamlMetadata struct.
type SsoSamlMetadata struct {
	AcsUrl               string         `json:"acs_url"`
	EntityId             string         `json:"entity_id"`
	LogoutUrl            string         `json:"logout_url"`
	Metadata             string         `json:"metadata"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for SsoSamlMetadata.
// It customizes the JSON marshaling process for SsoSamlMetadata objects.
func (s SsoSamlMetadata) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(s.toMap())
}

// toMap converts the SsoSamlMetadata object to a map representation for JSON marshaling.
func (s SsoSamlMetadata) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, s.AdditionalProperties)
	structMap["acs_url"] = s.AcsUrl
	structMap["entity_id"] = s.EntityId
	structMap["logout_url"] = s.LogoutUrl
	structMap["metadata"] = s.Metadata
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for SsoSamlMetadata.
// It customizes the JSON unmarshaling process for SsoSamlMetadata objects.
func (s *SsoSamlMetadata) UnmarshalJSON(input []byte) error {
	var temp tempSsoSamlMetadata
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	err = temp.validate()
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "acs_url", "entity_id", "logout_url", "metadata")
	if err != nil {
		return err
	}

	s.AdditionalProperties = additionalProperties
	s.AcsUrl = *temp.AcsUrl
	s.EntityId = *temp.EntityId
	s.LogoutUrl = *temp.LogoutUrl
	s.Metadata = *temp.Metadata
	return nil
}

// tempSsoSamlMetadata is a temporary struct used for validating the fields of SsoSamlMetadata.
type tempSsoSamlMetadata struct {
	AcsUrl    *string `json:"acs_url"`
	EntityId  *string `json:"entity_id"`
	LogoutUrl *string `json:"logout_url"`
	Metadata  *string `json:"metadata"`
}

func (s *tempSsoSamlMetadata) validate() error {
	var errs []string
	if s.AcsUrl == nil {
		errs = append(errs, "required field `acs_url` is missing for type `sso_saml_metadata`")
	}
	if s.EntityId == nil {
		errs = append(errs, "required field `entity_id` is missing for type `sso_saml_metadata`")
	}
	if s.LogoutUrl == nil {
		errs = append(errs, "required field `logout_url` is missing for type `sso_saml_metadata`")
	}
	if s.Metadata == nil {
		errs = append(errs, "required field `metadata` is missing for type `sso_saml_metadata`")
	}
	if len(errs) == 0 {
		return nil
	}
	return errors.New(strings.Join(errs, "\n"))
}
