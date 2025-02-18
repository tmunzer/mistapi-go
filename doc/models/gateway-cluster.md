
# Gateway Cluster

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayCluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Nodes` | [`[]models.GatewayClusterNode`](../../doc/models/gateway-cluster-node.md) | Required | When replacing a node, either mac has to remain the same as existing cluster<br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `2`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example (as JSON)

```json
{
  "nodes": [
    {
      "mac": "mac0",
      "exampleAdditionalProperty": {
        "key1": "val1",
        "key2": "val2"
      }
    }
  ],
  "exampleAdditionalProperty": {
    "key1": "val1",
    "key2": "val2"
  }
}
```

