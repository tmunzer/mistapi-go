
# Ap Usb

Legacy USB integration settings for an access point

- Note: if native imagotag is enabled, BLE will be disabled automatically
- Note: legacy, new config moved to ESL Config.

## Structure

`ApUsb`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cacert` | `models.Optional[string]` | Optional | Only if `type`==`imagotag`. CA certificate used to validate the Imagotag service certificate |
| `Channel` | `*int` | Optional | Only if `type`==`imagotag`, channel selection, not needed by default, required for manual channel override only |
| `Enabled` | `*bool` | Optional | Whether to enable any usb config |
| `Host` | `*string` | Optional | Only if `type`==`imagotag`. Imagotag service host or IP address contacted by the AP |
| `Port` | `*int` | Optional | Only if `type`==`imagotag`. TCP port used to reach the Imagotag service<br><br>**Default**: `0` |
| `Type` | [`*models.ApUsbTypeEnum`](../../doc/models/ap-usb-type-enum.md) | Optional | usb config type. enum: `hanshow`, `imagotag`, `solum` |
| `VerifyCert` | `*bool` | Optional | Only if `type`==`imagotag`, whether to turn on SSL verification |
| `VlanId` | `*int` | Optional | Only if `type`==`solum` or `type`==`hanshow`<br><br>**Default**: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apUsb := models.ApUsb{
        Cacert:               models.NewOptional(models.ToPointer("cacert2")),
        Channel:              models.ToPointer(3),
        Enabled:              models.ToPointer(false),
        Host:                 models.ToPointer("1.1.1.1"),
        Port:                 models.ToPointer(0),
        Type:                 models.ToPointer(models.ApUsbTypeEnum_IMAGOTAG),
        VlanId:               models.ToPointer(1),
    }

}
```

