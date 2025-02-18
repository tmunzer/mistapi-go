
# Psk Portal

*This model accepts additional fields of type interface{}.*

## Structure

`PskPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.PskPortalAuthEnum`](../../doc/models/psk-portal-auth-enum.md) | Optional | enum: `sponsor`, `sso`<br>**Default**: `"sso"` |
| `BgImageUrl` | `*string` | Optional | - |
| `CleanupPsk` | `*bool` | Optional | Used to cleanup exited psk when portal delete or ssid changed<br>**Default**: `false` |
| `CreatedTime` | `*float64` | Optional | When the object has been created, in epoch |
| `ExpireTime` | `*int` | Optional | unit min |
| `ExpiryNotificationTime` | `*int` | Optional | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire |
| `HidePsksCreatedByOtherAdmins` | `*bool` | Optional | Only if `type`==`admin`<br>**Default**: `false` |
| `Id` | `*uuid.UUID` | Optional | Unique ID of the object instance in the Mist Organnization |
| `MaxUsage` | `*int` | Optional | `max_usage`==`0` means unlimited<br>**Default**: `0` |
| `ModifiedTime` | `*float64` | Optional | When the object has been modified for the last time, in epoch |
| `Name` | `string` | Required | - |
| `NotificationRenewUrl` | `*string` | Optional | Optional, will include the link in the notification email the customer can either provide their own url or use the one generate from mist, or do a url shorterner against either |
| `NotifyExpiry` | `*bool` | Optional | If set to true, reminder notification will be sent when psk is about to expire |
| `NotifyOnCreateOrEdit` | `*bool` | Optional | **Default**: `false` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PassphraseRules` | [`*models.PskPortalPassphraseRules`](../../doc/models/psk-portal-passphrase-rules.md) | Optional | - |
| `RequiredFields` | `[]string` | Optional | what information to ask for (email is required by default) |
| `Role` | `*string` | Optional | - |
| `Ssid` | `string` | Required | intended SSID |
| `Sso` | [`*models.PskPortalSso`](../../doc/models/psk-portal-sso.md) | Optional | If `auth`==`sso` |
| `TemplateUrl` | `*string` | Optional | UI customization |
| `ThumbnailUrl` | `*string` | Optional | - |
| `Type` | [`*models.PskPortalTypeEnum`](../../doc/models/psk-portal-type-enum.md) | Optional | for personal psk portal. enum: `admin`, `byod` |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "auth": "sso",
  "cleanup_psk": false,
  "hide_psks_created_by_other_admins": false,
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "max_usage": 0,
  "name": "name0",
  "notification_renew_url": "https://custom-sso/url",
  "notify_on_create_or_edit": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "ssid": "ssid2",
  "bg_image_url": "bg_image_url8",
  "created_time": 207.6,
  "expire_time": 166,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

