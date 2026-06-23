
# Orggroup

MSP organization group containing related organizations

*This model accepts additional fields of type interface{}.*

## Structure

`Orggroup`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `ModifiedTime` | `*float64` | Optional, Read-only | When the object has been modified for the last time, in epoch |
| `MspId` | `*uuid.UUID` | Optional, Read-only | Managed service provider identifier |
| `Name` | `string` | Required | Display name of the organization group |
| `OrgIds` | `[]uuid.UUID` | Optional | List of organization identifiers included in an organization group |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orggroup := models.Orggroup{
        CreatedTime:          models.ToPointer(float64(12.46)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        ModifiedTime:         models.ToPointer(float64(66.5)),
        MspId:                models.ToPointer(uuid.MustParse("b9d42c2e-88ee-41f8-b798-f009ce7fe909")),
        Name:                 "name6",
        OrgIds:               []uuid.UUID{
            uuid.MustParse("00001d40-0000-0000-0000-000000000000"),
            uuid.MustParse("00001d41-0000-0000-0000-000000000000"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

