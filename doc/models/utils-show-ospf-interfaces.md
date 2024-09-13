
# Utils Show Ospf Interfaces

## Structure

`UtilsShowOspfInterfaces`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `PortId` | `*string` | Optional | the network interface |
| `Vrf` | `*string` | Optional | VRF name |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/3",
  "vrf": "lan",
  "node": "node0"
}
```

