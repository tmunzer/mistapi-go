
# Locate Switch

*This model accepts additional fields of type interface{}.*

## Structure

`LocateSwitch`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | minutes the leds should keep flashing<br>**Default**: `5`<br>**Constraints**: `>= 1`, `<= 120` |
| `Mac` | `*string` | Optional | for virtual chassis, the MAC of the member |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "duration": 5,
  "mac": "f01c2d4ff760",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

