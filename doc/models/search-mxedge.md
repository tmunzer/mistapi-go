
# Search Mxedge

Mist Edge record returned by search APIs

## Structure

`SearchMxedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Distro` | `*string` | Optional | Linux distribution codename running on the Mist Edge |
| `LastSeen` | `*float64` | Optional | Time when the Mist Edge was last observed, in epoch seconds |
| `Model` | `*string` | Optional | Mist Edge hardware or VM model |
| `MxclusterId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `MxedgeId` | `*uuid.UUID` | Optional, Read-only | Unique ID of the object instance in the Mist Organization |
| `Name` | `*string` | Optional | Display name of the Mist Edge |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `TuntermVersion` | `*string` | Optional | Tunnel termination service version running on the Mist Edge |
| `Uptime` | `*int` | Optional | Number of seconds the Mist Edge has been running since last boot |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    searchMxedge := models.SearchMxedge{
        Distro:               models.ToPointer("distro2"),
        LastSeen:             models.ToPointer(float64(2.52)),
        Model:                models.ToPointer("ME-VM"),
        MxclusterId:          models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        MxedgeId:             models.ToPointer(uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f")),
        Name:                 models.ToPointer("me-vm-1"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

