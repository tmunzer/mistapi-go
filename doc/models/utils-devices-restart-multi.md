
# Utils Devices Restart Multi

Request to restart multiple devices

*This model accepts additional fields of type interface{}.*

## Structure

`UtilsDevicesRestartMulti`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceIds` | `[]uuid.UUID` | Required | List of device identifiers to restart |
| `Node` | `*string` | Optional | Only for SSR: if node is not present, both nodes are restarted. For other devices: node should not be present |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    utilsDevicesRestartMulti := models.UtilsDevicesRestartMulti{
        DeviceIds:            []uuid.UUID{
            uuid.MustParse("000009eb-0000-0000-0000-000000000000"),
        },
        Node:                 models.ToPointer("node8"),
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

