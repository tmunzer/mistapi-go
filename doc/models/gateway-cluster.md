
# Gateway Cluster

## Structure

`GatewayCluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Nodes` | [`[]models.GatewayClusterNode`](../../doc/models/gateway-cluster-node.md) | Required | when replacing a node, either mac has to remain the same as existing cluster<br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `2`, *Unique Items Required* |

## Example (as JSON)

```json
{
  "nodes": [
    {
      "mac": "mac0"
    }
  ]
}
```

