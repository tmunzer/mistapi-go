
# Evpn Topology Switch Config Dhcpd Config

DHCP server enablement for an EVPN topology switch override

## Structure

`EvpnTopologySwitchConfigDhcpdConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | If DHCPD is enabled on the switch |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnTopologySwitchConfigDhcpdConfig := models.EvpnTopologySwitchConfigDhcpdConfig{
        Enabled:              models.ToPointer(false),
    }

}
```

