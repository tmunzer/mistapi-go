
# Org Setting Mgmt

Organization management connectivity settings

## Structure

`OrgSettingMgmt`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `MxtunnelIds` | `[]uuid.UUID` | Optional | Mist Tunnel IDs available for organization management connectivity |
| `UseMxtunnel` | `*bool` | Optional | Whether to use Mist Tunnel for mgmt connectivity, this takes precedence over use_wxtunnel<br><br>**Default**: `false` |
| `UseWxtunnel` | `*bool` | Optional | Whether to use wxtunnel for mgmt connectivity<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
    "github.com/google/uuid"
)

func main() {
    orgSettingMgmt := models.OrgSettingMgmt{
        MxtunnelIds:          []uuid.UUID{
            uuid.MustParse("0000201e-0000-0000-0000-000000000000"),
            uuid.MustParse("0000201f-0000-0000-0000-000000000000"),
            uuid.MustParse("00002020-0000-0000-0000-000000000000"),
        },
        UseMxtunnel:          models.ToPointer(false),
        UseWxtunnel:          models.ToPointer(false),
    }

}
```

