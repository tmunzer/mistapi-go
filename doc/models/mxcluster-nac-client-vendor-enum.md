
# Mxcluster Nac Client Vendor Enum

convention to be followed is : "<vendor>-<variant>", <variant> could be an os/platform/model/company. For ex: for cisco vendor, there could variants wrt os (such as ios, nxos etc), platforms (asa etc), or acquired companies (such as meraki, aironet) etc. enum: `aruba`, `cisco-aironet`, `cisco-dnac`, `cisco-ios`, `cisco-meraki`, `brocade`, `generic`, `juniper`, `paloalto`

## Enumeration

`MxclusterNacClientVendorEnum`

## Fields

| Name |
|  --- |
| `aruba` |
| `cisco-aironet` |
| `cisco-dnac` |
| `cisco-ios` |
| `cisco-meraki` |
| `brocade` |
| `generic` |
| `juniper` |
| `paloalto` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxclusterNacClientVendor := models.MxclusterNacClientVendorEnum_GENERIC

}
```

