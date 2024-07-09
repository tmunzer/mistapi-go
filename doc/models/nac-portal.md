
# Nac Portal

## Structure

`NacPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessType` | [`*models.NacPortalAccessTypeEnum`](../../doc/models/nac-portal-access-type-enum.md) | Optional | **Default**: `"wireless"` |
| `BgImageUrl` | `*string` | Optional | background image |
| `CertExpireTime` | `*int` | Optional | in days |
| `EnableTelemetry` | `*bool` | Optional | model, version, fingering, events (connecting, disconnect, roaming), which ap |
| `ExpiryNotificationTime` | `*int` | Optional | in days |
| `Name` | `*string` | Optional | - |
| `NotifyExpiry` | `*bool` | Optional | phase 2 |
| `Ssid` | `*string` | Optional | - |
| `Sso` | [`*models.NacPortalSso`](../../doc/models/nac-portal-sso.md) | Optional | - |
| `TemplateUrl` | `*string` | Optional | - |
| `ThumbnailUrl` | `*string` | Optional | - |
| `Tos` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "access_type": "wireless",
  "cert_expire_time": 365,
  "name": "get-wifi",
  "ssid": "Corp",
  "bg_image_url": "bg_image_url8",
  "enable_telemetry": false,
  "expiry_notification_time": 198
}
```

