
# Hours

Day-of-week operating hour filters using hour ranges such as 09:00-17:00

## Structure

`Hours`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Fri` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Mon` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Sat` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Sun` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Thu` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Tue` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |
| `Wed` | `*string` | Optional | Hour range of the day (e.g. `09:00-17:00`). If the hour is not defined then it's treated as 00:00-23:59. |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    hours := models.Hours{
        Fri:                  models.ToPointer("09:00-17:00"),
        Mon:                  models.ToPointer("09:00-17:00"),
        Sat:                  models.ToPointer("09:00-17:00"),
        Sun:                  models.ToPointer("09:00-17:00"),
        Thu:                  models.ToPointer("09:00-17:00"),
        Tue:                  models.ToPointer("09:00-17:00"),
        Wed:                  models.ToPointer("09:00-17:00"),
    }

}
```

