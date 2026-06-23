
# Response Self Subscription

Subscription record visible to the current account

## Structure

`ResponseSelfSubscription`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSelfSubscription := models.ResponseSelfSubscription{
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
    }

}
```

