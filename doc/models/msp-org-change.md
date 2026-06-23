
# Msp Org Change

Request to assign or unassign orgs for an MSP account

*This model accepts additional fields of type interface{}.*

## Structure

`MspOrgChange`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Op` | [`models.MspOrgChangeOperationEnum`](../../doc/models/msp-org-change-operation-enum.md) | Required | Assignment operation to apply to the listed org IDs. enum: `assign`, `unassign` |
| `OrgIds` | `[]string` | Required | List of org IDs to assign to or unassign from an MSP account |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mspOrgChange := models.MspOrgChange{
        Op:                   models.MspOrgChangeOperationEnum_ASSIGN,
        OrgIds:               []string{
            "org_ids0",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

