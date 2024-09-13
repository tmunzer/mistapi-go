
# Utils Show Ospf Database

## Structure

`UtilsShowOspfDatabase`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `SelfOriginate` | `*bool` | Optional | show originating info, default is false<br>**Default**: `false` |
| `Vrf` | `*string` | Optional | VRF name |

## Example (as JSON)

```json
{
  "self_originate": false,
  "vrf": "lan",
  "node": "node0"
}
```

