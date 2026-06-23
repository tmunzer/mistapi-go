
# Synthetictest Config Wan Speedtest

WAN speedtest scheduling settings for synthetic tests

## Structure

`SynthetictestConfigWanSpeedtest`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Enabled` | `*bool` | Optional | Whether scheduled WAN speedtests are enabled |
| `TimeOfDay` | `*string` | Optional | `any` / HH:MM (24-hour format)<br><br>**Default**: `"any"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    synthetictestConfigWanSpeedtest := models.SynthetictestConfigWanSpeedtest{
        Enabled:              models.ToPointer(false),
        TimeOfDay:            models.ToPointer("12:00"),
    }

}
```

