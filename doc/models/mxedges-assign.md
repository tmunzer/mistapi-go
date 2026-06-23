
# Mxedges Assign

Request to assign Mist Edges to a site

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgesAssign`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxedgeIds` | `[]uuid.UUID` | Required | Mist Edge identifiers to assign to a site |
| `SiteId` | `uuid.UUID` | Required | Identifier of the site that receives the Mist Edges |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxedgesAssign := models.MxedgesAssign{
        MxedgeIds:            []uuid.UUID{
            uuid.MustParse("0000189a-0000-0000-0000-000000000000"),
            uuid.MustParse("0000189b-0000-0000-0000-000000000000"),
        },
        SiteId:               uuid.MustParse("43e9c864-a7e4-4310-8031-d9817d2c5a43"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

