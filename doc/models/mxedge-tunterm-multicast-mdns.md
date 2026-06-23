
# Mxedge Tunterm Multicast Mdns

mDNS forwarding settings for tunnel termination

## Structure

`MxedgeTuntermMulticastMdns`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether mDNS forwarding is enabled for the configured VLANs |
| `VlanIds` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermMulticastMdns := models.MxedgeTuntermMulticastMdns{
        Enabled:              models.ToPointer(false),
        VlanIds:              []string{
            "vlan_ids2",
            "vlan_ids3",
        },
    }

}
```

