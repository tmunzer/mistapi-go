
# Edgeconnect Register Cmd

EdgeConnect device registration command response

## Structure

`EdgeconnectRegisterCmd`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RegistrationCode` | `*string` | Optional | Registration code used to adopt an EdgeConnect device into Mist |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    edgeconnectRegisterCmd := models.EdgeconnectRegisterCmd{
        RegistrationCode:     models.ToPointer("registration_code6"),
    }

}
```

