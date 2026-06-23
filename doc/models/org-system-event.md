
# Org System Event

Organization system event record

## Structure

`OrgSystemEvent`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ChangeCat` | `*string` | Optional | Category of configuration or administrative change for the event |
| `Metadata` | `*string` | Optional | JSON-encoded event metadata, such as affected object IDs and admin details |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Scope` | `*string` | Optional | Event scope, such as organization or site |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |
| `Type` | `*string` | Optional | System event type, such as `add-wlan` or `delete-wlan` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSystemEvent := models.OrgSystemEvent{
        ChangeCat:            models.ToPointer("admin_action"),
        Metadata:             models.ToPointer("metadata6"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Scope:                models.ToPointer("org"),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Type:                 models.ToPointer("delete-wlan"),
    }

}
```

