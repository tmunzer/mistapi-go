
# Auto Map Assignment

Request options for validating or starting automatic AP map assignment

## Structure

`AutoMapAssignment`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Dryrun` | `*bool` | Optional | If `true`, validates the site's APs without starting the map assignment process. Returns device validity and estimated runtime.<br><br>**Default**: `false` |
| `ForceCollection` | `*bool` | Optional | If `true`, forces data collection via orchestration. If `false`, attempts to use existing BLE data first.<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    autoMapAssignment := models.AutoMapAssignment{
        Dryrun:               models.ToPointer(false),
        ForceCollection:      models.ToPointer(false),
    }

}
```

