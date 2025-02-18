
# Stats Ap Usb Stat

*This model accepts additional fields of type interface{}.*

## Structure

`StatsApUsbStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `models.Optional[int]` | Optional | - |
| `Connected` | `models.Optional[bool]` | Optional | - |
| `LastActivity` | `models.Optional[int]` | Optional | - |
| `Type` | `models.Optional[string]` | Optional | - |
| `Up` | `models.Optional[bool]` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "channel": 3,
  "connected": true,
  "last_activity": 1586873254,
  "type": "imagotag",
  "up": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

