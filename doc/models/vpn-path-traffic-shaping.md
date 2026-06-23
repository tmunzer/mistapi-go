
# Vpn Path Traffic Shaping

Traffic shaping settings for a VPN path

## Structure

`VpnPathTrafficShaping`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ClassPercentage` | `[]int` | Optional | percentages for different class of traffic: high / medium / low / best-effort adding up to 100<br><br>**Constraints**: *Minimum Items*: `4`, *Maximum Items*: `4` |
| `Enabled` | `*bool` | Optional | Whether traffic shaping is enabled for this VPN path |
| `MaxTxKbps` | `models.Optional[int]` | Optional | Maximum transmit rate for this VPN path, in Kbps; `null` means no explicit limit |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vpnPathTrafficShaping := models.VpnPathTrafficShaping{
        ClassPercentage:      []int{
            71,
            72,
        },
        Enabled:              models.ToPointer(false),
        MaxTxKbps:            models.NewOptional(models.ToPointer(80)),
    }

}
```

