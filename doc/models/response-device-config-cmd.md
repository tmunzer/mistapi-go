
# Response Device Config Cmd

Device configuration command response

## Structure

`ResponseDeviceConfigCmd`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cmd` | `string` | Required | Configuration command returned for the device |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDeviceConfigCmd := models.ResponseDeviceConfigCmd{
        Cmd:                  "cmd0",
    }

}
```

