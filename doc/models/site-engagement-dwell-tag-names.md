
# Site Engagement Dwell Tag Names

Display labels for engagement dwell-time categories

## Structure

`SiteEngagementDwellTagNames`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Bounce` | `*string` | Optional | Display label for bounce visits<br><br>**Default**: `"Visitor"` |
| `Engaged` | `*string` | Optional | Display label for engaged visits<br><br>**Default**: `"Associates"` |
| `Passerby` | `*string` | Optional | Display label for passerby visits<br><br>**Default**: `"Passerby"` |
| `Stationed` | `*string` | Optional | Display label for stationed visits<br><br>**Default**: `"Assets"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteEngagementDwellTagNames := models.SiteEngagementDwellTagNames{
        Bounce:               models.ToPointer("Bounce"),
        Engaged:              models.ToPointer("Engaged"),
        Passerby:             models.ToPointer("Passer By"),
        Stationed:            models.ToPointer("Stationed"),
    }

}
```

