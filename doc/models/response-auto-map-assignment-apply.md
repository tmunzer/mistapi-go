
# Response Auto Map Assignment Apply

Result returned after applying accepted auto map assignments

## Structure

`ResponseAutoMapAssignmentApply`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcceptedMaps` | `[]uuid.UUID` | Required | List of map IDs that were successfully accepted |
| `Message` | `string` | Required | Human-readable description of the operation result |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseAutoMapAssignmentApply := models.ResponseAutoMapAssignmentApply{
        AcceptedMaps:         []uuid.UUID{
            uuid.MustParse("000013ff-0000-0000-0000-000000000000"),
            uuid.MustParse("000013fe-0000-0000-0000-000000000000"),
        },
        Message:              "message6",
    }

}
```

