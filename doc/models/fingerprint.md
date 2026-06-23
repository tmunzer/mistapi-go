
# Fingerprint

Client device fingerprint record returned by NAC fingerprint insights

## Structure

`Fingerprint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Family` | `*string` | Optional, Read-only | Device family or category inferred from client fingerprinting |
| `Mac` | `*string` | Optional, Read-only | Client device MAC address for the fingerprint record |
| `Mfg` | `*string` | Optional, Read-only | Manufacturer name inferred from client fingerprinting |
| `Model` | `*string` | Optional, Read-only | Device model inferred from client fingerprinting |
| `OrgId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist organization |
| `Os` | `*string` | Optional, Read-only | Operating system name and version inferred from client fingerprinting |
| `OsType` | `*string` | Optional, Read-only | Operating system family inferred from client fingerprinting |
| `RandomMac` | `*bool` | Optional, Read-only | Whether the client device uses a randomized MAC address |
| `SiteId` | `*uuid.UUID` | Optional, Read-only | Unique identifier of a Mist site |
| `Timestamp` | `*float64` | Optional, Read-only | Epoch timestamp, in seconds |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    fingerprint := models.Fingerprint{
        Family:               models.ToPointer("family8"),
        Mac:                  models.ToPointer("mac0"),
        Mfg:                  models.ToPointer("mfg4"),
        Model:                models.ToPointer("model4"),
        OrgId:                models.ToPointer(uuid.MustParse("a97c1b22-a4e9-411e-9bfd-d8695a0f9e61")),
        SiteId:               models.ToPointer(uuid.MustParse("441a1214-6928-442a-8e92-e1d34b8ec6a6")),
    }

}
```

