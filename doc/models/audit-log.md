
# Audit Log

Administrative audit log entry

## Structure

`AuditLog`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `uuid.UUID` | Required, Read-only | ID of the administrator |
| `AdminName` | `string` | Required | Name or email of the administrator associated with the audited action |
| `After` | `*interface{}` | Optional | Field values after the change |
| `Before` | `*interface{}` | Optional | Field values prior to the change |
| `ForSite` | `*bool` | Optional, Read-only | Whether this audit log entry is scoped to a site |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Message` | `string` | Required | Human-readable audit message describing the action |
| `OrgId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `uuid.UUID` | Required, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    auditLog := models.AuditLog{
        AdminId:              uuid.MustParse("456b7016-a916-a4b1-78dd-72b947c152b7"),
        AdminName:            "admin_name8",
        After:                models.ToPointer(interface{}("[key1, val1][key2, val2]")),
        Before:               models.ToPointer(interface{}("[key1, val1][key2, val2]")),
        ForSite:              models.ToPointer(false),
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Message:              "message0",
        OrgId:                uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61"),
        SiteId:               uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6"),
        Timestamp:            float64(210.88),
    }

}
```

