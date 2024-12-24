
# Wlan Portal

portal wlan settings

*This model accepts additional fields of type interface{}.*

## Structure

`WlanPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllowWlanIdRoam` | `*bool` | Optional | Optional if `amazon_enabled`==`true`. Whether to allow guest to connect to other Guest WLANs (with different `WLAN.ssid`) of same org without reauthentication (disable random_mac for seamless roaming)<br>**Default**: `false` |
| `AmazonClientId` | `models.Optional[string]` | Optional | Optional if `amazon_enabled`==`true`. Amazon OAuth2 client id. This is optional. If not provided, it will use a default one. |
| `AmazonClientSecret` | `models.Optional[string]` | Optional | Optional if `amazon_enabled`==`true`. Amazon OAuth2 client secret. If amazon_client_id was provided, provide a correspoinding value. Else leave blank. |
| `AmazonEmailDomains` | `[]string` | Optional | Optional if `amazon_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `AmazonEnabled` | `*bool` | Optional | whether amazon is enabled as a login method<br>**Default**: `false` |
| `AmazonExpire` | `models.Optional[int]` | Optional | Optional if `amazon_enabled`==`true`. Interval for which guest remains authorized using amazon auth (in minutes), if not provided, uses expire` |
| `Auth` | [`*models.WlanPortalAuthEnum`](../../doc/models/wlan-portal-auth-enum.md) | Optional | authentication scheme. enum: `amazon`, `azure`, `email`, `external`, `facebook`, `google`, `microsoft`, `multi`, `none`, `password`, `sponsor`, `sso`<br>**Default**: `"none"` |
| `AzureClientId` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`. Azure active directory app client id |
| `AzureClientSecret` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`. Azure active directory app client secret |
| `AzureEnabled` | `*bool` | Optional | whether Azure Active Directory is enabled as a login method<br>**Default**: `false` |
| `AzureExpire` | `models.Optional[int]` | Optional | interval for which guest remains authorized using azure auth (in minutes), if not provided, uses expire` |
| `AzureTenantId` | `models.Optional[string]` | Optional | Required if `azure_enabled`==`true`. Azure active directory tenant id. |
| `BroadnetPassword` | `*string` | Optional | Required if `sms_provider`==`broadnet` |
| `BroadnetSid` | `*string` | Optional | Required if `sms_provider`==`broadnet` |
| `BroadnetUserId` | `*string` | Optional | Required if `sms_provider`==`broadnet` |
| `BypassWhenCloudDown` | `*bool` | Optional | whether to bypass the guest portal when cloud not reachable (and apply the default policies)<br>**Default**: `false` |
| `ClickatellApiKey` | `*string` | Optional | Required if `sms_provider`==`clickatell` |
| `CrossSite` | `*bool` | Optional | whether to allow guest to roam between WLANs (with same `WLAN.ssid`, regardless of variables) of different sites of same org without reauthentication (disable random_mac for seamless roaming)<br>**Default**: `false` |
| `EmailEnabled` | `*bool` | Optional | whether email (access code verification) is enabled as a login method<br>**Default**: `false` |
| `Enabled` | `*bool` | Optional | whether guest portal is enabled<br>**Default**: `false` |
| `Expire` | `*int` | Optional | how long to remain authorized, in minutes<br>**Default**: `1440` |
| `ExternalPortalUrl` | `*string` | Optional | Required if `wlan_portal_auth`==`external`. External portal URL (e.g. https://host/url) where we can append our query parameters to |
| `FacebookClientId` | `models.Optional[string]` | Optional | Required if `facebook_enabled`==`true`. Facebook OAuth2 app id. This is optional. If not provided, it will use a default one. |
| `FacebookClientSecret` | `models.Optional[string]` | Optional | Required if `facebook_enabled`==`true`. Facebook OAuth2 app secret. If facebook_client_id was provided, provide a correspoinding value. Else leave blank. |
| `FacebookEmailDomains` | `[]string` | Optional | Optional if `facebook_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `FacebookEnabled` | `*bool` | Optional | whether facebook is enabled as a login method<br>**Default**: `false` |
| `FacebookExpire` | `models.Optional[int]` | Optional | Optional if `facebook_enabled`==`true`. Interval for which guest remains authorized using facebook auth (in minutes), if not provided, uses expire` |
| `Forward` | `*bool` | Optional | whether to forward the user to another URL after authorized<br>**Default**: `false` |
| `ForwardUrl` | `models.Optional[string]` | Optional | the URL to forward the user to |
| `GoogleClientId` | `models.Optional[string]` | Optional | Google OAuth2 app id. This is optional. If not provided, it will use a default one. |
| `GoogleClientSecret` | `models.Optional[string]` | Optional | Optional if `google_enabled`==`true`. Google OAuth2 app secret. If google_client_id was provided, provide a correspoinding value. Else leave blank. |
| `GoogleEmailDomains` | `[]string` | Optional | Optional if `google_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `GoogleEnabled` | `*bool` | Optional | whether google is enabled as login method<br>**Default**: `false` |
| `GoogleExpire` | `models.Optional[int]` | Optional | Optional if `google_enabled`==`true`. Interval for which guest remains authorized using google auth (in minutes), if not provided, uses expire` |
| `GupshupPassword` | `*string` | Optional | Required if `sms_provider`==`gupshup` |
| `GupshupUserid` | `*string` | Optional | Required if `sms_provider`==`gupshup` |
| `MicrosoftClientId` | `models.Optional[string]` | Optional | Optional if `microsoft_enabled`==`true`. Microsoft 365 OAuth2 client id. This is optional. If not provided, it will use a default one. |
| `MicrosoftClientSecret` | `models.Optional[string]` | Optional | Optional if `microsoft_enabled`==`true`. Microsoft 365 OAuth2 client secret. If microsoft_client_id was provided, provide a correspoinding value. Else leave blank. |
| `MicrosoftEmailDomains` | `[]string` | Optional | Optional if `microsoft_enabled`==`true`. Matches authenticated user email against provided domains. If null or [], all authenticated emails will be allowed. |
| `MicrosoftEnabled` | `*bool` | Optional | whether microsoft 365 is enabled as a login method<br>**Default**: `false` |
| `MicrosoftExpire` | `models.Optional[int]` | Optional | Optional if `microsoft_enabled`==`true`. Interval for which guest remains authorized using microsoft auth (in minutes), if not provided, uses expire` |
| `PassphraseEnabled` | `*bool` | Optional | Whether password is enabled<br>**Default**: `false` |
| `PassphraseExpire` | `models.Optional[int]` | Optional | Optional if `passphrase_enabled`==`true`. Interval for which guest remains authorized using passphrase auth (in minutes), if not provided, uses `expire` |
| `Password` | `models.Optional[string]` | Optional | Required if `passphrase_enabled`==`true`. |
| `PredefinedSponsorsEnabled` | `*bool` | Optional | whether to show list of sponsor emails mentioned in `sponsors` object as a dropdown. If both `sponsor_notify_all` and `predefined_sponsors_enabled` are false, behaviour is acc to `sponsor_email_domains`<br>**Default**: `true` |
| `PredefinedSponsorsHideEmail` | `*bool` | Optional | whether to hide sponsor’s email from list of sponsors<br>**Default**: `false` |
| `Privacy` | `*bool` | Optional | **Default**: `false` |
| `PuzzelPassword` | `*string` | Optional | Required if `sms_provider`==`puzzel` |
| `PuzzelServiceId` | `*string` | Optional | Required if `sms_provider`==`puzzel` |
| `PuzzelUsername` | `*string` | Optional | Required if `sms_provider`==`puzzel` |
| `SmsMessageFormat` | `*string` | Optional | Optional if `sms_enabled`==`true`. SMS Message format<br>**Default**: `"Code {{code}} expires in {{duration}} minutes."` |
| `SmsEnabled` | `*bool` | Optional | whether sms is enabled as a login method<br>**Default**: `false` |
| `SmsExpire` | `models.Optional[int]` | Optional | Optional if `sms_enabled`==`true`. Interval for which guest remains authorized using sms auth (in minutes), if not provided, uses expire` |
| `SmsProvider` | [`*models.WlanPortalSmsProviderEnum`](../../doc/models/wlan-portal-sms-provider-enum.md) | Optional | Optioanl if `sms_enabled`==`true`. enum: `broadnet`, `clickatell`, `gupshup`, `manual`, `puzzel`, `telstra`, `twilio`<br>**Default**: `"manual"` |
| `SponsorAutoApprove` | `*bool` | Optional | Optional if `sponsor_enabled`==`true`. Whether to automatically approve guest and allow sponsor to revoke guest access, needs predefined_sponsors_enabled enabled and sponsor_notify_all disabled<br>**Default**: `false` |
| `SponsorEmailDomains` | `[]string` | Optional | list of domain allowed for sponsor email. Required if `sponsor_enabled` is `true` and `sponsors` is empty. |
| `SponsorEnabled` | `*bool` | Optional | whether sponsor is enabled<br>**Default**: `false` |
| `SponsorExpire` | `models.Optional[int]` | Optional | Optional if `sponsor_enabled`==`true`. Interval for which guest remains authorized using sponsor auth (in minutes), if not provided, uses expire` |
| `SponsorLinkValidityDuration` | `*string` | Optional | Optional if `sponsor_enabled`==`true`. How long to remain valid sponsored guest request approve/deny link received in email, in minutes.<br>**Default**: `"60"` |
| `SponsorNotifyAll` | `*bool` | Optional | Optional if `sponsor_enabled`==`true`. whether to notify all sponsors that are mentioned in `sponsors` object. Both `sponsor_notify_all` and `predefined_sponsors_enabled` should be true in order to notify sponsors. If true, email sent to 10 sponsors in no particular order.<br>**Default**: `false` |
| `SponsorStatusNotify` | `*bool` | Optional | Optional if `sponsor_enabled`==`true`. If enabled, guest will get email about sponsor's action (approve/deny)<br>**Default**: `false` |
| `Sponsors` | [`*models.WlanPortalSponsors`](../../doc/models/containers/wlan-portal-sponsors.md) | Optional | object of allowed sponsors email with name. Required if `sponsor_enabled` is `true` and `sponsor_email_domains` is empty. Property key is the sponsor email, Property value is the sponsor name. List of email allowed for backward compatibility |
| `SsoDefaultRole` | `*string` | Optional | Optionl if `wlan_portal_auth`==`sso`, default role to assign if there’s no match. By default, an assertion is treated as invalid when there’s no role matched |
| `SsoForcedRole` | `*string` | Optional | Optionl if `wlan_portal_auth`==`sso` |
| `SsoIdpCert` | `*string` | Optional | Required if `wlan_portal_auth`==`sso`. IDP Cert (used to verify the signed response) |
| `SsoIdpSignAlgo` | [`*models.WlanPortalIdpSignAlgoEnum`](../../doc/models/wlan-portal-idp-sign-algo-enum.md) | Optional | Optioanl if `wlan_portal_auth`==`sso`, Signing algorithm for SAML Assertion. enum: `sha1`, `sha256`, `sha384`, `sha512`<br>**Default**: `"sha1"` |
| `SsoIdpSsoUrl` | `*string` | Optional | Required if `wlan_portal_auth`==`sso`, IDP Single-Sign-On URL |
| `SsoIssuer` | `*string` | Optional | Required if `wlan_portal_auth`==`sso`, IDP issuer URL |
| `SsoNameidFormat` | [`*models.WlanPortalSsoNameidFormatEnum`](../../doc/models/wlan-portal-sso-nameid-format-enum.md) | Optional | Optional if `wlan_portal_auth`==`sso`. enum: `email`, `unspecified`<br>**Default**: `"email"` |
| `TelstraClientId` | `*string` | Optional | Required if `sms_provider`==`telstra`, Client ID provided by Telstra |
| `TelstraClientSecret` | `*string` | Optional | Required if `sms_provider`==`telstra`, Client secret provided by Telstra |
| `TwilioAuthToken` | `models.Optional[string]` | Optional | Required if `sms_provider`==`twilio`, Auth token account with twilio account |
| `TwilioPhoneNumber` | `models.Optional[string]` | Optional | Required if `sms_provider`==`twilio`, Twilio phone number associated with the account. See example for accepted format. |
| `TwilioSid` | `models.Optional[string]` | Optional | Required if `sms_provider`==`twilio`, Account SID provided by Twilio |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "allow_wlan_id_roam": false,
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
  "email_enabled": false,
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
  "privacy": false,
  "smsMessageFormat": "Code {{code}} expires in {{duration}} minutes.",
  "sms_enabled": false,
  "sms_provider": "twilio",
  "sponsor_auto_approve": false,
  "sponsor_email_domains": [
    "reserved.net",
    "reserved.org"
  ],
  "sponsor_enabled": false,
  "sponsor_link_validity_duration": "30",
  "sponsor_notify_all": false,
  "sponsor_status_notify": false,
  "sso_idp_sign_algo": "sha1",
  "sso_nameid_format": "email",
  "twilio_auth_token": "af9dac44c344a875ab5d31cb7abcdefg",
  "twilio_phone_number": "+18548888888",
  "twilio_sid": "AC72ec6ba0ec5af30e6731c5e47abcdefgh",
  "amazon_client_id": "amazon_client_id4",
  "amazon_client_secret": "amazon_client_secret0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

