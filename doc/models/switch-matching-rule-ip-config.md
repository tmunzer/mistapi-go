
# Switch Matching Rule Ip Config

In-Band Management interface configuration

## Structure

`SwitchMatchingRuleIpConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Network` | `*string` | Optional | VLAN Name for the management interface |
| `Type` | [`*models.IpTypeEnum`](../../doc/models/ip-type-enum.md) | Optional | IP address assignment mode, either DHCP or static. enum: `dhcp`, `static`<br><br>**Default**: `"dhcp"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchMatchingRuleIpConfig := models.SwitchMatchingRuleIpConfig{
        Network:              models.ToPointer("network8"),
        Type:                 models.ToPointer(models.IpTypeEnum_STATIC),
    }

}
```

