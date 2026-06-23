
# Org Event

Event record generated at the organization level

## Structure

`OrgEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Text` | `*string` | Optional | Detailed human-readable message for the organization event |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | Event type code for this organization event |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgEvent := models.OrgEvent{
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Text:                 models.ToPointer("authentication failed, API key invalid"),
        Timestamp:            models.ToPointer(float64(14.8)),
        Type:                 models.ToPointer("CRADLEPOINT_SYNC_FAILED"),
    }

}
```

