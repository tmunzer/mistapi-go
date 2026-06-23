
# Site Setting Critical Url Monitoring

Critical URLs whose latency is measured and included in site health

## Structure

`SiteSettingCriticalUrlMonitoring`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether critical URL monitoring is enabled<br><br>**Default**: `true` |
| `Monitors` | [`[]models.SiteSettingCriticalUrlMonitoringMonitor`](../../doc/models/site-setting-critical-url-monitoring-monitor.md) | Optional | Critical URL monitor definitions for site health |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingCriticalUrlMonitoring := models.SiteSettingCriticalUrlMonitoring{
        Enabled:              models.ToPointer(true),
        Monitors:             []models.SiteSettingCriticalUrlMonitoringMonitor{
            models.SiteSettingCriticalUrlMonitoringMonitor{
                Url:                  models.ToPointer("url0"),
                VlanId:               models.ToPointer(models.VlanIdWithVariableContainer.FromString("String5")),
            },
            models.SiteSettingCriticalUrlMonitoringMonitor{
                Url:                  models.ToPointer("url0"),
                VlanId:               models.ToPointer(models.VlanIdWithVariableContainer.FromString("String5")),
            },
        },
    }

}
```

