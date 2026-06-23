
# Nac Client Coa

Change of Authorization request for a NAC client

## Structure

`NacClientCoa`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `CoaType` | [`*models.NacCoaTypeEnum`](../../doc/models/nac-coa-type-enum.md) | Optional | CoA type to send. enum: `reauth`, `disconnect`<br><br>**Default**: `"reauth"` |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    nacClientCoa := models.NacClientCoa{
        CoaType:              models.ToPointer(models.NacCoaTypeEnum_REAUTH),
    }

}
```

