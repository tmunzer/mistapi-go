
# Dhcpd Config Vendor Option

Vendor-encapsulated DHCP option value

## Structure

`DhcpdConfigVendorOption`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Type` | [`*models.DhcpdConfigVendorOptionTypeEnum`](../../doc/models/dhcpd-config-vendor-option-type-enum.md) | Optional | enum: `boolean`, `hex`, `int16`, `int32`, `ip`, `string`, `uint16`, `uint32` |
| `Value` | `*string` | Optional | Option value to send for this vendor option |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    dhcpdConfigVendorOption := models.DhcpdConfigVendorOption{
        Type:                 models.ToPointer(models.DhcpdConfigVendorOptionTypeEnum_IP),
        Value:                models.ToPointer("value4"),
    }

}
```

