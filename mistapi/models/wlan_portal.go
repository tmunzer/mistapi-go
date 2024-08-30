package models

import (
    "encoding/json"
)

// WlanPortal represents a WlanPortal struct.
// portal wlan settings
type WlanPortal struct {
    // amazon OAuth2 client id. This is optional. If not provided, it will use a default one.
    AmazonClientId              Optional[string]               `json:"amazon_client_id"`
    // amazon OAuth2 client secret. If amazon_client_id was provided, provide a correspoinding value. Else leave blank.
    AmazonClientSecret          Optional[string]               `json:"amazon_client_secret"`
    // Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
    AmazonEmailDomains          []string                       `json:"amazon_email_domains,omitempty"`
    // whether amazon is enabled as a login method
    AmazonEnabled               *bool                          `json:"amazon_enabled,omitempty"`
    // interval for which guest remains authorized using amazon auth (in minutes), if not provided, uses expire`
    AmazonExpire                Optional[float64]              `json:"amazon_expire"`
    // authentication scheme. enum: `external`, `none`, `sso`
    Auth                        *WlanPortalAuthEnum            `json:"auth,omitempty"`
    // Required if `azure_enabled`==`true`.
    // Azure active directory app client id
    AzureClientId               Optional[string]               `json:"azure_client_id"`
    // Required if `azure_enabled`==`true`.
    // Azure active directory app client secret
    AzureClientSecret           Optional[string]               `json:"azure_client_secret"`
    // whether Azure Active Directory is enabled as a login method
    AzureEnabled                *bool                          `json:"azure_enabled,omitempty"`
    // interval for which guest remains authorized using azure auth (in minutes), if not provided, uses expire`
    AzureExpire                 Optional[float64]              `json:"azure_expire"`
    // Required if `azure_enabled`==`true`.
    // Azure active directory tenant id.
    AzureTenantId               Optional[string]               `json:"azure_tenant_id"`
    // when `sms_provider`==`broadnet`
    BroadnetPassword            *string                        `json:"broadnet_password,omitempty"`
    // when `sms_provider`==`broadnet`
    BroadnetSid                 *string                        `json:"broadnet_sid,omitempty"`
    // when `sms_provider`==`broadnet`
    BroadnetUserId              *string                        `json:"broadnet_user_id,omitempty"`
    // whether to bypass the guest portal when cloud not reachable (and apply the default policies)
    BypassWhenCloudDown         *bool                          `json:"bypass_when_cloud_down,omitempty"`
    // when `sms_provider`==`clickatell`
    ClickatellApiKey            *string                        `json:"clickatell_api_key,omitempty"`
    // whether to allow guest to roam between WLANs (with same `WLAN.ssid`, regardless of variables) of different sites of same org without reauthentication (disable random_mac for seamless roaming)
    CrossSite                   *bool                          `json:"cross_site,omitempty"`
    // whether email (access code verification) is enabled as a login method
    EmailEnabled                *bool                          `json:"email_enabled,omitempty"`
    // whether guest portal is enabled
    Enabled                     *bool                          `json:"enabled,omitempty"`
    // how long to remain authorized, in minutes
    Expire                      *float64                       `json:"expire,omitempty"`
    // external portal URL (e.g. https://host/url) where we can append our query parameters to
    ExternalPortalUrl           *string                        `json:"external_portal_url,omitempty"`
    // Required if `facebook_enabled`==`true`.
    // Facebook OAuth2 app id. This is optional. If not provided, it will use a default one.
    FacebookClientId            Optional[string]               `json:"facebook_client_id"`
    // Required if `facebook_enabled`==`true`.
    // Facebook OAuth2 app secret. If facebook_client_id was provided, provide a correspoinding value. Else leave blank.
    FacebookClientSecret        Optional[string]               `json:"facebook_client_secret"`
    // Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
    FacebookEmailDomains        []string                       `json:"facebook_email_domains,omitempty"`
    // whether facebook is enabled as a login method
    FacebookEnabled             *bool                          `json:"facebook_enabled,omitempty"`
    // interval for which guest remains authorized using facebook auth (in minutes), if not provided, uses expire`
    FacebookExpire              Optional[float64]              `json:"facebook_expire"`
    // whether to forward the user to another URL after authorized
    Forward                     *bool                          `json:"forward,omitempty"`
    // the URL to forward the user to
    ForwardUrl                  Optional[string]               `json:"forward_url"`
    // Google OAuth2 app id. This is optional. If not provided, it will use a default one.
    GoogleClientId              Optional[string]               `json:"google_client_id"`
    // Google OAuth2 app secret. If google_client_id was provided, provide a correspoinding value. Else leave blank.
    GoogleClientSecret          Optional[string]               `json:"google_client_secret"`
    // Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
    GoogleEmailDomains          []string                       `json:"google_email_domains,omitempty"`
    // whether google is enabled as login method
    GoogleEnabled               *bool                          `json:"google_enabled,omitempty"`
    // interval for which guest remains authorized using google auth (in minutes), if not provided, uses expire`
    GoogleExpire                Optional[float64]              `json:"google_expire"`
    // when `sms_provider`==`gupshup`
    GupshupPassword             *string                        `json:"gupshup_password,omitempty"`
    // when `sms_provider`==`gupshup`
    GupshupUserid               *string                        `json:"gupshup_userid,omitempty"`
    // microsoft 365 OAuth2 client id. This is optional. If not provided, it will use a default one.
    MicrosoftClientId           Optional[string]               `json:"microsoft_client_id"`
    // microsoft 365 OAuth2 client secret. If microsoft_client_id was provided, provide a correspoinding value. Else leave blank.
    MicrosoftClientSecret       Optional[string]               `json:"microsoft_client_secret"`
    // Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed.
    MicrosoftEmailDomains       []string                       `json:"microsoft_email_domains,omitempty"`
    // whether microsoft 365 is enabled as a login method
    MicrosoftEnabled            *bool                          `json:"microsoft_enabled,omitempty"`
    // interval for which guest remains authorized using microsoft auth (in minutes), if not provided, uses expire`
    MicrosoftExpire             Optional[float64]              `json:"microsoft_expire"`
    // whether password is enabled
    PassphraseEnabled           *bool                          `json:"passphrase_enabled,omitempty"`
    // interval for which guest remains authorized using passphrase auth (in minutes), if not provided, uses `expire`
    PassphraseExpire            Optional[float64]              `json:"passphrase_expire"`
    // passphrase
    Password                    Optional[string]               `json:"password"`
    // whether to show list of sponsor emails mentioned in `sponsors` object as a dropdown. If both `sponsor_notify_all` and `predefined_sponsors_enabled` are false, behaviour is acc to `sponsor_email_domains`
    PredefinedSponsorsEnabled   *bool                          `json:"predefined_sponsors_enabled,omitempty"`
    // whether to hide sponsor’s email from list of sponsors
    PredefinedSponsorsHideEmail *bool                          `json:"predefined_sponsors_hide_email,omitempty"`
    Privacy                     *bool                          `json:"privacy,omitempty"`
    // when `sms_provider`==`puzzel`
    PuzzelPassword              *string                        `json:"puzzel_password,omitempty"`
    // when `sms_provider`==`puzzel`
    PuzzelServiceId             *string                        `json:"puzzel_service_id,omitempty"`
    // when `sms_provider`==`puzzel`
    PuzzelUsername              *string                        `json:"puzzel_username,omitempty"`
    SmsMessageFormat            *string                        `json:"smsMessageFormat,omitempty"`
    // whether sms is enabled as a login method
    SmsEnabled                  *bool                          `json:"sms_enabled,omitempty"`
    // interval for which guest remains authorized using sms auth (in minutes), if not provided, uses expire`
    SmsExpire                   Optional[float64]              `json:"sms_expire"`
    // enum: `broadnet`, `clickatell`, `gupshup`, `manual`, `puzzel`, `telstra`, `twilio`
    SmsProvider                 *WlanPortalSmsProviderEnum     `json:"sms_provider,omitempty"`
    // whether to automatically approve guest and allow sponsor to revoke guest access, needs predefined_sponsors_enabled enabled and sponsor_notify_all disabled
    SponsorAutoApprove          *bool                          `json:"sponsor_auto_approve,omitempty"`
    // list of domain allowed for sponsor email. Required if `sponsor_enabled` is `true` and `sponsors` is empty.
    SponsorEmailDomains         []string                       `json:"sponsor_email_domains,omitempty"`
    // whether sponsor is enabled
    SponsorEnabled              *bool                          `json:"sponsor_enabled,omitempty"`
    // interval for which guest remains authorized using sponsor auth (in minutes), if not provided, uses expire`
    SponsorExpire               Optional[float64]              `json:"sponsor_expire"`
    // how long to remain valid sponsored guest request approve/deny link received in email, in minutes.
    SponsorLinkValidityDuration *string                        `json:"sponsor_link_validity_duration,omitempty"`
    // whether to notify all sponsors that are mentioned in `sponsors` object. Both `sponsor_notify_all` and `predefined_sponsors_enabled` should be true in order to notify sponsors. If true, email sent to 10 sponsors in no particular order.
    SponsorNotifyAll            *bool                          `json:"sponsor_notify_all,omitempty"`
    // if enabled, guest will get email about sponsor's action (approve/deny)
    SponsorStatusNotify         *bool                          `json:"sponsor_status_notify,omitempty"`
    // object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty.
    // Property key is the sponsor email, Property value is the sponsor name
    Sponsors                    map[string]string              `json:"sponsors,omitempty"`
    // default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched
    SsoDefaultRole              *string                        `json:"sso_default_role,omitempty"`
    SsoForcedRole               *string                        `json:"sso_forced_role,omitempty"`
    // IDP Cert (used to verify the signed response)
    SsoIdpCert                  *string                        `json:"sso_idp_cert,omitempty"`
    // signing algorithm for SAML Assertion
    SsoIdpSignAlgo              *string                        `json:"sso_idp_sign_algo,omitempty"`
    // IDP Single-Sign-On URL
    SsoIdpSsoUrl                *string                        `json:"sso_idp_sso_url,omitempty"`
    // IDP issuer URL
    SsoIssuer                   *string                        `json:"sso_issuer,omitempty"`
    // enum: `email`, `unspecified`
    SsoNameidFormat             *WlanPortalSsoNameidFormatEnum `json:"sso_nameid_format,omitempty"`
    // when `sms_provider`==`telstra`, Client ID provided by Telstra
    TelstraClientId             *string                        `json:"telstra_client_id,omitempty"`
    // when `sms_provider`==`telstra`, Client secret provided by Telstra
    TelstraClientSecret         *string                        `json:"telstra_client_secret,omitempty"`
    // when `sms_provider`==`twilio`, Auth token account with twilio account
    TwilioAuthToken             Optional[string]               `json:"twilio_auth_token"`
    // when `sms_provider`==`twilio`, Twilio phone number associated with the account. See example for accepted format.
    TwilioPhoneNumber           Optional[string]               `json:"twilio_phone_number"`
    // when `sms_provider`==`twilio`, Account SID provided by Twilio
    TwilioSid                   Optional[string]               `json:"twilio_sid"`
    AdditionalProperties        map[string]any                 `json:"_"`
}

