
# Nac Portal

*This model accepts additional fields of type interface{}.*

## Structure

`NacPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AccessType` | [`*models.NacPortalAccessTypeEnum`](../../doc/models/nac-portal-access-type-enum.md) | Optional | if `type`==`marvis_client`. enum: `wireless`, `wireless+wired`<br><br>**Default**: `"wireless"` |
| `BgImageUrl` | `*string` | Optional | Background image |
| `CertExpireTime` | `*int` | Optional | In days |
| `EapType` | [`*models.NacPortalEapTypeEnum`](../../doc/models/nac-portal-eap-type-enum.md) | Optional | enum: `wpa2`, `wpa3`<br><br>**Default**: `"wpa2"` |
| `EnableTelemetry` | `*bool` | Optional | Model, version, fingering, events (connecting, disconnect, roaming), which ap |
| `ExpiryNotificationTime` | `*int` | Optional | In days |
| `GuestPortalConfig` | [`*models.WlanPortal`](../../doc/models/wlan-portal.md) | Optional | Portal wlan settings |
| `Name` | `*string` | Optional | - |
| `NotifyExpiry` | `*bool` | Optional | phase 2 |
| `Ssid` | `*string` | Optional | - |
| `Sso` | [`*models.NacPortalSso`](../../doc/models/nac-portal-sso.md) | Optional | - |
| `TemplateUrl` | `*string` | Optional | - |
| `ThumbnailUrl` | `*string` | Optional | - |
| `Tos` | `*string` | Optional | - |
| `Type` | [`*models.NacPortalTypeEnum`](../../doc/models/nac-portal-type-enum.md) | Optional | enum: `guest`, `marvis_client` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "access_type": "wireless",
  "cert_expire_time": 365,
  "eap_type": "wpa2",
  "name": "get-wifi",
  "ssid": "Corp",
  "bg_image_url": "bg_image_url2",
  "enable_telemetry": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

