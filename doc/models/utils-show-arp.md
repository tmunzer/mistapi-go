
# Utils Show Arp

ARP table lookup request for device command output

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowArp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Refresh duration in seconds; set only when `interval` is nonzero<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 300` |
| `Interval` | `*int` | Optional | Refresh interval in seconds for repeated command output<br><br>**Default**: `0`<br><br>**Constraints**: `>= 0`, `<= 10` |
| `Ip` | `*string` | Optional | Address filter for the ARP table lookup |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `PortId` | `*string` | Optional | Device port identifier filter for the ARP table lookup |
| `Vrf` | `*string` | Optional | Routing instance or VRF filter for the ARP table lookup |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsShowArp := models.UtilsShowArp{
        Duration:             models.ToPointer(0),
        Interval:             models.ToPointer(0),
        Ip:                   models.ToPointer("192.168.30.7"),
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        PortId:               models.ToPointer("ge-0/0/0.0"),
        Vrf:                  models.ToPointer("guest"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

