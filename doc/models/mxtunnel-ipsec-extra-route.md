
# Mxtunnel Ipsec Extra Route

Extra route advertised for an IPsec tunnel

## Structure

`MxtunnelIpsecExtraRoute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dest` | `*string` | Optional | Route destination CIDR for the IPsec extra route |
| `NextHop` | `*string` | Optional | Route next-hop IPv4 address for the IPsec extra route |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxtunnelIpsecExtraRoute := models.MxtunnelIpsecExtraRoute{
        Dest:                 models.ToPointer("dest4"),
        NextHop:              models.ToPointer("next_hop4"),
    }

}
```

