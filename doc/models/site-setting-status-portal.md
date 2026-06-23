
# Site Setting Status Portal

End-user status portal settings for the site

## Structure

`SiteSettingStatusPortal`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether the site status portal is enabled<br><br>**Default**: `false` |
| `Hostnames` | `[]string` | Optional | Hostnames served by the site status portal |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingStatusPortal := models.SiteSettingStatusPortal{
        Enabled:              models.ToPointer(false),
        Hostnames:            []string{
            "hostnames4",
        },
    }

}
```

