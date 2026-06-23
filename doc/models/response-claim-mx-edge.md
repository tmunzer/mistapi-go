
# Response Claim Mx Edge

Mist Edge claim response

## Structure

`ResponseClaimMxEdge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Magic` | `string` | Required | Claim magic token returned for the Mist Edge |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseClaimMxEdge := models.ResponseClaimMxEdge{
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Magic:                "magic2",
    }

}
```