// MarshalJSON implements the json.Marshaler interface for WlanPortal.
// It customizes the JSON marshaling process for WlanPortal objects.
func (w WlanPortal) MarshalJSON() (
    []byte,
    error) {
    return json.Marshal(w.toMap())
}

// toMap converts the WlanPortal object to a map representation for JSON marshaling.
func (w WlanPortal) toMap() map[string]any {
    structMap := make(map[string]any)
    MapAdditionalProperties(structMap, w.AdditionalProperties)
    if w.AmazonClientId.IsValueSet() {
        if w.AmazonClientId.Value() != nil {
            structMap["amazon_client_id"] = w.AmazonClientId.Value()
        } else {
            structMap["amazon_client_id"] = nil
        }
    }
    if w.AmazonClientSecret.IsValueSet() {
        if w.AmazonClientSecret.Value() != nil {
            structMap["amazon_client_secret"] = w.AmazonClientSecret.Value()
        } else {
            structMap["amazon_client_secret"] = nil
        }
    }
    if w.AmazonEmailDomains != nil {
        structMap["amazon_email_domains"] = w.AmazonEmailDomains
    }
    if w.AmazonEnabled != nil {
        structMap["amazon_enabled"] = w.AmazonEnabled
    }
    if w.AmazonExpire.IsValueSet() {
        if w.AmazonExpire.Value() != nil {
            structMap["amazon_expire"] = w.AmazonExpire.Value()
        } else {
            structMap["amazon_expire"] = nil
        }
    }
    if w.Auth != nil {
        structMap["auth"] = w.Auth
    }
    if w.AzureClientId.IsValueSet() {
        if w.AzureClientId.Value() != nil {
            structMap["azure_client_id"] = w.AzureClientId.Value()
        } else {
            structMap["azure_client_id"] = nil
        }
    }
    if w.AzureClientSecret.IsValueSet() {
        if w.AzureClientSecret.Value() != nil {
            structMap["azure_client_secret"] = w.AzureClientSecret.Value()
        } else {
            structMap["azure_client_secret"] = nil
        }
    }
    if w.AzureEnabled != nil {
        structMap["azure_enabled"] = w.AzureEnabled
    }
    if w.AzureExpire.IsValueSet() {
        if w.AzureExpire.Value() != nil {
            structMap["azure_expire"] = w.AzureExpire.Value()
        } else {
            structMap["azure_expire"] = nil
        }
    }
    if w.AzureTenantId.IsValueSet() {
        if w.AzureTenantId.Value() != nil {
            structMap["azure_tenant_id"] = w.AzureTenantId.Value()
        } else {
            structMap["azure_tenant_id"] = nil
        }
    }
    if w.BroadnetPassword != nil {
        structMap["broadnet_password"] = w.BroadnetPassword
    }
    if w.BroadnetSid != nil {
        structMap["broadnet_sid"] = w.BroadnetSid
    }
    if w.BroadnetUserId != nil {
        structMap["broadnet_user_id"] = w.BroadnetUserId
    }
    if w.BypassWhenCloudDown != nil {
        structMap["bypass_when_cloud_down"] = w.BypassWhenCloudDown
    }
    if w.ClickatellApiKey != nil {
        structMap["clickatell_api_key"] = w.ClickatellApiKey
    }
    if w.CrossSite != nil {
        structMap["cross_site"] = w.CrossSite
    }
    if w.EmailEnabled != nil {
        structMap["email_enabled"] = w.EmailEnabled
    }
    if w.Enabled != nil {
        structMap["enabled"] = w.Enabled
    }
    if w.Expire != nil {
        structMap["expire"] = w.Expire
    }
    if w.ExternalPortalUrl != nil {
        structMap["external_portal_url"] = w.ExternalPortalUrl
    }
    if w.FacebookClientId.IsValueSet() {
        if w.FacebookClientId.Value() != nil {
            structMap["facebook_client_id"] = w.FacebookClientId.Value()
        } else {
            structMap["facebook_client_id"] = nil
        }
    }
    if w.FacebookClientSecret.IsValueSet() {
        if w.FacebookClientSecret.Value() != nil {
            structMap["facebook_client_secret"] = w.FacebookClientSecret.Value()
        } else {
            structMap["facebook_client_secret"] = nil
        }
    }
    if w.FacebookEmailDomains != nil {
        structMap["facebook_email_domains"] = w.FacebookEmailDomains
    }
    if w.FacebookEnabled != nil {
        structMap["facebook_enabled"] = w.FacebookEnabled
    }
    if w.FacebookExpire.IsValueSet() {
        if w.FacebookExpire.Value() != nil {
            structMap["facebook_expire"] = w.FacebookExpire.Value()
        } else {
            structMap["facebook_expire"] = nil
        }
    }
    if w.Forward != nil {
        structMap["forward"] = w.Forward
    }
    if w.ForwardUrl.IsValueSet() {
        if w.ForwardUrl.Value() != nil {
            structMap["forward_url"] = w.ForwardUrl.Value()
        } else {
            structMap["forward_url"] = nil
        }
    }
    if w.GoogleClientId.IsValueSet() {
        if w.GoogleClientId.Value() != nil {
            structMap["google_client_id"] = w.GoogleClientId.Value()
        } else {
            structMap["google_client_id"] = nil
        }
    }
    if w.GoogleClientSecret.IsValueSet() {
        if w.GoogleClientSecret.Value() != nil {
            structMap["google_client_secret"] = w.GoogleClientSecret.Value()
        } else {
            structMap["google_client_secret"] = nil
        }
    }
    if w.GoogleEmailDomains != nil {
        structMap["google_email_domains"] = w.GoogleEmailDomains
    }
    if w.GoogleEnabled != nil {
        structMap["google_enabled"] = w.GoogleEnabled
    }
    if w.GoogleExpire.IsValueSet() {
        if w.GoogleExpire.Value() != nil {
            structMap["google_expire"] = w.GoogleExpire.Value()
        } else {
            structMap["google_expire"] = nil
        }
    }
    if w.GupshupPassword != nil {
        structMap["gupshup_password"] = w.GupshupPassword
    }
    if w.GupshupUserid != nil {
        structMap["gupshup_userid"] = w.GupshupUserid
    }
    if w.MicrosoftClientId.IsValueSet() {
        if w.MicrosoftClientId.Value() != nil {
            structMap["microsoft_client_id"] = w.MicrosoftClientId.Value()
        } else {
            structMap["microsoft_client_id"] = nil
        }
    }
    if w.MicrosoftClientSecret.IsValueSet() {
        if w.MicrosoftClientSecret.Value() != nil {
            structMap["microsoft_client_secret"] = w.MicrosoftClientSecret.Value()
        } else {
            structMap["microsoft_client_secret"] = nil
        }
    }
    if w.MicrosoftEmailDomains != nil {
        structMap["microsoft_email_domains"] = w.MicrosoftEmailDomains
    }
    if w.MicrosoftEnabled != nil {
        structMap["microsoft_enabled"] = w.MicrosoftEnabled
    }
    if w.MicrosoftExpire.IsValueSet() {
        if w.MicrosoftExpire.Value() != nil {
            structMap["microsoft_expire"] = w.MicrosoftExpire.Value()
        } else {
            structMap["microsoft_expire"] = nil
        }
    }
    if w.PassphraseEnabled != nil {
        structMap["passphrase_enabled"] = w.PassphraseEnabled
    }
    if w.PassphraseExpire.IsValueSet() {
        if w.PassphraseExpire.Value() != nil {
            structMap["passphrase_expire"] = w.PassphraseExpire.Value()
        } else {
            structMap["passphrase_expire"] = nil
        }
    }
    if w.Password.IsValueSet() {
        if w.Password.Value() != nil {
            structMap["password"] = w.Password.Value()
        } else {
            structMap["password"] = nil
        }
    }
    if w.PredefinedSponsorsEnabled != nil {
        structMap["predefined_sponsors_enabled"] = w.PredefinedSponsorsEnabled
    }
    if w.PredefinedSponsorsHideEmail != nil {
        structMap["predefined_sponsors_hide_email"] = w.PredefinedSponsorsHideEmail
    }
    if w.Privacy != nil {
        structMap["privacy"] = w.Privacy
    }
    if w.PuzzelPassword != nil {
        structMap["puzzel_password"] = w.PuzzelPassword
    }
    if w.PuzzelServiceId != nil {
        structMap["puzzel_service_id"] = w.PuzzelServiceId
    }
    if w.PuzzelUsername != nil {
        structMap["puzzel_username"] = w.PuzzelUsername
    }
    if w.SmsMessageFormat != nil {
        structMap["smsMessageFormat"] = w.SmsMessageFormat
    }
    if w.SmsEnabled != nil {
        structMap["sms_enabled"] = w.SmsEnabled
    }
    if w.SmsExpire.IsValueSet() {
        if w.SmsExpire.Value() != nil {
            structMap["sms_expire"] = w.SmsExpire.Value()
        } else {
            structMap["sms_expire"] = nil
        }
    }
    if w.SmsProvider != nil {
        structMap["sms_provider"] = w.SmsProvider
    }
    if w.SponsorAutoApprove != nil {
        structMap["sponsor_auto_approve"] = w.SponsorAutoApprove
    }
    if w.SponsorEmailDomains != nil {
        structMap["sponsor_email_domains"] = w.SponsorEmailDomains
    }
    if w.SponsorEnabled != nil {
        structMap["sponsor_enabled"] = w.SponsorEnabled
    }
    if w.SponsorExpire.IsValueSet() {
        if w.SponsorExpire.Value() != nil {
            structMap["sponsor_expire"] = w.SponsorExpire.Value()
        } else {
            structMap["sponsor_expire"] = nil
        }
    }
    if w.SponsorLinkValidityDuration != nil {
        structMap["sponsor_link_validity_duration"] = w.SponsorLinkValidityDuration
    }
    if w.SponsorNotifyAll != nil {
        structMap["sponsor_notify_all"] = w.SponsorNotifyAll
    }
    if w.SponsorStatusNotify != nil {
        structMap["sponsor_status_notify"] = w.SponsorStatusNotify
    }
    if w.Sponsors != nil {
        structMap["sponsors"] = w.Sponsors
    }
    if w.SsoDefaultRole != nil {
        structMap["sso_default_role"] = w.SsoDefaultRole
    }
    if w.SsoForcedRole != nil {
        structMap["sso_forced_role"] = w.SsoForcedRole
    }
    if w.SsoIdpCert != nil {
        structMap["sso_idp_cert"] = w.SsoIdpCert
    }
    if w.SsoIdpSignAlgo != nil {
        structMap["sso_idp_sign_algo"] = w.SsoIdpSignAlgo
    }
    if w.SsoIdpSsoUrl != nil {
        structMap["sso_idp_sso_url"] = w.SsoIdpSsoUrl
    }
    if w.SsoIssuer != nil {
        structMap["sso_issuer"] = w.SsoIssuer
    }
    if w.SsoNameidFormat != nil {
        structMap["sso_nameid_format"] = w.SsoNameidFormat
    }
    if w.TelstraClientId != nil {
        structMap["telstra_client_id"] = w.TelstraClientId
    }
    if w.TelstraClientSecret != nil {
        structMap["telstra_client_secret"] = w.TelstraClientSecret
    }
    if w.TwilioAuthToken.IsValueSet() {
        if w.TwilioAuthToken.Value() != nil {
            structMap["twilio_auth_token"] = w.TwilioAuthToken.Value()
        } else {
            structMap["twilio_auth_token"] = nil
        }
    }
    if w.TwilioPhoneNumber.IsValueSet() {
        if w.TwilioPhoneNumber.Value() != nil {
            structMap["twilio_phone_number"] = w.TwilioPhoneNumber.Value()
        } else {
            structMap["twilio_phone_number"] = nil
        }
    }
    if w.TwilioSid.IsValueSet() {
        if w.TwilioSid.Value() != nil {
            structMap["twilio_sid"] = w.TwilioSid.Value()
        } else {
            structMap["twilio_sid"] = nil
        }
    }
    return structMap
}

