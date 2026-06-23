
# Ap Airista

Airista RTLS integration settings for an AP

## Structure

`ApAirista`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable Airista config<br><br>**Default**: `false` |
| `Host` | `models.Optional[string]` | Optional | Required if enabled, Airista server host |
| `Port` | `models.Optional[int]` | Optional | Optional if enabled, Airista server port. Defaults to 1144<br><br>**Default**: `1144` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apAirista := models.ApAirista{
        Enabled:              models.ToPointer(false),
        Host:                 models.NewOptional(models.ToPointer("airista.pvt.net")),
        Port:                 models.NewOptional(models.ToPointer(1144)),
    }

}
```

