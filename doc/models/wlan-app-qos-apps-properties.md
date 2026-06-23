
# Wlan App Qos Apps Properties

QoS rewrite settings for traffic matching a named application

## Structure

`WlanAppQosAppsProperties`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | [`*models.Dscp`](../../doc/models/containers/dscp.md) | Optional | DSCP value range between 0 and 63 |
| `DstSubnet` | `*string` | Optional | Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |
| `SrcSubnet` | `*string` | Optional | Subnet filter is not required but helps AP to only inspect certain traffic (thus reducing AP load) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAppQosAppsProperties := models.WlanAppQosAppsProperties{
        Dscp:                 models.ToPointer(models.DscpContainer.FromString("String3")),
        DstSubnet:            models.ToPointer("dst_subnet6"),
        SrcSubnet:            models.ToPointer("src_subnet4"),
    }

}
```

