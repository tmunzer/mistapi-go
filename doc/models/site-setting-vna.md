
# Site Setting Vna

Virtual Network Assistant settings for AP, switch, and gateway experiences at a site

## Structure

`SiteSettingVna`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Enable Virtual Network Assistant (using SUB-VNA license). This applied to AP / Switch / Gateway<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingVna := models.SiteSettingVna{
        Enabled:              models.ToPointer(false),
    }

}
```

