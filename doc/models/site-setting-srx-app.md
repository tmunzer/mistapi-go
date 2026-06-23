
# Site Setting Srx App

Juniper SRX application visibility settings for the site

## Structure

`SiteSettingSrxApp`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether Juniper SRX application visibility is enabled<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteSettingSrxApp := models.SiteSettingSrxApp{
        Enabled:              models.ToPointer(false),
    }

}
```

