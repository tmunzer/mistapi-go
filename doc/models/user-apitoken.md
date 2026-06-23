
# User Apitoken

User API token metadata

*This model accepts additional fields of type interface{}.*

## Structure

`UserApitoken`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CreatedTime` | `*float64` | Optional, Read-only | When the object has been created, in epoch |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Key` | `*string` | Optional, Read-only | Secret API token value returned for this token |
| `LastUsed` | `models.Optional[int]` | Optional, Read-only | Time when this user API token was last used, in epoch seconds; null if never used |
| `Name` | `*string` | Optional | Display label for this user API token |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    userApitoken := models.UserApitoken{
        CreatedTime:          models.ToPointer(float64(42.78)),
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Key:                  models.ToPointer("1qkb...QQCL"),
        LastUsed:             models.NewOptional(models.ToPointer(1690115110)),
        Name:                 models.ToPointer("org_token_xyz"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

