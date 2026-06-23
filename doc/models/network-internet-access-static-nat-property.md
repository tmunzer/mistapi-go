
# Network Internet Access Static Nat Property

Direct-internet static NAT rule target settings

## Structure

`NetworkInternetAccessStaticNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | The Static NAT destination IP address. Must be an IP address (i.e. "192.168.70.3") or a Variable (i.e. "{{myvar}}") |
| `Name` | `*string` | Optional | Label for this direct internet static NAT rule |
| `WanName` | `*string` | Optional | SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity. Can be a Variable (i.e. "{{myvar}}") |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkInternetAccessStaticNatProperty := models.NetworkInternetAccessStaticNatProperty{
        InternalIp:           models.ToPointer("192.168.70.3"),
        Name:                 models.ToPointer("pos_station-1"),
        WanName:              models.ToPointer("wan0"),
    }

}
```

