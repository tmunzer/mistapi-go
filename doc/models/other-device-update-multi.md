
# Other Device Update Multi

Bulk site assignment update for other devices

*This model accepts additional fields of type interface{}.*

## Structure

`OtherDeviceUpdateMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Macs` | `[]string` | Optional | Other-device MAC addresses included in the bulk update |
| `Op` | [`models.OtherDeviceUpdateOperationEnum`](../../doc/models/other-device-update-operation-enum.md) | Required | The operation being performed. enum: `assign`, `unassign` |
| `SiteId` | `*uuid.UUID` | Optional | Site ID used when assigning other devices |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    otherDeviceUpdateMulti := models.OtherDeviceUpdateMulti{
        Macs:                 []string{
            "macs1",
        },
        Op:                   models.OtherDeviceUpdateOperationEnum_ASSIGN,
        SiteId:               models.ToPointer(uuid.MustParse("000024c2-0000-0000-0000-000000000000")),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

