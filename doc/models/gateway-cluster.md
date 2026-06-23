
# Gateway Cluster

Gateway HA cluster request or response body

*This model accepts additional fields of type interface{}.*

## Structure

`GatewayCluster`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Nodes` | [`[]models.GatewayClusterNode`](../../doc/models/gateway-cluster-node.md) | Required | Gateway cluster nodes. When replacing a node, one node MAC address must remain the same as the existing cluster<br><br>**Constraints**: *Minimum Items*: `1`, *Maximum Items*: `2`, *Unique Items Required* |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    gatewayCluster := models.GatewayCluster{
        Nodes:                []models.GatewayClusterNode{
            models.GatewayClusterNode{
                Mac:                  "mac0",
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

