
# Push Policy Push Window

If enabled, new config will only be pushed to device within the specified time window

## Structure

`PushPolicyPushWindow`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether configuration pushes are limited to the configured push window<br><br>**Default**: `false` |
| `Hours` | [`*models.Hours`](../../doc/models/hours.md) | Optional | Day-of-week operating hour filters using hour ranges such as 09:00-17:00 |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    pushPolicyPushWindow := models.PushPolicyPushWindow{
        Enabled:              models.ToPointer(false),
        Hours:                models.ToPointer(models.Hours{
            Fri:                  models.ToPointer("fri2"),
            Mon:                  models.ToPointer("mon8"),
            Sat:                  models.ToPointer("sat0"),
            Sun:                  models.ToPointer("sun6"),
            Thu:                  models.ToPointer("thu6"),
        }),
    }

}
```

