
# Device Other Stats Vendor Specific

when `vendor`==`cradlepoint`

## Structure

`DeviceOtherStatsVendorSpecific`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.DeviceOtherStatsVendorSpecificPort`](../../doc/models/device-other-stats-vendor-specific-port.md) | Optional | - |
| `TargetVersion` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "ports": {
    "mdm-4d0e073b": {
      "bytes_in": 33004879,
      "bytes_out": 41103393,
      "health_category": "",
      "health_score": 0,
      "id": "101027967",
      "mode": "wan",
      "model": "Internal 5GB (SIM1)",
      "state": "READY",
      "type": "5G",
      "uptime": 252371.34149021498
    }
  },
  "target_version": "7.23.40"
}
```

