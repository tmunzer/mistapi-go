
# Org Setting Mist Nac Server Cert

RADIUS server certificate presented by Mist NAC during EAP-TLS

## Structure

`OrgSettingMistNacServerCert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `*string` | Optional | PEM-encoded RADIUS server certificate presented during EAP-TLS |
| `Key` | `*string` | Optional | Private key paired with the Mist NAC RADIUS server certificate |
| `Password` | `*string` | Optional | Optional password for the private key |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingMistNacServerCert := models.OrgSettingMistNacServerCert{
        Cert:                 models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
        Key:                  models.ToPointer("-----BEGIN PRI..."),
        Password:             models.ToPointer("password4"),
    }

}
```

