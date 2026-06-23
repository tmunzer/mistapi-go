
# Response Ssr Upgrade Counts

Device counts grouped by SSR upgrade status

## Structure

`ResponseSsrUpgradeCounts`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Failed` | `int` | Required | Number of SSR devices that failed the upgrade |
| `Queued` | `int` | Required | Number of SSR devices queued for upgrade |
| `Success` | `int` | Required | Number of SSR devices successfully upgraded |
| `Upgrading` | `int` | Required | Number of SSR devices currently upgrading |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSsrUpgradeCounts := models.ResponseSsrUpgradeCounts{
        Failed:               132,
        Queued:               56,
        Success:              124,
        Upgrading:            146,
    }

}
```

