
# Stats Device Other

## Structure

`StatsDeviceOther`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigStatus` | `*string` | Optional | - |
| `LastConfig` | `*int` | Optional | - |
| `LastSeen` | `*int` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `Vendor` | `*string` | Optional | - |
| `VendorSpecific` | [`*models.StatsDeviceOtherVendorSpecific`](../../doc/models/stats-device-other-vendor-specific.md) | Optional | when `vendor`==`cradlepoint` |
| `Version` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "config_status": "synced",
  "last_config": 1675392788,
  "last_seen": 1675843629,
  "mac": "5c5b35000018",
  "status": "online",
  "uptime": 20296,
  "vendor": "cradlepoint",
  "version": "7.22.70"
}
```

