
# Mxedge Tunterm Multicast Config

Multicast forwarding settings for tunnel termination

## Structure

`MxedgeTuntermMulticastConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Mdns` | [`*models.MxedgeTuntermMulticastMdns`](../../doc/models/mxedge-tunterm-multicast-mdns.md) | Optional | mDNS forwarding settings for tunnel termination |
| `Ssdp` | [`*models.MxedgeTuntermMulticastSsdp`](../../doc/models/mxedge-tunterm-multicast-ssdp.md) | Optional | SSDP forwarding settings for tunnel termination |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermMulticastConfig := models.MxedgeTuntermMulticastConfig{
        Mdns:                 models.ToPointer(models.MxedgeTuntermMulticastMdns{
            Enabled:              models.ToPointer(false),
            VlanIds:              []string{
                "vlan_ids4",
                "vlan_ids5",
            },
        }),
        Ssdp:                 models.ToPointer(models.MxedgeTuntermMulticastSsdp{
            Enabled:              models.ToPointer(false),
            VlanIds:              []string{
                "vlan_ids2",
                "vlan_ids3",
                "vlan_ids4",
            },
        }),
    }

}
```

