
# Idp User Cert Lookup Field Enum

allow customer to choose the EAP-TLS client certificate's field. To use for IDP User Groups lookup. enum: `automatic`, `cn`, `email`, `upn`

## Enumeration

`IdpUserCertLookupFieldEnum`

## Fields

| Name |
|  --- |
| `automatic` |
| `cn` |
| `email` |
| `upn` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    idpUserCertLookupField := models.IdpUserCertLookupFieldEnum_AUTOMATIC

}
```

