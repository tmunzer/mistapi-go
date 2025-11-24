
# Nac Portal

*This model accepts additional fields of type interface{}.*

## Structure

`NacPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessType` | [`*models.NacPortalAccessTypeEnum`](../../doc/models/nac-portal-access-type-enum.md) | Optional | if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`<br><br>**Default**: `"wireless"` |
| `AdditionalCacerts` | `[]string` | Optional | Optional list of additional CA certificates to be used |
| `BgImageUrl` | `*string` | Optional | Background image |
| `CertExpireTime` | `*int` | Optional | In days |
| `EapType` | [`*models.NacPortalEapTypeEnum`](../../doc/models/nac-portal-eap-type-enum.md) | Optional | enum: `wpa2`, `wpa3`<br><br>**Default**: `"wpa2"` |
| `EnableTelemetry` | `*bool` | Optional | Model, version, fingering, events (connecting, disconnect, roaming), which ap |
| `ExpiryNotificationTime` | `*int` | Optional | In days |
| `Name` | `*string` | Optional | - |
| `NotifyExpiry` | `*bool` | Optional | phase 2 |
| `Portal` | [`*models.NacPortalGuestPortal`](../../doc/models/nac-portal-guest-portal.md) | Optional | Guest portal configuration when `type`==`guest_portal`. If<br><br>* `auth`==`none`, the user is presented with a terms of service and can click and continue.<br>* `auth`==`external`, the user is redirected to an external URL for authentication.<br>* `auth`==`multi`, the user is presented with a choice of authentication methods:<br>  - social logins: facebook / google / amazon / microsoft / azure<br>  - sponsor<br>  - sms: supported provider: twillio<br>  - email<br>  - sso<br>  - userpass: pre created guest list |
| `PortalAuthorizeJwtSecret` | `*string` | Optional | If `type`==`guest_portal` and `auth`==`external`, the `portal_authorize_jwt_secret` will be generated |
| `PortalAuthorizeUrl` | `*string` | Optional | If `type`==`guest_portal` and `auth`==`external`, the `portal_authorize_url` will be generated |
| `PortalSsoUrl` | `*string` | Optional | If `type`==`guest_portal` or `type`==`guest_admin` and ans SSO is enabled, the `portal_sso_url` will be generated (which needs to be configured in your IDP |
| `Ssid` | `*string` | Optional | - |
| `Sso` | [`*models.NacPortalSso`](../../doc/models/nac-portal-sso.md) | Optional | - |
| `TemplateUrl` | `*string` | Optional | - |
| `ThumbnailUrl` | `*string` | Optional | - |
| `Tos` | `*string` | Optional | - |
| `Type` | [`*models.NacPortalTypeEnum`](../../doc/models/nac-portal-type-enum.md) | Optional | enum:<br><br>* `guest_admin`: NAC-Based Portal Admin for Pre Created Guest Authentication<br>* `guest_portal`: NAC-Based Guest Portal<br>* `marvis_client` |
| `UiUrl` | `*string` | Optional | If `auth`==`guest_admin`, the URL to the guest admin portal |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "access_type": "wireless",
  "additional_cacerts": [
    "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"
  ],
  "cert_expire_time": 365,
  "eap_type": "wpa2",
  "name": "get-wifi",
  "portal_authorize_jwt_secret": "1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef",
  "portal_authorize_url": "https://guest-mistnac.mist.com/callback/be22bba7-8e22-e1cf-5185-b880816fe2cf/authorize",
  "portal_sso_url": "https://guest-mistnac.mist.com/callback/be22bba7-8e22-e1cf-5185-b880816fe2cf/acs",
  "ssid": "Corp",
  "ui_url": "https://guest-mistnac.mist.com/admin/51908ea7-dea7-4581-a578-f7320c4d5216/login",
  "bg_image_url": "bg_image_url2",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

