
# Utils Service Ping

Service ping request for SSR devices

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsServicePing`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Count` | `*int` | Optional | Number of ICMP echo requests to send for the service ping<br><br>**Default**: `10` |
| `Host` | `string` | Required | Destination IPv4 address for the service ping |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `Service` | `string` | Required | Ping packet takes the same path as the service |
| `Size` | `*int` | Optional | ICMP payload size, in bytes<br><br>**Default**: `56`<br><br>**Constraints**: `>= 56`, `<= 65535` |
| `Tenant` | `*string` | Optional | Routing tenant context in which the packet is sent |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsServicePing := models.UtilsServicePing{
        Count:                models.ToPointer(10),
        Host:                 "host6",
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        Service:              "service8",
        Size:                 models.ToPointer(56),
        Tenant:               models.ToPointer("tenant8"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

