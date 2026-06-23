
# Response Ssr Upgrade Status Targets

SSR device IDs grouped by upgrade status

## Structure

`ResponseSsrUpgradeStatusTargets`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Failed` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Queued` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Success` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |
| `Upgrading` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSsrUpgradeStatusTargets := models.ResponseSsrUpgradeStatusTargets{
        Failed:               []string{
            "failed4",
        },
        Queued:               []string{
            "queued6",
        },
        Success:              []string{
            "success0",
            "success1",
            "success2",
        },
        Upgrading:            []string{
            "upgrading7",
        },
    }

}
```

