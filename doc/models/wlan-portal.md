
# Wlan Portal

Guest portal settings for the WLAN

## Structure

`WlanPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowWlanIdRoam` | `*bool` | Optional | Optional if `amazon_enabled`==`true`. Whether to allow guest to connect to other Guest WLANs (with different `WLAN.ssid`) of same org without reauthentication (disable random_mac for seamless roaming)<br><br>**Default**: `false` |
| `AmazonClientId` | `models.Optional[string]` | Optional | Optional if `amazon_enabled`==`true`. Amazon OAuth2 client id. This is optional. If not provided, it will use a default one. |
| `AmazonClientSecret` | `models.Optional[string]` | Optional | Optional if `amazon_enabled`==`true`. Amazon OAuth2 client secret. If amazon_client_id was provided, provide a corresponding value. Else leave blank. |
| `AmazonEmailDomains` | `[]string` | Optional | Optional if `amazon_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `AmazonEnabled` | `*bool` | Optional | Whether amazon is enabled as a login method<br><br>**Default**: `false` |
| `AmazonExpire` | `models.Optional[int]` | Optional | Optional if `amazon_enabled`==`true`. Interval for which guest remains authorized using amazon auth (in minutes), if not provided, uses expire` |
| `Auth` | [`*models.WlanPortalAuthEnum`](../../doc/models/wlan-portal-auth-enum.md) | Optional | authentication scheme. enum: `amazon`, `azure`, `email`, `external`, `facebook`, `google`, `microsoft`, `multi`, `none`, `password`, `sms`, `sponsor`, `sso`<br><br>**Default**: `"none"` |
| `AzureClientId` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`. Azure active directory app client id |
| `AzureClientSecret` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`. Azure active directory app client secret |
| `AzureEnabled` | `*bool` | Optional | Whether Azure Active Directory is enabled as a login method<br><br>**Default**: `false` |
| `AzureExpire` | `models.Optional[int]` | Optional | Interval for which guest remains authorized using azure auth (in minutes), if not provided, uses expire` |
| `AzureTenantId` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`. Azure active directory tenant id. |
| `BroadnetPassword` | `*string` | Optional | Required if `sms_provider`==`broadnet`. Password for the Broadnet SMS provider account |
| `BroadnetSid` | `*string` | Optional | Required if `sms_provider`==`broadnet`. SID for the Broadnet SMS provider account |
| `BroadnetUserId` | `*string` | Optional | Required if `sms_provider`==`broadnet`. User ID for the Broadnet SMS provider account |
| `BypassWhenCloudDown` | `*bool` | Optional | Whether to bypass the guest portal when cloud not reachable (and apply the default policies)<br><br>**Default**: `false` |
| `ClickatellApiKey` | `*string` | Optional | Required if `sms_provider`==`clickatell`. API key for the Clickatell SMS provider account |
| `CrossSite` | `*bool` | Optional | Whether to allow guest to roam between WLANs (with same `WLAN.ssid`, regardless of variables) of different sites of same org without reauthentication (disable random_mac for seamless roaming)<br><br>**Default**: `false` |
| `EmailEnabled` | `*bool` | Optional | Whether email (access code verification) is enabled as a login method<br><br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | Whether guest portal is enabled<br><br>**Default**: `false` |
| `Expire` | `*int` | Optional | How long to remain authorized, in minutes<br><br>**Default**: `1440` |
| `ExternalPortalUrl` | `*string` | Optional | Required if `wlan_portal_auth`==`external`. External portal URL (e.g. https://host/url) where we can append our query parameters to |
| `FacebookClientId` | `models.Optional[string]` | Optional | Required if `facebook_enabled`==`true`. Facebook OAuth2 app id. This is optional. If not provided, it will use a default one. |
| `FacebookClientSecret` | `models.Optional[string]` | Optional | Required if `facebook_enabled`==`true`. Facebook OAuth2 app secret. If facebook_client_id was provided, provide a corresponding value. Else leave blank. |
| `FacebookEmailDomains` | `[]string` | Optional | Optional if `facebook_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `FacebookEnabled` | `*bool` | Optional | Whether facebook is enabled as a login method<br><br>**Default**: `false` |
| `FacebookExpire` | `models.Optional[int]` | Optional | Optional if `facebook_enabled`==`true`. Interval for which guest remains authorized using facebook auth (in minutes), if not provided, uses expire` |
| `Forward` | `*bool` | Optional | Whether to forward the user to another URL after authorized<br><br>**Default**: `false` |
| `ForwardUrl` | `models.Optional[string]` | Optional | URL to forward the user to |
| `GoogleClientId` | `models.Optional[string]` | Optional | Google OAuth2 app id. This is optional. If not provided, it will use a default one. |
| `GoogleClientSecret` | `models.Optional[string]` | Optional | Optional if `google_enabled`==`true`. Google OAuth2 app secret. If google_client_id was provided, provide a corresponding value. Else leave blank. |
| `GoogleEmailDomains` | `[]string` | Optional | Optional if `google_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `GoogleEnabled` | `*bool` | Optional | Whether Google is enabled as login method<br><br>**Default**: `false` |
| `GoogleExpire` | `models.Optional[int]` | Optional | Optional if `google_enabled`==`true`. Interval for which guest remains authorized using Google Auth (in minutes), if not provided, uses expire` |
| `GupshupPassword` | `*string` | Optional | Required if `sms_provider`==`gupshup`. Password for the Gupshup SMS provider account |
| `GupshupUserid` | `*string` | Optional | Required if `sms_provider`==`gupshup`. User ID for the Gupshup SMS provider account |
| `MicrosoftClientId` | `models.Optional[string]` | Optional | Optional if `microsoft_enabled`==`true`. Microsoft 365 OAuth2 client id. This is optional. If not provided, it will use a default one. |
| `MicrosoftClientSecret` | `models.Optional[string]` | Optional | Optional if `microsoft_enabled`==`true`. Microsoft 365 OAuth2 client secret. If microsoft_client_id was provided, provide a corresponding value. Else leave blank. |
| `MicrosoftEmailDomains` | `[]string` | Optional | Optional if `microsoft_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `MicrosoftEnabled` | `*bool` | Optional | Whether microsoft 365 is enabled as a login method<br><br>**Default**: `false` |
| `MicrosoftExpire` | `models.Optional[int]` | Optional | Optional if `microsoft_enabled`==`true`. Interval for which guest remains authorized using microsoft auth (in minutes), if not provided, uses expire` |
| `PassphraseEnabled` | `*bool` | Optional | Whether password is enabled<br><br>**Default**: `false` |
| `PassphraseExpire` | `models.Optional[int]` | Optional | Optional if `passphrase_enabled`==`true`. Interval for which guest remains authorized using passphrase auth (in minutes), if not provided, uses `expire` |
| `Password` | `models.Optional[string]` | Optional | Required if `passphrase_enabled`==`true`. Passphrase guests must enter when passphrase authentication is enabled |
| `PredefinedSponsorsEnabled` | `*bool` | Optional | Whether to show list of sponsor emails mentioned in `sponsors` object as a dropdown. If both `sponsor_notify_all` and `predefined_sponsors_enabled` are false, behavior is acc to `sponsor_email_domains`<br><br>**Default**: `true` |
| `PredefinedSponsorsHideEmail` | `*bool` | Optional | Whether to hide sponsor’s email from list of sponsors<br><br>**Default**: `false` |
| `Privacy` | `*bool` | Optional | Whether to show the privacy policy in the WLAN guest portal<br><br>**Default**: `false` |
| `PuzzelPassword` | `*string` | Optional | Required if `sms_provider`==`puzzel`. Password for the Puzzel SMS provider account |
| `PuzzelServiceId` | `*string` | Optional | Required if `sms_provider`==`puzzel`. Service ID for the Puzzel SMS provider account |
| `PuzzelUsername` | `*string` | Optional | Required if `sms_provider`==`puzzel`. Username for the Puzzel SMS provider account |
| `SmsMessageFormat` | `*string` | Optional | Optional if `sms_enabled`==`true`. SMS Message format<br><br>**Default**: `"Code {{code}} expires in {{duration}} minutes."` |
| `SmsEnabled` | `*bool` | Optional | Whether sms is enabled as a login method<br><br>**Default**: `false` |
| `SmsExpire` | `models.Optional[int]` | Optional | Optional if `sms_enabled`==`true`. Interval for which guest remains authorized using sms auth (in minutes), if not provided, uses expire` |
| `SmsProvider` | [`*models.WlanPortalSmsProviderEnum`](../../doc/models/wlan-portal-sms-provider-enum.md) | Optional | Optional if `sms_enabled`==`true`. enum: `broadnet`, `clickatell`, `gupshup`, `manual`, `puzzel`, `smsglobal`, `telstra`, `twilio`<br><br>**Default**: `"manual"` |
| `SmsglobalApiKey` | `*string` | Optional | Required if `sms_provider`==`smsglobal`, Client API Key |
| `SmsglobalApiSecret` | `*string` | Optional | Required if `sms_provider`==`smsglobal`, Client secret |
| `SmsglobalSender` | `*string` | Optional | Optional sender's number or sender ID for SMSGlobal. If not provided, uses the default number associated with the account |
| `SponsorAutoApprove` | `*bool` | Optional | Optional if `sponsor_enabled`==`true`. Whether to automatically approve guest and allow sponsor to revoke guest access, needs predefined_sponsors_enabled enabled and sponsor_notify_all disabled<br><br>**Default**: `false` |
| `SponsorEmailDomains` | `[]string` | Optional | List of domain allowed for sponsor email. Required if `sponsor_enabled` is `true` and `sponsors` is empty. |
| `SponsorEnabled` | `*bool` | Optional | Whether sponsor is enabled<br><br>**Default**: `false` |
| `SponsorExpire` | `models.Optional[int]` | Optional | Optional if `sponsor_enabled`==`true`. Interval for which guest remains authorized using sponsor auth (in minutes), if not provided, uses expire` |
| `SponsorLinkValidityDuration` | [`*models.SponsorLinkValidityDuration`](../../doc/models/containers/sponsor-link-validity-duration.md) | Optional | Optional if `sponsor_enabled`==`true`. How long to remain valid sponsored guest request approve/deny link received in email, in minutes. Value is between 5 and 60. |
| `SponsorNotifyAll` | `*bool` | Optional | Optional if `sponsor_enabled`==`true`. whether to notify all sponsors that are mentioned in `sponsors` object. Both `sponsor_notify_all` and `predefined_sponsors_enabled` should be true in order to notify sponsors. If true, email sent to 10 sponsors in no particular order.<br><br>**Default**: `false` |
| `SponsorStatusNotify` | `*bool` | Optional | Optional if `sponsor_enabled`==`true`. If enabled, guest will get email about sponsor's action (approve/deny)<br><br>**Default**: `false` |
| `Sponsors` | [`*models.WlanPortalSponsors`](../../doc/models/containers/wlan-portal-sponsors.md) | Optional | Object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty. Property key is the sponsor email, Property value is the sponsor name. List of email allowed for backward compatibility |
| `SsoDefaultRole` | `*string` | Optional | Optional if `wlan_portal_auth`==`sso`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched |
| `SsoForcedRole` | `*string` | Optional | Optional if `wlan_portal_auth`==`sso`. Role assigned to authenticated users when guest SSO is used |
| `SsoIdpCert` | `*string` | Optional | Required if `wlan_portal_auth`==`sso`. IDP Cert (used to verify the signed response) |
| `SsoIdpSignAlgo` | [`*models.WlanPortalIdpSignAlgoEnum`](../../doc/models/wlan-portal-idp-sign-algo-enum.md) | Optional | Optional if `wlan_portal_auth`==`sso`, Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`<br><br>**Default**: `"sha1"` |
| `SsoIdpSsoUrl` | `*string` | Optional | Required if `wlan_portal_auth`==`sso`, IDP Single-Sign-On URL |
| `SsoIssuer` | `*string` | Optional | Required if `wlan_portal_auth`==`sso`, IDP issuer URL |
| `SsoNameidFormat` | [`*models.WlanPortalSsoNameidFormatEnum`](../../doc/models/wlan-portal-sso-nameid-format-enum.md) | Optional | Optional if `wlan_portal_auth`==`sso`. enum: `email`, `unspecified`<br><br>**Default**: `"email"` |
| `TelstraClientId` | `*string` | Optional | Required if `sms_provider`==`telstra`, Client ID provided by Telstra |
| `TelstraClientSecret` | `*string` | Optional | Required if `sms_provider`==`telstra`, Client secret provided by Telstra |
| `TwilioAuthToken` | `models.Optional[string]` | Optional | Required if `sms_provider`==`twilio`, Auth token account with twilio account |
| `TwilioPhoneNumber` | `models.Optional[string]` | Optional | Required if `sms_provider`==`twilio`, Twilio phone number associated with the account. See example for accepted format. |
| `TwilioSid` | `models.Optional[string]` | Optional | Required if `sms_provider`==`twilio`, Account SID provided by Twilio |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanPortal := models.WlanPortal{
        AllowWlanIdRoam:             models.ToPointer(false),
        AmazonClientId:              models.NewOptional(models.ToPointer("amazon_client_id8")),
        AmazonClientSecret:          models.NewOptional(models.ToPointer("amazon_client_secret2")),
        AmazonEmailDomains:          nil,
        AmazonEnabled:               models.ToPointer(false),
        Auth:                        models.ToPointer(models.WlanPortalAuthEnum_NONE),
        AzureEnabled:                models.ToPointer(false),
        BroadnetPassword:            models.ToPointer("password"),
        BroadnetSid:                 models.ToPointer("MIST"),
        BroadnetUserId:              models.ToPointer("juniper"),
        BypassWhenCloudDown:         models.ToPointer(false),
        CrossSite:                   models.ToPointer(false),
        EmailEnabled:                models.ToPointer(false),
        Enabled:                     models.ToPointer(false),
        Expire:                      models.ToPointer(1440),
        FacebookEmailDomains:        nil,
        FacebookEnabled:             models.ToPointer(false),
        Forward:                     models.ToPointer(false),
        ForwardUrl:                  models.NewOptional(models.ToPointer("https://abc.com/promotions")),
        GoogleEmailDomains:          []string{
            "mydomain.edu",
            "mydomain.org",
        },
        GoogleEnabled:               models.ToPointer(false),
        MicrosoftEmailDomains:       nil,
        MicrosoftEnabled:            models.ToPointer(false),
        PassphraseEnabled:           models.ToPointer(false),
        Password:                    models.NewOptional(models.ToPointer("let me in")),
        PredefinedSponsorsEnabled:   models.ToPointer(true),
        PredefinedSponsorsHideEmail: models.ToPointer(false),
        Privacy:                     models.ToPointer(false),
        SmsMessageFormat:            models.ToPointer("Code {{code}} expires in {{duration}} minutes."),
        SmsEnabled:                  models.ToPointer(false),
        SmsProvider:                 models.ToPointer(models.WlanPortalSmsProviderEnum_TWILIO),
        SponsorAutoApprove:          models.ToPointer(false),
        SponsorEmailDomains:         []string{
            "reserved.net",
            "reserved.org",
        },
        SponsorEnabled:              models.ToPointer(false),
        SponsorNotifyAll:            models.ToPointer(false),
        SponsorStatusNotify:         models.ToPointer(false),
        SsoIdpSignAlgo:              models.ToPointer(models.WlanPortalIdpSignAlgoEnum_SHA1),
        SsoNameidFormat:             models.ToPointer(models.WlanPortalSsoNameidFormatEnum_EMAIL),
        TwilioAuthToken:             models.NewOptional(models.ToPointer("af9dac44c344a875ab5d31cb7abcdefg")),
        TwilioPhoneNumber:           models.NewOptional(models.ToPointer("+18548888888")),
        TwilioSid:                   models.NewOptional(models.ToPointer("af9dac44c344a875ab5d31cb7abcdefg")),
    }

}
```

