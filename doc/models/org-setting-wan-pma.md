
# Org Setting Wan Pma

PMA feature settings for WAN Assurance

## Structure

`OrgSettingWanPma`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether PMA is enabled for WAN Assurance<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingWanPma := models.OrgSettingWanPma{
        Enabled:              models.ToPointer(false),
    }

}
```

