
# Site Setting Wired Vna

Wired Virtual Network Assistant settings for the site

## Structure

`SiteSettingWiredVna`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Wired VNA is enabled for the site<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingWiredVna := models.SiteSettingWiredVna{
        Enabled:              models.ToPointer(false),
    }

}
```

