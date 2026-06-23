
# Mxcluster Nac Client Ip

Mist NAC client settings for a RADIUS client IP or subnet

## Structure

`MxclusterNacClientIp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RequireMessageAuthenticator` | `*bool` | Optional | Whether to require Message-Authenticator in requests<br><br>**Default**: `false` |
| `Secret` | `*string` | Optional | Client-specific shared secret, if different from the cluster default |
| `SiteId` | `*uuid.UUID` | Optional | Present only for third-party clients, identifies the associated site |
| `Vendor` | [`*models.MxclusterNacClientVendorEnum`](../../doc/models/mxcluster-nac-client-vendor-enum.md) | Optional | convention to be followed is : "<vendor>-<variant>", <variant> could be an os/platform/model/company. For ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, aironet) etc. enum: `aruba`, `cisco-aironet`, `cisco-dnac`, `cisco-ios`, `cisco-meraki`, `brocade`, `generic`, `juniper`, `paloalto` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxclusterNacClientIp := models.MxclusterNacClientIp{
        RequireMessageAuthenticator: models.ToPointer(false),
        Secret:                      models.ToPointer("secret4"),
        SiteId:                      models.ToPointer(uuid.MustParse("00000000-0000-0000-1234-000000000000")),
        Vendor:                      models.ToPointer(models.MxclusterNacClientVendorEnum_CISCOIOS),
    }

}
```

