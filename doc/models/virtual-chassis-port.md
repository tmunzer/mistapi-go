
# Virtual Chassis Port

Request to set or delete Virtual Chassis ports on members

*This model accepts additional fields of type interface{}.*

## Structure

`VirtualChassisPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Members` | [`[]models.ConfigVcPortMember`](../../doc/models/config-vc-port-member.md) | Required | Virtual Chassis member port selections for a port operation<br><br>**Constraints**: *Unique Items Required* |
| `Op` | [`models.VirtualChassisPortOperationEnum`](../../doc/models/virtual-chassis-port-operation-enum.md) | Required | Action to perform on the specified Virtual Chassis ports. enum: `delete`, `set`<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    virtualChassisPort := models.VirtualChassisPort{
        Members:              []models.ConfigVcPortMember{
            models.ConfigVcPortMember{
                Member:               float64(188.64),
                VcPorts:              []string{
                    "vc_ports2",
                    "vc_ports3",
                    "vc_ports4",
                },
            },
        },
        Op:                   models.VirtualChassisPortOperationEnum_DELETE,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

