
# Site Setting Wan Vna

WAN Virtual Network Assistant settings for the site

## Structure

`SiteSettingWanVna`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether WAN VNA is enabled for the site<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingWanVna := models.SiteSettingWanVna{
        Enabled:              models.ToPointer(false),
    }

}
```

