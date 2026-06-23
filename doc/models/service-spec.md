
# Service Spec

Protocol and port match rule for a custom service

## Structure

`ServiceSpec`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `PortRange` | `*string` | Optional | Port number, port range, or variable |
| `Protocol` | `*string` | Optional | `https`/ `tcp` / `udp` / `icmp` / `gre` / `any` / `:protocol_number`, `protocol_number` is between 1-254<br><br>**Default**: `"any"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    serviceSpec := models.ServiceSpec{
        PortRange:            models.ToPointer("8080,8443"),
        Protocol:             models.ToPointer("tcp"),
    }

}
```

