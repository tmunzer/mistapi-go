
# Org Setting Wireless Pma

PMA feature settings for Wireless Assurance

## Structure

`OrgSettingWirelessPma`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether PMA is enabled for Wireless Assurance<br><br>**Default**: `true` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    orgSettingWirelessPma := models.OrgSettingWirelessPma{
        Enabled:              models.ToPointer(true),
    }

}
```

