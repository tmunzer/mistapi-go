
# Stats Mxedge Tunterm Port Config

Tunnel termination port role configuration reported by a Mist Edge

## Structure

`StatsMxedgeTuntermPortConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DownstreamPorts` | `[]string` | Optional | Downstream tunnel termination port identifiers |
| `SeparateUpstreamDownstream` | `*bool` | Optional | Whether separate port sets are used for upstream and downstream tunnel termination traffic |
| `UpstreamPorts` | `[]string` | Optional | Upstream tunnel termination port identifiers |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    statsMxedgeTuntermPortConfig := models.StatsMxedgeTuntermPortConfig{
        DownstreamPorts:            []string{
            "0",
            "1",
        },
        SeparateUpstreamDownstream: models.ToPointer(false),
        UpstreamPorts:              []string{
            "0",
            "1",
        },
    }

}
```

