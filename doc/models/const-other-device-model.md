
# Const Other Device Model

Other-device model definition returned by the constants API

## Structure

`ConstOtherDeviceModel`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `VendorModelId` | `*string` | Optional | Vendor-specific model identifier for this other-device model |
| `Display` | `*string` | Optional | User-facing model name shown for the other-device model |
| `Model` | `*string` | Optional | Device model identifier for this other-device definition |
| `Type` | `*string` | Optional | Device category for this other-device model |
| `Vendor` | `*string` | Optional | Manufacturer name for this other-device model |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constOtherDeviceModel := models.ConstOtherDeviceModel{
        VendorModelId:        models.ToPointer("_vendor_model_id8"),
        Display:              models.ToPointer("display0"),
        Model:                models.ToPointer("model6"),
        Type:                 models.ToPointer("type2"),
        Vendor:               models.ToPointer("vendor4"),
    }

}
```

