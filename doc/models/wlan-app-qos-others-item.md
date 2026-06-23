
# Wlan App Qos Others Item

Custom QoS rewrite rule for traffic not matched by a named application

## Structure

`WlanAppQosOthersItem`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dscp` | [`*models.Dscp`](../../doc/models/containers/dscp.md) | Optional | DSCP value range between 0 and 63 |
| `DstSubnet` | `*string` | Optional | Destination subnet filter for this custom QoS rule |
| `PortRanges` | `*string` | Optional | TCP or UDP port ranges matched by this custom QoS rule |
| `Protocol` | `*string` | Optional | IP protocol matched by this custom QoS rule |
| `SrcSubnet` | `*string` | Optional | Source subnet filter for this custom QoS rule |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanAppQosOthersItem := models.WlanAppQosOthersItem{
        Dscp:                 models.ToPointer(models.DscpContainer.FromString("String5")),
        DstSubnet:            models.ToPointer("10.2.0.0/16"),
        PortRanges:           models.ToPointer("80,1024-6553"),
        Protocol:             models.ToPointer("udp"),
        SrcSubnet:            models.ToPointer("10.2.0.0/16"),
    }

}
```

