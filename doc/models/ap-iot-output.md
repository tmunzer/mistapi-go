
# Ap Iot Output

IoT output AP settings

## Structure

`ApIotOutput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether to enable a pin<br>**Default**: `false` |
| `Name` | `*string` | Optional | optional; descriptive pin name |
| `Output` | `*bool` | Optional | whether the pin is configured as an output. DO and A1-A4 can be repurposed by changing |
| `Pullup` | [`*models.ApIotOutputPullupEnum`](../../doc/models/ap-iot-output-pullup-enum.md) | Optional | the type of pull-up the pin uses (internal, external, none), default none |
| `Value` | `*int` | Optional | output pin signal level, default 0 |

## Example (as JSON)

```json
{
  "enabled": false,
  "name": "motion",
  "output": false,
  "pullup": "external",
  "value": 178
}
```

