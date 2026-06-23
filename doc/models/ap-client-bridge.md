
# Ap Client Bridge

AP client bridge mode configuration

## Structure

`ApClientBridge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Auth` | [`*models.ApClientBridgeAuth`](../../doc/models/ap-client-bridge-auth.md) | Optional | Authentication settings for the AP client bridge uplink |
| `Enabled` | `*bool` | Optional | When acted as client bridge:<br><br>* only 5G radio can be used<br>* will not serve as AP on any radios<br><br>**Default**: `false` |
| `Ssid` | `*string` | Optional | Uplink SSID used by the AP when client bridge mode is enabled<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apClientBridge := models.ApClientBridge{
        Auth:                 models.ToPointer(models.ApClientBridgeAuth{
            Psk:                  models.ToPointer("psk4"),
            Type:                 models.ToPointer(models.ApClientBridgeAuthTypeEnum_OPEN),
        }),
        Enabled:              models.ToPointer(false),
        Ssid:                 models.ToPointer("Uplink-SSID"),
    }

}
```

