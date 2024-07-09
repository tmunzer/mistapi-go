
# Response Mxedge Upgrade

## Structure

`ResponseMxedgeUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Counts` | [`models.MxedgeUpgradeResponseCounts`](../../doc/models/mxedge-upgrade-response-counts.md) | Required | - |
| `Id` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Status` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Strategy` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Versions` | `interface{}` | Required | - |

## Example (as JSON)

```json
{
  "channel": "channel2",
  "counts": {
    "failed": 166,
    "queued": 234,
    "success": 90,
    "upgrading": 112
  },
  "id": "id2",
  "status": "status4",
  "strategy": "strategy2",
  "versions": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

