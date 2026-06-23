
# Utils Show Forwarding Table

Forwarding table lookup request for device command output

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowForwardingTable`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `Prefix` | `*string` | Optional | IPv4 or IPv6 prefix filter for forwarding table entries |
| `ServiceIp` | `*string` | Optional | Only supported with SSR |
| `ServiceName` | `*string` | Optional | Only supported with SSR |
| `ServicePort` | `*int` | Optional | Only supported with SSR |
| `ServiceProtocol` | `*string` | Optional | Only supported with SSR |
| `ServiceTenant` | `*string` | Optional | Only supported with SSR |
| `Vrf` | `*string` | Optional | Routing instance or VRF filter for forwarding table entries |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowForwardingTable := models.UtilsShowForwardingTable{
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        Prefix:               models.ToPointer("3.1.1.0/24"),
        ServiceIp:            models.ToPointer("3.1.1.10"),
        ServiceName:          models.ToPointer("internet-wan_and_lte"),
        ServicePort:          models.ToPointer(32768),
        ServiceProtocol:      models.ToPointer("udp"),
        ServiceTenant:        models.ToPointer("branch1-wifi-mgt"),
        Vrf:                  models.ToPointer("guest"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

