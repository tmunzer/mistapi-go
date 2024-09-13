
# Utils Show Ospf Neighbors

## Structure

`UtilsShowOspfNeighbors`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Neighbor` | `*string` | Optional | Neighbor IP Address |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `PortId` | `*string` | Optional | the network interface |
| `Vrf` | `*string` | Optional | VRF name |

## Example (as JSON)

```json
{
  "neighbor": "10.1.1.1",
  "port_id": "ge-0/0/3",
  "vrf": "lan",
  "node": "node0"
}
```

