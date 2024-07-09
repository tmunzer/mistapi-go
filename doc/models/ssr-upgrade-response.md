
# Ssr Upgrade Response

## Structure

`SsrUpgradeResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Counts` | [`models.SsrUpgradeResponseCounts`](../../doc/models/ssr-upgrade-response-counts.md) | Required | - |
| `DeviceType` | `string` | Required | - |
| `Id` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Status` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Strategy` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Versions` | `map[string]string` | Required | - |

## Example (as JSON)

```json
{
  "channel": "channel6",
  "counts": {
    "failed": 166,
    "queued": 234,
    "success": 90,
    "upgrading": 112
  },
  "device_type": "device_type8",
  "id": "id8",
  "status": "status0",
  "strategy": "strategy8",
  "versions": {
    "key0": "versions3",
    "key1": "versions4"
  }
}
```

