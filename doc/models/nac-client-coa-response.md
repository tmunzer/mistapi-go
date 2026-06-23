
# Nac Client Coa Response

Response returned after sending a NAC client CoA command

## Structure

`NacClientCoaResponse`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DeviceMac` | `*string` | Optional | Target AP or switch MAC address for the CoA command |
| `DeviceType` | [`*models.DeviceTypeEnum`](../../doc/models/device-type-enum.md) | Optional | enum: `ap`, `gateway`, `switch` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacClientCoaResponse := models.NacClientCoaResponse{
        DeviceMac:            models.ToPointer("device_mac4"),
        DeviceType:           models.ToPointer(models.DeviceTypeEnum_ENUMSWITCH),
    }

}
```

