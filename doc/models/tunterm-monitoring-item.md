
# Tunterm Monitoring Item

Monitoring check for tunnel termination reachability

## Structure

`TuntermMonitoringItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Can be ip, ipv6, hostname<br><br>**Constraints**: *Minimum Length*: `1` |
| `Port` | `*int` | Optional | When `protocol`==`tcp`, TCP port checked by the monitoring probe |
| `Protocol` | [`*models.TuntermMonitoringProtocolEnum`](../../doc/models/tunterm-monitoring-protocol-enum.md) | Optional | enum: `arp`, `ping`, `tcp`<br><br>**Constraints**: *Minimum Length*: `1` |
| `SrcVlanId` | `*int` | Optional | Optional source for the monitoring check, vlan_id configured in tunterm_other_ip_configs |
| `Timeout` | `*int` | Optional | Maximum time for this monitoring check, in seconds<br><br>**Default**: `300` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tuntermMonitoringItem := models.TuntermMonitoringItem{
        Host:                 models.ToPointer("10.2.8.15"),
        Port:                 models.ToPointer(80),
        Protocol:             models.ToPointer(models.TuntermMonitoringProtocolEnum_TCP),
        SrcVlanId:            models.ToPointer(5),
        Timeout:              models.ToPointer(300),
    }

}
```