// UnmarshalJSON implements the json.Unmarshaler interface for WlanPortal.
// It customizes the JSON unmarshaling process for WlanPortal objects.
func (w *WlanPortal) UnmarshalJSON(input []byte) error {
    var temp tempWlanPortal
    err := json.Unmarshal(input, &temp)
    if err != nil {
    	return err
    }
    additionalProperties, err := UnmarshalAdditionalProperties(input, "amazon_client_id", "amazon_client_secret", "amazon_email_domains", "amazon_enabled", "amazon_expire", "auth", "azure_client_id", "azure_client_secret", "azure_enabled", "azure_expire", "azure_tenant_id", "broadnet_password", "broadnet_sid", "broadnet_user_id", "bypass_when_cloud_down", "clickatell_api_key", "cross_site", "email_enabled", "enabled", "expire", "external_portal_url", "facebook_client_id", "facebook_client_secret", "facebook_email_domains", "facebook_enabled", "facebook_expire", "forward", "forward_url", "google_client_id", "google_client_secret", "google_email_domains", "google_enabled", "google_expire", "gupshup_password", "gupshup_userid", "microsoft_client_id", "microsoft_client_secret", "microsoft_email_domains", "microsoft_enabled", "microsoft_expire", "passphrase_enabled", "passphrase_expire", "password", "predefined_sponsors_enabled", "predefined_sponsors_hide_email", "privacy", "puzzel_password", "puzzel_service_id", "puzzel_username", "smsMessageFormat", "sms_enabled", "sms_expire", "sms_provider", "sponsor_auto_approve", "sponsor_email_domains", "sponsor_enabled", "sponsor_expire", "sponsor_link_validity_duration", "sponsor_notify_all", "sponsor_status_notify", "sponsors", "sso_default_role", "sso_forced_role", "sso_idp_cert", "sso_idp_sign_algo", "sso_idp_sso_url", "sso_issuer", "sso_nameid_format", "telstra_client_id", "telstra_client_secret", "twilio_auth_token", "twilio_phone_number", "twilio_sid")
    if err != nil {
    	return err
    }
    
    w.AdditionalProperties = additionalProperties
    w.AmazonClientId = temp.AmazonClientId
    w.AmazonClientSecret = temp.AmazonClientSecret
    w.AmazonEmailDomains = temp.AmazonEmailDomains
    w.AmazonEnabled = temp.AmazonEnabled
    w.AmazonExpire = temp.AmazonExpire
    w.Auth = temp.Auth
    w.AzureClientId = temp.AzureClientId
    w.AzureClientSecret = temp.AzureClientSecret
    w.AzureEnabled = temp.AzureEnabled
    w.AzureExpire = temp.AzureExpire
    w.AzureTenantId = temp.AzureTenantId
    w.BroadnetPassword = temp.BroadnetPassword
    w.BroadnetSid = temp.BroadnetSid
    w.BroadnetUserId = temp.BroadnetUserId
    w.BypassWhenCloudDown = temp.BypassWhenCloudDown
    w.ClickatellApiKey = temp.ClickatellApiKey
    w.CrossSite = temp.CrossSite
    w.EmailEnabled = temp.EmailEnabled
    w.Enabled = temp.Enabled
    w.Expire = temp.Expire
    w.ExternalPortalUrl = temp.ExternalPortalUrl
    w.FacebookClientId = temp.FacebookClientId
    w.FacebookClientSecret = temp.FacebookClientSecret
    w.FacebookEmailDomains = temp.FacebookEmailDomains
    w.FacebookEnabled = temp.FacebookEnabled
    w.FacebookExpire = temp.FacebookExpire
    w.Forward = temp.Forward
    w.ForwardUrl = temp.ForwardUrl
    w.GoogleClientId = temp.GoogleClientId
    w.GoogleClientSecret = temp.GoogleClientSecret
    w.GoogleEmailDomains = temp.GoogleEmailDomains
    w.GoogleEnabled = temp.GoogleEnabled
    w.GoogleExpire = temp.GoogleExpire
    w.GupshupPassword = temp.GupshupPassword
    w.GupshupUserid = temp.GupshupUserid
    w.MicrosoftClientId = temp.MicrosoftClientId
    w.MicrosoftClientSecret = temp.MicrosoftClientSecret
    w.MicrosoftEmailDomains = temp.MicrosoftEmailDomains
    w.MicrosoftEnabled = temp.MicrosoftEnabled
    w.MicrosoftExpire = temp.MicrosoftExpire
    w.PassphraseEnabled = temp.PassphraseEnabled
    w.PassphraseExpire = temp.PassphraseExpire
    w.Password = temp.Password
    w.PredefinedSponsorsEnabled = temp.PredefinedSponsorsEnabled
    w.PredefinedSponsorsHideEmail = temp.PredefinedSponsorsHideEmail
    w.Privacy = temp.Privacy
    w.PuzzelPassword = temp.PuzzelPassword
    w.PuzzelServiceId = temp.PuzzelServiceId
    w.PuzzelUsername = temp.PuzzelUsername
    w.SmsMessageFormat = temp.SmsMessageFormat
    w.SmsEnabled = temp.SmsEnabled
    w.SmsExpire = temp.SmsExpire
    w.SmsProvider = temp.SmsProvider
    w.SponsorAutoApprove = temp.SponsorAutoApprove
    w.SponsorEmailDomains = temp.SponsorEmailDomains
    w.SponsorEnabled = temp.SponsorEnabled
    w.SponsorExpire = temp.SponsorExpire
    w.SponsorLinkValidityDuration = temp.SponsorLinkValidityDuration
    w.SponsorNotifyAll = temp.SponsorNotifyAll
    w.SponsorStatusNotify = temp.SponsorStatusNotify
    w.Sponsors = temp.Sponsors
    w.SsoDefaultRole = temp.SsoDefaultRole
    w.SsoForcedRole = temp.SsoForcedRole
    w.SsoIdpCert = temp.SsoIdpCert
    w.SsoIdpSignAlgo = temp.SsoIdpSignAlgo
    w.SsoIdpSsoUrl = temp.SsoIdpSsoUrl
    w.SsoIssuer = temp.SsoIssuer
    w.SsoNameidFormat = temp.SsoNameidFormat
    w.TelstraClientId = temp.TelstraClientId
    w.TelstraClientSecret = temp.TelstraClientSecret
    w.TwilioAuthToken = temp.TwilioAuthToken
    w.TwilioPhoneNumber = temp.TwilioPhoneNumber
    w.TwilioSid = temp.TwilioSid
    return nil
}

