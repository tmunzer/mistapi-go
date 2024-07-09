
# Mxedge Stats Service Stat

## Structure

`MxedgeStatsServiceStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ExtIp` | `*string` | Optional | external IP from ep-terminator’s point of view. valid only for service having its own cloud connection |
| `LastSeen` | `*int` | Optional | timestamp when the last stats is seen (cloud unix time, in second). valid only for service having its own stats or whole system (last among last_seen of all services) |
| `PackageState` | `*string` | Optional | package/service installation state. |
| `PackageVersion` | `*string` | Optional | package/service installation state. |
| `RunningState` | `*string` | Optional | service running state. |
| `Uptime` | `*int` | Optional | service uptime. |

## Example (as JSON)

```json
{
  "ext_ip": "ext_ip6",
  "last_seen": 0,
  "package_state": "package_state2",
  "package_version": "package_version0",
  "running_state": "running_state0"
}
```

