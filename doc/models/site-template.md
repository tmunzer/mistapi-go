
# Site Template

Site template containing auto-upgrade settings and template variables

*This model accepts additional fields of type interface{}.*

## Structure

`SiteTemplate`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `AutoUpgrade` | [`*models.SiteTemplateAutoUpgrade`](../../doc/models/site-template-auto-upgrade.md) | Optional | Automatic upgrade settings applied by a site template |
| `Name` | `*string` | Optional | Display name of the site template |
| `Vars` | `map[string]string` | Optional | Dictionary of name->value, the vars can then be used in Wlans. This can overwrite those from Site Vars |
| `AdditionalProperties` | `map[string]interface{}` | Optional | - |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteTemplate := models.SiteTemplate{
        AutoUpgrade:          models.ToPointer(models.SiteTemplateAutoUpgrade{
            DayOfWeek:            models.ToPointer(models.DayOfWeekEnum_SUN),
            Enabled:              models.ToPointer(false),
            TimeOfDay:            models.ToPointer("time_of_day0"),
            Version:              models.ToPointer("version2"),
        }),
        Name:                 models.ToPointer("name6"),
        Vars:                 map[string]string{
            "RADIUS_IP1": "172.31.2.5",
            "RADIUS_SECRET": "11s64632d",
        },
        AdditionalProperties: map[string]interface{}{
            "exampleAdditionalProperty": interface{}("[key1, val1][key2, val2]"),
        },
    }

}
```

