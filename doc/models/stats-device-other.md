
# Stats Device Other

*This model accepts additional fields of type interface{}.*

## Structure

`StatsDeviceOther`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CachedStats` | `*bool` | Optional | - |
| `ConfigStatus` | `*string` | Optional | - |
| `ConnectedDevices` | [`map[string]models.StatsDeviceOtherConnectedDevice`](../../doc/models/stats-device-other-connected-device.md) | Optional | Property key is the connected device MAC Address |
| `Interfaces` | [`map[string]models.StatsDeviceOtherInterface`](../../doc/models/stats-device-other-interface.md) | Optional | Property key is the interface name |
| `LastConfig` | `*int` | Optional | - |
| `LastSeen` | `models.Optional[float64]` | Optional | Last seen timestamp |
| `LldpEnabled` | `*bool` | Optional | - |
| `Mac` | `*string` | Optional | - |
| `Status` | `*string` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `Vendor` | `*string` | Optional | - |
| `VendorSpecific` | [`*models.StatsDeviceOtherVendorSpecific`](../../doc/models/stats-device-other-vendor-specific.md) | Optional | When `vendor`==`cradlepoint` |
| `Version` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "config_status": "synced",
  "last_config": 1675392788,
  "last_seen": 1470417522,
  "mac": "5c5b35000018",
  "status": "online",
  "uptime": 20296,
  "vendor": "cradlepoint",
  "version": "7.22.70",
  "cached_stats": false,
  "connected_devices": {
    "key0": {
      "mac": "mac2",
      "name": "name8",
      "port_id": "port_id8",
      "type": "type8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "mac": "mac2",
      "name": "name8",
      "port_id": "port_id8",
      "type": "type8",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "interfaces": {
    "key0": {
      "bytes_in": 48,
      "bytes_out": 188,
      "carrier": "carrier6",
      "imei": "imei8",
      "imsi": "imsi2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    },
    "key1": {
      "bytes_in": 48,
      "bytes_out": 188,
      "carrier": "carrier6",
      "imei": "imei8",
      "imsi": "imsi2",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

