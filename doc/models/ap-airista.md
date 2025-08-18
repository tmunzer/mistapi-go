
# Ap Airista

*This model accepts additional fields of type interface{}.*

## Structure

`ApAirista`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable Airista config<br><br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | Required if enabled, Airista server host |
| `Port` | `models.Optional[int]` | Optional | **Default**: `1144` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "host": "airista.pvt.net",
  "port": 1144,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

