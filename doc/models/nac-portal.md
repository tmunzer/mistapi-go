
# Nac Portal

## Structure

`NacPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessType` | [`*models.NacPortalAccessTypeEnum`](../../doc/models/nac-portal-access-type-enum.md) | Optional | if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`<br>**Default**: `"wireless"` |
| `BgImageUrl` | `*string` | Optional | background image |
| `CertExpireTime` | `*int` | Optional | in days |
| `EnableTelemetry` | `*bool` | Optional | model, version, fingering, events (connecting, disconnect, roaming), which ap |
| `ExpiryNotificationTime` | `*int` | Optional | in days |
| `GuestPortalConfig` | [`*models.WlanPortal`](../../doc/models/wlan-portal.md) | Optional | portal wlan settings |
| `Name` | `*string` | Optional | - |
| `NotifyExpiry` | `*bool` | Optional | phase 2 |
| `Ssid` | `*string` | Optional | - |
| `Sso` | [`*models.NacPortalSso`](../../doc/models/nac-portal-sso.md) | Optional | - |
| `TemplateUrl` | `*string` | Optional | - |
| `ThumbnailUrl` | `*string` | Optional | - |
| `Tos` | `*string` | Optional | - |
| `Type` | [`*models.NacPortalTypeEnum`](../../doc/models/nac-portal-type-enum.md) | Optional | - |

## Example (as JSON)

```json
{
  "access_type": "wireless",
  "cert_expire_time": 365,
  "name": "get-wifi",
  "ssid": "Corp",
  "bg_image_url": "bg_image_url2",
  "enable_telemetry": false,
  "expiry_notification_time": 2
}
```

