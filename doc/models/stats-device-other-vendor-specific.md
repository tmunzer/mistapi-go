
# Stats Device Other Vendor Specific

When `vendor`==`cradlepoint`

*This model accepts additional fields of type interface{}.*

## Structure

`StatsDeviceOtherVendorSpecific`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ports` | [`map[string]models.StatsDeviceOtherVendorSpecificPort`](../../doc/models/stats-device-other-vendor-specific-port.md) | Optional | - |
| `TargetVersion` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ports": {
    "mdm-4d0e073b": {
      "bytes_in": 5623096929,
      "bytes_out": 12372750366,
      "carrier": "Orange",
      "imei": "866401234567893",
      "imsi": "2080101234567893",
      "ip": "10.134.237.57",
      "link": true,
      "mode": "wan",
      "rsrp": -108,
      "rsrq": -14,
      "rssi": -74,
      "service_mode": "5G NSA",
      "sinr": -1.2,
      "state": "READY",
      "type": "mdm",
      "uptime": 2095779,
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "target_version": "7.23.40",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

