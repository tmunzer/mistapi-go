// Package mistapi
// Copyright \xA9 2024 Juniper Networks, Inc. All rights reserved
package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingScepResponse represents a OrgSettingScepResponse struct.
type OrgSettingScepResponse struct {
    // List of SCEP cert providers, e.g. `intune`, `jamf`, `byod`
    CertProviders        []OrgSettingScepCertProviderEnum `json:"cert_providers,omitempty"`
    Enabled              *bool                            `json:"enabled,omitempty"`
    IntuneScepUrl        *string                          `json:"intune_scep_url,omitempty"`
    JamfAccessToken      *string                          `json:"jamf_access_token,omitempty"`
    JamfScepUrl          *string                          `json:"jamf_scep_url,omitempty"`
    JamfWebhookUrl       *string                          `json:"jamf_webhook_url,omitempty"`
    // Whether SCEP is suspended for this org
    Suspended            *bool                            `json:"suspended,omitempty"`
    AdditionalProperties map[string]interface{}           `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingScepResponse,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingScepResponse) String() string {
    return fmt.Sprintf(
    	"OrgSettingScepResponse[CertProviders=%v, Enabled=%v, IntuneScepUrl=%v, JamfAccessToken=%v, JamfScepUrl=%v, JamfWebhookUrl=%v, Suspended=%v, AdditionalProperties=%v]",
    	o.CertProviders, o.Enabled, o.IntuneScepUrl, o.JamfAccessToken, o.JamfScepUrl, o.JamfWebhookUrl, o.Suspended, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingScepResponse.
// It customizes the JSON marshaling process for OrgSettingScepResponse objects.
func (o OrgSettingScepResponse) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "cert_providers", "enabled", "intune_scep_url", "jamf_access_token", "jamf_scep_url", "jamf_webhook_url", "suspended"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingScepResponse object to a map representation for JSON marshaling.
func (o OrgSettingScepResponse) toMap() map[string]any {
    structMap := make(map[string]any)
    MergeAdditionalProperties(structMap, o.AdditionalProperties)
    if o.CertProviders != nil {
        structMap["cert_providers"] = o.CertProviders
    }
    if o.Enabled != nil {
        structMap["enabled"] = o.Enabled
    }
    if o.IntuneScepUrl != nil {
        structMap["intune_scep_url"] = o.IntuneScepUrl
    }
    if o.JamfAccessToken != nil {
        structMap["jamf_access_token"] = o.JamfAccessToken
    }
    if o.JamfScepUrl != nil {
        structMap["jamf_scep_url"] = o.JamfScepUrl
    }
    if o.JamfWebhookUrl != nil {
        structMap["jamf_webhook_url"] = o.JamfWebhookUrl
    }
    if o.Suspended != nil {
        structMap["suspended"] = o.Suspended
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingScepResponse.
// It customizes the JSON unmarshaling process for OrgSettingScepResponse objects.
func (o *OrgSettingScepResponse) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingScepResponse
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert_providers", "enabled", "intune_scep_url", "jamf_access_token", "jamf_scep_url", "jamf_webhook_url", "suspended")
    if err != nil {
    	return err
    }
    o.AdditionalProperties = additionalProperties
    
    o.CertProviders = temp.CertProviders
    o.Enabled = temp.Enabled
    o.IntuneScepUrl = temp.IntuneScepUrl
    o.JamfAccessToken = temp.JamfAccessToken
    o.JamfScepUrl = temp.JamfScepUrl
    o.JamfWebhookUrl = temp.JamfWebhookUrl
    o.Suspended = temp.Suspended
    return nil
}

// tempOrgSettingScepResponse is a temporary struct used for validating the fields of OrgSettingScepResponse.
type tempOrgSettingScepResponse  struct {
    CertProviders   []OrgSettingScepCertProviderEnum `json:"cert_providers,omitempty"`
    Enabled         *bool                            `json:"enabled,omitempty"`
    IntuneScepUrl   *string                          `json:"intune_scep_url,omitempty"`
    JamfAccessToken *string                          `json:"jamf_access_token,omitempty"`
    JamfScepUrl     *string                          `json:"jamf_scep_url,omitempty"`
    JamfWebhookUrl  *string                          `json:"jamf_webhook_url,omitempty"`
    Suspended       *bool                            `json:"suspended,omitempty"`
}
