
# Utils Clear Bgp

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsClearBgp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Neighbor` | `string` | Required | neighbor ip-address or 'all' |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `Type` | [`models.UtilsClearBgpTypeEnum`](../../doc/models/utils-clear-bgp-type-enum.md) | Required | enum: `hard`, `in`, `out`, `soft`<br>**Default**: `"hard"` |
| `Vrf` | `*string` | Optional | vrf name |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "neighbor": "neighbor0",
  "type": "hard",
  "node": "node0",
  "vrf": "vrf4",
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

