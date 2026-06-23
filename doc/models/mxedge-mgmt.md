
# Mxedge Mgmt

Management settings for a Mist Edge appliance

## Structure

`MxedgeMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `ConfigAutoRevert` | `*bool` | Optional | Whether the Mist Edge automatically reverts configuration changes if connectivity is lost<br><br>**Default**: `false` |
| `FipsEnabled` | `*bool` | Optional | Whether FIPS mode is enabled on the Mist Edge<br><br>**Default**: `false` |
| `MistPassword` | `*string` | Optional | Password for the Mist service account on the Mist Edge |
| `OobIpType` | [`*models.MxedgeMgmtOobIpTypeEnum`](../../doc/models/mxedge-mgmt-oob-ip-type-enum.md) | Optional | enum: `dhcp`, `disabled`, `static`<br><br>**Default**: `"dhcp"` |
| `OobIpType6` | [`*models.MxedgeMgmtOobIpType6Enum`](../../doc/models/mxedge-mgmt-oob-ip-type-6-enum.md) | Optional | enum: `autoconf`, `dhcp`, `disabled`, `static`<br><br>**Default**: `"autoconf"` |
| `RootPassword` | `*string` | Optional | Root account password for the Mist Edge |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeMgmt := models.MxedgeMgmt{
        ConfigAutoRevert:     models.ToPointer(false),
        FipsEnabled:          models.ToPointer(false),
        MistPassword:         models.ToPointer("MIST_PASSWORD"),
        OobIpType:            models.ToPointer(models.MxedgeMgmtOobIpTypeEnum_DHCP),
        OobIpType6:           models.ToPointer(models.MxedgeMgmtOobIpType6Enum_AUTOCONF),
        RootPassword:         models.ToPointer("ROOT_PASSWORD"),
    }

}
```

