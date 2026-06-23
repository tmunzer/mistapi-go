
# Org Setting Gateway Mgmt Overlay Ip

Overlay IP configuration used for gateway management traffic

## Structure

`OrgSettingGatewayMgmtOverlayIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | When it's going overlay, a routable IP to overlay will be required |
| `Node1Ip` | `*string` | Optional | For SSR HA cluster, another IP for node1 will be required, too |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingGatewayMgmtOverlayIp := models.OrgSettingGatewayMgmtOverlayIp{
        Ip:                   models.ToPointer("ip4"),
        Node1Ip:              models.ToPointer("node1_ip2"),
    }

}
```

