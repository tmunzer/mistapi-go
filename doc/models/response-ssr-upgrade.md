
# Response Ssr Upgrade

SSR firmware upgrade job summary

## Structure

`ResponseSsrUpgrade`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Channel` | `string` | Required | Firmware release channel used for the SSR upgrade<br><br>**Constraints**: *Minimum Length*: `1` |
| `Counts` | [`models.ResponseSsrUpgradeCounts`](../../doc/models/response-ssr-upgrade-counts.md) | Required | Device counts grouped by SSR upgrade status |
| `DeviceType` | `string` | Required | Type of devices targeted by the SSR upgrade |
| `Id` | `uuid.UUID` | Required, Read-only | Unique ID of the object instance in the Mist Organization |
| `Status` | `string` | Required | Current status of the SSR upgrade job<br><br>**Constraints**: *Minimum Length*: `1` |
| `Strategy` | `string` | Required | Upgrade strategy used by the SSR upgrade job<br><br>**Constraints**: *Minimum Length*: `1` |
| `Versions` | `map[string]string` | Required | SSR firmware versions included in the upgrade job |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    responseSsrUpgrade := models.ResponseSsrUpgrade{
        Channel:              "channel6",
        Counts:               models.ResponseSsrUpgradeCounts{
            Failed:               166,
            Queued:               234,
            Success:              90,
            Upgrading:            112,
        },
        DeviceType:           "device_type2",
        Id:                   uuid.MustParse("53f10664-3ce8-4c27-b382-0ef66432349f"),
        Status:               "status0",
        Strategy:             "strategy8",
        Versions:             map[string]string{
            "key0": "versions3",
        },
    }

}
```

