package models

import (
    "encoding/json"
    "fmt"
    "github.com/google/uuid"
)

// IdpConfig represents a IdpConfig struct.
type IdpConfig struct {
    AlertOnly            *bool                  `json:"alert_only,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    // org_level IDP Profile can be used, this takes precedence over `profile`
    IdpprofileId         *uuid.UUID             `json:"idpprofile_id,omitempty"`
    // enum: `Custom`, `strict` (default), `standard` or keys from from idp_profiles
    Profile              *string                `json:"profile,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for IdpConfig,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (i IdpConfig) String() string {
    return fmt.Sprintf(
    	"IdpConfig[AlertOnly=%v, Enabled=%v, IdpprofileId=%v, Profile=%v, AdditionalProperties=%v]",
    	i.AlertOnly, i.Enabled, i.IdpprofileId, i.Profile, i.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for IdpConfig.
// It customizes the JSON marshaling process for IdpConfig objects.
func (i IdpConfig) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(i.AdditionalProperties,
        "alert_only", "enabled", "idpprofile_id", "profile"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(i.toMap())
}

// toMap converts the IdpConfig object to a map representation for JSON marshaling.
func (i IdpConfig) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, i.AdditionalProperties)
    if i.AlertOnly != nil {
        structMap["alert_only"] = i.AlertOnly
    }
    if i.Enabled != nil {
        structMap["enabled"] = i.Enabled
    }
    if i.IdpprofileId != nil {
        structMap["idpprofile_id"] = i.IdpprofileId
    }
    if i.Profile != nil {
        structMap["profile"] = i.Profile
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for IdpConfig.
// It customizes the JSON unmarshaling process for IdpConfig objects.
func (i *IdpConfig) UnmarshalJSON(input []byte) error {
    var temp tempIdpConfig
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "alert_only", "enabled", "idpprofile_id", "profile")
    if err != nil {
    	return err
    }
    i.AdditionalProperties = additionalProperties
    
    i.AlertOnly = temp.AlertOnly
    i.Enabled = temp.Enabled
    i.IdpprofileId = temp.IdpprofileId
    i.Profile = temp.Profile
    return nil
}

// tempIdpConfig is a temporary struct used for validating the fields of IdpConfig.
type tempIdpConfig  struct {
    AlertOnly    *bool      `json:"alert_only,omitempty"`
    Enabled      *bool      `json:"enabled,omitempty"`
    IdpprofileId *uuid.UUID `json:"idpprofile_id,omitempty"`
    Profile      *string    `json:"profile,omitempty"`
}
