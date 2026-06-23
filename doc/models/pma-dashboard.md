
# Pma Dashboard

PMA dashboard metadata and redirect URL

## Structure

`PmaDashboard`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Description` | `*string` | Optional | Human-readable summary of the PMA dashboard |
| `Label` | `*string` | Optional | Group label that categorizes the PMA dashboard |
| `Name` | `*string` | Optional | Display name of the PMA dashboard |
| `Url` | `*string` | Optional | Access URL that redirects the user to the PMA dashboard |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pmaDashboard := models.PmaDashboard{
        Description:          models.ToPointer("Dashboard 1 description"),
        Label:                models.ToPointer("Wireless"),
        Name:                 models.ToPointer("dashboard_1"),
        Url:                  models.ToPointer("https://api.mist.com/api/v1/forward/looker?jwt=..."),
    }

}
```

