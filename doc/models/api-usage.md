
# Api Usage

API rate-limit usage status for the current user or API token

## Structure

`ApiUsage`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `RequestLimit` | `int` | Required, Read-only | Maximum API requests allowed in the current hourly rate-limit window<br><br>**Default**: `5000` |
| `Requests` | `int` | Required, Read-only | Number of API requests made in the current hourly rate-limit window |
| `Seconds` | `*float64` | Optional | Time remaining, in seconds, before the current hourly rate-limit window resets |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    apiUsage := models.ApiUsage{
        RequestLimit:         5000,
        Requests:             28,
        Seconds:              models.ToPointer(float64(31.84)),
    }

}
```

