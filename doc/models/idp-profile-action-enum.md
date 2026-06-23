
# Idp Profile Action Enum

IDP profile action. enum: `alert`, `close`, `drop`. `alert` is the default, `drop` silently drops packets, and `close` asks the client/server to close the connection

## Enumeration

`IdpProfileActionEnum`

## Fields

| Name |
|  --- |
| `alert` |
| `close` |
| `drop` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    idpProfileAction := models.IdpProfileActionEnum_ALERT

}
```

