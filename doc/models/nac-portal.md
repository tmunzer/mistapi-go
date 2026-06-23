
# Nac Portal

NAC portal configuration for 802.1X onboarding, guest access, or Marvis client certificate provisioning

*This model accepts additional fields of type interface{}.*

## Structure

`NacPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessType` | [`*models.NacPortalAccessTypeEnum`](../../doc/models/nac-portal-access-type-enum.md) | Optional | if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`<br><br>**Default**: `"wireless"` |
| `AdditionalCacerts` | `[]string` | Optional | Optional list of additional CA certificates to be used |
| `AdditionalNacServerName` | `[]string` | Optional | Optional list of additional NAC server names |
| `BgImageUrl` | `*string` | Optional | URL of the NAC portal background image |
| `CertExpireTime` | `*int` | Optional | Validity duration for portal-issued client certificates, in days |
| `EapType` | [`*models.NacPortalEapTypeEnum`](../../doc/models/nac-portal-eap-type-enum.md) | Optional | EAP mode used when onboarding wireless clients through the NAC portal. enum: `wpa2`, `wpa3`<br><br>**Default**: `"wpa2"` |
| `EnableTelemetry` | `*bool` | Optional | Model, version, fingering, events (connecting, disconnect, roaming), which ap |
| `ExpiryNotificationTime` | `*int` | Optional | Number of days before certificate expiration to start sending reminder notifications |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | Human-readable name of the NAC portal |
| `NotifyExpiry` | `*bool` | Optional | Whether to send reminder notifications before portal-issued certificates expire |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Portal` | [`*models.NacPortalGuestPortal`](../../doc/models/nac-portal-guest-portal.md) | Optional | Guest portal configuration when `type`==`guest_portal`. If<br><br>* `auth`==`none`, the user is presented with a terms of service and can click and continue.<br>* `auth`==`external`, the user is redirected to an external URL for authentication.<br>* `auth`==`multi`, the user is presented with a choice of authentication methods:<br>  - social logins: facebook / google / amazon / microsoft / azure<br>  - sponsor<br>  - sms: supported provider: twillio<br>  - email<br>  - sso<br>  - userpass: pre created guest list |
| `PortalAuthorizeJwtSecret` | `*string` | Optional, Read-only | If `type`==`guest_portal` and `auth`==`external`, the `portal_authorize_jwt_secret` will be generated |
| `PortalAuthorizeUrl` | `*string` | Optional, Read-only | If `type`==`guest_portal` and `auth`==`external`, the `portal_authorize_url` will be generated |
| `PortalSsoUrl` | `*string` | Optional, Read-only | If `type`==`guest_portal` or `type`==`guest_admin` and ans SSO is enabled, the `portal_sso_url` will be generated (which needs to be configured in your IDP |
| `Ssid` | `*string` | Optional | Wireless SSID associated with the NAC portal |
| `Sso` | [`*models.NacPortalSso`](../../doc/models/nac-portal-sso.md) | Optional | SAML SSO configuration for a NAC portal |
| `TemplateUrl` | `*string` | Optional | URL for the NAC portal template customization resource |
| `ThumbnailUrl` | `*string` | Optional, Read-only | Read-only URL of the NAC portal background image thumbnail |
| `Tos` | `*string` | Optional | Terms of service text shown in the NAC portal |
| `Type` | [`*models.NacPortalTypeEnum`](../../doc/models/nac-portal-type-enum.md) | Optional | enum:<br><br>* `guest_admin`: NAC-Based Portal Admin for Pre Created Guest Authentication<br>* `guest_portal`: NAC-Based Guest Portal<br>* `marvis_client` |
| `UiUrl` | `*string` | Optional, Read-only | If `auth`==`guest_admin`, the URL to the guest admin portal |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    nacPortal := models.NacPortal{
        AccessType:               models.ToPointer(models.NacPortalAccessTypeEnum_WIRELESS),
        AdditionalCacerts:        []string{
            "-----BEGIN CERTIFICATE-----\\<redacted>\\n-----END CERTIFICATE-----",
        },
        AdditionalNacServerName:  []string{
            "nac1.corp.com",
            "nac2.corp.com",
        },
        BgImageUrl:               models.ToPointer("bg_image_url4"),
        CertExpireTime:           models.ToPointer(365),
        EapType:                  models.ToPointer(models.NacPortalEapTypeEnum_WPA2),
        Id:                       models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                     models.ToPointer("get-wifi"),
        OrgId:                    models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PortalAuthorizeJwtSecret: models.ToPointer("1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef"),
        PortalAuthorizeUrl:       models.ToPointer("https://guest-mistnac.mist.com/callback/be22bba7-8e22-e1cf-5185-b880816fe2cf/authorize"),
        PortalSsoUrl:             models.ToPointer("https://guest-mistnac.mist.com/callback/be22bba7-8e22-e1cf-5185-b880816fe2cf/acs"),
        Ssid:                     models.ToPointer("Corp"),
        UiUrl:                    models.ToPointer("https://guest-mistnac.mist.com/admin/51908ea7-dea7-4581-a578-f7320c4d5216/login"),
        AdditionalProperties:     map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

