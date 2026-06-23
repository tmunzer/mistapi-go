
# Tunterm Port Config

Ethernet port configuration for tunnel termination interfaces

## Structure

`TuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | List of ports to be used for downstream (to AP) purpose |
| `SeparateUpstreamDownstream` | `*bool` | Optional | Whether to separate upstream / downstream ports. default is false where all ports will be used.<br><br>**Default**: `false` |
| `UpstreamPortVlanId` | [`*models.TuntermPortConfigUpstreamPortVlanId`](../../doc/models/containers/tunterm-port-config-upstream-port-vlan-id.md) | Optional | Native VLAN ID for upstream ports |
| `UpstreamPorts` | `[]string` | Optional | List of ports to be used for upstream purpose (to LAN) |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    tuntermPortConfig := models.TuntermPortConfig{
        DownstreamPorts:            []string{
            "2",
            "3",
        },
        SeparateUpstreamDownstream: models.ToPointer(false),
        UpstreamPortVlanId:         models.ToPointer(models.TuntermPortConfigUpstreamPortVlanIdContainer.FromString("String3")),
        UpstreamPorts:              []string{
            "0",
            "1",
        },
    }

}
```

