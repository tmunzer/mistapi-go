
# Utils Clear Arp

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearArp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Ip` | `*string` | Optional | The IP address for which to clear an ARP entry. port_id must be specified. |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `PortId` | `*string` | Optional | The device interface on which to clear the ARP cache. |
| `Vlan` | `*int` | Optional | The VLAN on which to clear the ARP cache. port_id must be specified. |
| `Vrf` | `*string` | Optional | The vrf for which to clear an ARP entry. applicable for switch. |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "ip": "10.1.1.1",
  "port_id": "wan",
  "vlan": 1000,
  "vrf": "guest",
  "node": "node0",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

