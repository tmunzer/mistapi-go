
# Utils Show Ospf Summary

## Structure

`UtilsShowOspfSummary`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Vrf` | `*string` | Optional | VRF name |

## Example (as JSON)

```json
{
  "vrf": "lan",
  "node": "node0"
}
```

