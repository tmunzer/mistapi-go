
# Utils Zigbee Join

Request body for temporarily allowing Zigbee end-device joins

## Structure

`UtilsZigbeeJoin`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Duration` | `*int` | Optional | Number of seconds to permit new Zigbee end-device joins; range is 30-3600<br><br>**Default**: `600`<br><br>**Constraints**: `>= 30`, `<= 3600` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsZigbeeJoin := models.UtilsZigbeeJoin{
        Duration:             models.ToPointer(600),
    }

}
```

