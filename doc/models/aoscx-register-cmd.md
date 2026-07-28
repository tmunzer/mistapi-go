
# Aoscx Register Cmd

AOSCX Brownfield Registration Commands

## Structure

`AoscxRegisterCmd`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CliCommands` | `*string` | Optional | AOSCX-specific CLI commands that can be copied and pasted directly into an AOSCX device to register it with Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    aoscxRegisterCmd := models.AoscxRegisterCmd{
        CliCommands:          models.ToPointer("cli_commands2"),
    }

}
```

