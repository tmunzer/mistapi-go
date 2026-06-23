
# Ap Iot

Digital and analog IoT port settings applied to an AP or AP profile

## Structure

`ApIot`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `A1` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `A2` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `A3` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `A4` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |
| `DI1` | [`*models.ApIotInput`](../../doc/models/ap-iot-input.md) | Optional | IoT Input AP settings |
| `DI2` | [`*models.ApIotInput`](../../doc/models/ap-iot-input.md) | Optional | IoT Input AP settings |
| `DO` | [`*models.ApIotOutput`](../../doc/models/ap-iot-output.md) | Optional | IoT output AP settings |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apIot := models.ApIot{
        A1:                   models.ToPointer(models.ApIotOutput{
            Enabled:              models.ToPointer(false),
            Name:                 models.ToPointer("name0"),
            Output:               models.ToPointer(false),
            Pullup:               models.ToPointer(models.ApIotPullupEnum_INTERNAL),
            Value:                models.ToPointer(242),
        }),
        A2:                   models.ToPointer(models.ApIotOutput{
            Enabled:              models.ToPointer(false),
            Name:                 models.ToPointer("name8"),
            Output:               models.ToPointer(false),
            Pullup:               models.ToPointer(models.ApIotPullupEnum_NONE),
            Value:                models.ToPointer(180),
        }),
        A3:                   models.ToPointer(models.ApIotOutput{
            Enabled:              models.ToPointer(false),
            Name:                 models.ToPointer("name6"),
            Output:               models.ToPointer(false),
            Pullup:               models.ToPointer(models.ApIotPullupEnum_INTERNAL),
            Value:                models.ToPointer(118),
        }),
        A4:                   models.ToPointer(models.ApIotOutput{
            Enabled:              models.ToPointer(false),
            Name:                 models.ToPointer("name8"),
            Output:               models.ToPointer(false),
            Pullup:               models.ToPointer(models.ApIotPullupEnum_INTERNAL),
            Value:                models.ToPointer(88),
        }),
        DI1:                  models.ToPointer(models.ApIotInput{
            Enabled:              models.ToPointer(false),
            Name:                 models.ToPointer("name0"),
            Pullup:               models.ToPointer(models.ApIotPullupEnum_INTERNAL),
        }),
    }

}
```

