
# Response Ssr Upgrade Counts

*This model accepts additional fields of type interface{}.*

## Structure

`ResponseSsrUpgradeCounts`

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
  "failed": 200,
  "queued": 244,
  "success": 56,
  "upgrading": 78,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

