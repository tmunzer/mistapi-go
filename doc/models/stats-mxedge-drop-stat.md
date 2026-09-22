
# Stats Mxedge Drop Stat

Packet drop counters reported by the Mist Edge tunnel termination service. Counters not listed here may be reported as additional properties.

*This model accepts additional fields of type interface{}.*

## Structure

`StatsMxedgeDropStat`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DropIp4MustFragment` | `*int` | Optional | Packets dropped because the IPv4 packet required fragmentation |
| `DropIpsecNoSa` | `*int` | Optional | Packets dropped because no matching IPsec security association was found |
| `DropIpsecUnknownSpi` | `*int` | Optional | Packets dropped because the IPsec SPI was unknown |
| `DropL2tpNoSession` | `*int` | Optional | Packets dropped because no matching L2TP session was found |
| `DropProtectedSvi` | `*int` | Optional | Packets dropped because the destination was a protected SVI |
| `DropVlanNotOnPort` | `*int` | Optional | Packets dropped because the VLAN is not configured on the port |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeDropStat := models.StatsMxedgeDropStat{
        DropIp4MustFragment:  models.ToPointer(2),
        DropIpsecNoSa:        models.ToPointer(1),
        DropIpsecUnknownSpi:  models.ToPointer(0),
        DropL2tpNoSession:    models.ToPointer(0),
        DropProtectedSvi:     models.ToPointer(0),
        DropVlanNotOnPort:    models.ToPointer(0),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

