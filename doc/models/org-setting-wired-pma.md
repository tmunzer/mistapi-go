
# Org Setting Wired Pma

PMA feature settings for Wired Assurance

## Structure

`OrgSettingWiredPma`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether PMA is enabled for Wired Assurance<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingWiredPma := models.OrgSettingWiredPma{
        Enabled:              models.ToPointer(false),
    }

}
```

