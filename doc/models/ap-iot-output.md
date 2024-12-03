
# Ap Iot Output

IoT output AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApIotOutput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether to enable a pin<br>**Default**: `false` |
| `Name` | `*string` | Optional | optional; descriptive pin name |
| `Output` | `*bool` | Optional | whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing |
| `Pullup` | [`*models.ApIotPullupEnum`](../../doc/models/ap-iot-pullup-enum.md) | Optional | the type of pull-up the pin uses. enum: `external`, `internal`, `none`<br>**Default**: `"none"` |
| `Value` | `*int` | Optional | output pin signal level, default 0 |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "name": "motion",
  "pullup": "none",
  "output": false,
  "value": 116,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

