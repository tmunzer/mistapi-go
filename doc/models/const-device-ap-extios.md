
# Const Device Ap Extios

External I/O port capabilities for an AP model

## Structure

`ConstDeviceApExtios`

## Fields

| Name | Type | Tags | Description |
|  --- | --- | --- | --- |
| `DefaultDir` | [`*models.ConstDeviceApExtiosDefaultDirEnum`](../../doc/models/const-device-ap-extios-default-dir-enum.md) | Optional | Default direction for this external I/O port. enum: `IN`, `OUT` |
| `Input` | `*bool` | Optional | Whether this external I/O port supports input mode |
| `Output` | `*bool` | Optional | Whether this external I/O port supports output mode |

## Example

```go
package main

import (
    "mistapi/models"
)

func main() {
    constDeviceApExtios := models.ConstDeviceApExtios{
        DefaultDir:           models.ToPointer(models.ConstDeviceApExtiosDefaultDirEnum_IN),
        Input:                models.ToPointer(false),
        Output:               models.ToPointer(false),
    }

}
```

