
# Site Setting Ap Synthetic Test

AP Synthetic Test configuration

## Structure

`SiteSettingApSyntheticTest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AdditionalVlanIds` | [`*models.AdditionalVlanIds`](../../doc/models/containers/additional-vlan-ids.md) | Optional | List or Comma separated list of additional VLAN IDs (on the LAN side or from other WLANs) should we be forwarding bonjour queries/responses |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingApSyntheticTest := models.SiteSettingApSyntheticTest{
        AdditionalVlanIds:    models.ToPointer(models.AdditionalVlanIdsContainer.FromString("String7")),
    }

}
```

