
# Aos Register Cmd

AOS Brownfield Registration Commands

## Structure

`AosRegisterCmd`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CliCommands` | `*string` | Optional | AOS-specific CLI commands that can be copied and pasted directly into an AOS device to register it with Mist. Includes registration code and configuration commands. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    aosRegisterCmd := models.AosRegisterCmd{
        CliCommands:          models.ToPointer("cli_commands2"),
    }

}
```

