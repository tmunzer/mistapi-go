
# Utils Devices Restart Node Enum

only for SRX/SSR: if node is not present, both nodes are restarted. For other devices: node should not be present. enum: `node0`, `node1`

## Enumeration

`UtilsDevicesRestartNodeEnum`

## Fields

| Name |
|  --- |
| `node0` |
| `node1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsDevicesRestartNode := models.UtilsDevicesRestartNodeEnum_NODE0

}
```

