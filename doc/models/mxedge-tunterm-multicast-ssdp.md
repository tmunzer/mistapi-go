
# Mxedge Tunterm Multicast Ssdp

SSDP forwarding settings for tunnel termination

## Structure

`MxedgeTuntermMulticastSsdp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether SSDP forwarding is enabled for the configured VLANs |
| `VlanIds` | `[]string` | Optional | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeTuntermMulticastSsdp := models.MxedgeTuntermMulticastSsdp{
        Enabled:              models.ToPointer(false),
        VlanIds:              []string{
            "vlan_ids4",
        },
    }

}
```

