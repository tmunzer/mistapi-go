
# Latlng Br

When `type`==`google`, latitude and longitude of the bottom-right corner

## Structure

`LatlngBr`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Lat` | `*string` | Optional | Bottom-right latitude for the Google map bounds |
| `Lng` | `*string` | Optional | Bottom-right longitude for the Google map bounds |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    latlngBr := models.LatlngBr{
        Lat:                  models.ToPointer("lat4"),
        Lng:                  models.ToPointer("lng2"),
    }

}
```

