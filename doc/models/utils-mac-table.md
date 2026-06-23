
# Utils Mac Table

MAC table lookup filters

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsMacTable`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MacAddress` | `*string` | Optional | Client MAC address filter for the MAC table lookup |
| `PortId` | `*string` | Optional | Interface identifier filter for the MAC table lookup |
| `VlanId` | `*string` | Optional | VLAN identifier filter for the MAC table lookup |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsMacTable := models.UtilsMacTable{
        MacAddress:           models.ToPointer("f8c1165c6400"),
        PortId:               models.ToPointer("ge-0/0/0.0"),
        VlanId:               models.ToPointer("ge-0/0/0.0"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

