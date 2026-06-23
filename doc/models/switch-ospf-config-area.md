
# Switch Ospf Config Area

Settings for a single OSPF area on a switch

*This model accepts additional fields of type interface{}.*

## Structure

`SwitchOspfConfigArea`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `NoSummary` | `*bool` | Optional | Disable OSPF summary routes for this area<br><br>**Default**: `false` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchOspfConfigArea := models.SwitchOspfConfigArea{
        NoSummary:            models.ToPointer(false),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

