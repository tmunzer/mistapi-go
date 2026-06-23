
# Switch Iot Port

Switch IOT port configuration

## Structure

`SwitchIotPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AlarmClass` | [`*models.SwitchIotPortAlarmClassEnum`](../../doc/models/switch-iot-port-alarm-class-enum.md) | Optional | Alarm class for the switch iot port in. enum: `minor`, `major`<br><br>**Default**: `"minor"` |
| `Enabled` | `*bool` | Optional | Whether this switch IOT port is enabled<br><br>**Default**: `false` |
| `InputSrc` | [`*models.SwitchIotPortInputSrcEnum`](../../doc/models/switch-iot-port-input-src-enum.md) | Optional | Only for "OUT" ports, input source for the switch iot port out. enum: `IN0`, `IN1`<br><br>**Default**: `"IN0"` |
| `Name` | `*string` | Optional | Display name for the switch IOT port |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchIotPort := models.SwitchIotPort{
        AlarmClass:           models.ToPointer(models.SwitchIotPortAlarmClassEnum_MINOR),
        Enabled:              models.ToPointer(false),
        InputSrc:             models.ToPointer(models.SwitchIotPortInputSrcEnum_IN0),
        Name:                 models.ToPointer("name6"),
    }

}
```

