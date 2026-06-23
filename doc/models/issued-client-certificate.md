
# Issued Client Certificate

Issued Mist SCEP client certificate metadata

## Structure

`IssuedClientCertificate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CertProvider` | `*string` | Optional | Certificate provider that issued this client certificate |
| `CommonName` | `*string` | Optional | Subject common name encoded in the issued client certificate |
| `CreatedTime` | `*int` | Optional | Certificate issuance time, in epoch seconds |
| `DeviceId` | `*uuid.UUID` | Optional, Read-only | Read-only UUID identifier for a device |
| `ExpireTime` | `*int` | Optional | Certificate expiry time, in epoch seconds |
| `SerialNumber` | `*string` | Optional | Certificate serial number used for lookup or revocation |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    issuedClientCertificate := models.IssuedClientCertificate{
        CertProvider:         models.ToPointer("byod"),
        CommonName:           models.ToPointer("john@corp.com"),
        CreatedTime:          models.ToPointer(124),
        DeviceId:             models.ToPointer(uuid.MustParse("00000000-0000-0000-1000-d8695a0f9e61")),
        ExpireTime:           models.ToPointer(66),
        SerialNumber:         models.ToPointer("91984382552102771A2B3C4E5F224719956718003374658"),
    }

}
```

