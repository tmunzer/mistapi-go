
# Wlan Dns Server Rewrite

For radius_group-based DNS server (rewrite DNS request depending on the Group RADIUS server returns)

## Structure

`WlanDnsServerRewrite`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether DNS server rewrite by RADIUS group is enabled for this WLAN<br><br>**Default**: `false` |
| `RadiusGroups` | `map[string]string` | Optional | Map between radius_group and the desired DNS server (IPv4 only). Property key is the RADIUS group, property value is the desired DNS Server |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanDnsServerRewrite := models.WlanDnsServerRewrite{
        Enabled:              models.ToPointer(false),
        RadiusGroups:         map[string]string{
            "contractor": "172.1.1.1",
            "guest": "8.8.8.8",
        },
    }

}
```

