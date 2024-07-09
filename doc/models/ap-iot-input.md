
# Ap Iot Input

IoT Input AP settings

## Structure

`ApIotInput`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | whether to enable a pin<br>**Default**: `false` |
| `Name` | `*string` | Optional | optional; descriptive pin name |
| `Pullup` | [`*models.ApIotInputPullupEnum`](../../doc/models/ap-iot-input-pullup-enum.md) | Optional | the type of pull-up the pin uses (internal, external, none), default none |

## Example (as JSON)

```json
{
  "enabled": false,
  "name": "name8",
  "pullup": "external"
}
```

