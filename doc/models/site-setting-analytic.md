
# Site Setting Analytic

Advanced analytics feature settings for a site

## Structure

`SiteSettingAnalytic`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Enable Advanced Analytic feature (using SUB-ANA license)<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingAnalytic := models.SiteSettingAnalytic{
        Enabled:              models.ToPointer(false),
    }

}
```

