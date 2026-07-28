
# Response Ssr Upgrade Status

Detailed status for an SSR firmware upgrade job

## Structure

`ResponseSsrUpgradeStatus`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | Firmware release channel used for the SSR upgrade<br><br>**Constraints**: *Minimum Length*: `1` |
| `DeviceType` | `*string` | Optional | Type of devices targeted by the SSR upgrade |
| `Force` | `*bool` | Optional | Whether the upgrade was forced even when the requested version matched the running version<br><br>**Default**: `false` |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Status` | `string` | Required | Current status of the SSR upgrade job<br><br>**Constraints**: *Minimum Length*: `1` |
| `Strategy` | `*string` | Optional | Upgrade strategy used by the SSR upgrade job<br><br>**Constraints**: *Minimum Length*: `1` |
| `Targets` | [`models.ResponseSsrUpgradeStatusTargets`](../../doc/models/response-ssr-upgrade-status-targets.md) | Required | SSR device IDs grouped by upgrade status |
| `Versions` | `interface{}` | Required | SSR firmware versions included in the upgrade job |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSsrUpgradeStatus := models.ResponseSsrUpgradeStatus{
        Channel:              "channel2",
        DeviceType:           models.ToPointer("device_type8"),
        Force:                models.ToPointer(false),
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Status:               "status4",
        Strategy:             models.ToPointer("strategy2"),
        Targets:              models.ResponseSsrUpgradeStatusTargets{
            Failed:               []string{
                "failed6",
            },
            Queued:               []string{
                "queued8",
            },
            Success:              []string{
                "success2",
                "success3",
                "success4",
            },
            Upgrading:            []string{
                "upgrading9",
            },
        },
        Versions:             interface{}("[key1, val1][key2, val2]"),
    }

}
```

