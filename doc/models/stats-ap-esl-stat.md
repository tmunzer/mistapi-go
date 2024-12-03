
# Stats Ap Esl Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApEslStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `models.Optional[int]` | Optional | - |
| `Connected` | `models.Optional[bool]` | Optional | - |
| `Type` | `models.Optional[string]` | Optional | - |
| `Up` | `models.Optional[bool]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": 76,
  "connected": false,
  "type": "type0",
  "up": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

