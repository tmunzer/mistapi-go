
# Site Wids

Wireless intrusion detection settings for a site

## Structure

`SiteWids`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RepeatedAuthFailures` | [`*models.SiteWidsRepeatedAuthFailures`](../../doc/models/site-wids-repeated-auth-failures.md) | Optional | Detection settings for repeated authentication failures |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteWids := models.SiteWids{
        RepeatedAuthFailures: models.ToPointer(models.SiteWidsRepeatedAuthFailures{
            Duration:             models.ToPointer(58),
            Threshold:            models.ToPointer(170),
        }),
    }

}
```

