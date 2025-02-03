
# Utils Show Forwarding Table

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsShowForwardingTable`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Prefix` | `*string` | Optional | IP Prefix |
| `ServiceIp` | `*string` | Optional | Only supported with SSR |
| `ServiceName` | `*string` | Optional | Only supported with SSR |
| `ServicePort` | `*int` | Optional | Only supported with SSR |
| `ServiceProtocol` | `*string` | Optional | Only supported with SSR |
| `ServiceTenant` | `*string` | Optional | Only supported with SSR |
| `Vrf` | `*string` | Optional | VRF Name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "prefix": "3.1.1.0/24",
  "service_ip": "3.1.1.10",
  "service_name": "internet-wan_and_lte",
  "service_port": 32768,
  "service_protocol": "udp",
  "service_tenant": "branch1-wifi-mgt",
  "vrf": "guest",
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

