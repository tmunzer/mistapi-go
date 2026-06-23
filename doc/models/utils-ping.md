
# Utils Ping

Request body for running a ping utility command

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | Number of echo requests to send<br><br>**Default**: `10` |
| `EgressInterface` | `*string` | Optional | Interface through which ping packets should egress |
| `Host` | `string` | Required | Destination IP address, IPv6 address, or hostname to ping |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `Size` | `*int` | Optional | ICMP payload size, in bytes<br><br>**Default**: `56`<br><br>**Constraints**: `>= 56`, `<= 65535` |
| `UseIpv6` | `*bool` | Optional | applicable when host is hostname<br><br>**Default**: `false` |
| `Vrf` | `*string` | Optional | Routing instance or VRF through which ping packets are sent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsPing := models.UtilsPing{
        Count:                models.ToPointer(10),
        EgressInterface:      models.ToPointer("egress_interface8"),
        Host:                 "1.1.1.1",
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        Size:                 models.ToPointer(56),
        UseIpv6:              models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

