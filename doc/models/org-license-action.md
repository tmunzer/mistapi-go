
# Org License Action

Request to amend, annotate, delete, or unamend an organization license

*This model accepts additional fields of type interface{}.*

## Structure

`OrgLicenseAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AmendmentId` | `*uuid.UUID` | Optional | If `op`==`unamend`, the ID of the operation to cancel |
| `DstOrgId` | `*uuid.UUID` | Optional | If `op`==`amend`, the id of the org where the license is moved |
| `Notes` | `*string` | Optional | If `op`==`annotate`, note text to attach to the license |
| `Op` | [`models.OrgLicenseActionOperationEnum`](../../doc/models/org-license-action-operation-enum.md) | Required | to move a license, use the `amend` operation. enum: `amend`, `annotate`, `delete`, `unamend` |
| `Quantity` | `*int` | Optional | If `op`==`amend`, the number of licenses to move |
| `SubscriptionId` | `*string` | Optional | If `op`==`amend` or `op`==`delete`, the ID of the subscription to use |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgLicenseAction := models.OrgLicenseAction{
        AmendmentId:          models.ToPointer(uuid.MustParse("00001b48-0000-0000-0000-000000000000")),
        DstOrgId:             models.ToPointer(uuid.MustParse("000026d6-0000-0000-0000-000000000000")),
        Notes:                models.ToPointer("notes8"),
        Op:                   models.OrgLicenseActionOperationEnum_AMEND,
        Quantity:             models.ToPointer(22),
        SubscriptionId:       models.ToPointer("subscription_id8"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

