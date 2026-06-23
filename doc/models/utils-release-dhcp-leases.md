
# Utils Release Dhcp Leases

Note:

* valid combinations for Junos:
  * `port_id`
  * `macs` + `network`
* valid combinations for SSR:
  * `port_id`
  * `macs` + `network`
  * `port_id` + `network`
  * `network`
* if network or port_id is specified and macs is empty, it means all clients under network or port_id

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsReleaseDhcpLeases`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Macs` | `[]string` | Optional | Client MAC addresses whose DHCP leases should be released |
| `Network` | `*string` | Optional | DHCP network containing the leases to release |
| `Node` | [`*models.HaClusterNodeEnum`](../../doc/models/ha-cluster-node-enum.md) | Optional | HA cluster node selector. enum: `node0`, `node1` |
| `PortId` | `string` | Required | Network interface containing the DHCP leases to release<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsReleaseDhcpLeases := models.UtilsReleaseDhcpLeases{
        Macs:                 []string{
            "90ec77aabbcc",
            "90ec77aabbdd",
        },
        Network:              models.ToPointer("guest"),
        Node:                 models.ToPointer(models.HaClusterNodeEnum_NODE0),
        PortId:               "ge-0/0/1.10",
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

