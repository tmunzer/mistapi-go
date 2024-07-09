
# Gateway Cluster Form

## Structure

`GatewayClusterForm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Nodes` | [`[]models.GatewayClusterFormNode`](../../doc/models/gateway-cluster-form-node.md) | Required | **Constraints**: *Minimum Items*: `1`, *Unique Items Required* |

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

