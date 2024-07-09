
# Utils Ping

## Structure

`UtilsPing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | **Default**: `10` |
| `EgressInterface` | `*string` | Optional | Interface through which packet needs to egress |
| `Host` | `string` | Required | - |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA |
| `Size` | `*int` | Optional | **Default**: `56`<br>**Constraints**: `>= 56`, `<= 65535` |

## Example (as JSON)

```json
{
  "count": 10,
  "host": "1.1.1.1",
  "size": 56,
  "egress_interface": "egress_interface8",
  "node": "node0"
}
```
