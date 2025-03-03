
# Response Mxedge Upgrade

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseMxedgeUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Counts` | [`models.MxedgeUpgradeResponseCounts`](../../doc/models/mxedge-upgrade-response-counts.md) | Required | - |
| `Id` | `uuid.UUID` | Required | Unique ID of the object instance in the Mist Organization |
| `Status` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Strategy` | `string` | Required | **Constraints**: *Minimum Length*: `1` |
| `Versions` | `interface{}` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": "channel0",
  "counts": {
    "failed": 166,
    "queued": 234,
    "success": 90,
    "upgrading": 112,
    "exampleAdditionalProperty": {
      "key1": "val1",
      "key2": "val2"
    }
  },
  "id": "53f10664-3ce8-4c27-b382-0ef66432349f",
  "status": "status4",
  "strategy": "strategy4",
  "versions": {
    "key1": "val1",
    "key2": "val2"
  },
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

