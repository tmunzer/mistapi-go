
# Site Setting Skyatp

Sky ATP threat intelligence settings for the site

## Structure

`SiteSettingSkyatp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Sky ATP is enabled for the site |
| `SendIpMacMapping` | `*bool` | Optional | Whether IP-to-MAC mappings are sent to Sky ATP<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingSkyatp := models.SiteSettingSkyatp{
        Enabled:              models.ToPointer(false),
        SendIpMacMapping:     models.ToPointer(false),
    }

}
```

