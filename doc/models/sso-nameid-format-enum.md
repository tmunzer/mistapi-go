
# Sso Nameid Format Enum

if `idp_type`==`saml`. enum: `email`, `unspecified`

## Enumeration

`SsoNameidFormatEnum`

## Fields

| Name |
|  --- |
| `email` |
| `unspecified` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    ssoNameidFormat := models.SsoNameidFormatEnum_EMAIL

}
```

