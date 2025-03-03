
# Stats Device Other Interface

*This model accepts additional fields of type interface{}.*

## Structure

`StatsDeviceOtherInterface`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `BytesIn` | `*int64` | Optional | - |
| `BytesOut` | `*int64` | Optional | - |
| `Carrier` | `*string` | Optional | - |
| `Imei` | `*string` | Optional | - |
| `Imsi` | `*string` | Optional | - |
| `Ip` | `*string` | Optional | - |
| `Link` | `*bool` | Optional | - |
| `Mode` | `*string` | Optional | - |
| `Rsrp` | `*float64` | Optional | - |
| `Rsrq` | `*float64` | Optional | - |
| `Rssi` | `*int` | Optional | - |
| `ServiceMode` | `*string` | Optional | - |
| `Sinr` | `*float64` | Optional | - |
| `State` | `*string` | Optional | - |
| `Type` | `*string` | Optional | - |
| `Uptime` | `*int` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
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
```

