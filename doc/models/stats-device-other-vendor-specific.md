
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
      "bytes_in": 33004879,
      "bytes_out": 41103393,
      "health_category": "",
      "health_score": 0,
      "id": "101027967",
      "mode": "wan",
      "model": "Internal 5GB (SIM1)",
      "state": "READY",
      "type": "5G",
      "uptime": 252371.34149021498,
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

