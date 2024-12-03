
# Ssr Upgrade Response Counts

*This model accepts additional fields of type interface{}.*

## Structure

`SsrUpgradeResponseCounts`

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
  "failed": 56,
  "queued": 132,
  "success": 200,
  "upgrading": 222,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

