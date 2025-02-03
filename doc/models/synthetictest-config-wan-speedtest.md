
# Synthetictest Config Wan Speedtest

*This model accepts additional fields of type interface{}.*

## Structure

`SynthetictestConfigWanSpeedtest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | - |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format)<br>**Default**: `"any"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "time_of_day": "12:00",
  "enabled": false,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

