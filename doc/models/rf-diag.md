
# Rf Diag

RF diagnostic recording request

*This model accepts additional fields of type interface{}.*

## Structure

`RfDiag`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Recording length, in seconds; maximum and default are 180<br><br>**Default**: `180`<br><br>**Constraints**: `<= 180` |
| `Mac` | `*string` | Optional | If `type`==`client` or `asset`, MAC address of the device being recorded |
| `Name` | `string` | Required | Recording name; the SDK client name is a good default for SDK-client recordings |
| `SdkclientId` | `*uuid.UUID` | Optional | If `type`==`sdkclient`, SDK client identifier for this recording |
| `Type` | [`models.RfClientTypeEnum`](../../doc/models/rf-client-type-enum.md) | Required | enum: `asset`, `client`, `sdkclient` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    rfDiag := models.RfDiag{
        Duration:             models.ToPointer(180),
        Mac:                  models.ToPointer("mac0"),
        Name:                 "name6",
        SdkclientId:          models.ToPointer(uuid.MustParse("00000928-0000-0000-0000-000000000000")),
        Type:                 models.RfClientTypeEnum_CLIENT,
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

