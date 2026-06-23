
# Wxlan Tunnel Ipsec

IPSec-related configurations; requires DMVPN be enabled

## Structure

`WxlanTunnelIpsec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether ipsec is enabled, requires DMVPN be enabled<br><br>**Default**: `false` |
| `Psk` | `string` | Required | Pre-shared key used for IPsec on this WxLAN tunnel |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wxlanTunnelIpsec := models.WxlanTunnelIpsec{
        Enabled:              models.ToPointer(false),
        Psk:                  "psk0",
    }

}
```

