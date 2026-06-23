
# Org Ssl Proxy Cert

SSL proxy certificate returned for the organization

## Structure

`OrgSslProxyCert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `*string` | Optional | PEM-encoded SSL proxy certificate for the organization |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSslProxyCert := models.OrgSslProxyCert{
        Cert:                 models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
    }

}
```

