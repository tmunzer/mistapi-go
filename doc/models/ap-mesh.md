
# Ap Mesh

Wireless mesh settings for an access point

## Structure

`ApMesh`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bands` | [`[]models.Dot11BandEnum`](../../doc/models/dot-11-band-enum.md) | Optional | List of bands that mesh should use. For relay, the first viable one will be picked. enum: `24`, `5`, `6` |
| `Enabled` | `*bool` | Optional | Whether mesh is enabled on this AP<br><br>**Default**: `false` |
| `Group` | `models.Optional[int]` | Optional | Mesh group, base AP(s) will only allow remote AP(s) in the same mesh group to join, 1-9, optional<br><br>**Constraints**: `>= 1`, `<= 9` |
| `Role` | [`*models.ApMeshRoleEnum`](../../doc/models/ap-mesh-role-enum.md) | Optional | Mesh role for this AP, either base or remote. enum: `base`, `remote` |
| `UseWpa3On5` | `*bool` | Optional | Whether to use WPA3 on the 5 GHz band for mesh links<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apMesh := models.ApMesh{
        Bands:                []models.Dot11BandEnum{
            models.Dot11BandEnum_ENUM6DEDICATED,
            models.Dot11BandEnum_ENUM6SELECTABLE,
        },
        Enabled:              models.ToPointer(false),
        Group:                models.NewOptional(models.ToPointer(1)),
        Role:                 models.ToPointer(models.ApMeshRoleEnum_BASE),
        UseWpa3On5:           models.ToPointer(false),
    }

}
```

