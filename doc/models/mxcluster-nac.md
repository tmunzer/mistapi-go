
# Mxcluster Nac

Mist NAC RADIUS settings for a Mist Edge cluster. Used when the Mist Edge Cluster is used as a RADIUS Proxy between the local devices and the Mist NAC

## Structure

`MxclusterNac`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AcctServerPort` | `*int` | Optional | RADIUS accounting port used by Mist NAC on the cluster<br><br>**Default**: `1813` |
| `AuthServerPort` | `*int` | Optional | RADIUS authentication port used by Mist NAC on the cluster<br><br>**Default**: `1812` |
| `ClientIps` | [`map[string]models.MxclusterNacClientIp`](../../doc/models/mxcluster-nac-client-ip.md) | Optional | Property key is the RADIUS Client IP/Subnet. |
| `Enabled` | `*bool` | Optional | Whether Mist NAC is enabled on the cluster<br><br>**Default**: `false` |
| `Secret` | `*string` | Optional | Shared RADIUS secret used by Mist NAC clients |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxclusterNac := models.MxclusterNac{
        AcctServerPort:       models.ToPointer(1813),
        AuthServerPort:       models.ToPointer(1812),
        ClientIps:            map[string]models.MxclusterNacClientIp{
            "key0": models.MxclusterNacClientIp{
                RequireMessageAuthenticator: models.ToPointer(false),
                Secret:                      models.ToPointer("secret4"),
                SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
            },
            "key1": models.MxclusterNacClientIp{
                RequireMessageAuthenticator: models.ToPointer(false),
                Secret:                      models.ToPointer("secret4"),
                SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
            },
            "key2": models.MxclusterNacClientIp{
                RequireMessageAuthenticator: models.ToPointer(false),
                Secret:                      models.ToPointer("secret4"),
                SiteId:                      models.ToPointer(uuid.MustParse("0000197c-0000-0000-0000-000000000000")),
                Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOAIRONET),
            },
        },
        Enabled:              models.ToPointer(false),
        Secret:               models.ToPointer("testing123"),
    }

}
```

