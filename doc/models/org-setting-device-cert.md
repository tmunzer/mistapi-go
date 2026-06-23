
# Org Setting Device Cert

Optional common device certificate configuration for organization settings

## Structure

`OrgSettingDeviceCert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Cert` | `*string` | Optional | PEM-encoded common device certificate used by organization settings |
| `Key` | `*string` | Optional | Private key paired with the common device certificate |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingDeviceCert := models.OrgSettingDeviceCert{
        Cert:                 models.ToPointer("-----BEGIN CERTIFICATE-----\\nMIIFZjCCA06gAwIBAgIIP61/1qm/uDowDQYJKoZIhvcNAQELBQE\\n-----END CERTIFICATE-----"),
        Key:                  models.ToPointer("-----BEGIN PRI..."),
    }

}
```

