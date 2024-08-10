
# Wlan Portal

portal wlan settings

## Structure

`WlanPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmazonClientId` | `models.Optional[string]` | Optional | amazon OAuth2 client id. This is optional. If not provided, it will use a default one. |
| `AmazonClientSecret` | `models.Optional[string]` | Optional | amazon OAuth2 client secret. If amazon_client_id was provided, provide a correspoinding value. Else leave blank. |
| `AmazonEmailDomains` | `[]string` | Optional | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `AmazonEnabled` | `*bool` | Optional | whether amazon is enabled as a login method<br>**Default**: `false` |
| `AmazonExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using amazon auth (in minutes), if not provided, uses expire` |
| `Auth` | [`*models.WlanPortalAuthEnum`](../../doc/models/wlan-portal-auth-enum.md) | Optional | authentication scheme. enum: `external`, `none`, `sso`<br>**Default**: `"none"` |
| `AzureClientId` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`.<br>Azure active directory app client id |
| `AzureClientSecret` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`.<br>Azure active directory app client secret |
| `AzureEnabled` | `*bool` | Optional | whether Azure Active Directory is enabled as a login method<br>**Default**: `false` |
| `AzureExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using azure auth (in minutes), if not provided, uses expire` |
| `AzureTenantId` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`.<br>Azure active directory tenant id. |
| `BroadnetPassword` | `*string` | Optional | when `sms_provider`==`broadnet` |
| `BroadnetSid` | `*string` | Optional | when `sms_provider`==`broadnet`<br>**Default**: `"MIST"` |
| `BroadnetUserId` | `*string` | Optional | when `sms_provider`==`broadnet` |
| `BypassWhenCloudDown` | `*bool` | Optional | whether to bypass the guest portal when cloud not reachable (and apply the default policies)<br>**Default**: `false` |
| `ClickatellApiKey` | `*string` | Optional | when `sms_provider`==`clickatell` |
| `CrossSite` | `*bool` | Optional | whether to allow guest to roam between WLANs (with same `WLAN.ssid`, regardless of variables) of different sites of same org without reauthentication (disable random_mac for seamless roaming)<br>**Default**: `false` |
| `EmailEnabled` | `*bool` | Optional | whether email (access code verification) is enabled as a login method |
| `Enabled` | `*bool` | Optional | whether guest portal is enabled<br>**Default**: `false` |
| `Expire` | `*float64` | Optional | how long to remain authorized, in minutes<br>**Default**: `1440` |
| `ExternalPortalUrl` | `*string` | Optional | external portal URL (e.g. https://host/url) where we can append our query parameters to |
| `FacebookClientId` | `models.Optional[string]` | Optional | Required if `facebook_enabled`==`true`.<br>Facebook OAuth2 app id. This is optional. If not provided, it will use a default one. |
| `FacebookClientSecret` | `models.Optional[string]` | Optional | Required if `facebook_enabled`==`true`.<br>Facebook OAuth2 app secret. If facebook_client_id was provided, provide a correspoinding value. Else leave blank. |
| `FacebookEmailDomains` | `[]string` | Optional | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `FacebookEnabled` | `*bool` | Optional | whether facebook is enabled as a login method<br>**Default**: `false` |
| `FacebookExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using facebook auth (in minutes), if not provided, uses expire` |
| `Forward` | `*bool` | Optional | whether to forward the user to another URL after authorized<br>**Default**: `false` |
| `ForwardUrl` | `models.Optional[string]` | Optional | the URL to forward the user to |
| `GoogleClientId` | `models.Optional[string]` | Optional | Google OAuth2 app id. This is optional. If not provided, it will use a default one. |
| `GoogleClientSecret` | `models.Optional[string]` | Optional | Google OAuth2 app secret. If google_client_id was provided, provide a correspoinding value. Else leave blank. |
| `GoogleEmailDomains` | `[]string` | Optional | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `GoogleEnabled` | `*bool` | Optional | whether google is enabled as login method<br>**Default**: `false` |
| `GoogleExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using google auth (in minutes), if not provided, uses expire` |
| `GupshupPassword` | `*string` | Optional | when `sms_provider`==`gupshup` |
| `GupshupUserid` | `*string` | Optional | when `sms_provider`==`gupshup` |
| `MicrosoftClientId` | `models.Optional[string]` | Optional | microsoft 365 OAuth2 client id. This is optional. If not provided, it will use a default one. |
| `MicrosoftClientSecret` | `models.Optional[string]` | Optional | microsoft 365 OAuth2 client secret. If microsoft_client_id was provided, provide a correspoinding value. Else leave blank. |
| `MicrosoftEmailDomains` | `[]string` | Optional | Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `MicrosoftEnabled` | `*bool` | Optional | whether microsoft 365 is enabled as a login method<br>**Default**: `false` |
| `MicrosoftExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using microsoft auth (in minutes), if not provided, uses expire` |
| `PassphraseEnabled` | `*bool` | Optional | whether password is enabled<br>**Default**: `false` |
| `PassphraseExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using passphrase auth (in minutes), if not provided, uses `expire` |
| `Password` | `models.Optional[string]` | Optional | passphrase |
| `PredefinedSponsorsEnabled` | `*bool` | Optional | whether to show list of sponsor emails mentioned in `sponsors` object as a dropdown. If both `sponsor_notify_all` and `predefined_sponsors_enabled` are false, behaviour is acc to `sponsor_email_domains`<br>**Default**: `true` |
| `PredefinedSponsorsHideEmail` | `*bool` | Optional | whether to hide sponsor’s email from list of sponsors<br>**Default**: `false` |
| `Privacy` | `*bool` | Optional | **Default**: `true` |
| `PuzzelPassword` | `*string` | Optional | when `sms_provider`==`puzzel` |
| `PuzzelServiceId` | `*string` | Optional | when `sms_provider`==`puzzel` |
| `PuzzelUsername` | `*string` | Optional | when `sms_provider`==`puzzel` |
| `SmsMessageFormat` | `*string` | Optional | - |
| `SmsEnabled` | `*bool` | Optional | whether sms is enabled as a login method<br>**Default**: `false` |
| `SmsExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using sms auth (in minutes), if not provided, uses expire` |
| `SmsProvider` | [`*models.WlanPortalSmsProviderEnum`](../../doc/models/wlan-portal-sms-provider-enum.md) | Optional | enum: `broadnet`, `clickatell`, `gupshup`, `manual`, `puzzel`, `telstra`, `twilio`<br>**Default**: `"manual"` |
| `SponsorAutoApprove` | `*bool` | Optional | whether to automatically approve guest and allow sponsor to revoke guest access, needs predefined_sponsors_enabled enabled and sponsor_notify_all disabled<br>**Default**: `false` |
| `SponsorEmailDomains` | `[]string` | Optional | list of domain allowed for sponsor email. Required if `sponsor_enabled` is `true` and `sponsors` is empty. |
| `SponsorEnabled` | `*bool` | Optional | whether sponsor is enabled<br>**Default**: `false` |
| `SponsorExpire` | `models.Optional[float64]` | Optional | interval for which guest remains authorized using sponsor auth (in minutes), if not provided, uses expire` |
| `SponsorLinkValidityDuration` | `*int` | Optional | how long to remain valid sponsored guest request approve/deny link received in email, in minutes.<br>**Default**: `60` |
| `SponsorNotifyAll` | `*bool` | Optional | whether to notify all sponsors that are mentioned in `sponsors` object. Both `sponsor_notify_all` and `predefined_sponsors_enabled` should be true in order to notify sponsors. If true, email sent to 10 sponsors in no particular order.<br>**Default**: `false` |
| `SponsorStatusNotify` | `*bool` | Optional | if enabled, guest will get email about sponsor's action (approve/deny)<br>**Default**: `false` |
| `Sponsors` | `map[string]string` | Optional | object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty.<br>Property key is the sponsor email, Property value is the sponsor name |
| `SsoDefaultRole` | `*string` | Optional | default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched |
| `SsoForcedRole` | `*string` | Optional | - |
| `SsoIdpCert` | `*string` | Optional | IDP Cert (used to verify the signed response) |
| `SsoIdpSignAlgo` | `*string` | Optional | signing algorithm for SAML Assertion |
| `SsoIdpSsoUrl` | `*string` | Optional | IDP Single-Sign-On URL |
| `SsoIssuer` | `*string` | Optional | IDP issuer URL |
| `SsoNameidFormat` | [`*models.WlanPortalSsoNameidFormatEnum`](../../doc/models/wlan-portal-sso-nameid-format-enum.md) | Optional | enum: `email`, `unspecified`<br>**Default**: `"email"` |
| `TelstraClientId` | `*string` | Optional | when `sms_provider`==`telstra`, Client ID provided by Telstra |
| `TelstraClientSecret` | `*string` | Optional | when `sms_provider`==`telstra`, Client secret provided by Telstra |
| `TwilioAuthToken` | `models.Optional[string]` | Optional | when `sms_provider`==`twilio`, Auth token account with twilio account |
| `TwilioPhoneNumber` | `models.Optional[string]` | Optional | when `sms_provider`==`twilio`, Twilio phone number associated with the account. See example for accepted format. |
| `TwilioSid` | `models.Optional[string]` | Optional | when `sms_provider`==`twilio`, Account SID provided by Twilio |

## Example (as JSON)

```json
{
  "amazon_email_domains": [
    "amazon_email_domains4",
    "amazon_email_domains5",
    "amazon_email_domains6"
  ],
  "amazon_enabled": false,
  "auth": "none",
  "azure_enabled": false,
  "broadnet_password": "password",
  "broadnet_sid": "MIST",
  "broadnet_user_id": "juniper",
  "bypass_when_cloud_down": false,
  "cross_site": false,
  "enabled": false,
  "expire": 1440,
  "facebook_email_domains": null,
  "facebook_enabled": false,
  "forward": false,
  "forward_url": "http://abc.com/promotions",
  "google_email_domains": [
    "mydomain.edu",
    "mydomain.org"
  ],
  "google_enabled": false,
  "microsoft_email_domains": null,
  "microsoft_enabled": false,
  "passphrase_enabled": false,
  "password": "let me in",
  "predefined_sponsors_enabled": true,
  "predefined_sponsors_hide_email": false,
  "privacy": true,
  "sms_enabled": false,
  "sms_provider": "twilio",
  "sponsor_auto_approve": false,
  "sponsor_email_domains": [
    "reserved.net",
    "reserved.org"
  ],
  "sponsor_enabled": false,
  "sponsor_link_validity_duration": 30,
  "sponsor_notify_all": false,
  "sponsor_status_notify": false,
  "sponsors": {
    "sponsor1@company.com": "FirstName1 LastName1",
    "sponsor2@company.com": "FirstName2 LastName2"
  },
  "sso_nameid_format": "email",
  "twilio_auth_token": "af9dac44c344a875ab5d31cb7abcdefg",
  "twilio_phone_number": "+18548888888",
  "twilio_sid": "AC72ec6ba0ec5af30e6731c5e47abcdefgh",
  "amazon_client_id": "amazon_client_id4",
  "amazon_client_secret": "amazon_client_secret0",
  "amazon_expire": 137.94
}
```

