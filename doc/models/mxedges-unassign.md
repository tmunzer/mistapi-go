
# Mxedges Unassign

Request to unassign Mist Edges from their site

*This model accepts additional fields of type interface{}.*

## Structure

`MxedgesUnassign`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxedgeIds` | `[]uuid.UUID` | Required | Mist Edge identifiers to unassign from their site |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxedgesUnassign := models.MxedgesUnassign{
        MxedgeIds:            []uuid.UUID{
            uuid.MustParse("0000264c-0000-0000-0000-000000000000"),
            uuid.MustParse("0000264d-0000-0000-0000-000000000000"),
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

