
# Latlng Tl

When `type`==`google`, latitude and longitude of the top-left corner

## Structure

`LatlngTl`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Lat` | `*string` | Optional | Top-left latitude for the Google map bounds |
| `Lng` | `*string` | Optional | Top-left longitude for the Google map bounds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    latlngTl := models.LatlngTl{
        Lat:                  models.ToPointer("lat0"),
        Lng:                  models.ToPointer("lng6"),
    }

}
```

