
# User Mac

Organization user MAC entry

*This model accepts additional fields of type interface{}.*

## Structure

`UserMac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Labels` | `[]string` | Optional | Labels applied to a user MAC entry |
| `Mac` | `string` | Required | Client MAC address for this entry. Only non-local-admin MAC addresses are accepted |
| `Name` | `*string` | Optional | Display name for this user MAC entry |
| `Notes` | `*string` | Optional | Free-form notes about this user MAC entry |
| `RadiusGroup` | `*string` | Optional | RADIUS group associated with this user MAC entry |
| `Vlan` | `*string` | Optional | Network VLAN value associated with this user MAC entry |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    userMac := models.UserMac{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Labels:               []string{
            "byod",
            "flr1",
        },
        Mac:                  "5684dae9ac8b",
        Name:                 models.ToPointer("Printer2"),
        Notes:                models.ToPointer("MAC address refers to Canon printers"),
        RadiusGroup:          models.ToPointer("VIP"),
        Vlan:                 models.ToPointer("30"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

