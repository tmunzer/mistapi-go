
# Site Zone Occupancy Alert

Zone occupancy alert settings for a site

## Structure

`SiteZoneOccupancyAlert`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `EmailNotifiers` | `[]string` | Optional | List of email addresses to send email notifications when the alert threshold is reached |
| `Enabled` | `*bool` | Optional | Indicate whether zone occupancy alert is enabled for the site<br><br>**Default**: `false` |
| `Threshold` | `*int` | Optional | Sending zone-occupancy-alert webhook message only if a zone stays non-compliant (i.e. actual occupancy > occupancy_limit) for a minimum duration specified in the threshold, in minutes<br><br>**Default**: `5`<br><br>**Constraints**: `>= 0`, `<= 30` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteZoneOccupancyAlert := models.SiteZoneOccupancyAlert{
        EmailNotifiers:       []string{
            "foo@juniper.net",
            "bar@juniper.net",
        },
        Enabled:              models.ToPointer(false),
        Threshold:            models.ToPointer(5),
    }

}
```

