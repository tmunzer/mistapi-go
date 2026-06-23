
# Site Setting Tunterm Multicast Config

Multicast forwarding settings for tunnel termination at the site

## Structure

`SiteSettingTuntermMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mdns` | [`*models.SiteSettingTuntermMulticastConfigMdns`](../../doc/models/site-setting-tunterm-multicast-config-mdns.md) | Optional | mDNS multicast forwarding settings for tunneled VLANs |
| `MulticastAll` | `*bool` | Optional | Whether all multicast traffic is forwarded through tunnel termination<br><br>**Default**: `false` |
| `Ssdp` | [`*models.SiteSettingTuntermMulticastConfigSsdp`](../../doc/models/site-setting-tunterm-multicast-config-ssdp.md) | Optional | SSDP multicast forwarding settings for tunneled VLANs |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingTuntermMulticastConfig := models.SiteSettingTuntermMulticastConfig{
        Mdns:                 models.ToPointer(models.SiteSettingTuntermMulticastConfigMdns{
            Enabled:              models.ToPointer(false),
            VlanIds:              []int{
                246,
                247,
            },
        }),
        MulticastAll:         models.ToPointer(false),
        Ssdp:                 models.ToPointer(models.SiteSettingTuntermMulticastConfigSsdp{
            Enabled:              models.ToPointer(false),
            VlanIds:              []int{
                236,
                237,
                238,
            },
        }),
    }

}
```

