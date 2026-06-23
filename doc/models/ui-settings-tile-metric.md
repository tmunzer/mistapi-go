
# Ui Settings Tile Metric

Metric selected for a site UI databoard tile

## Structure

`UiSettingsTileMetric`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ApiName` | `*string` | Optional | Metric API name requested by this tile |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    uiSettingsTileMetric := models.UiSettingsTileMetric{
        ApiName:              models.ToPointer("client_dhcp_latency"),
    }

}
```

