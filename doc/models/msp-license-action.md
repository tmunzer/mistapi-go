
# Msp License Action

License operation request for an MSP account

*This model accepts additional fields of type interface{}.*

## Structure

`MspLicenseAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmendmentId` | `*string` | Optional | Required if `op`==`unamend`; identifies the license amendment to undo |
| `DstOrgId` | `*uuid.UUID` | Optional | Required if `op`==`amend`; destination org ID that receives the amended license quantity |
| `Notes` | `*string` | Optional | Required if `op`==`annotate`; note text to attach to the license action |
| `Op` | [`models.MspLicenseActionOperationEnum`](../../doc/models/msp-license-action-operation-enum.md) | Required | enum: `amend`, `annotate`, `delete`, `unamend`<br><br>**Constraints**: *Minimum Length*: `1` |
| `Quantity` | `*float64` | Optional | Required if `op`==`amend`; license quantity to move to the destination org |
| `SubscriptionId` | `*string` | Optional | Required if `op`==`annotate`; subscription ID for the license action<br><br>**Constraints**: *Minimum Length*: `1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mspLicenseAction := models.MspLicenseAction{
        AmendmentId:          models.ToPointer("amendment_id0"),
        DstOrgId:             models.ToPointer(uuid.MustParse("000011a4-0000-0000-0000-000000000000")),
        Notes:                models.ToPointer("notes4"),
        Op:                   models.MspLicenseActionOperationEnum_DELETE,
        Quantity:             models.ToPointer(float64(182)),
        SubscriptionId:       models.ToPointer("subscription_id4"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

