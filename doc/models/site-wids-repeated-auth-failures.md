
# Site Wids Repeated Auth Failures

Detection settings for repeated authentication failures

## Structure

`SiteWidsRepeatedAuthFailures`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Window where a trigger will be detected and action to be taken (in seconds) |
| `Threshold` | `*int` | Optional | Count of events to trigger |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    siteWidsRepeatedAuthFailures := models.SiteWidsRepeatedAuthFailures{
        Duration:             models.ToPointer(60),
        Threshold:            models.ToPointer(192),
    }

}
```

