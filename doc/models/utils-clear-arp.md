
# Utils Clear Arp

Request to clear ARP entries on a device

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearArp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | The IP address for which to clear an ARP entry. port_id must be specified. |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `PortId` | `*string` | Optional | The device interface on which to clear the ARP cache. |
| `Vlan` | `*int` | Optional | The VLAN on which to clear the ARP cache. port_id must be specified. |
| `Vrf` | `*string` | Optional | The vrf for which to clear an ARP entry. applicable for switch. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsClearArp := models.UtilsClearArp{
        Ip:                   models.ToPointer("10.1.1.1"),
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        PortId:               models.ToPointer("wan"),
        Vlan:                 models.ToPointer(1000),
        Vrf:                  models.ToPointer("guest"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

