
# Response Set Devices Map

Result of assigning APs to a site map

## Structure

`ResponseSetDevicesMap`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Locked` | `[]string` | Optional | AP MAC addresses that were locked during map assignment |
| `Moved` | `[]string` | Optional | AP MAC addresses moved during map assignment |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseSetDevicesMap := models.ResponseSetDevicesMap{
        Locked:               []string{
            "locked4",
            "locked5",
            "locked6",
        },
        Moved:                []string{
            "moved1",
            "moved0",
            "moved9",
        },
    }

}
```

