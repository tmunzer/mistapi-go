
# Stats Mxedge Service Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeServiceStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | External IP from ep-terminator’s point of view. valid only for service having its own cloud connection |
| `LastSeen` | `*float64` | Optional | Timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services) |
| `PackageState` | `*string` | Optional | Package/service installation state. |
| `PackageVersion` | `*string` | Optional | Package/service installation state. |
| `RunningState` | `*string` | Optional | Service running state. |
| `Uptime` | `*int` | Optional | Service uptime. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ext_ip": "ext_ip0",
  "last_seen": 97.42,
  "package_state": "package_state6",
  "package_version": "package_version4",
  "running_state": "running_state6",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

