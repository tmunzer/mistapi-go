
# Utils Traceroute

Request body for running a traceroute utility command

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsTraceroute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Destination IP address, IPv6 address, or hostname for traceroute |
| `Network` | `*string` | Optional | For SSR, source network from which to initiate traceroute<br><br>**Default**: `"internal"` |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `Port` | `*int` | Optional | When `protocol`==`udp`, not supported in SSR. The udp port to use<br><br>**Default**: `33434` |
| `Protocol` | [`*models.UtilsTracerouteProtocolEnum`](../../doc/models/utils-traceroute-protocol-enum.md) | Optional | enum: `icmp` (Only supported by AP/MxEdge), `udp`<br><br>**Default**: `"udp"` |
| `Timeout` | `*int` | Optional | Not supported in SSR. Maximum time in seconds to wait for the response<br><br>**Default**: `60` |
| `UseIpv6` | `*bool` | Optional | Whether to resolve hostname targets over IPv6<br><br>**Default**: `false` |
| `Vrf` | `*string` | Optional | For SRX, routing instance or VRF from which to initiate traceroute; master VRF/RI is used by default |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsTraceroute := models.UtilsTraceroute{
        Host:                 models.ToPointer("host0"),
        Network:              models.ToPointer("internal"),
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        Port:                 models.ToPointer(33434),
        Protocol:             models.ToPointer(models.UtilsTracerouteProtocolEnum_UDP),
        Timeout:              models.ToPointer(60),
        UseIpv6:              models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

