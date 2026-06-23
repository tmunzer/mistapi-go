
# Response Router Ssr Register Cmd

SSR registration token and commands used to register the router with Mist

## Structure

`ResponseRouterSsrRegisterCmd`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConductorCmd` | `*string` | Optional | Command to run in the SSR conductor CLI for Mist registration |
| `RegistrationCode` | `*string` | Optional | Registration token used by the SSR registration commands |
| `RouterShellCmd` | `*string` | Optional | Shell command to run on the SSR router for Mist registration |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseRouterSsrRegisterCmd := models.ResponseRouterSsrRegisterCmd{
        ConductorCmd:         models.ToPointer("conductor_cmd0"),
        RegistrationCode:     models.ToPointer("registration_code6"),
        RouterShellCmd:       models.ToPointer("router_shell_cmd0"),
    }

}
```

