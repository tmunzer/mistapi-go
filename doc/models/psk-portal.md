
# Psk Portal

Self-service portal configuration for issuing personal PSKs

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.PskPortalAuthEnum`](../../doc/models/psk-portal-auth-enum.md) | Optional | Portal access method for guest authentication, either `sponsor` or `sso`. enum: `sponsor`, `sso`<br><br>**Default**: `"sso"` |
| `BgImageUrl` | `*string` | Optional | URL of the background image used by the PSK portal |
| `CleanupPsk` | `*bool` | Optional | Whether to clean up existing PSKs when the portal is deleted or its SSID changes<br><br>**Default**: `false` |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `ExpireTime` | `*int` | Optional | PSK lifetime, in minutes, for keys created through this portal |
| `ExpiryNotificationTime` | `*int` | Optional | Number of days before PSK expiration to start sending reminder notifications |
| `HidePsksCreatedByOtherAdmins` | `*bool` | Optional | Only if `type`==`admin`, hide PSKs created by other PSK admins<br><br>**Default**: `false` |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MaxUsage` | `*int` | Optional | Maximum concurrent clients for each PSK created through this portal; `0` means unlimited<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0` |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | Display name of the PSK portal |
| `NotificationRenewUrl` | `*string` | Optional | Optional renewal URL included in notification emails. Use a custom URL, a Mist-generated URL, or a shortened URL pointing to either |
| `NotifyExpiry` | `*bool` | Optional | Whether to send reminder notifications before PSKs created through this portal expire |
| `NotifyOnCreateOrEdit` | `*bool` | Optional | Whether to send notifications when a PSK is created or edited through this portal<br><br>**Default**: `false` |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PassphraseRules` | [`*models.PskPortalPassphraseRules`](../../doc/models/psk-portal-passphrase-rules.md) | Optional | Passphrase generation rules for PSKs created through a portal |
| `RequiredFields` | `[]string` | Optional | User information fields requested by the portal; email is required by default |
| `Role` | `*string` | Optional | Client role assigned to PSKs created through this portal |
| `Ssid` | `string` | Required | WLAN SSID for PSKs created through this portal |
| `Sso` | [`*models.PskPortalSso`](../../doc/models/psk-portal-sso.md) | Optional | Single sign-on settings used when `auth`==`sso` |
| `TemplateUrl` | `*string` | Optional | URL of the UI customization template for this portal |
| `ThumbnailUrl` | `*string` | Optional | URL of the thumbnail image used by the PSK portal |
| `Type` | [`*models.PskPortalTypeEnum`](../../doc/models/psk-portal-type-enum.md) | Optional | for personal psk portal. enum: `admin`, `byod` |
| `UiUrl` | `*string` | Optional | Public URL where users access the PSK portal |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN ID, either numeric or expressed as a template variable string |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    pskPortal := models.PskPortal{
        Auth:                         models.ToPointer(models.PskPortalAuthEnum_SSO),
        BgImageUrl:                   models.ToPointer("bg_image_url4"),
        CleanupPsk:                   models.ToPointer(false),
        CreatedTime:                  models.ToPointer(float64(96.64)),
        ExpireTime:                   models.ToPointer(254),
        HidePsksCreatedByOtherAdmins: models.ToPointer(false),
        Id:                           models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MaxUsage:                     models.ToPointer(0),
        Name:                         "name4",
        NotificationRenewUrl:         models.ToPointer("https://custom-sso/url"),
        NotifyOnCreateOrEdit:         models.ToPointer(false),
        OrgId:                        models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Ssid:                         "ssid8",
        AdditionalProperties:         map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

