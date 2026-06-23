
# Alarm

Additional information per alarm type

## Structure

`Alarm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AckAdminId` | `*uuid.UUID` | Optional, Read-only | UUID of the admin who acked the alarm |
| `AckAdminName` | `*string` | Optional | Name & Email ID of the admin who acked the alarm |
| `Acked` | `*bool` | Optional | Whether the alarm is acked or not |
| `AckedTime` | `*int` | Optional, Read-only | Epoch (seconds) when the alarm was acked |
| `Aps` | `[]string` | Optional | additional information: List of MACs of the APs |
| `Bssids` | `[]string` | Optional | Wireless BSSIDs related to this alarm |
| `Count` | `int` | Required, Read-only | Number of incident within an alarm window |
| `Gateways` | `[]string` | Optional | additional information: List of MACs of the gateways |
| `Group` | `string` | Required | Category group for this alarm, such as certificate_expiry, infrastructure, marvis, or security |
| `Hostnames` | `[]string` | Optional | additional information: List of Hostnames of the devices (AP/Switch/Gateway) |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `LastSeen` | `float64` | Required, Read-only | Epoch (seconds) of the last incident/alarm within an alarm window |
| `Note` | `*string` | Optional | Text describing the alarm |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `ResolvedTime` | `*int` | Optional | Epoch (seconds) of the resolved_time for the alarm |
| `Severity` | `string` | Required | Impact level assigned to this alarm, such as critical, warn, or info |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Ssids` | `[]string` | Optional | Wireless SSIDs related to this alarm |
| `Status` | [`*models.AlarmStatusEnum`](../../doc/models/alarm-status-enum.md) | Optional | Current lifecycle status of this alarm. enum: `open`, `resolved` |
| `Switches` | `[]string` | Optional | additional information: List of MACs of the switches |
| `Timestamp` | `float64` | Required, Read-only | Epoch timestamp, in seconds |
| `Type` | `string` | Required, Read-only | Key-name of the alarm type |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    alarm := models.Alarm{
        AckAdminId:           models.ToPointer(uuid.MustParse("456b7016-a916-a4b1-78dd-72b947c152b7")),
        AckAdminName:         models.ToPointer("Joe"),
        Acked:                models.ToPointer(true),
        AckedTime:            models.ToPointer(1711031352),
        Aps:                  []string{
            "ffeeddccbbaa",
            "ffeeddccbbab",
        },
        Count:                2,
        Gateways:             []string{
            "ffeeddccbbaa",
            "ffeeddccbbab",
        },
        Group:                "security",
        Hostnames:            []string{
            "MC_DavidL",
            "MCM_AP_33_Nishant",
        },
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        LastSeen:             float64(1711031774),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        Severity:             "critical",
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
        Switches:             []string{
            "ffeeddccbbaa",
            "ffeeddccbbab",
        },
        Timestamp:            float64(237.82),
        Type:                 "rogue_client",
    }

}
```

