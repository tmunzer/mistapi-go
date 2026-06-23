
# Nac Auth Type Enum

enum: `cert`, `device-auth`, `eap-teap`, `eap-tls`, `eap-ttls`, `idp`, `mab`, `eap-peap`

## Enumeration

`NacAuthTypeEnum`

## Fields

| Name |
|  --- |
| `cert` |
| `device-auth` |
| `eap-teap` |
| `eap-tls` |
| `eap-ttls` |
| `idp` |
| `mab` |
| `eap-peap` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacAuthType := models.NacAuthTypeEnum_CERT

}
```

