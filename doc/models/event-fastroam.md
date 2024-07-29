
# Event Fastroam

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
| `Timestamp` | `float64` | Required | timestamp of the event in nsec |
| `Type` | [`*models.EventFastroamTypeEnum`](../../doc/models/event-fastroam-type-enum.md) | Optional | enum: `fail`, `none`, `pingpong`, `poor`, `slow`, `success` |

## Example (as JSON)

```json
{
  "ap_mac": "ap_mac0",
  "client_mac": "client_mac6",
  "fromap": "fromap6",
  "latency": 43.52,
  "ssid": "ssid6",
  "subtype": "subtype0",
  "timestamp": 209.26,
  "type": "slow"
}
```

