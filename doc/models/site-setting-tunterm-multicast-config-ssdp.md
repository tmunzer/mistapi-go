
# Site Setting Tunterm Multicast Config Ssdp

SSDP multicast forwarding settings for tunneled VLANs

## Structure

`SiteSettingTuntermMulticastConfigSsdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether SSDP multicast forwarding is enabled<br><br>**Default**: `false` |
| `VlanIds` | `[]int` | Optional | VLAN IDs where SSDP forwarding is enabled |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingTuntermMulticastConfigSsdp := models.SiteSettingTuntermMulticastConfigSsdp{
        Enabled:              models.ToPointer(false),
        VlanIds:              []int{
            2,
            3,
            5,
        },
    }

}
```

