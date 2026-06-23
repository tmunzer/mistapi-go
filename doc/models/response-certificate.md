
# Response Certificate

If the current Org CA certificate is set to expire within 30 days, a pending certificate will be returned along with the expected auto-renewal timestamp.

## Structure

`ResponseCertificate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `string` | Required | Current organization CA certificate in PEM format |
| `PendingCert` | `*string` | Optional | Pending auto-renewed organization CA certificate returned when the current certificate nears expiry |
| `PendingCertExpiry` | `*int` | Optional | Epoch timestamp when the pending certificate is expected to replace the current certificate |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    responseCertificate := models.ResponseCertificate{
        Cert:                 "cert6",
        PendingCert:          models.ToPointer("pending_cert8"),
        PendingCertExpiry:    models.ToPointer(248),
    }

}
```

