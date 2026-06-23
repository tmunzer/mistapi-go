
# Idp Machine Cert Lookup Field Enum

allow customer to choose the EAP-TLS client certificate's field to use for IDP Machine Groups lookup. enum: `automatic`, `cn`, `dns`

## Enumeration

`IdpMachineCertLookupFieldEnum`

## Fields

| Name |
|  --- |
| `automatic` |
| `cn` |
| `dns` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    idpMachineCertLookupField := models.IdpMachineCertLookupFieldEnum_AUTOMATIC

}
```

