
# Mist Nacedge

Mist NAC Site Survivability settings for the site

## Structure

`MistNacedge`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AuthTtl` | `*int` | Optional | Cache of last auth result; in seconds<br><br>**Default**: `604800`<br><br>**Constraints**: `>= 60`, `<= 2592000` |
| `CachingSiteIds` | `[]uuid.UUID` | Optional | Site UUIDs whose auth requests are cached by NAC Edges in the cluster |
| `DefaultDot1xVlan` | `*string` | Optional | Default vlan for all dot1x devices, if different from default_vlan |
| `DefaultVlan` | `*string` | Optional | Default vlan to assign for devices not in the cache |
| `Enabled` | `*bool` | Optional | Whether Mist Site Survivability is enabled for the site |
| `MxedgeHosts` | `[]string` | Optional | List of NAC Edges used for the Site Survivability feature |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    mistNacedge := models.MistNacedge{
        AuthTtl:              models.ToPointer(604800),
        CachingSiteIds:       []uuid.UUID{
            uuid.MustParse("4ac1dcf4-9d8b-7910-ac87-6ad873648a5c"),
            uuid.MustParse("7dc1acf4-9d8b-7910-ac87-6ad873648a5c"),
        },
        DefaultDot1xVlan:     models.ToPointer("20"),
        DefaultVlan:          models.ToPointer("test_vlan"),
        Enabled:              models.ToPointer(false),
        MxedgeHosts:          []string{
            "mxedge1.local",
        },
    }

}
```

