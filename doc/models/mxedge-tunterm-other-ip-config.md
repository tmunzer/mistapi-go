
# Mxedge Tunterm Other Ip Config

Additional IP configuration for a Mist Tunnel VLAN interface

## Structure

`MxedgeTuntermOtherIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `string` | Required | Address for the additional Mist Tunnel interface, in IPv4 format |
| `Netmask` | `string` | Required | Subnet mask for the additional Mist Tunnel IPv4 address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermOtherIpConfig := models.MxedgeTuntermOtherIpConfig{
        Ip:                   "ip2",
        Netmask:              "netmask8",
    }

}
```

