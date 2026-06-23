
# Radsec Server

External RadSec server settings

## Structure

`RadsecServer`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Host` | `*string` | Optional | Address or hostname of the RadSec server |
| `Port` | `*int` | Optional | TCP port used by the RadSec server<br><br>**Constraints**: `>= 1`, `<= 65535` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    radsecServer := models.RadsecServer{
        Host:                 models.ToPointer("1.1.1.1"),
        Port:                 models.ToPointer(1812),
    }

}
```

