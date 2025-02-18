
# Ap Led

LED AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApLed`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Brightness` | `*int` | Optional | **Default**: `255`<br>**Constraints**: `>= 0`, `<= 255` |
| `Enabled` | `*bool` | Optional | **Default**: `true` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "brightness": 255,
  "enabled": true,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

