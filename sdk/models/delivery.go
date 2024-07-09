package models

import (
    "encoding/json"
    "errors"
    "strings"
)

// Delivery represents a Delivery struct.
// Delivery object to configure the alarm delivery
type Delivery struct {
    // List of additional email string to deliver the alarms via emails
    AdditionalEmails     []string       `json:"additional_emails,omitempty"`
    // Whether to enable the alarm delivery via emails or not
    Enabled              bool           `json:"enabled"`
    // Whether to deliver the alarms via emails to Org admins or not
    ToOrgAdmins          *bool          `json:"to_org_admins,omitempty"`
    // Whether to deliver the alarms via emails to Site admins or not
    ToSiteAdmins         *bool          `json:"to_site_admins,omitempty"`
    AdditionalProperties map[string]any `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for Delivery.
// It customizes the JSON marshaling process for Delivery objects.
func (d Delivery) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(d.toMap())
}

// toMap converts the Delivery object to a map representation for JSON marshaling.
func (d Delivery) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, d.AdditionalProperties)
    if d.AdditionalEmails != nil {
        structMap["additional_emails"] = d.AdditionalEmails
    }
    structMap["enabled"] = d.Enabled
    if d.ToOrgAdmins != nil {
        structMap["to_org_admins"] = d.ToOrgAdmins
    }
    if d.ToSiteAdmins != nil {
        structMap["to_site_admins"] = d.ToSiteAdmins
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for Delivery.
// It customizes the JSON unmarshaling process for Delivery objects.
func (d *Delivery) UnmarshalJSON(input []byte) error {
    var temp delivery
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    err = temp.validate()
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "additional_emails", "enabled", "to_org_admins", "to_site_admins")
    if err != nil {
    	return err
    }
    
    d.AdditionalProperties = additionalProperties
    d.AdditionalEmails = temp.AdditionalEmails
    d.Enabled = *temp.Enabled
    d.ToOrgAdmins = temp.ToOrgAdmins
    d.ToSiteAdmins = temp.ToSiteAdmins
    return nil
}

// delivery is a temporary struct used for validating the fields of Delivery.
type delivery  struct {
    AdditionalEmails []string `json:"additional_emails,omitempty"`
    Enabled          *bool    `json:"enabled"`
    ToOrgAdmins      *bool    `json:"to_org_admins,omitempty"`
    ToSiteAdmins     *bool    `json:"to_site_admins,omitempty"`
}

func (d *delivery) validate() error {
    var errs []string
    if d.Enabled == nil {
        errs = append(errs, "required field `enabled` is missing for type `Delivery`")
    }
    if len(errs) == 0 {
        return nil
    }
    return errors.New(strings.Join(errs, "\n"))
}
