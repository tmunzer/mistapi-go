
# Auto Map Assignment Request

Request body for accepting or clearing pending map assignments

## Structure

`AutoMapAssignmentRequest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MapIds` | `[]uuid.UUID` | Optional | Optional list of specific map IDs to apply/clear. If not provided or empty, all pending map assignments are accepted/rejected. |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    autoMapAssignmentRequest := models.AutoMapAssignmentRequest{
        MapIds:               []uuid.UUID{
            uuid.MustParse("000007e8-0000-0000-0000-000000000000"),
            uuid.MustParse("000007e9-0000-0000-0000-000000000000"),
        },
    }

}
```

