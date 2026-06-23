
# Skyatp List Domain

Domain entry in a Sky ATP SecIntel list

## Structure

`SkyatpListDomain`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Comment` | `*string` | Optional | Optional note describing the domain list entry |
| `Value` | `string` | Required | Domain name included in the Sky ATP SecIntel list |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    skyatpListDomain := models.SkyatpListDomain{
        Comment:              models.ToPointer("restricted"),
        Value:                "unsafe.com",
    }

}
```

