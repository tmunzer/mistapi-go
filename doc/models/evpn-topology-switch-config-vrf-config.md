
# Evpn Topology Switch Config Vrf Config

VRF enablement for an EVPN topology switch override

## Structure

`EvpnTopologySwitchConfigVrfConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether to enable VRF (when supported on the device) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    evpnTopologySwitchConfigVrfConfig := models.EvpnTopologySwitchConfigVrfConfig{
        Enabled:              models.ToPointer(false),
    }

}
```

