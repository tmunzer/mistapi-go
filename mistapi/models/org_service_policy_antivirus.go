package models

import (
	"encoding/json"
	"github.com/google/uuid"
)

// OrgServicePolicyAntivirus represents a OrgServicePolicyAntivirus struct.
// for SRX-only
type OrgServicePolicyAntivirus struct {
	// org-level AV Profile can be used, this takes precendence over 'profile'
	AvprofileId *uuid.UUID `json:"avprofile_id,omitempty"`
	Enabled     *bool      `json:"enabled,omitempty"`
	// default / noftp / httponly / or keys from av_profiles
	Profile              *string        `json:"profile,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgServicePolicyAntivirus.
// It customizes the JSON marshaling process for OrgServicePolicyAntivirus objects.
func (o OrgServicePolicyAntivirus) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgServicePolicyAntivirus object to a map representation for JSON marshaling.
func (o OrgServicePolicyAntivirus) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.AvprofileId != nil {
		structMap["avprofile_id"] = o.AvprofileId
	}
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	if o.Profile != nil {
		structMap["profile"] = o.Profile
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgServicePolicyAntivirus.
// It customizes the JSON unmarshaling process for OrgServicePolicyAntivirus objects.
func (o *OrgServicePolicyAntivirus) UnmarshalJSON(input []byte) error {
	var temp tempOrgServicePolicyAntivirus
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "avprofile_id", "enabled", "profile")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.AvprofileId = temp.AvprofileId
	o.Enabled = temp.Enabled
	o.Profile = temp.Profile
	return nil
}

// tempOrgServicePolicyAntivirus is a temporary struct used for validating the fields of OrgServicePolicyAntivirus.
type tempOrgServicePolicyAntivirus struct {
	AvprofileId *uuid.UUID `json:"avprofile_id,omitempty"`
	Enabled     *bool      `json:"enabled,omitempty"`
	Profile     *string    `json:"profile,omitempty"`
}
