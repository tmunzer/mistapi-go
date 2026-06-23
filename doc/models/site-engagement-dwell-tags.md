
# Site Engagement Dwell Tags

Visit duration ranges in seconds used to assign engagement tags

## Structure

`SiteEngagementDwellTags`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bounce` | `models.Optional[string]` | Optional | Visit duration range for bounce visits, in seconds<br><br>**Default**: `"301-14400"` |
| `Engaged` | `models.Optional[string]` | Optional | Visit duration range for engaged visits, in seconds<br><br>**Default**: `"14401-28800"` |
| `Passerby` | `models.Optional[string]` | Optional | Visit duration range for passerby visits, in seconds<br><br>**Default**: `"1-300"` |
| `Stationed` | `models.Optional[string]` | Optional | Visit duration range for stationed visits, in seconds<br><br>**Default**: `"28801-42000"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteEngagementDwellTags := models.SiteEngagementDwellTags{
        Bounce:               models.NewOptional(models.ToPointer("301-14400")),
        Engaged:              models.NewOptional(models.ToPointer("14401-28800")),
        Passerby:             models.NewOptional(models.ToPointer("1-300")),
        Stationed:            models.NewOptional(models.ToPointer("28801-42000")),
    }

}
```

