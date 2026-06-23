
# Site Setting Critical Url Monitoring Monitor

Critical URL monitor definition for site health

## Structure

`SiteSettingCriticalUrlMonitoringMonitor`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Url` | `*string` | Optional | Monitored HTTP or HTTPS URL used for site health latency |
| `VlanId` | [`*models.VlanIdWithVariable`](../../doc/models/containers/vlan-id-with-variable.md) | Optional | VLAN ID, either numeric or expressed as a template variable string |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingCriticalUrlMonitoringMonitor := models.SiteSettingCriticalUrlMonitoringMonitor{
        Url:                  models.ToPointer("http://50.1.3.5:8080"),
        VlanId:               models.ToPointer(models.VlanIdWithVariableContainer.FromString("String9")),
    }

}
```

