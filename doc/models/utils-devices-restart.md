
# Utils Devices Restart

Request to restart a device or device node

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsDevicesRestart`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Member` | `*int` | Optional | Virtual Chassis member number to restart<br><br>**Constraints**: `>= 0`, `<= 9` |
| `Node` | [`*models.UtilsDevicesRestartNodeEnum`](../../doc/models/utils-devices-restart-node-enum.md) | Optional | only for SRX/SSR: if node is not present, both nodes are restarted. For other devices: node should not be present. enum: `node0`, `node1` |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsDevicesRestart := models.UtilsDevicesRestart{
        Member:               models.ToPointer(9),
        Node:                 models.ToPointer(models.UtilsDevicesRestartNodeEnum_NODE0),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

