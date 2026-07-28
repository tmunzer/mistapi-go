
# Org Setting Cacerts Config

Per-issuer CA certificate configuration used to verify client certificates

## Structure

`OrgSettingCacertsConfig`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `string` | Required | PEM-encoded CA certificate |
| `CrlEnabled` | `*bool` | Optional | Whether CRL checks are enabled. When true, CRL from AIA is used if available unless `crl_url` is set.<br><br>**Default**: `true` |
| `CrlUrl` | `*string` | Optional | Optional override URL for the certificate CRL distribution point |
| `Name` | `*string` | Optional | Optional user-friendly label for the CA issuer configuration |
| `OcspEnabled` | `*bool` | Optional | Whether OCSP checks are enabled. When true, OCSP responder from AIA is used if available unless `ocsp_url` is set.<br><br>**Default**: `true` |
| `OcspUrl` | `*string` | Optional | Optional override URL for the OCSP responder |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingCacertsConfig := models.OrgSettingCacertsConfig{
        Cert:                 "-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----",
        CrlEnabled:           models.ToPointer(true),
        CrlUrl:               models.ToPointer("https://crl.example.com/issuer1.crl"),
        Name:                 models.ToPointer("Issuer 1"),
        OcspEnabled:          models.ToPointer(true),
        OcspUrl:              models.ToPointer("https://ocsp.example.com"),
    }

}
```

