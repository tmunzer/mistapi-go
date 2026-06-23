
# Site Engagement

Engagement analytics dwell-time rules for classifying site visits. If hours is omitted, rules apply every day from 00:00 to 23:59. Multiple ranges for the same day are not supported.

## Structure

`SiteEngagement`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DwellTagNames` | [`*models.SiteEngagementDwellTagNames`](../../doc/models/site-engagement-dwell-tag-names.md) | Optional | Display labels for engagement dwell-time categories |
| `DwellTags` | [`*models.SiteEngagementDwellTags`](../../doc/models/site-engagement-dwell-tags.md) | Optional | Visit duration ranges in seconds used to assign engagement tags |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | Day-of-week operating hour filters using hour ranges such as 09:00-17:00 |
| `MaxDwell` | `*int` | Optional | Maximum dwell time in seconds considered by engagement analytics<br><br>**Default**: `43200`<br><br>**Constraints**: `>= 1`, `<= 68400` |
| `MinDwell` | `*int` | Optional | Minimum dwell time in seconds for engagement analytics<br><br>**Constraints**: `>= 0` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteEngagement := models.SiteEngagement{
        DwellTagNames:        models.ToPointer(models.SiteEngagementDwellTagNames{
            Bounce:               models.ToPointer("bounce0"),
            Engaged:              models.ToPointer("engaged2"),
            Passerby:             models.ToPointer("passerby6"),
            Stationed:            models.ToPointer("stationed4"),
        }),
        DwellTags:            models.ToPointer(models.SiteEngagementDwellTags{
            Bounce:               models.NewOptional(models.ToPointer("bounce0")),
            Engaged:              models.NewOptional(models.ToPointer("engaged2")),
            Passerby:             models.NewOptional(models.ToPointer("passerby6")),
            Stationed:            models.NewOptional(models.ToPointer("stationed6")),
        }),
        Hours:                models.ToPointer(models.Hours{
            Fri:                  models.ToPointer("fri2"),
            Mon:                  models.ToPointer("mon8"),
            Sat:                  models.ToPointer("sat0"),
            Sun:                  models.ToPointer("sun6"),
            Thu:                  models.ToPointer("thu6"),
        }),
        MaxDwell:             models.ToPointer(43200),
        MinDwell:             models.ToPointer(38),
    }

}
```

