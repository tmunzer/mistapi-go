
# Utils Clear Bgp

## Structure

`UtilsClearBgp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Neighbor` | `string` | Required | neighbor ip-address or 'all' |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Type` | [`models.UtilsClearBgpTypeEnum`](../../doc/models/utils-clear-bgp-type-enum.md) | Required | enum: `hard`, `in`, `out`, `soft`<br>**Default**: `"hard"` |
| `Vrf` | `*string` | Optional | vrf name |

## Example (as JSON)

```json
{
  "neighbor": "neighbor0",
  "type": "hard",
  "node": "node0",
  "vrf": "vrf4"
}
```

