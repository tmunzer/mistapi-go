
# Response Device Config Cli

Device configuration rendered as CLI commands

## Structure

`ResponseDeviceConfigCli`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cli` | `[]string` | Required | Unique string values returned or accepted by this schema<br><br>**Constraints**: *Unique Items Required* |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseDeviceConfigCli := models.ResponseDeviceConfigCli{
        Cli:                  []string{
            "cli0",
            "cli9",
            "cli8",
        },
    }

}
```

