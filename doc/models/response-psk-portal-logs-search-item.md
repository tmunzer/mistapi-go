
# Response Psk Portal Logs Search Item

PSK Portal log entry returned by organization log search

## Structure

`ResponsePskPortalLogsSearchItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Id` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Message` | `*string` | Optional | Human-readable message describing the PSK Portal action |
| `NameId` | `*string` | Optional | SSO NameID value associated with the PSK Portal action |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `PskId` | `*uuid.UUID` | Optional | Identifier of the PSK associated with the log entry |
| `PskName` | `*string` | Optional | Display name of the PSK associated with the log entry |
| `PskportalId` | `*uuid.UUID` | Optional | Identifier of the PSK Portal associated with the log entry |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responsePskPortalLogsSearchItem := models.ResponsePskPortalLogsSearchItem{
        Id:                   models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Message:              models.ToPointer("Rotate PSK test@mist.com"),
        NameId:               models.ToPointer("test@mist.com"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        PskId:                models.ToPointer(uuid.MustParse("608fe603-f9f0-4ce9-9473-04ef6c6ea749")),
        PskName:              models.ToPointer("test@mist.com"),
        PskportalId:          models.ToPointer(uuid.MustParse("c1742c09-af35-4161-96ef-7dc65c6d5674")),
    }

}
```

