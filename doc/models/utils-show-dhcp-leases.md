
# Utils Show Dhcp Leases

## Structure

`UtilsShowDhcpLeases`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Network` | `string` | Required | DHCP network for the leases, returns full table if not specified |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |

## Example (as JSON)

```json
{
  "network": "guest",
  "node": "node0"
}
```

