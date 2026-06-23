
# Switch Matching Rule Oob Ip Config

Out-of-Band Management interface configuration

## Structure

`SwitchMatchingRuleOobIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |
| `UseMgmtVrf` | `*bool` | Optional | If supported on the platform. If enabled, DNS will be using this routing-instance, too<br><br>**Default**: `false` |
| `UseMgmtVrfForHostOut` | `*bool` | Optional | For host-out traffic (NTP/TACPLUS/RADIUS/SYSLOG/SNMP), if alternative source network/ip is desired<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMatchingRuleOobIpConfig := models.SwitchMatchingRuleOobIpConfig{
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
        UseMgmtVrf:           models.ToPointer(false),
        UseMgmtVrfForHostOut: models.ToPointer(false),
    }

}
```

