
# Wlan Qos

QoS override settings for WLAN client traffic

## Structure

`WlanQos`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Class` | [`*models.WlanQosClassEnum`](../../doc/models/wlan-qos-class-enum.md) | Optional | enum: `background`, `best_effort`, `video`, `voice`<br><br>**Default**: `"best_effort"` |
| `Overwrite` | `*bool` | Optional | Whether to overwrite QoS<br><br>**Default**: `false` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    wlanQos := models.WlanQos{
        Class:                models.ToPointer(models.WlanQosClassEnum_BESTEFFORT),
        Overwrite:            models.ToPointer(false),
    }

}
```

