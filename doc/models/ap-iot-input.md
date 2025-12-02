
# Ap Iot Input

IoT Input AP settings

## Structure

`ApIotInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable a pin<br><br>**Default**: `false` |
| `Name` | `*string` | Optional | Optional; descriptive pin name |
| `Pullup` | [`*models.ApIotPullupEnum`](../../doc/models/ap-iot-pullup-enum.md) | Optional | the type of pull-up the pin uses. enum: `external`, `internal`, `none`<br><br>**Default**: `"none"` |

## Example (as JSON)

```json
{
  "enabled": false,
  "name": "motion",
  "pullup": "none"
}
```

