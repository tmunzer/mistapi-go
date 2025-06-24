package models

import (
    "encoding/json"
    "fmt"
)

// OrgSettingScep represents a OrgSettingScep struct.
type OrgSettingScep struct {
    CertProviders        []string               `json:"cert_providers,omitempty"`
    Enabled              *bool                  `json:"enabled,omitempty"`
    IntuneScepUrl        *string                `json:"intune_scep_url,omitempty"`
    JamfAccessToken      *string                `json:"jamf_access_token,omitempty"`
    JamfScepUrl          *string                `json:"jamf_scep_url,omitempty"`
    JamfWebhookUrl       *string                `json:"jamf_webhook_url,omitempty"`
    AdditionalProperties map[string]interface{} `json:"_"`
}

// String implements the fmt.Stringer interface for OrgSettingScep,
// providing a human-readable string representation useful for logging, debugging or displaying information.
func (o OrgSettingScep) String() string {
    return fmt.Sprintf(
    	"OrgSettingScep[CertProviders=%v, Enabled=%v, IntuneScepUrl=%v, JamfAccessToken=%v, JamfScepUrl=%v, JamfWebhookUrl=%v, AdditionalProperties=%v]",
    	o.CertProviders, o.Enabled, o.IntuneScepUrl, o.JamfAccessToken, o.JamfScepUrl, o.JamfWebhookUrl, o.AdditionalProperties)
}

// MarshalJSON implements the json.Marshaler interface for OrgSettingScep.
// It customizes the JSON marshaling process for OrgSettingScep objects.
func (o OrgSettingScep) MarshalJSON() (
    []byte,
    error) {
    if err := DetectConflictingProperties(o.AdditionalProperties,
        "cert_providers", "enabled", "intune_scep_url", "jamf_access_token", "jamf_scep_url", "jamf_webhook_url"); err != nil {
        return []byte{}, err
    }
    return json.Marshal(o.toMap())
}

// toMap converts the OrgSettingScep object to a map representation for JSON marshaling.
func (o OrgSettingScep) toMap() map[string]any {
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
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for OrgSettingScep.
// It customizes the JSON unmarshaling process for OrgSettingScep objects.
func (o *OrgSettingScep) UnmarshalJSON(input []byte) error {
    var temp tempOrgSettingScep
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := ExtractAdditionalProperties[interface{}](input, "cert_providers", "enabled", "intune_scep_url", "jamf_access_token", "jamf_scep_url", "jamf_webhook_url")
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
    return nil
}

// tempOrgSettingScep is a temporary struct used for validating the fields of OrgSettingScep.
type tempOrgSettingScep  struct {
    CertProviders   []string `json:"cert_providers,omitempty"`
    Enabled         *bool    `json:"enabled,omitempty"`
    IntuneScepUrl   *string  `json:"intune_scep_url,omitempty"`
    JamfAccessToken *string  `json:"jamf_access_token,omitempty"`
    JamfScepUrl     *string  `json:"jamf_scep_url,omitempty"`
    JamfWebhookUrl  *string  `json:"jamf_webhook_url,omitempty"`
}
