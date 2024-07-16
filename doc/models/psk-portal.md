
# Psk Portal

## Structure

`PskPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.PskPortalAuthEnum`](../../doc/models/psk-portal-auth-enum.md) | Optional | Note: `sponsor` not yet available<br>**Default**: `"sso"` |
| `BgImageUrl` | `*string` | Optional | - |
| `CleanupPsk` | `*bool` | Optional | used to cleanup exited psk when portal delete or ssid changed<br>**Default**: `false` |
| `CreatedTime` | `*float64` | Optional | - |
| `ExpireTime` | `*int` | Optional | unit min |
| `ExpiryNotificationTime` | `*int` | Optional | Number of days before psk is expired. Used as to when to start sending reminder notification when the psk is about to expire |
| `HidePsksCreatedByOtherAdmins` | `*bool` | Optional | only if `type`==`admin`<br>**Default**: `false` |
| `Id` | `*string` | Optional | - |
| `MaxUsage` | `*int` | Optional | `max_usage`==`0` means unlimited<br>**Default**: `0` |
| `ModifiedTime` | `*float64` | Optional | - |
| `Name` | `string` | Required | - |
| `NotificationRenewUrl` | `*string` | Optional | optional, will include the link in the notification email the customer can either provide their own url or use the one generate from mist, or do a url shorterner against either |
| `NotifyExpiry` | `*bool` | Optional | If set to true, reminder notification will be sent when psk is about to expire |
| `NotifyOnCreateOrEdit` | `*bool` | Optional | **Default**: `false` |
| `OrgId` | `*uuid.UUID` | Optional | - |
| `PassphraseRules` | [`*models.PskPortalPassphraseRules`](../../doc/models/psk-portal-passphrase-rules.md) | Optional | - |
| `RequiredFields` | `[]string` | Optional | what information to ask for (email is required by default) |
| `Role` | `*string` | Optional | - |
| `Ssid` | `string` | Required | intended SSID |
| `Sso` | [`*models.PskPortalSso`](../../doc/models/psk-portal-sso.md) | Optional | if `auth`==`sso` |
| `TemplateUrl` | `*string` | Optional | UI customization |
| `ThumbnailUrl` | `*string` | Optional | - |
| `Type` | [`*models.PskPortalTypeEnum`](../../doc/models/psk-portal-type-enum.md) | Optional | for personal psk portal |
| `VlanId` | `*int` | Optional | - |

## Example (as JSON)

```json
{
  "auth": "sso",
  "cleanup_psk": false,
  "hide_psks_created_by_other_admins": false,
  "max_usage": 0,
  "name": "name6",
  "notification_renew_url": "https://custom-sso/url",
  "notify_on_create_or_edit": false,
  "org_id": "a97c1b22-a4e9-411e-9bfd-d8695a0f9e61",
  "ssid": "ssid6",
  "bg_image_url": "bg_image_url2",
  "created_time": 38.26,
  "expire_time": 204
}
```

