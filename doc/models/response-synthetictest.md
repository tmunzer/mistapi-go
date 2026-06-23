
# Response Synthetictest

Result returned after a site synthetic test is queued

## Structure

`ResponseSynthetictest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Message` | `*string` | Optional | Human-readable queuing result returned by the API |
| `Status` | `*string` | Optional | Queue status for the synthetic test request |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSynthetictest := models.ResponseSynthetictest{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Message:              models.ToPointer("Successfully queued synthetic test for the site."),
        Status:               models.ToPointer("success"),
    }

}
```

