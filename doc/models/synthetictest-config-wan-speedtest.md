
# Synthetictest Config Wan Speedtest

WAN speedtest scheduling settings for synthetic tests

## Structure

`SynthetictestConfigWanSpeedtest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Disabled` | `*bool` | Optional | Whether scheduled WAN speedtests are disabled. Defaults to `false` (enabled); set this to `true` to disable speedtests.<br><br>**Default**: `false` |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format)<br><br>**Default**: `"any"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestConfigWanSpeedtest := models.SynthetictestConfigWanSpeedtest{
        Disabled:             models.ToPointer(false),
        TimeOfDay:            models.ToPointer("12:00"),
    }

}
```

