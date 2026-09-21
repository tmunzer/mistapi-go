
# User Macs Count

User MACs count response

## Structure

`UserMacsCount`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `End` | `*int` | Optional | Query end timestamp for user MAC counts |
| `Limit` | `*int` | Optional | Maximum number of distinct count results to return |
| `Results` | [`[]models.UserMac`](../../doc/models/user-mac.md) | Optional | User MAC entries returned by the count query |
| `Start` | `*int` | Optional | Query start timestamp for user MAC counts |
| `Total` | `*int` | Optional | Overall number of user MAC count results |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    userMacsCount := models.UserMacsCount{
        End:                  models.ToPointer(224),
        Limit:                models.ToPointer(54),
        Results:              []models.UserMac{
            models.UserMac{
                Labels:               []string{
                    "labels2",
                    "labels1",
                },
                Mac:                  "mac0",
                Name:                 models.ToPointer("name6"),
                Notes:                models.ToPointer("notes6"),
                RadiusGroup:          models.ToPointer("radius_group8"),
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
            models.UserMac{
                Labels:               []string{
                    "labels2",
                    "labels1",
                },
                Mac:                  "mac0",
                Name:                 models.ToPointer("name6"),
                Notes:                models.ToPointer("notes6"),
                RadiusGroup:          models.ToPointer("radius_group8"),
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
            models.UserMac{
                Labels:               []string{
                    "labels2",
                    "labels1",
                },
                Mac:                  "mac0",
                Name:                 models.ToPointer("name6"),
                Notes:                models.ToPointer("notes6"),
                RadiusGroup:          models.ToPointer("radius_group8"),
                AdditionalProperties: map[string]interface{}{
                    "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
                },
            },
        },
        Start:                models.ToPointer(182),
        Total:                models.ToPointer(216),
    }

}
```

