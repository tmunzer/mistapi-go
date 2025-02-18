
# Event Fastroam

*This model accepts additional fields of type interface{}.*

## Structure

`EventFastroam`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApMac` | `string` | Required | - |
| `ClientMac` | `string` | Required | - |
| `Fromap` | `string` | Required | - |
| `Latency` | `float64` | Required | - |
| `Ssid` | `string` | Required | - |
| `Subtype` | `*string` | Optional | - |
| `Timestamp` | `float64` | Required | Timestamp of the event in nsec |
| `Type` | [`*models.EventFastroamTypeEnum`](../../doc/models/event-fastroam-type-enum.md) | Optional | enum: `fail`, `none`, `pingpong`, `poor`, `slow`, `success` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac2",
  "client_mac": "client_mac8",
  "fromap": "fromap2",
  "latency": 124.3,
  "ssid": "ssid2",
  "subtype": "subtype2",
  "timestamp": 128.48,
  "type": "fail",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

