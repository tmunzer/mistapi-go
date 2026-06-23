
# Dhcpd Config Option

Custom DHCP option value

## Structure

`DhcpdConfigOption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.DhcpdConfigOptionTypeEnum`](../../doc/models/dhcpd-config-option-type-enum.md) | Optional | enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32` |
| `Value` | `*string` | Optional | Option value to send for this DHCP option |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdConfigOption := models.DhcpdConfigOption{
        Type:                 models.ToPointer(models.DhcpdConfigOptionTypeEnum_ENUMUINT16),
        Value:                models.ToPointer("value2"),
    }

}
```

