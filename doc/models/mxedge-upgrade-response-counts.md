
# Mxedge Upgrade Response Counts

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgeUpgradeResponseCounts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Failed` | `int` | Required | - |
| `Queued` | `int` | Required | - |
| `Success` | `int` | Required | - |
| `Upgrading` | `int` | Required | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "failed": 68,
  "queued": 120,
  "success": 188,
  "upgrading": 210,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

