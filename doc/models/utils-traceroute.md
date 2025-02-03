
# Utils Traceroute

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsTraceroute`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Host name |
| `Network` | `*string` | Optional | For SSR, optional, the source to initiate traceroute from<br>**Default**: `"internal"` |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Port` | `*int` | Optional | When `protocol`==`udp`, not supported in SSR. The udp port to use<br>**Default**: `33434` |
| `Protocol` | [`*models.UtilsTracerouteProtocolEnum`](../../doc/models/utils-traceroute-protocol-enum.md) | Optional | enum: `icmp` (Only suported by AP/MxEdge), `udp`<br>**Default**: `"udp"` |
| `Timeout` | `*int` | Optional | Not supported in SSR. Maximum time in seconds to wait for the response<br>**Default**: `60` |
| `Vrf` | `*string` | Optional | For SRX, optional, the source to initiate traceroute from. by default, master VRF/RI is assumed |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "network": "internal",
  "port": 33434,
  "protocol": "udp",
  "timeout": 60,
  "host": "host2",
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

