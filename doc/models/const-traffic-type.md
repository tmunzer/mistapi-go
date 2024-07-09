
# Const Traffic Type

## Structure

`ConstTrafficType`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Display` | `*string` | Optional | - |
| `Dscp` | `*int` | Optional | - |
| `FailoverPolicy` | `*string` | Optional | - |
| `MaxJitter` | `*int` | Optional | - |
| `MaxLatency` | `*int` | Optional | - |
| `MaxLoss` | `*int` | Optional | - |
| `Name` | `*string` | Optional | - |
| `TrafficClass` | `*string` | Optional | - |

## Example (as JSON)

```json
{
  "display": "VoIP Video",
  "dscp": 32,
  "failover_policy": "non_revertible",
  "max_jitter": 250,
  "max_latency": 1500,
  "max_loss": 35,
  "name": "voip_video",
  "traffic_class": "medium"
}
```

