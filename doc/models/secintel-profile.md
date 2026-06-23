
# Secintel Profile

SecIntel profile containing category-specific threat intelligence actions

*This model accepts additional fields of type interface{}.*

## Structure

`SecintelProfile`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Name` | `*string` | Optional | Display name of the SecIntel profile |
| `Profiles` | [`[]models.SecintelProfileProfile`](../../doc/models/secintel-profile-profile.md) | Optional | Category-specific SecIntel action settings |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    secintelProfile := models.SecintelProfile{
        Name:                 models.ToPointer("secintel-custom"),
        Profiles:             []models.SecintelProfileProfile{
            models.SecintelProfileProfile{
                Action:               models.ToPointer(models.SecintelProfileProfileActionEnum_STRICT),
                Category:             models.ToPointer(models.SecintelProfileProfileCategoryEnum_IH),
            },
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

