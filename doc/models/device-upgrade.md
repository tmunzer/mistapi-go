
# Device Upgrade

Device firmware upgrade request options

*This model accepts additional fields of type interface{}.*

## Structure

`DeviceUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Reboot` | `*bool` | Optional | For Switches and Gateways only (APs are automatically rebooted). Reboot device immediately after upgrade is completed<br><br>**Default**: `false` |
| `RebootAt` | `*int` | Optional | For Switches and Gateways only and if `reboot`==`true`. Reboot start time in epoch seconds, default is `start_time` |
| `Snapshot` | `*bool` | Optional | For Junos devices only. Perform recovery snapshot after device is rebooted<br><br>**Default**: `false` |
| `StartTime` | `*int` | Optional | Firmware download start time in epoch |
| `Version` | `string` | Required | Specific version / `stable`, default is to use the latest<br><br>**Default**: `"stable"` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    deviceUpgrade := models.DeviceUpgrade{
        Reboot:               models.ToPointer(false),
        RebootAt:             models.ToPointer(74),
        Snapshot:             models.ToPointer(false),
        StartTime:            models.ToPointer(144),
        Version:              "stable",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

