
# Switch Network Multicast

Multicast settings for a switch network (VLAN)

## Structure

`SwitchNetworkMulticast`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable IGMP snooping on this VLAN<br><br>**Default**: `false` |
| `IgmpVersion` | [`*models.SwitchNetworkMulticastIgmpVersionEnum`](../../doc/models/switch-network-multicast-igmp-version-enum.md) | Optional | IGMP version. '2' (default, ASM/IGMPv2) / '3' (SSM/IGMPv3)<br><br>**Default**: `"2"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    switchNetworkMulticast := models.SwitchNetworkMulticast{
        Enabled:              models.ToPointer(false),
        IgmpVersion:          models.ToPointer(models.SwitchNetworkMulticastIgmpVersionEnum_ENUM2),
    }

}
```

