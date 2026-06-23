
# Dhcp Snooping

DHCP snooping security settings

## Structure

`DhcpSnooping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AllNetworks` | `*bool` | Optional | Whether DHCP snooping applies to all configured networks |
| `EnableArpSpoofCheck` | `*bool` | Optional | Enable for dynamic ARP inspection check |
| `EnableIpSourceGuard` | `*bool` | Optional | Enable for check for forging source IP address |
| `Enabled` | `*bool` | Optional | Whether DHCP snooping is enabled |
| `Networks` | `[]string` | Optional | If `all_networks`==`false`, list of network with DHCP snooping enabled |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpSnooping := models.DhcpSnooping{
        AllNetworks:          models.ToPointer(false),
        EnableArpSpoofCheck:  models.ToPointer(false),
        EnableIpSourceGuard:  models.ToPointer(false),
        Enabled:              models.ToPointer(false),
        Networks:             []string{
            "networks4",
        },
    }

}
```

