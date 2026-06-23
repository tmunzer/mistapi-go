
# Ap Aeroscout

AeroScout location integration settings applied to an AP or AP profile

## Structure

`ApAeroscout`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable aeroscout config<br><br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | Required if enabled, aeroscout server host |
| `LocateConnected` | `*bool` | Optional | Whether to enable the feature to allow wireless clients data received and sent to AES server for location calculation<br><br>**Default**: `false` |
| `Port` | `models.Optional[int]` | Optional | Optional if enabled, Aeroscout server port. Defaults to 1144<br><br>**Default**: `1144` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apAeroscout := models.ApAeroscout{
        Enabled:              models.ToPointer(false),
        Host:                 models.NewOptional(models.ToPointer("aero.pvt.net")),
        LocateConnected:      models.ToPointer(false),
        Port:                 models.NewOptional(models.ToPointer(1144)),
    }

}
```

