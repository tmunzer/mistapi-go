
# Utils Zigbee Kick

Request body for kicking one or more Zigbee clients from an AP

## Structure

`UtilsZigbeeKick`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Macs` | `[]string` | Required | One or more Zigbee EUI-64 (8-byte) MACs. Accepts colon-separated (`00:17:7a:01:06:0c:ae:9f`) or plain hex (`00177a01060cae9f`). Must be non-empty.<br><br>**Constraints**: *Minimum Items*: `1` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    utilsZigbeeKick := models.UtilsZigbeeKick{
        Macs:                 []string{
            "00177a01060cae9f",
            "00177a01060caea1",
        },
    }

}
```