// tempWlanPortal is a temporary struct used for validating the fields of WlanPortal.
type tempWlanPortal  struct {
    AmazonClientId              Optional[string]               `json:"amazon_client_id"`
    AmazonClientSecret          Optional[string]               `json:"amazon_client_secret"`
    AmazonEmailDomains          []string                       `json:"amazon_email_domains,omitempty"`
    AmazonEnabled               *bool                          `json:"amazon_enabled,omitempty"`
    AmazonExpire                Optional[float64]              `json:"amazon_expire"`
    Auth                        *WlanPortalAuthEnum            `json:"auth,omitempty"`
    AzureClientId               Optional[string]               `json:"azure_client_id"`
    AzureClientSecret           Optional[string]               `json:"azure_client_secret"`
    AzureEnabled                *bool                          `json:"azure_enabled,omitempty"`
    AzureExpire                 Optional[float64]              `json:"azure_expire"`
    AzureTenantId               Optional[string]               `json:"azure_tenant_id"`
    BroadnetPassword            *string                        `json:"broadnet_password,omitempty"`
    BroadnetSid                 *string                        `json:"broadnet_sid,omitempty"`
    BroadnetUserId              *string                        `json:"broadnet_user_id,omitempty"`
    BypassWhenCloudDown         *bool                          `json:"bypass_when_cloud_down,omitempty"`
    ClickatellApiKey            *string                        `json:"clickatell_api_key,omitempty"`
    CrossSite                   *bool                          `json:"cross_site,omitempty"`
    EmailEnabled                *bool                          `json:"email_enabled,omitempty"`
    Enabled                     *bool                          `json:"enabled,omitempty"`
    Expire                      *float64                       `json:"expire,omitempty"`
    ExternalPortalUrl           *string                        `json:"external_portal_url,omitempty"`
    FacebookClientId            Optional[string]               `json:"facebook_client_id"`
    FacebookClientSecret        Optional[string]               `json:"facebook_client_secret"`
    FacebookEmailDomains        []string                       `json:"facebook_email_domains,omitempty"`
    FacebookEnabled             *bool                          `json:"facebook_enabled,omitempty"`
    FacebookExpire              Optional[float64]              `json:"facebook_expire"`
    Forward                     *bool                          `json:"forward,omitempty"`
    ForwardUrl                  Optional[string]               `json:"forward_url"`
    GoogleClientId              Optional[string]               `json:"google_client_id"`
    GoogleClientSecret          Optional[string]               `json:"google_client_secret"`
    GoogleEmailDomains          []string                       `json:"google_email_domains,omitempty"`
    GoogleEnabled               *bool                          `json:"google_enabled,omitempty"`
    GoogleExpire                Optional[float64]              `json:"google_expire"`
    GupshupPassword             *string                        `json:"gupshup_password,omitempty"`
    GupshupUserid               *string                        `json:"gupshup_userid,omitempty"`
    MicrosoftClientId           Optional[string]               `json:"microsoft_client_id"`
    MicrosoftClientSecret       Optional[string]               `json:"microsoft_client_secret"`
    MicrosoftEmailDomains       []string                       `json:"microsoft_email_domains,omitempty"`
    MicrosoftEnabled            *bool                          `json:"microsoft_enabled,omitempty"`
    MicrosoftExpire             Optional[float64]              `json:"microsoft_expire"`
    PassphraseEnabled           *bool                          `json:"passphrase_enabled,omitempty"`
    PassphraseExpire            Optional[float64]              `json:"passphrase_expire"`
    Password                    Optional[string]               `json:"password"`
    PredefinedSponsorsEnabled   *bool                          `json:"predefined_sponsors_enabled,omitempty"`
    PredefinedSponsorsHideEmail *bool                          `json:"predefined_sponsors_hide_email,omitempty"`
    Privacy                     *bool                          `json:"privacy,omitempty"`
    PuzzelPassword              *string                        `json:"puzzel_password,omitempty"`
    PuzzelServiceId             *string                        `json:"puzzel_service_id,omitempty"`
    PuzzelUsername              *string                        `json:"puzzel_username,omitempty"`
    SmsMessageFormat            *string                        `json:"smsMessageFormat,omitempty"`
    SmsEnabled                  *bool                          `json:"sms_enabled,omitempty"`
    SmsExpire                   Optional[float64]              `json:"sms_expire"`
    SmsProvider                 *WlanPortalSmsProviderEnum     `json:"sms_provider,omitempty"`
    SponsorAutoApprove          *bool                          `json:"sponsor_auto_approve,omitempty"`
    SponsorEmailDomains         []string                       `json:"sponsor_email_domains,omitempty"`
    SponsorEnabled              *bool                          `json:"sponsor_enabled,omitempty"`
    SponsorExpire               Optional[float64]              `json:"sponsor_expire"`
    SponsorLinkValidityDuration *string                        `json:"sponsor_link_validity_duration,omitempty"`
    SponsorNotifyAll            *bool                          `json:"sponsor_notify_all,omitempty"`
    SponsorStatusNotify         *bool                          `json:"sponsor_status_notify,omitempty"`
    Sponsors                    map[string]string              `json:"sponsors,omitempty"`
    SsoDefaultRole              *string                        `json:"sso_default_role,omitempty"`
    SsoForcedRole               *string                        `json:"sso_forced_role,omitempty"`
    SsoIdpCert                  *string                        `json:"sso_idp_cert,omitempty"`
    SsoIdpSignAlgo              *string                        `json:"sso_idp_sign_algo,omitempty"`
    SsoIdpSsoUrl                *string                        `json:"sso_idp_sso_url,omitempty"`
    SsoIssuer                   *string                        `json:"sso_issuer,omitempty"`
    SsoNameidFormat             *WlanPortalSsoNameidFormatEnum `json:"sso_nameid_format,omitempty"`
    TelstraClientId             *string                        `json:"telstra_client_id,omitempty"`
    TelstraClientSecret         *string                        `json:"telstra_client_secret,omitempty"`
    TwilioAuthToken             Optional[string]               `json:"twilio_auth_token"`
    TwilioPhoneNumber           Optional[string]               `json:"twilio_phone_number"`
    TwilioSid                   Optional[string]               `json:"twilio_sid"`
}
