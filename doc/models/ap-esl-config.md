
# Ap Esl Config

Electronic shelf label integration settings for an AP

## Structure

`ApEslConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cacert` | `*string` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Channel` | `*int` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Enabled` | `*bool` | Optional | usb_config is ignored if esl_config enabled<br><br>**Default**: `false` |
| `Host` | `*string` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Port` | `*int` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `Type` | [`*models.ApEslTypeEnum`](../../doc/models/ap-esl-type-enum.md) | Optional | note: ble_config will be ignored if esl_config is enabled and with native mode. enum: `hanshow`, `imagotag`, `native`, `solum` |
| `VerifyCert` | `*bool` | Optional | Only if `type`==`imagotag` or `type`==`native` |
| `VlanId` | `*int` | Optional | Only if `type`==`solum` or `type`==`hanshow`<br><br>**Default**: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apEslConfig := models.ApEslConfig{
        Cacert:               models.ToPointer("--BEGIN CERTIFICATE--\nMIIDXTCCAkWgAwIBAgIJAL5b1z4f3k2TMA0GCSqGSIb3DQEBCwUAMIGVMQsw\n"),
        Channel:              models.ToPointer(3),
        Enabled:              models.ToPointer(false),
        Host:                 models.ToPointer("1.1.1.1"),
        Port:                 models.ToPointer(0),
        Type:                 models.ToPointer(models.ApEslTypeEnum_IMAGOTAG),
        VerifyCert:           models.ToPointer(true),
        VlanId:               models.ToPointer(1),
    }

}
```

