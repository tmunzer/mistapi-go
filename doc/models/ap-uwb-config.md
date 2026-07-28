
# Ap Uwb Config

Ultra-wideband (UWB) RTLS / OMLOX asset-visibility integration settings for an access point. The device-level value overrides the device profile value, which in turn overrides the site-level setting.

## Structure

`ApUwbConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether UWB RTLS integration is enabled<br><br>**Default**: `false` |
| `Host` | `*string` | Optional | RTLS server hostname or IP address |
| `Port` | `*int` | Optional | RTLS server port number<br><br>**Constraints**: `>= 1`, `<= 65535` |
| `Slot` | `*int` | Optional | UWB time slot assigned to this AP, 0–15<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 15` |
| `Type` | [`*models.ApUwbConfigTypeEnum`](../../doc/models/ap-uwb-config-type-enum.md) | Optional | UWB integration type. enum: `zigpos` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apUwbConfig := models.ApUwbConfig{
        Enabled:              models.ToPointer(false),
        Host:                 models.ToPointer("coriva.example.com"),
        Port:                 models.ToPointer(9000),
        Slot:                 models.ToPointer(0),
        Type:                 models.ToPointer(models.ApUwbConfigTypeEnum_ZIGPOS),
    }

}
```

