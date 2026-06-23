
# Issued Client Certificates Results

Issued client certificate search results wrapper

## Structure

`IssuedClientCertificatesResults`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Limit` | `*int` | Optional | Maximum number of results requested |
| `Page` | `*int` | Optional | Current page number |
| `Results` | [`[]models.IssuedClientCertificate`](../../doc/models/issued-client-certificate.md) | Optional | Issued client certificates returned by a SCEP certificate query |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    issuedClientCertificatesResults := models.IssuedClientCertificatesResults{
        Limit:                models.ToPointer(84),
        Page:                 models.ToPointer(198),
        Results:              []models.IssuedClientCertificate{
            models.IssuedClientCertificate{
                CertProvider:         models.ToPointer("cert_provider6"),
                CommonName:           models.ToPointer("common_name4"),
                CreatedTime:          models.ToPointer(208),
                DeviceId:             models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000")),
                ExpireTime:           models.ToPointer(238),
            },
            models.IssuedClientCertificate{
                CertProvider:         models.ToPointer("cert_provider6"),
                CommonName:           models.ToPointer("common_name4"),
                CreatedTime:          models.ToPointer(208),
                DeviceId:             models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000")),
                ExpireTime:           models.ToPointer(238),
            },
            models.IssuedClientCertificate{
                CertProvider:         models.ToPointer("cert_provider6"),
                CommonName:           models.ToPointer("common_name4"),
                CreatedTime:          models.ToPointer(208),
                DeviceId:             models.ToPointer(uuid.MustParse("00001510-0000-0000-0000-000000000000")),
                ExpireTime:           models.ToPointer(238),
            },
        },
    }

}
```

