package models

import (
	"encoding/json"
)

// OrgServicePoliciesSecintel represents a OrgServicePoliciesSecintel struct.
// For SRX Only
type OrgServicePoliciesSecintel struct {
	Enabled *bool `json:"enabled,omitempty"`
	// enum: `default`, `standard`, `strict`
	Profile *SecintelProfileProfileActionEnum `json:"profile,omitempty"`
	// org-level secintel Profile can be used, this takes precendence over 'profile'
	SecintelprofileId    *string        `json:"secintelprofile_id,omitempty"`
	AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for OrgServicePoliciesSecintel.
// It customizes the JSON marshaling process for OrgServicePoliciesSecintel objects.
func (o OrgServicePoliciesSecintel) MarshalJSON() (
	[]byte,
	error) {
	return json.Marshal(o.toMap())
}

// toMap converts the OrgServicePoliciesSecintel object to a map representation for JSON marshaling.
func (o OrgServicePoliciesSecintel) toMap() map[string]any {
	structMap := make(map[string]any)
	MapAdditionalProperties(structMap, o.AdditionalProperties)
	if o.Enabled != nil {
		structMap["enabled"] = o.Enabled
	}
	if o.Profile != nil {
		structMap["profile"] = o.Profile
	}
	if o.SecintelprofileId != nil {
		structMap["secintelprofile_id"] = o.SecintelprofileId
	}
	return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgServicePoliciesSecintel.
// It customizes the JSON unmarshaling process for OrgServicePoliciesSecintel objects.
func (o *OrgServicePoliciesSecintel) UnmarshalJSON(input []byte) error {
	var temp tempOrgServicePoliciesSecintel
	err := json.Unmarshal(input, &temp)
	if err != nil {
		return err
	}
	additionalProperties, err := UnmarshalAdditionalProperties(input, "enabled", "profile", "secintelprofile_id")
	if err != nil {
		return err
	}

	o.AdditionalProperties = additionalProperties
	o.Enabled = temp.Enabled
	o.Profile = temp.Profile
	o.SecintelprofileId = temp.SecintelprofileId
	return nil
}

// tempOrgServicePoliciesSecintel is a temporary struct used for validating the fields of OrgServicePoliciesSecintel.
type tempOrgServicePoliciesSecintel struct {
	Enabled           *bool                             `json:"enabled,omitempty"`
	Profile           *SecintelProfileProfileActionEnum `json:"profile,omitempty"`
	SecintelprofileId *string                           `json:"secintelprofile_id,omitempty"`
}
