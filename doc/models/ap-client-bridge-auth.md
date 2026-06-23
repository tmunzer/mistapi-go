
# Ap Client Bridge Auth

Authentication settings for the AP client bridge uplink

## Structure

`ApClientBridgeAuth`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Psk` | `*string` | Optional | Pre-shared key used when `type`==`psk` for client bridge authentication<br><br>**Constraints**: *Minimum Length*: `8`, *Maximum Length*: `63` |
| `Type` | [`*models.ApClientBridgeAuthTypeEnum`](../../doc/models/ap-client-bridge-auth-type-enum.md) | Optional | wpa2-AES/CCMPp is assumed when `type`==`psk`. enum: `open`, `psk`<br><br>**Default**: `"psk"`<br><br>**Constraints**: *Minimum Length*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apClientBridgeAuth := models.ApClientBridgeAuth{
        Psk:                  models.ToPointer("foryoureyesonly"),
        Type:                 models.ToPointer(models.ApClientBridgeAuthTypeEnum_PSK),
    }

}
```

