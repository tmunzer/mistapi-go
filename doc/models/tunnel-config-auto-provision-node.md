
# Tunnel Config Auto Provision Node

Auto-provisioned tunnel endpoint settings

*This model accepts additional fields of type interface{}.*

## Structure

`TunnelConfigAutoProvisionNode`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ProbeIps` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `WanNames` | `[]string` | Optional | Optional, only needed if `vars_only`==`false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigAutoProvisionNode := models.TunnelConfigAutoProvisionNode{
        ProbeIps:             []string{
            "probe_ips1",
        },
        WanNames:             []string{
            "wan_names2",
            "wan_names3",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

