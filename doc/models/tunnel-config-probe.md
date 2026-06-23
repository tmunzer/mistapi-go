
# Tunnel Config Probe

Tunnel health probe settings

## Structure

`TunnelConfigProbe`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Interval` | `*int` | Optional | How often to trigger the probe |
| `Threshold` | `*int` | Optional | Number of consecutive misses before declaring the tunnel down |
| `Timeout` | `*int` | Optional | Time within which to complete the connectivity check |
| `Type` | [`*models.TunnelConfigProbeTypeEnum`](../../doc/models/tunnel-config-probe-type-enum.md) | Optional | Protocol used by the custom IPsec tunnel health probe. enum: `http`, `icmp`<br><br>**Default**: `"icmp"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tunnelConfigProbe := models.TunnelConfigProbe{
        Interval:             models.ToPointer(38),
        Threshold:            models.ToPointer(98),
        Timeout:              models.ToPointer(162),
        Type:                 models.ToPointer(models.TunnelConfigProbeTypeEnum_ICMP),
    }

}
```

