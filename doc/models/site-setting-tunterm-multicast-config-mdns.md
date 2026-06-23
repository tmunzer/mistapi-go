
# Site Setting Tunterm Multicast Config Mdns

mDNS multicast forwarding settings for tunneled VLANs

## Structure

`SiteSettingTuntermMulticastConfigMdns`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether mDNS multicast forwarding is enabled<br><br>**Default**: `false` |
| `VlanIds` | `[]int` | Optional | VLAN IDs where mDNS forwarding is enabled |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingTuntermMulticastConfigMdns := models.SiteSettingTuntermMulticastConfigMdns{
        Enabled:              models.ToPointer(false),
        VlanIds:              []int{
            2,
            3,
            5,
        },
    }

}
```

