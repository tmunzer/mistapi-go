
# Mxedge Vm Params

Mist Edge VM parameters

## Structure

`MxedgeVmParams`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `Model` | `*string` | Optional | Mist Edge VM SKU or model to deploy |
| `Name` | `*string` | Optional | User given name (optional) |
| `UserData` | `*string` | Optional | Base64 encoded user data |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    mxedgeVmParams := models.MxedgeVmParams{
        Model:                models.ToPointer("ME-VM"),
        Name:                 models.ToPointer("name6"),
        UserData:             models.ToPointer("user_data0"),
    }

}
```

