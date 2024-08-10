
# Utils Release Dhcp

## Structure

`UtilsReleaseDhcp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | only for HA. enum: `node0`, `node1` |
| `PortId` | `string` | Required | The nework interface on which to release the current DHCP release<br>**Constraints**: *Minimum Length*: `1` |

## Example (as JSON)

```json
{
  "port_id": "ge-0/0/1.10",
  "node": "node0"
}
```

