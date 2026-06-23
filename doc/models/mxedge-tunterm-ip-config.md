
# Mxedge Tunterm Ip Config

IP configuration for the Mist Tunnel interface

## Structure

`MxedgeTuntermIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Gateway` | `string` | Required | IPv4 gateway for the Mist Tunnel interface |
| `Gateway6` | `*string` | Optional | IPv6 gateway for the Mist Tunnel interface |
| `Ip` | `string` | Required | Address on the untagged Mist Tunnel interface, in IPv4 format |
| `Ip6` | `*string` | Optional | Address on the Mist Tunnel interface, in IPv6 format |
| `Netmask` | `string` | Required | Subnet mask for the Mist Tunnel IPv4 address |
| `Netmask6` | `*string` | Optional | Prefix length for the Mist Tunnel IPv6 address |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermIpConfig := models.MxedgeTuntermIpConfig{
        Gateway:              "10.2.1.254",
        Gateway6:             models.ToPointer("2001:1010:1010:1010::1"),
        Ip:                   "10.2.1.1",
        Ip6:                  models.ToPointer("2001:1010:1010:1010::2"),
        Netmask:              "255.255.255.0",
        Netmask6:             models.ToPointer("/64"),
    }

}
```

