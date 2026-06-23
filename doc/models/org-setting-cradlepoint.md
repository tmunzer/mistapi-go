
# Org Setting Cradlepoint

Read-only Cradlepoint integration settings stored for the organization

## Structure

`OrgSettingCradlepoint`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CpApiId` | `*string` | Optional, Read-only | Cradlepoint API ID used by Mist for the integration |
| `CpApiKey` | `*string` | Optional, Read-only | Cradlepoint API key paired with the Cradlepoint API ID |
| `EcmApiId` | `*string` | Optional, Read-only | Cradlepoint ECM API ID used by Mist for the integration |
| `EcmApiKey` | `*string` | Optional, Read-only | Cradlepoint ECM API key paired with the ECM API ID |
| `EnableLldp` | `*bool` | Optional, Read-only | Whether Mist uses Cradlepoint LLDP data to link routers to Mist sites and devices |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingCradlepoint := models.OrgSettingCradlepoint{
        CpApiId:              models.ToPointer("84446d61-2206-4ea5-855a-0043f980be54"),
        CpApiKey:             models.ToPointer("79c329da9893e34099c7d8ad5cb9c941"),
        EcmApiId:             models.ToPointer("73446d61-2206-4ea5-855a-0043f980be62"),
        EcmApiKey:            models.ToPointer("68b329da9893e34099c7d8ad5cb9c940"),
        EnableLldp:           models.ToPointer(false),
    }

}
```

