
# Vc Port

Request to change the Virtual Chassis port mode on a switch

*This model accepts additional fields of type interface{}.*

## Structure

`VcPort`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mode` | [`*models.VcPortModeEnum`](../../doc/models/vc-port-mode-enum.md) | Optional | enum: `network`, `vcp-higig`, `vcp-hgoe` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vcPort := models.VcPort{
        Mode:                 models.ToPointer(models.VcPortModeEnum_VCPHGOE),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

