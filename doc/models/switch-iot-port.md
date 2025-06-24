
# Switch Iot Port

Switch IOT port configuration

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchIotPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmClass` | [`*models.SwitchIotPortAlarmClassEnum`](../../doc/models/switch-iot-port-alarm-class-enum.md) | Optional | Alarm class for the switch iot port in. enum: `minor`, `major`<br><br>**Default**: `"minor"` |
| `Enabled` | `*bool` | Optional | **Default**: `false` |
| `InputSrc` | [`*models.SwitchIotPortInputSrcEnum`](../../doc/models/switch-iot-port-input-src-enum.md) | Optional | Only for "OUT" ports, input source for the switch iot port out. enum: `IN0`, `IN1`<br><br>**Default**: `"IN0"` |
| `Name` | `*string` | Optional | - |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "alarm_class": "minor",
  "enabled": false,
  "input_src": "IN0",
  "name": "name0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

