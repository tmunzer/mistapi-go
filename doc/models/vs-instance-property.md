
# Vs Instance Property

EX9200 virtual-switch instance settings

## Structure

`VsInstanceProperty`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Networks` | `[]string` | Optional | Network names included in a virtual-switch instance |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    vsInstanceProperty := models.VsInstanceProperty{
        Networks:             []string{
            "networks8",
            "networks9",
        },
    }

}
```

