
# Response Auto Map Assignment Clear

Result returned after clearing auto map assignment candidates

## Structure

`ResponseAutoMapAssignmentClear`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Message` | `string` | Required | Human-readable description of the operation result |
| `RejectedMaps` | `[]uuid.UUID` | Required | List of map IDs that were successfully rejected |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseAutoMapAssignmentClear := models.ResponseAutoMapAssignmentClear{
        Message:              "message8",
        RejectedMaps:         []uuid.UUID{
            uuid.MustParse("0000172a-0000-0000-0000-000000000000"),
            uuid.MustParse("0000172b-0000-0000-0000-000000000000"),
            uuid.MustParse("0000172c-0000-0000-0000-000000000000"),
        },
    }

}
```

