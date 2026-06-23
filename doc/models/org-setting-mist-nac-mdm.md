
# Org Setting Mist Nac Mdm

MDM (Mobile Device Management) CoA configuration

## Structure

`OrgSettingMistNacMdm`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaType` | [`*models.NacCoaTypeEnum`](../../doc/models/nac-coa-type-enum.md) | Optional | CoA type to send. enum: `reauth`, `disconnect`<br><br>**Default**: `"reauth"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingMistNacMdm := models.OrgSettingMistNacMdm{
        CoaType:              models.ToPointer(models.NacCoaTypeEnum_REAUTH),
    }

}
```

