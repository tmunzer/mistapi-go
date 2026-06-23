
# Mxcluster Nacedge

NAC Edge survivability settings for a Mist Edge cluster. Requires `mist_nac` to be enabled on the cluster

## Structure

`MxclusterNacedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthTtl` | `*int` | Optional | Cache TTL for last auth result in seconds<br><br>**Default**: `604800`<br><br>**Constraints**: `>= 60`, `<= 2592000` |
| `CachingSiteIds` | `[]uuid.UUID` | Optional | Site UUIDs whose auth requests are cached by NAC Edges in the cluster |
| `DefaultDot1xVlan` | `*string` | Optional | Default VLAN for all dot1x devices, if different from default_vlan |
| `DefaultVlan` | `*string` | Optional | Default VLAN to assign for devices not in the cache |
| `Enabled` | `*bool` | Optional | Whether NAC Edge survivability is enabled for this cluster<br><br>**Default**: `false` |
| `NacEdgeHosts` | `[]string` | Optional | List of Mist NAC Edge hosts for AP survivability authentication |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mxclusterNacedge := models.MxclusterNacedge{
        AuthTtl:              models.ToPointer(604800),
        CachingSiteIds:       []uuid.UUID{
            uuid.MustParse("4ac1dcf4-9d8b-7910-ac87-6ad873648a5c"),
            uuid.MustParse("7dc1acf4-9d8b-7910-ac87-6ad873648a5c"),
        },
        DefaultDot1xVlan:     models.ToPointer("20"),
        DefaultVlan:          models.ToPointer("testVlan"),
        Enabled:              models.ToPointer(false),
        NacEdgeHosts:         []string{
            "nac-west-1.corp.com",
            "nac-west-2.corp.com",
        },
    }

}
```

