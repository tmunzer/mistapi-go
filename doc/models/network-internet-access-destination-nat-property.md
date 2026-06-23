
# Network Internet Access Destination Nat Property

Direct-internet destination NAT rule target settings

## Structure

`NetworkInternetAccessDestinationNatProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `InternalIp` | `*string` | Optional | The Destination NAT destination IP address. Must be an IP (i.e. "192.168.70.30") or a Variable (i.e. "{{myvar}}") |
| `Name` | `*string` | Optional | Label for this direct internet destination NAT rule |
| `Port` | `*string` | Optional | The Destination NAT destination IP address. Must be a Port (i.e. "443") or a Variable (i.e. "{{myvar}}") |
| `WanName` | `*string` | Optional | SRX Only. If not set, we configure the nat policies against all WAN ports for simplicity |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    networkInternetAccessDestinationNatProperty := models.NetworkInternetAccessDestinationNatProperty{
        InternalIp:           models.ToPointer("192.168.70.30"),
        Name:                 models.ToPointer("web server"),
        Port:                 models.ToPointer("443"),
        WanName:              models.ToPointer("wan0"),
    }

}
```

