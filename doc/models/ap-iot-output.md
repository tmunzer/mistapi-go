
# Ap Iot Output

IoT output AP settings

*This model accepts additional fields of type interface{}.*

## Structure

`ApIotOutput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable a pin<br><br>**Default**: `false` |
| `Name` | `*string` | Optional | Optional; descriptive pin name |
| `Output` | `*bool` | Optional | Whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing |
| `Pullup` | [`*models.ApIotPullupEnum`](../../doc/models/ap-iot-pullup-enum.md) | Optional | the type of pull-up the pin uses. enum: `external`, `internal`, `none`<br><br>**Default**: `"none"` |
| `Value` | `*int` | Optional | Output pin signal level, default 0 |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "enabled": false,
  "name": "motion",
  "output": true,
  "pullup": "none",
  "value": 0,
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

