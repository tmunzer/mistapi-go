
# Marvis Config Action

A Marvis-injected config action record

## Structure

`MarvisConfigAction`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdminId` | `*uuid.UUID` | Optional | Admin UUID associated with the config action |
| `Id` | `*uuid.UUID` | Optional | UUID of the config action |
| `Mac` | `*string` | Optional | Device MAC address |
| `Op` | `*string` | Optional | Operation type (e.g. disable_port, enable_port, update_mtu, add_vlans_to_port) |
| `OrgId` | `*uuid.UUID` | Optional | Organization UUID |
| `PortId` | `*string` | Optional | Port identifier (e.g. ge-0/0/13) |
| `Reason` | `*string` | Optional | Reason for the config action (e.g. rogue_dhcp_server_detected) |
| `SiteId` | `*uuid.UUID` | Optional | Site UUID |
| `Src` | `*string` | Optional | Source of the config action (e.g. marvis) |
| `Timestamp` | `*float64` | Optional | Timestamp when the config action was recorded, in epoch seconds |
| `Type` | `*string` | Optional | Config type (e.g. wired) |
| `VlanIds` | `[]int` | Optional | List of VLAN IDs involved in the config action |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    marvisConfigAction := models.MarvisConfigAction{
        AdminId:              models.ToPointer(uuid.MustParse("000000d4-0000-0000-0000-000000000000")),
        Id:                   models.ToPointer(uuid.MustParse("0000247a-0000-0000-0000-000000000000")),
        Mac:                  models.ToPointer("mac2"),
        Op:                   models.ToPointer("op8"),
        OrgId:                models.ToPointer(uuid.MustParse("00002552-0000-0000-0000-000000000000")),
    }

}
```

